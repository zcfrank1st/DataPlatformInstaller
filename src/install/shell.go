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

func installModule(moduleName string) {
    // todo install modules
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
            slaves = strings.Split(readConsole(), ",")
            checkRes := util.CheckIfInLicencedIps(slaves)
            if checkRes {
                fmt.Println(slaves)
                break
            } else {
                fmt.Println(aurora.Red("[Warning] not supported IP(s)! Please retry "))
            }
        }

        module.InstallHadoop()
    case "Zookeeper":
    case "Hbase":
    case "Kafka":
    case "Storm":

    case "DPMonitor":
    case "DPMetrics":
    }
}

func addModule(moduleName string) {
    // todo add modules
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
|____/|_|  |___| v0.1`))

    fmt.Println()


    //初始安装licence，增加节点licence
    fmt.Println(aurora.BgGray(aurora.Black("@@ Please input install licence :")))
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
        // todo wget parcel, then unzip

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

        extraModule := "DPMonitor"
        installPhase(6, extraModule)

        extraModule = "DPMetrics"
        installPhase(7, extraModule)

    } else if 1 == typ {
        fmt.Println(aurora.Blue("Add Node Process Start"))

        if !util.IfFirstInstallDirsExists() {
            fmt.Println(aurora.Red("[ERROR] Can not find the first install parcel ! "))
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

    util.SaveConfigToLocal()
    fmt.Println(aurora.Magenta("All done, goodbye !"))
}
