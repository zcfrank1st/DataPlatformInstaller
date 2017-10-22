package util

import "os"

func IfFirstInstallDirsExists() bool{
    _, err := os.Stat("/opt/DPI-" + VERSION)

    if err != nil {
        return false
    } else {
        return true
    }
}