package utils

import (
	"os"
	"strings"
)

var args []string
var shoulduwuify bool = true
var noascii bool = false
var usepng bool = false
var hascolor bool = true

func Initargs() {
	args = os.Args[1:]
	for _, argument := range args {
		if strings.HasPrefix(argument, "--") {
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
func Asciioverwrite(ascii []string) []string {
	if noascii {
		literallynothing := []string{"", ""}
		return (literallynothing)
	} else {
		return ascii
	}
}
