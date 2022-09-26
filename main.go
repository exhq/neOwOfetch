package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/exhq/neowofetch/data"
	"github.com/exhq/neowofetch/utils"
)

var isuwuified bool = true

func handleConfig() {
	_, folder := os.Stat(filepath.Dir(data.GetConfigFile()))
	_, file := os.Stat(data.GetConfigFile())
	if os.IsNotExist(folder) {
		os.Mkdir(filepath.Dir(data.GetConfigFile()), os.ModePerm)
	}

	if os.IsNotExist(file) {
		println("config was not found. a default config file has been generated in '~/.config/neowofetch/conf'")
		f, _ := os.Create(data.GetConfigFile())
		_, _ = f.WriteString("println green neOwOfetch 🔥\ninfo magenta username\nprint blue @\ninfoln blue hostname\nprint white uptime:     \ninfoln red uptime\nprint white shell:      \ninfoln blue shell\nprint white distro:     \ninfoln blue distro\nprint white terminal:   \ninfoln blue terminal\nprint white memory:     \ninfo blue memoryUsed\nprint white /\ninfoln blue memoryAll")
	} else {
		body, _ := ioutil.ReadFile(data.GetConfigFile())
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
			utils.CutePrint(data.GetUsername(), format)
		case "uptime":
			no, _ := strconv.Atoi(data.GetUptime())
			utils.CutePrint(data.FormatTime(no), format)
		case "hostname":
			utils.CutePrint(data.GetHostname(), format)
		case "GPU":
			utils.CutePrint(data.GetGPU(), format)
		case "shell":
			utils.CutePrint(data.GetShell(), format)
		case "terminal":
			utils.CutePrint(data.GetTerminal(), format)
		case "memoryAll":
			utils.CutePrint(data.GetMemory(false), format)
		case "memoryUsed":
			utils.CutePrint(data.GetMemory(true), format)
		default:
			print("{UNKNOWN KEYWORD}")
		}
		if action == "infoln" {
			utils.CuteNewLine()
		}
	}
}

func main() {
	utils.Initargs()
	utils.Initcolor()
	utils.CutePrintInit()
	handleConfig()
	utils.CutePrintEnd()
}
