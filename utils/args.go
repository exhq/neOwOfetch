package utils

import (
	"os"
	"os/exec"
	"strings"
)

var args []string
var shoulduwuify bool = true
var noascii bool = false
var usepng bool = false
var hascolor bool = true
var Customascii = false
var asciidir string

func Initargs() {
	args = os.Args[1:]
	for _, argument := range args {
		if strings.HasPrefix(argument, "--ascii=") {
			cmd := exec.Command("whoami")
			shell, _ := cmd.Output()
			funny := strings.Replace(string(shell), "\n", "", -1)
			Customascii = true
			asciidir = strings.ReplaceAll(argument[8:], "~", "/home/"+funny)
		} else if strings.HasPrefix(argument, "--") {
			switch argument {
			case "--nouwu":
				shoulduwuify = false
			case "--noascii":
				noascii = true
			case "--usepng":
				usepng = true
			case "--nocolor":
				hascolor = false
			}
		}
	}
}

func Woulduwuify() bool {
	return shoulduwuify
}
