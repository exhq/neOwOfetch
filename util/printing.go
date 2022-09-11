package util

import (
	"fmt"
	"math/rand"
	"strings"
)

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
	"*twearks*",
	"*sweats*",
}

var logoIndex = 0
var isInProgressLine = false
var logo = `        /\      
       /  \    
      /\   \    
     / > ω <\   
    /   __   \  
   / __|  |__-\ 
  /_-''    ''-_\  
`
var logoLines []string
var logoWidth int
var shouldUwuify = true

func InitUwuPrinter() {
	logoLines = strings.Split(logo, "\n")
	logoWidth = 0
	for _, v := range logoLines {
		lineLength := len([]rune(v))
		if lineLength > logoWidth {
			logoWidth = lineLength
		}
	}
}

func initLine() {
	if !isInProgressLine {
		isInProgressLine = true
		if logoIndex < len(logoLines) {
			logoLine := logoLines[logoIndex]
			logoLineLength := len([]rune(logoLine))
			padding := strings.Repeat(" ", logoWidth-logoLineLength)
			print(logoLine, padding)
			logoIndex += 1
		} else {
			print(strings.Repeat(" ", logoWidth))
		}
	}
}

func UwuPrint(message string, noUwuOverride bool, whole string) {
	//will add color eventually, my brain hurts
	var hadAnyContent bool
	var wholeword string
	var checkspaces int
	isuwu := true
	initLine()
	if noUwuOverride || !shouldUwuify {
		isuwu = false
		wholeword = message
	}
	notuwuified := ""
	if isuwu {
		words := strings.Split(message, " ")
		hadAnyContent = false
		for _, word := range words {
			notuwuified += word
			if word == "" {
				checkspaces += 1
				continue
			}
			word = strings.ReplaceAll(word, "r", "w")
			word = strings.ReplaceAll(word, "i", "iy")
			word = strings.ReplaceAll(word, "iyy", "iy")
			word = strings.ReplaceAll(word, "l", "w")

			if strings.HasSuffix(word, "!") {
				word = word[:len(word)-1] + "1!11!1"
			}

			if strings.Contains(word, "u") &&
				!strings.Contains(word, "uwu") &&
				!strings.Contains(word, "owo") {
				word = strings.ReplaceAll(word, "u", "uwu")
			}
			hadAnyContent = true
			wholeword += word
		}
	}

	if hadAnyContent && rand.Intn(5) == 0 {
		print(uwuEmotes[rand.Intn(len(uwuEmotes))])
	}
	wholeword = wholeword + strings.Repeat(" ", checkspaces)

	if strings.Contains(notuwuified, "blue") {
		handlecolor(wholeword)
	} else {
		print(wholeword)
	}
}

func handlecolor(wholeword string) {
	print(wholeword + "bruh")
}

func UwuNewline() {
	initLine()
	isInProgressLine = false
	fmt.Println()
}

func UwuPrintRest() {
	for logoIndex < len(logoLines) {
		UwuNewline()
	}
}
