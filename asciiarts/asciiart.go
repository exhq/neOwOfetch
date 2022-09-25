package asciiarts

import (
	_ "embed"
	"strings"
)

//go:embed alpine.txt
var Alpine string

//go:embed unknown.txt
var unknown string

func GetAsciiInternal(distroID string) string {
	switch distroID {
	case "Alpine":
		return Alpine
	default:
		return unknown
	}
}
func GetAscii(distroID string) string {
	ascii := GetAsciiInternal(distroID)
	ascii = strings.ReplaceAll(ascii, "{WHITE}", "")
	ascii = strings.ReplaceAll(ascii, "{YELLOW}", "")
	ascii = strings.ReplaceAll(ascii, "{BLUE}", "")
	return ascii
}
