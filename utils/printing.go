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
var color_map = map[string]string{
	"black":   "30",
	"red":     "31",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"magenta": "35",
	"cyan":    "36",
	"white":   "37",
	"*":       "37",
}

var namechanges = map[string]string{
	"linux":  "linuwu",
	"arch":   "nyarch",
	"ubuntu": "uwuntu",
}

var logoIndex = 0
var isInProgressLine = false
var logoLines []string
var logoWidth int

var defaultcolor = "red 255 38 116 \ngreen 16 210 117\nblue 104 174 212\nwhite 250 253 255"
var pngWidth = 12
var pngHeight = 12
var pngData []byte

var colorconf = os.Getenv("HOME") + "/.config/neowofetch/colors"
var folderconf = filepath.Dir(colorconf)
var _, existcolorconf = os.Stat(colorconf)
var _, existfolderconf = os.Stat(folderconf)

func Initcolor() {

	if os.IsNotExist(existfolderconf) {
		os.Mkdir(folderconf, os.ModePerm)
	}
	if os.IsNotExist(existcolorconf) {
		f, _ := os.Create(colorconf)
		_, _ = f.WriteString(defaultcolor)
	}

	c, _ := os.ReadFile(colorconf)
	content := string(c)
	if Defaultcolor {
		content = defaultcolor
	}
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
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
	if Asciiforced {
		logo = asciiarts.GetAscii(Forceddistro)
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
		word = strings.ToLower(word)

		if namechanges[word] != "" {
			word = namechanges[word]
			if i != 0 {
				ret += " "
			}
			ret += word
			continue
		}

		if len(word) > 5 {
			if !strings.Contains(word, "owo") {
				word = strings.ReplaceAll(word, "o", "OwO")
			} else if !strings.Contains(word, "uwu") {
				word = strings.ReplaceAll(word, "u", "UwU")
			}
			word = strings.ReplaceAll(strings.ReplaceAll(word, "r", "w"), "l", "w")
		}

		if i != 0 {
			ret += " "
		}
		ret += word
	}
	return ret
}

type Format struct {
	spaces         int
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
			if strings.HasPrefix(v, "space") {
				parsedFormat.spaces, _ = strconv.Atoi(v[6:])
				return
			}
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
	if !Asciiforced {
		return data.GetDistroVariable("ID")
	} else {
		return Forceddistro
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
	neededspaces := parsedFormat.spaces - len(message)
	if neededspaces < 0 {
		neededspaces = 0
	}
	if !colorold {
		fmt.Printf("%s%s\x1b[0m%s", parsedFormat.colorFormat, message, strings.Repeat(" ", neededspaces))
	} else {
		if os.IsNotExist(existcolorconf) {
			f, _ := os.Create(colorconf)
			_, _ = f.WriteString("red 255 0 0 \ngreen 0 255 0\nblue 0 0 255\nwhite 255 255 255")
		}
		if colorold && hascolor {
			for k, v := range color_map {
				if strings.Contains(format, k) {
					fmt.Printf("\033[1;%sm%s\033[m", v, message)
					break
				}
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
