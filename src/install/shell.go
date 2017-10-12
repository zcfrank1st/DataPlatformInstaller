package main

import (
    "bufio"
    "os"
    "fmt"
    "github.com/logrusorgru/aurora"
    "module"
    "strings"
)

var scanner *bufio.Scanner

func init () {
    scanner = bufio.NewScanner(os.Stdin)
}

func readConsole() string{
    scanner.Scan()
    return scanner.Text()
}

func printStep(step int, moduleName string) {
    fmt.Println(aurora.Green(fmt.Sprintf("%d) Install %s? (Yy/Nn)", step, moduleName)))
}

func printAlert() {
    fmt.Println(aurora.Red("not support input, please retry"))
}

func installModule(moduleName string) {
    // todo
    switch moduleName {
    case "Hadoop":
        fmt.Println("ips: ")
        ips := readConsole()
        fmt.Println(ips)
        module.InstallHadoop()
    case "Zookeeper":
    case "Hbase":
    case "Kafka":
    case "Storm":

    case "DPMonitor":
    case "DPMetrics":
    }
}

func installLoop(moduleName string) {
    for {
        inputIgnoreCase := strings.ToLower(readConsole())
        if "y" == inputIgnoreCase {
            installModule(moduleName)
            break
        } else if "n" == inputIgnoreCase {
            break
        } else {
            printAlert()
        }
    }
}

func installPhase (step int, moduleName string) {
    printStep(step, moduleName)
    installLoop(moduleName)
}


func main() {
    fmt.Println(aurora.Blue(`
 ____  ____ ___
|  _ \|  _ \_ _|
| | | | |_) | |
| |_| |  __/| |
|____/|_|  |___| v0.1`))

    fmt.Println()

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

    // todo: write config to let ui and monitor use
}
