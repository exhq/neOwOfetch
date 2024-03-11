package main

import (
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

func checkforconfigfolder() {
	_, folder := os.Stat(data.GetHome() + "/.config")
	if os.IsNotExist(folder) {
		os.Mkdir(data.GetHome()+"/.config", os.ModePerm)
	}
}

func getShellVersion(shellCommand string) string {
	// https://www.youtube.com/watch?v=YPN0qhSyWy8
	cmd := shellCommand + " --version | grep -o -E \"([0-9]\\.?)*\" | head -n1"
	out, _ := exec.Command("bash", "-c", cmd).Output()
	return filepath.Base(shellCommand) + " " + strings.ReplaceAll(string(out), "\n", "")
}

func handleConfig() {
	_, folder := os.Stat(filepath.Dir(data.GetConfigFile()))
	_, file := os.Stat(data.GetConfigFile())
	if os.IsNotExist(folder) {
		os.Mkdir(filepath.Dir(data.GetConfigFile()), os.ModePerm)
	}
	defaultconfig := "println green neOwOfetchh 🔥\ninfo white username\nprint blue @\ninfoln blue hostname\nprint white|space=12 uptime:\ninfoln red uptime\nprint white|space=12 shell:\ninfoln blue shell\nprint white|space=12 distro:\ninfoln blue distro\nprint white|space=12 terminal:\ninfoln blue terminal\nprint white|space=12 WM:\ninfoln blue wm\nprint white|space=12 memory:\ninfo blue memoryUsed\nprint white /\ninfoln blue memoryAll"
	if os.IsNotExist(file) {
		f, _ := os.Create(data.GetConfigFile())
		_, _ = f.WriteString(defaultconfig)
	}
	body, _ := ioutil.ReadFile(data.GetConfigFile())
	sbody := (string(body))
	if utils.Defaultconf {
		sbody = defaultconfig
	}
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

func handlePrint(action, format string, rest string) {
	if action == "print" {
		utils.CutePrint(rest, format)
	} else if action == "println" {
		utils.CutePrint(rest, format)
		utils.CuteNewLine()
	} else if action == "info" || action == "infoln" {
		infoGetters := map[string]func() string{
			"distro": func() string {
				if utils.Asciiforced {
					return utils.Forceddistro
				}
				return data.GetDistro()
			},
			"username": data.GetUsername,
			"uptime": func() string {
				no, _ := strconv.Atoi(data.GetUptime())
				return data.FormatTime(no)
			},
			"hostname":   func() string { return data.Unamebs("-n") },
			"kernelname": func() string { return data.Unamebs("-srm") },
			//"CPU":        data.GetCPU,
			"GPU":        data.GetGPU,
			"shell":      func() string { return getShellVersion(data.GetShell()) },
			"terminal":   data.GetTerminal,
			"memoryAll":  func() string { return data.GetMemory(false) },
			"memoryUsed": func() string { return data.GetMemory(true) },
			"wm":         data.GetWM,
			"ip":         data.GetLocalIP,
		}

		getter, ok := infoGetters[rest]
		if ok {
			utils.CutePrint(getter(), format)
		} else {
			utils.CutePrint("{UNKNOWN KEYWORD: "+rest+"}", format)
		}

		if action == "infoln" {
			utils.CuteNewLine()
		}
	}
}

func runpage() {
	if utils.Ishelp {
		println(`neowofetch version x (idk how to implement this)
--nouwu           turns off uwuifying
--usepng          uses a png (only supports arch and ubuntu rn)
--noascii         turns off the ascii
--nocolor         no color.
--noconf          uses builtin config instead of looking for one in the filesystem
--nocolorconf     same thing but for color config :P
--16color         uses escape codes instead of RGB
--help            YOU LITERALLY FUCKING RAN IT RIGHT NOW`)
		os.Exit(0)
	}
}

func main() {
	utils.Initargs()
	runpage()
	checkforconfigfolder()
	utils.Initcolor()
	utils.CutePrintInit()
	handleConfig()
	utils.CutePrintEnd()
}
