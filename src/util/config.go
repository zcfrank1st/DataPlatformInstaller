package util

import (
    "encoding/json"
    "io/ioutil"
)

var (
    RoleMapper map[string]string  // {"NameNode": "192.168.33.1", }
)


func SaveConfigToLocal() {
    //config, to let ui and monitor use
    jsonString, _ := json.Marshal(RoleMapper)
    ioutil.WriteFile("/etc/dpi.conf", jsonString, 0777)
}