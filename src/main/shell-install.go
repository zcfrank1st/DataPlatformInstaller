package main

import (
    "bufio"
    "os"
    "fmt"
    "github.com/logrusorgru/aurora"
    "module"
    "strings"
    "util"
)

var scanner *bufio.Scanner

func init () {
    scanner = bufio.NewScanner(os.Stdin)
}

func readConsole() string{
    for {
        scanner.Scan()
        text := scanner.Text()
        if "" != text {
            return text
        }
    }
}

func printAlert() {
    fmt.Println(aurora.Red("[ERROR] Not support input, please retry !"))
}

func printInstallStep(step int, moduleName string) {
    fmt.Println(aurora.Green(fmt.Sprintf("%d) Install %s? (Yy/Nn)", step, moduleName)))
}

func printAddNodeStep(step int, moduleName string) {
    fmt.Println(aurora.Green(fmt.Sprintf("%d) Add %s node ? (Yy/Nn)", step, moduleName)))
}

func printNodes() {
    fmt.Println(aurora.Magenta(fmt.Sprintf("IPs from %s", util.Nodes)))
}

// todo
func installModule(moduleName string) {
    printNodes()
    switch moduleName {
    case "Hadoop":
        master := ""
        slaves := []string{}

        for {
            fmt.Println(aurora.Green("[Master]: "))
            master = readConsole()
            checkRes := util.CheckIfInLicencedIps([]string{master})
            if checkRes {
                fmt.Println("master: "  + master)
                break
            } else {
                fmt.Println(aurora.Red("[Warning] not supported IP(s)! Please retry "))
            }
        }

        for {
            fmt.Println(aurora.Green("[Slaves](use , to split): "))
            slaves = strings.Split(strings.Replace(readConsole(), " ", "", -1), ",")
            checkRes := util.CheckIfInLicencedIps(slaves)
            if checkRes {
                fmt.Println(slaves)
                break
            } else {
                fmt.Println(aurora.Red("[Warning] not supported IP(s)! Please retry "))
            }
        }

        err := module.InstallHadoop(master, slaves)
        if err != nil {
            fmt.Println(aurora.Red(err))
            os.Exit(10)
        }
    case "Zookeeper":
        zkNodes := []string{}

        for {
            fmt.Println(aurora.Green("[Nodes](use , to split): "))
            zkNodes = strings.Split(strings.Replace(readConsole(), " ", "", -1), ",")
            checkRes := util.CheckIfInLicencedIps(zkNodes)
            if checkRes {
                fmt.Println(zkNodes)
                break
            } else {
                fmt.Println(aurora.Red("[Warning] not supported IP(s)! Please retry "))
            }
        }

        err := module.InstallZookeeper(zkNodes)
        if err != nil {
            fmt.Println(aurora.Red(err))
            os.Exit(11)
        }

    case "Hbase":
    case "Kafka":
    case "Storm":
    }
}

// todo
func addModule(moduleName string) {
    printNodes()
    switch moduleName {
    case "Hadoop":
    case "Zookeeper":
    case "Hbase":
    case "Kafka":
    case "Storm":
    }
}

func installOrAddLoop(moduleName string, f func (string)) {
    for {
        inputIgnoreCase := strings.ToLower(readConsole())
        if "y" == inputIgnoreCase {
            f(moduleName)
            break
        } else if "n" == inputIgnoreCase {
            break
        } else {
            printAlert()
        }
    }
}

func installPhase (step int, moduleName string) {
    printInstallStep(step, moduleName)
    installOrAddLoop(moduleName, installModule)
}

func addNodePhase (step int, moduleName string) {
    printAddNodeStep(step, moduleName)
    installOrAddLoop(moduleName, addModule)
}


func main() {
    fmt.Println(aurora.Blue(`
 ____  ____ ___
|  _ \|  _ \_ _|
| | | | |_) | |
| |_| |  __/| |
|____/|_|  |___| v` + util.VERSION))

    fmt.Println()


    //初始安装licence，增加节点licence
    fmt.Println(aurora.Cyan("@@[Licence] Please input install licence :"))
    typ, err, num := util.CheckLicence(readConsole())
    if err != nil {
        fmt.Println(aurora.Red("[ERROR] Invalid Licence !"))
        os.Exit(1)
    }

    fmt.Println(aurora.Green("Please input Nodes IPs (use , to split) :"))
    err1 := util.CheckLicencedIPs(num, readConsole())
    if err1 != nil {
        fmt.Println(aurora.Red("[ERROR] Invalid Input IPs !"))
        os.Exit(2)
    }

    if 0 == typ {
        fmt.Println(aurora.Magenta("Loading dependencies from the repo, Please waiting ... "))
        // DPI-0.1.0.box
        wgetErr := util.Wget(util.RESOURCE_URL + "/DPI-" + util.VERSION + ".box", "/opt")
        if wgetErr != nil {
            fmt.Println(aurora.Red("[ERROR] fetch dependency box error!"))
            os.Exit(3)
        }
        unTarErr := util.UnTar("/opt/DPI-" + util.VERSION + ".box")
        if unTarErr != nil {
            fmt.Println(aurora.Red("[ERROR] release modules error!"))
            os.Exit(4)
        }

        fmt.Println(aurora.Blue("Install Process Start"))

        moduleName := "Hadoop"
        installPhase(1, moduleName)

        moduleName = "Zookeeper"
        installPhase(2, moduleName)

        moduleName = "Hbase"
        installPhase(3, moduleName)

        moduleName = "Kafka"
        installPhase(4, moduleName)

        moduleName = "Storm"
        installPhase(5, moduleName)

        // default: install DPMonitor and DPMetrics

    } else if 1 == typ {
        fmt.Println(aurora.Blue("Add Node Process Start"))

        if !util.IfFirstInstallDirsExists() {
            fmt.Println(aurora.Red("[ERROR] Can not find the install box ! "))
            os.Exit(3)
        }

        moduleName := "Hadoop"
        addNodePhase(1, moduleName)

        moduleName = "Zookeeper"
        addNodePhase(2, moduleName)

        moduleName = "Hbase"
        addNodePhase(3, moduleName)

        moduleName = "Kafka"
        addNodePhase(4, moduleName)

        moduleName = "Storm"
        addNodePhase(5, moduleName)
    }

    util.SaveOrUpdateConfigToLocal()
    fmt.Println(aurora.Magenta("All done, goodbye !"))
}
