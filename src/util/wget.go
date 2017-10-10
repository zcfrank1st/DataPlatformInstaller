package util

import (
    "os/exec"
    "os"
)

const DIR = "/Users/zcfrank1st/Desktop/"

func Wget(url string) {
    cmd := exec.Command("wget", "-P", DIR, url)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}
