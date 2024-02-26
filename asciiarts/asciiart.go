package asciiarts

import (
	"embed"
	"regexp"
)

//go:embed *.txt
var asciifiles embed.FS

func GetAscii(distroID string) string {
	regex, _ := regexp.Compile("{[[:upper:]]+}")
	asciiBuff, error := asciifiles.ReadFile(distroID + ".txt")
	if error != nil {
		asciiBuff, _ = asciifiles.ReadFile("unknown.txt")
	}
	ascii := string(asciiBuff)
	ascii = regex.ReplaceAllString(ascii, "")
	return ascii
}
