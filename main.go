package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var isuwuified bool = true
var tempbool bool

func cprint(input string, newline bool, uwuoverwrite bool) {
	endings := [15]string{
		"owo",
		"UwU",
		">w<",
		"^w^",
		"●w●",
		"☆w☆",
		"𝗨𝘄𝗨",
		"(´꒳`)",
		"♥(。U ω U。)",
		"(˘ε˘)",
		"( ˘ᴗ˘ )",
		"(*ฅ́˘ฅ̀*)",
		"*screams*",
		"*twearks*",
		"*sweats*",
	}
	ninput := ""
	if isuwuified && !uwuoverwrite {
		bruh := strings.Split(input, " ")
		for _, word := range bruh {
			if word == "" {
				continue
			}
			word = strings.ReplaceAll(word, "r", "w")
			word = strings.ReplaceAll(word, "i", "iy")
			word = strings.ReplaceAll(word, "l", "w")
			if strings.HasSuffix(word, "!") {
				word = word[0:len(word)-1] + "1!11!1"
			}
			if strings.Contains(word, "u") && !strings.Contains(word, "uwu") && !strings.Contains(word, "owo") {
				word = strings.ReplaceAll(word, "u", "uwu")
			}

			ninput += word + " "

		}
		if rand.Intn(5-1)+1 == 2 {
			ninput += endings[rand.Intn(len(endings))]

		}

	} else {
		ninput = input
	}

	if newline == false {
		fmt.Print(ninput)
	} else {
		fmt.Print(ninput + "\n")
	}
}

func handleConfig() {
	_, folder := os.Stat("/home/" + getUsername() + "/.config/neowofetch")
	_, file := os.Stat("/home/" + getUsername() + "/.config/neowofetch/conf")
	if os.IsNotExist(folder) {
		os.Mkdir("/home/"+getUsername()+"/.config/neowofetch", os.ModePerm)
	}

	if os.IsNotExist(file) {
		println("bruh you aint got tha file? bruh fr fr bruh wtf")
		f, _ := os.Create("/home/" + getUsername() + "/.config/neowofetch/conf")
		_, _ = f.WriteString(
			`test
among`)

	} else {
		body, _ := ioutil.ReadFile("/home/" + getUsername() + "/.config/neowofetch/conf")
		fbody := strings.Split(string(body), "\n")
		for _, s := range fbody {
			w := strings.Split(s, " ")
			if len(w) == 1 {
				continue
			}
			declr := w[0]
			inf := w[1]
			if declr == "nn-prin" {
				cprint(strings.Join(w[1:], " "), false, false)
			}
			if declr == "prin" {
				cprint(strings.Join(w[1:], " "), true, false)
			}
			if declr == "nn-info" || declr == "info" {
				if declr == "info" {
					tempbool = true
				} else {
					tempbool = false
				}
				if inf == "username" {
					cprint(getUsername(), tempbool, false)
				} else if inf == "hostname" {
					cprint(getHostname(), tempbool, false)
				} else if inf == "uptime" {
					among, _ := strconv.Atoi(getUptime())
					cprint(formatTime(among), tempbool, true)
				}
			}
		}
	}

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
	_, exists := os.LookupEnv("TERM")
	_, existprgm := os.LookupEnv("TERM_PROGRAM")
	return ("exists=" + strconv.FormatBool(exists) + "existprgm=" + strconv.FormatBool(existprgm))
}
func getCPU() {
	mem, err := os.Open("/proc/cpuinfo")
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
	handleArgs()
	handleConfig()

	//if isuwuified {
	//	fmt.Print("\n shit will be uwuified\n")
	//} else {
	//	fmt.Print("shit will be NOT uwuified\n")
	//}
}
