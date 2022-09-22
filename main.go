package main

import (
	"fmt"
	"io/ioutil"
	"neowofetch/asciiart"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var isuwuified bool = true
var linearch []string
var aa int

func getHome() string {
	return os.Getenv("HOME")

}
func incrementaa() {
	aa += 1
}

func getConfigFile() string {
	return getHome() + "/.config/neowofetch/conf"
}

func initascii() {
	linearch = asciiart.Getascii(getDistro())
	aa = 0
	print(linearch[aa])
	aa = aa + 1
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

func handlePrint(action, colour string, rest string) {
	if action == "print" {
		Cprint(colour, rest, true)
	} else if action == "println" {
		Cprint(colour, rest, true)
		print("\n")
		if aa < len(linearch) {
			print(linearch[aa])
		}
		if aa == len(linearch) || aa == len(linearch)-1 {
			print(strings.Repeat(" ", len(linearch[1])))
		}

		if aa > len(linearch) {
			print(strings.Repeat(" ", len(linearch[1])))
		}
		incrementaa()
	} else if action == "info" || action == "infoln" {
		switch rest {
		case "distro":
			Cprint(colour, getDistro(), true)
		case "username":
			Cprint(colour, getUsername(), true)
		case "uptime":
			no, _ := strconv.Atoi(getUptime())
			Cprint(colour, formatTime(no), false)
		case "hostname":
			Cprint(colour, getHostname(), true)
		case "GPU":
			Cprint(colour, getGPU(), true)
		case "shell":
			Cprint(colour, getShell(), true)
		case "terminal":
			Cprint(colour, getTerminal(), true)
		case "memoryAll":
			Cprint(colour, getMemory(false), true)
		case "memoryUsed":
			Cprint(colour, getMemory(true), true)
		default:
			print("{UNKNOWN KEYWORD}")
		}
	}
	if action == "infoln" {

		print("\n")
		if aa < len(linearch) {
			print(linearch[aa])
		} else {
			print(strings.Repeat(" ", len(linearch[1])))
		}
		incrementaa()
	}
}

func Cprint(colour string, message string, uwu bool) {
	nouwu := len(os.Args) == 2 && os.Args[1] == "-nouwu"

	if uwu && !nouwu {
		message = uwuify(message)
	}
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	black := color.New(color.FgBlack).SprintFunc()
	switch colour {
	case "yellow":
		print(yellow(message))
	case "white":
		print(message)
	case "magenta":
		print(magenta(message))
	case "red":
		print(red(message))
	case "blue":
		print(blue(message))
	case "black":
		print(black(message))
	case "green":
		print(green(message))
	}
}

func uwuify(message string) string {
	var answer string
	var sentence []string
	var hasspace bool
	if strings.Contains(message, " ") {
		sentence = strings.Split(message, " ")
		hasspace = true
	} else {
		sentence = strings.Split(message, " ")
		hasspace = false
	}

	for _, word := range sentence {
		if !strings.Contains(strings.ToLower(word), "uwu") {
			word = strings.Replace(word, "u", "UwU", -1)

			if strings.Contains(strings.ToLower(word), "owo") {
				word = strings.Replace(word, "o", "OwO", -1)
			}
			word = strings.Replace(word, "r", "w", -1)

		}
		if hasspace {
			answer += word + " "
		} else {
			answer += word
		}
	}
	return answer
}

func handleArgs() {
	if len(os.Args) == 1 {
		return
	} else if len(os.Args) > 1 {
		args := os.Args
		for _, arg := range args {
			if arg == "nouwu" {
				isuwuified = false
			}
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
	_, existprgm := os.LookupEnv("TERM_PROGRAM")
	if !existprgm {
		return os.Getenv("TERM")
	} else {
		return os.Getenv("TERM_PROGRAM")
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

func handleremainingascii() {
	if aa < len(linearch) {
		for i := 0; i < len(linearch)-aa; i++ {
			print("\n", linearch[aa])
			incrementaa()

		}
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
	initascii()
	handleArgs()
	handleConfig()
	handleremainingascii()
	print("\n")
}
