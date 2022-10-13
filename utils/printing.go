package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/exhq/neowofetch/asciiarts"
	"github.com/exhq/neowofetch/data"
	"github.com/exhq/neowofetch/images"
)

func rgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[38:2::%d:%d:%dm", r, g, b)
}

var colors = make(map[string]string)
var oldcolors = make(map[string]int)

var logoIndex = 0
var isInProgressLine = false
var logoLines []string
var logoWidth int

var pngWidth = 12
var pngHeight = 12
var pngData []byte

func Initcolor() {
	colorconf := os.Getenv("HOME") + "/.config/neowofetch/colors"
	folderconf := filepath.Dir(colorconf)

	_, existcolorconf := os.Stat(colorconf)
	_, existfolderconf := os.Stat(folderconf)

	if os.IsNotExist(existfolderconf) {
		os.Mkdir(folderconf, os.ModePerm)
	}
	if os.IsNotExist(existcolorconf) {
		println("color was not found. a default config file has been generated in '~/.config/neowofetch/colors'. rerun the program")
		f, _ := os.Create(colorconf)
		_, _ = f.WriteString("red 255 0 0 \ngreen 0 255 0\nblue 0 0 255\nwhite 255 255 255")
		os.Exit(0)
	}

	c, _ := os.ReadFile(colorconf)
	content := string(c)
	if Defaultcolor {
		content = "red 255 0 0 \ngreen 0 255 0\nblue 0 0 255\nwhite 255 255 255"
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		word := strings.Split(line, " ")

		R, _ := strconv.Atoi(word[1])
		G, _ := strconv.Atoi(word[2])
		B, _ := strconv.Atoi(word[3])
		colors[word[0]] = rgb(R, G, B)

	}
}

func CutePrintInit() {
	dist := data.GetDistroVariable("ID")
	logo := asciiarts.GetAscii(dist)
	if asciiforced {
		logo = asciiarts.GetAscii(forceddistro)
	}
	if noascii {
		logo = ""
	}
	if usepng {
		pngData = images.DistroImages[dist]
		logo = strings.Repeat(" ", pngWidth) + " " + strings.Repeat("\n", pngHeight)
	}
	if Customascii {
		body, _ := ioutil.ReadFile(asciidir)
		logo = (string(body))
	}
	logoLines = strings.Split(logo, "\n")
	logoWidth = 0
	for _, v := range logoLines {
		lineLength := len([]rune(v)) + 2
		if lineLength > logoWidth {
			logoWidth = lineLength
		}
	}
}

func printLogoIfAtBeginningOfNewLine() {
	if !isInProgressLine {
		isInProgressLine = true
		if logoIndex < len(logoLines) {
			logoLine := logoLines[logoIndex]
			logoLineLength := len([]rune(logoLine))
			padding := strings.Repeat(" ", logoWidth-logoLineLength)
			fmt.Printf("%s%s", logoLine, padding)
		} else {
			fmt.Printf("%s", strings.Repeat(" ", logoWidth))
		}
		logoIndex += 1
	}
}

func uwuify(message string) string {
	sentence := strings.Split(message, " ")
	ret := ""
	for i, word := range sentence {
		if !strings.Contains(strings.ToLower(word), "uwu") {
			word = strings.ReplaceAll(word, "u", "UwU")

			if !strings.Contains(strings.ToLower(word), "owo") {
				word = strings.ReplaceAll(word, "o", "OwO")
			}
			word = strings.ReplaceAll(word, "r", "w")

		}
		if i != 0 {
			ret += " "
		}
		ret += word
	}
	return ret
}

type Format struct {
	noUwuOverride  bool
	colorFormat    string
	oldcolorFormat int
}

func parseFormat(format string) (parsedFormat Format) {
	for _, v := range strings.Split(format, "|") {
		colorFormat, isColor := colors[v]
		parsedFormat.colorFormat += colorFormat
		if isColor && hascolor {
			parsedFormat.colorFormat += colorFormat
		} else {
			switch v {
			case "italic":
				parsedFormat.colorFormat += "\x1b[3m"
			case "bold":
				parsedFormat.colorFormat += "\x1b1"
			case "nouwu":
				parsedFormat.noUwuOverride = true
			case "*":
			default:
				if hascolor && !colorold {
					println("Unknown format code: ", v)
				}
			}
		}
	}
	return parsedFormat
}

func getcustomizeddistro() string {
	if !asciiforced {
		return data.GetDistroVariable("ID")
	} else {
		return forceddistro
	}
}

func CutePrint(
	message string,
	format string,
) {
	printLogoIfAtBeginningOfNewLine()
	parsedFormat := parseFormat(format)
	willUwuify := shoulduwuify && !parsedFormat.noUwuOverride
	if willUwuify {
		message = uwuify(message)
	}
	if !colorold {
		fmt.Printf("%s%s\x1b[0m", parsedFormat.colorFormat, message)
	} else {

		if colorold && hascolor {
			switch {
			case strings.Contains(format, "black"):
				fmt.Printf("\033[1;30m%s\033[0m", message)
			case strings.Contains(format, "red"):
				fmt.Printf("\033[1;31m%s\033[0m", message)
			case strings.Contains(format, "green"):
				fmt.Printf("\033[1;32m%s\033[0m", message)
			case strings.Contains(format, "yellow"):
				fmt.Printf("\033[1;33m%s\033[0m", message)
			case strings.Contains(format, "blue"):
				fmt.Printf("\033[1;34m%s\033[0m", message)
			case strings.Contains(format, "magenta"):
				fmt.Printf("\033[1;35m%s\033[0m", message)
			case strings.Contains(format, "cyan"):
				fmt.Printf("\033[1;36m%s\033[0m", message)
			case strings.Contains(format, "white"):
				fmt.Printf("\033[1;37m%s\033[0m", message)
			case strings.Contains(format, "*"):
				fmt.Printf("\033[1;37m%s\033[0m", message)
			}
		}
	}

}

func CuteNewLine() {
	printLogoIfAtBeginningOfNewLine()
	isInProgressLine = false
	fmt.Println()
}

func CutePrintEnd() {
	for logoIndex < len(logoLines) {
		CuteNewLine()
	}
	if usepng {
		fmt.Printf("\x1b[%dA", logoIndex)
		fmt.Printf("\x1b]1337;File=inline=1;width=%d;height=%d:", pngWidth, pngHeight)
		enc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
		enc.Write(images.DistroImages[getcustomizeddistro()])
		enc.Close()
		fmt.Println("\a")
	}
}
