package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/exhq/neowofetch/data"
	"github.com/exhq/neowofetch/utils"
)

var isuwuified bool = true

func getHome() string {
	return os.Getenv("HOME")

}

func getConfigFile() string {
	return getHome() + "/.config/neowofetch/conf"
}

func handleConfig() {
	_, folder := os.Stat(filepath.Dir(getConfigFile()))
	_, file := os.Stat(getConfigFile())
	if os.IsNotExist(folder) {
		os.Mkdir(filepath.Dir(getConfigFile()), os.ModePerm)
	}

	if os.IsNotExist(file) {
		println("config was not found. a default config file has been generated in '~/.config/neowofetch/conf'")
		f, _ := os.Create(getConfigFile())
		_, _ = f.WriteString("println green neOwOfetch 🔥\ninfo magenta username\nprint blue @\ninfoln blue hostname\nprint white uptime:     \ninfoln red uptime\nprint white shell:      \ninfoln blue shell\nprint white distro:     \ninfoln blue distro\nprint white terminal:   \ninfoln blue terminal\nprint white memory:     \ninfo blue memoryUsed\nprint white /\ninfoln blue memoryAll")
	} else {
		body, _ := ioutil.ReadFile(getConfigFile())
		sbody := (string(body))
		fbody := strings.Split(sbody, "\n")
		for _, line := range fbody {
			word := strings.Split(line, " ")
			if len(word) < 3 {
				continue
			}
			action := word[0]
			color := word[1]
			rest := strings.Join(word[2:], " ")
			handlePrint(action, color, rest)
		}

	}

}

func handlePrint(action, format string, rest string) {
	if action == "print" {
		utils.CutePrint(rest, format)
	} else if action == "println" {
		utils.CutePrint(rest, format)
		utils.CuteNewLine()
	} else if action == "info" || action == "infoln" {
		switch rest {
		case "distro":
			utils.CutePrint(data.GetDistro(), format)
		case "username":
			utils.CutePrint(getUsername(), format)
		case "uptime":
			no, _ := strconv.Atoi(getUptime())
			utils.CutePrint(formatTime(no), format)
		case "hostname":
			utils.CutePrint(getHostname(), format)
		case "GPU":
			utils.CutePrint(getGPU(), format)
		case "shell":
			utils.CutePrint(getShell(), format)
		case "terminal":
			utils.CutePrint(getTerminal(), format)
		case "memoryAll":
			utils.CutePrint(getMemory(false), format)
		case "memoryUsed":
			utils.CutePrint(getMemory(true), format)
		default:
			print("{UNKNOWN KEYWORD}")
		}
		if action == "infoln" {
			utils.CuteNewLine()
		}
	}
}

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
	content, _ := os.ReadFile("/proc/uptime")
	return (string(content[0:strings.Index(string(content), ".")]))

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
	a, existprgm := os.LookupEnv("TERM_PROGRAM")
	if !existprgm {
		return os.Getenv("TERM")
	} else {
		return a
	}

}
func getCPU() {
	mem, _ := os.Open("/proc/cpuinfo")
	mem_info := make([]byte, 1024)
	mem.Read(mem_info)
	mem.Close()
	// mem_list := strings.Split(string(mem_info), "\n")
	// mem_map := make(map[string]string)
	print(mem_info)
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

func formatTime(seconds int) string {
	minutes := seconds / 60
	secondsre := strconv.Itoa(seconds % 60)
	hour := strconv.Itoa(minutes / 60)
	minutesre := strconv.Itoa(minutes % 60)
	return (hour + "h " + minutesre + "m " + secondsre + "s")
}

func getColorPalette() {
}

func main() {
	utils.Initargs()
	utils.Initcolor()
	utils.CutePrintInit()
	handleConfig()
	utils.CutePrintEnd()
}
