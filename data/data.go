package data

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetGPU() string {
	cmd := exec.Command("lspci", "-v")
	shell, err := cmd.Output()
	_ = err
	var bruh string
	for _, line := range strings.Split(strings.TrimSuffix(string(shell), "\n"), "\n") {
		if strings.Contains(line, "VGA") {
			bruh += line[strings.Index(line, ": ")+2 : strings.Index(line, " (")]
		}
	}
	return bruh
}

func GetUsername() string {
	cmd := exec.Command("whoami")
	shell, _ := cmd.Output()
	return strings.Replace(string(shell), "\n", "", -1)
}

func GetUptime() string {
	content, _ := os.ReadFile("/proc/uptime")
	return (string(content[0:strings.Index(string(content), ".")]))

}

func GetKernel() string {
	cmd := exec.Command("uname", "-r")
	kernel, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return "fuck you"
	}
	return string(kernel)
}

func Unamebs(gamering string) string {
	cmd := exec.Command("uname", gamering)
	shell, _ := cmd.Output()
	return strings.Replace(string(shell), "\n", "", -1)
}

func GetDistro() string {
	return GetDistroVariable("PRETTY_NAME")
}
func GetMemory(used bool) string {

	//
	// The coolest part about this function unlike neofetch is that it also takes account of the basic os operations
	// into account thus not giving an illusion that your os is completely having fun
	// honestly idk lmao
	//

	mem, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	mem_info := make([]byte, 1024)
	mem.Read(mem_info)
	mem.Close()
	mem_list := strings.Split(string(mem_info), "\n")
	mem_map := make(map[string]string)
	for _, v := range mem_list {
		if strings.Contains(v, ":") {
			kv := strings.Split(v, ":")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			kv[1] = strings.Replace(kv[1], "kB", "", 3)
			kv[1] = strings.TrimSpace(kv[1])
			mem_map[kv[0]] = kv[1]
		}
	}
	mem_free, _ := strconv.Atoi(mem_map["MemFree"])
	mem_total, _ := strconv.Atoi(mem_map["MemTotal"])
	mem_used := mem_total - mem_free
	if used {
		return formatmem(mem_used / 1024)
	} else {
		return formatmem(mem_total / 1024)
	}
}
func GetHome() string {
	return os.Getenv("HOME")

}
func GetConfigFile() string {
	return GetHome() + "/.config/neowofetch/conf"
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
func FormatTime(seconds int) string {
	minutes := seconds / 60
	secondsre := strconv.Itoa(seconds % 60)
	hour := strconv.Itoa(minutes / 60)
	minutesre := strconv.Itoa(minutes % 60)
	return (hour + "h " + minutesre + "m " + secondsre + "s")
}
func GetCPU() {
	mem, _ := os.Open("/proc/cpuinfo")
	mem_info := make([]byte, 1024)
	mem.Read(mem_info)
	mem.Close()
	// mem_list := strings.Split(string(mem_info), "\n")
	// mem_map := make(map[string]string)
	print(mem_info)
}
func GetTerminal() string {
	a, existprgm := os.LookupEnv("TERM_PROGRAM")
	if !existprgm {
		return os.Getenv("TERM")
	} else {
		return a
	}

}
func formatmem(input int) string {
	return strconv.Itoa(input) + "MiB"
}
func GetShell() string {
	return os.Getenv("SHELL")
}
func GetWM() string {
	return os.Getenv("XDG_CURRENT_DESKTOP")
}
func getPackages() {
}
func getResolution() {
}

func getTheme() {
}
func getIcons() {
}
func getColorPalette() {
}
