package util

import (
    "os/exec"
    "os"
)

const WGET_DIR = "/Users/zcfrank1st/Desktop/"

func commandRun (name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func Wget(url string) {
    commandRun("wget", "-P", WGET_DIR, url)
}

func Scp(hosts string) {
    // todo scp
    commandRun("scp", "")
}

func Cat() {}

func Grep() {}