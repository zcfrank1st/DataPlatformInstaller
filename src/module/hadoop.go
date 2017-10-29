package module

import (
    "util"
    "io/ioutil"
    "strings"
)

var HADOOP_CONFIG_DIR = util.DPI_DIR + "/hadoop/etc/hadoop"

// ${DPI_DIR}/hadoop/etc/hadoop
func InstallHadoop(master string, slaves []string) error {
    configErr := generateHadoopConfig(master, slaves)
    if configErr != nil {
        return configErr
    }

    syncErr := syncHadoop(master, slaves)
    if syncErr != nil {
        return syncErr
    }

    return nil
}

// todo
func AddHadoopNode() {}

func generateHadoopConfig(master string, slaves []string) error {
    filenames := []string {"core-site.xml", "hdfs-site.xml", "mapred-site.xml", "yarn-site.xml"}
    for _, ele := range filenames {
        commonErr := generateCommonConfig(master, ele)
        if commonErr != nil {
            return commonErr
        }
    }

    slaveErr := generateSlavesConfig(slaves)
    if slaveErr != nil {
        return slaveErr
    }

    return nil
}

func generateCommonConfig(master string, filename string) error {
    commonFile := HADOOP_CONFIG_DIR + "/" + filename
    byteContent, readErr := ioutil.ReadFile(commonFile)
    if readErr != nil {
        return readErr
    }
    content := strings.Replace(strings.Replace(string(byteContent), "${M}", master, -1), "${V}", util.VERSION, -1)
    writeErr := ioutil.WriteFile(commonFile, []byte(content), 0777)
    if writeErr != nil {
        return writeErr
    }
    return nil
}

func generateSlavesConfig(slaves []string) error {
    content := ""
    for _, slave := range slaves {
        content = content + slave + "\n"
    }
    return ioutil.WriteFile(HADOOP_CONFIG_DIR + "/slaves", []byte(content), 0777)
}

func syncHadoop(master string, slaves []string) error {
    return util.Scp(append(slaves, master), "hadoop")
}