package util

import (
    "errors"
    "strings"
)

var Nodes []string

func CheckLicence(key string) (int, error, int) {
    // todo check install licence
    return 0, nil , 3
}

func CheckLicencedIPs(num int, ips string) error {
    ipArray := strings.Split(ips, ",")
    if num == len(ipArray) {
        Nodes = ipArray
        return nil
    } else {
        return errors.New("ip numbers not match")
    }
}