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

func checkforconfigfolder() {
	_, folder := os.Stat(data.GetHome() + "/.config")
	if os.IsNotExist(folder) {
		os.Mkdir(data.GetHome()+"/.config", os.ModePerm)
	}
}

func handleConfig() {
	_, folder := os.Stat(filepath.Dir(data.GetConfigFile()))
	_, file := os.Stat(data.GetConfigFile())
	if os.IsNotExist(folder) {
		os.Mkdir(filepath.Dir(data.GetConfigFile()), os.ModePerm)
	}
	defaultconfig := "println green neOwOfetchh 🔥\ninfo white username\nprint blue @\ninfoln blue hostname\nprint white uptime:   \ninfoln red uptime\nprint white shell:      \ninfoln blue shell\nprint white distro:   \ninfoln blue distro\nprint white terminal:   \ninfoln blue terminal\nprint white memory:   \ninfo blue memoryUsed\nprint white /\ninfoln blue memoryAll"
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
		switch rest {
		case "distro":
			utils.CutePrint(data.GetDistro(), format)
		case "username":
			utils.CutePrint(data.GetUsername(), format)
		case "uptime":
			no, _ := strconv.Atoi(data.GetUptime())
			utils.CutePrint(data.FormatTime(no), format)
		case "hostname":
			utils.CutePrint(data.Unamebs("-n"), format)
		case "kernelname":
			utils.CutePrint(data.Unamebs("-s"), format)
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
	checkforconfigfolder()
	utils.Initargs()
	utils.Initcolor()
	utils.CutePrintInit()
	handleConfig()
	utils.CutePrintEnd()
}
