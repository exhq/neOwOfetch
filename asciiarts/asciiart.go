package asciiarts

import (
	"embed"
	"strings"
)

//go:embed *.txt
var asciifiles embed.FS

func GetAscii(distroID string) string {
	asciiBuff, error := asciifiles.ReadFile(distroID + ".txt")
	if error != nil {
		asciiBuff, _ = asciifiles.ReadFile("unknown.txt")
	}
	ascii := string(asciiBuff)
	ascii = strings.ReplaceAll(ascii, "{WHITE}", "")
	ascii = strings.ReplaceAll(ascii, "{YELLOW}", "")
	ascii = strings.ReplaceAll(ascii, "{BLUE}", "")
	ascii = strings.ReplaceAll(ascii, "{LPINK}", "")
	ascii = strings.ReplaceAll(ascii, "{BACKGROUND_GREEN}", "")
	ascii = strings.ReplaceAll(ascii, "{NORMAL}", "")
	ascii = strings.ReplaceAll(ascii, "{BLACK}", "")
	ascii = strings.ReplaceAll(ascii, "{GREEN}", "")
	ascii = strings.ReplaceAll(ascii, "{RED}", "")
	ascii = strings.ReplaceAll(ascii, "{PINK}", "")
	ascii = strings.ReplaceAll(ascii, "{MAGENTA}", "")
	ascii = strings.ReplaceAll(ascii, "{CYAN}", "")
	return ascii
}
