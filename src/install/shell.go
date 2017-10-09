package main

import (
    "bufio"
    "os"
    "fmt"
    "github.com/logrusorgru/aurora"
)

// tips: all machines need to be ssh

var scanner *bufio.Scanner

func init () {
    scanner = bufio.NewScanner(os.Stdin)
}

func readConsole() string{
    scanner.Scan()
    return scanner.Text()
}

func printStep(step int, moduleName string) {
    fmt.Println(aurora.Green(fmt.Sprintf("%d) Install %s? (Y/N)", step, moduleName)))
}

func printAlert() {
    fmt.Println(aurora.Red("not support input, please retry"))
}

func installLoop(moduleName string) {
    for {
        input := readConsole()
        if "Y" == input {
            // todo input ip(s) to install modules
            break
        } else if "N" == input {
            break
        } else {
            printAlert()
        }
    }
}

func installHadoop() {}
func installSpark(){}
func installFlume(){}
func installSqoop(){}

func installModule (step int, moduleName string) {
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
    installModule(1, moduleName)

    moduleName = "Spark"
    installModule(2, moduleName)

    moduleName = "Sqoop"
    installModule(3, moduleName)

    moduleName = "Flume"
    installModule(4, moduleName)
}
