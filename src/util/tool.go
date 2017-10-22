package util

import "os"

func IfFirstInstallDirsExists() bool{
    _, err := os.Stat(DPI_DIR)

    if err != nil {
        return false
    } else {
        return true
    }
}