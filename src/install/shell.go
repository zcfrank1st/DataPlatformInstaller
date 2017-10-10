package main

import (
    "bufio"
    "os"
    "fmt"
    "github.com/logrusorgru/aurora"
    "module"
    "strings"
)

// tips:
// all machines need to be ssh
// java install
// scala install

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
    // todo 传参问题，角色问题
    switch moduleName {
    case "Hadoop":
        fmt.Println("ips: ")
        ips := readConsole()
        fmt.Println(ips)
        module.InstallHadoop()
    case "Spark":
        module.InstallSpark()
    case "Flume":
        module.InstallFlume()
    // todo install hive, then use spark sql，it depends
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

    moduleName = "Spark"
    installPhase(2, moduleName)

    moduleName = "Flume"
    installPhase(3, moduleName)
}
