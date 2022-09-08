package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getHostname() string {
	cmd := exec.Command("uname", "-n")
	shell, _ := cmd.Output()
	return strings.Replace(string(shell), "\n", "", -1)
}
func getUsername() string {
	cmd := exec.Command("whoami")
	shell, _ := cmd.Output()
	return strings.Replace(string(shell), "\n", "", -1)
}
func getLogo() {
}
func getDistro() string {
	distro, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
	distro_info := make([]byte, 1024)
	distro.Read(distro_info)
	distro.Close()
	distro_list := strings.Split(string(distro_info), "\n")
	distro_tuples := make(map[string]string)
	for _, v := range distro_list {
		if strings.Contains(v, "=") {
			kv := strings.Split(v, "=")
			kv[0] = strings.TrimSpace(kv[0])
			kv[1] = strings.TrimSpace(kv[1])
			distro_tuples[kv[0]] = kv[1]
		}
	}
	return strings.Trim(distro_tuples["PRETTY_NAME"], "\"")
}
func getKernel() string {
	cmd := exec.Command("uname", "-r")
	kernel, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return "fuck you"
	}
	return string(kernel)
}
func getUptime() string {
	cmd := exec.Command("cat", "/proc/uptime")
	uptime, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return "fuck you"
	}
	intup, _ := strconv.Atoi(string(uptime[:strings.Count(string(uptime), ".")+3]))
	return literallyBasicMath(intup)
}
func getPackages() {
}
func getShell() string {
	return os.Getenv("SHELL")
}
func getResolution() {

}
func getWM() {
}
func getTheme() {
}
func getIcons() {
}
func getTerminal() string {
	_, exists := os.LookupEnv("TERM")
	_, existprgm := os.LookupEnv("TERM_PROGRAM")
	return ("exists=" + strconv.FormatBool(exists) + "existprgm=" + strconv.FormatBool(existprgm))
}
func getCPU() {
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
	_, _ = mem_list, mem_map
}
func getGPU() string {
	cmd := exec.Command("lspci", "-v")
	shell, err := cmd.Output()
	_ = err
	var bruh string
	//return strings.Replace(string(shell), "\n", "", -1)
	//return string(shell)
	for _, line := range strings.Split(strings.TrimSuffix(string(shell), "\n"), "\n") {
		if strings.Contains(line, "VGA") {
			bruh += line[strings.Index(line, ": ")+2 : strings.Index(line, " (")]
		}
	}
	return bruh
}
func getMemory(used bool) string {

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
	//memory := fmt.Sprintf("%d/%d", mem_used/1024, mem_total/1024)
	if used {
		return strconv.Itoa(mem_used / 1024)
	} else {
		return strconv.Itoa(mem_total / 1024)
	}
}

func literallyBasicMath(seconds int) string {
	minutes := seconds / 60
	secondsre := strconv.Itoa(seconds % 60)
	hour := strconv.Itoa(minutes / 60)
	minutesre := strconv.Itoa(minutes % 60)
	return (hour + "h " + minutesre + "m " + secondsre + "s")
}

func getColorPalette() {
}

func main() {
	fmt.Print(getUptime(), "\n")
	//fmt.Print(getTerminal(), "\n")
	fmt.Print(getShell(), "\n")
	fmt.Print(getHostname(), "@", getUsername(), "\n")
	fmt.Print(getDistro(), "\n")
	fmt.Print(getGPU(), "\n")
	fmt.Print(getMemory(true), "\n")
	fmt.Print(getKernel(), "\n")
	fmt.Print(13/2, 13%2, "\n")
}
