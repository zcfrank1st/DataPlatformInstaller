package util

import (
    "os/exec"
    "os"
    "errors"
)

func commandRun (name string, arg ...string) error {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

func Wget(url string, targetDir string) error {
    err := commandRun("wget", "-P", targetDir, url)
    if err != nil {
        return errors.New("wget install box error")
    }

    return nil
}

func Scp(hosts []string, moduleName string) error {
    // scp -r /opt/soft/test root@10.6.159.147:/opt/soft/scptest
    for _, host := range hosts {
        err := commandRun("scp", "-r", DPI_DIR + "/" + moduleName, "root@" + host + ":/opt")
        if err != nil {
            return errors.New(moduleName + " scp [" + host + "] error")
        }
    }

    return nil
}

func UnTar(packageName string) error {
    err := commandRun("tar", "-xvf", packageName) // DPI-0.1.0
    if err != nil {
        return errors.New("unzip install box error")
    }
    return nil
}