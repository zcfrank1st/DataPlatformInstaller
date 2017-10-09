package main

import (
    "bufio"
    "os"
    "fmt"
)

var scanner *bufio.Scanner

func init () {
    scanner = bufio.NewScanner(os.Stdin)
}

// tips: all machines need to be ssh
func readConsole() string{
    scanner.Scan()
    return scanner.Text()
}


func main() {
    // TODO  install modules
    fmt.Println(`
 ____  ____ ___
|  _ \|  _ \_ _|
| | | | |_) | |
| |_| |  __/| |
|____/|_|  |___| v0.1`)

    fmt.Println()
    fmt.Println("1) Install Hadoop? (Y/N)")

    for {
        hadoopInput := readConsole()
        if "Y" == hadoopInput {
            fmt.Println(hadoopInput)
            fmt.Println("install")
            break
        } else if "N" == hadoopInput {
            fmt.Println("not install")
            break
        } else {
            fmt.Println("not support input, please retry")
        }
    }


    fmt.Println("2) Install Spark? (Y/N)")
    fmt.Println("3) Install Sqoop? (Y/N)")
    fmt.Println("4) Install Flume? (Y/N)")
}
