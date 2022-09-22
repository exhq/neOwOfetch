package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/exhq/neowofetch/data"
	"github.com/exhq/neowofetch/images"
)

func rgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[38:2::%d:%d:%dm", r, g, b)
}

var colors = map[string]string{
	"green": rgb(0, 255, 0),
	"blue":  rgb(0, 0, 255),
	"red":   rgb(255, 0, 0),
}

var uwuEmotes = [15]string{
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
	"*twerks*",
	"*sweats*",
}
var logoIndex = 0
var isInProgressLine = false
var logoLines []string
var logoWidth int

var pngWidth = 12
var pngHeight = 12
var pngData []byte

func CutePrintInit() {
	dist := data.GetDistroVariable("ID")
	logo := Getascii(dist)
	if noascii {
		logo = ""
	}
	if usepng {
		pngData = images.DistroImages[dist]
		logo = strings.Repeat(" ", pngWidth) + strings.Repeat("\n", pngHeight)
	}
	logoLines = strings.Split(logo, "\n")
	logoWidth = 0
	for _, v := range logoLines {
		lineLength := len([]rune(v))
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

func uwuify(message string) (ret string) {
	sentence := strings.Split(message, " ")
	ret = ""
	for i, word := range sentence {
		if !strings.Contains(strings.ToLower(word), "uwu") {
			word = strings.ReplaceAll(word, "u", "UwU")

			if strings.Contains(strings.ToLower(word), "owo") {
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
	noUwuOverride bool
	colorFormat   string
}

func parseFormat(format string) (parsedFormat Format) {
	for _, v := range strings.Split(format, "|") {
		colorFormat, isColor := colors[v]
		if isColor {
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
				//println("Unknown format code: ", v)
			}
		}
	}
	return parsedFormat
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
	fmt.Printf("%s%s\x1b[0m", parsedFormat.colorFormat, message)
}

func CuteNewLine() {
	printLogoIfAtBeginningOfNewLine()
	if rand.Intn(5) == 0 && shoulduwuify {
		fmt.Printf(" %s", uwuEmotes[rand.Intn(len(uwuEmotes))])
	}
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
		enc.Write(images.DistroImages["arch"])
		enc.Close()
		fmt.Println("\a")
	}
}
