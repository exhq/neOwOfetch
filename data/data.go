package data

import (
	"fmt"
	"os"
	"strings"
)

func GetDistro() string {
	return GetDistroVariable("PRETTY_NAME")
}
func GetDistroVariable(varname string) string {
	distro, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	distro_list := strings.Split(string(distro), "\n")
	distro_tuples := make(map[string]string)
	for _, v := range distro_list {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			distro_tuples[kv[0]] = kv[1]
		}
	}
	return strings.Trim(distro_tuples[varname], "\"")
}
