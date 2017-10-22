package util

import (
    "os/exec"
    "os"
)

func commandRun (name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}

func Wget(url string, targetDir string) {
    commandRun("wget", "-P", targetDir, url)
}

func Scp(hosts []string) {
    // scp -r /opt/soft/test root@10.6.159.147:/opt/soft/scptest
    for idx := range hosts {
        commandRun("scp", "-r", DPI_DIR, "root@" + hosts[idx] + ":/opt")
    }
}

func UnTar(packageName string) {
    commandRun("tar", "-xvf", packageName) // DPI-0.1.0
}