package asciiarts

import (
	_ "embed"
	"strings"
)

//go:embed alpine.txt
var alpine string

//go:embed unknown.txt
var unknown string

//go:embed amogos.txt
var amogos string

//go:embed android.txt
var android string

//go:embed arch.txt
var arch string

//go:embed arcolinux.txt
var arcolinux string

//go:embed artix.txt
var artix string

//go:embed debian.txt
var debian string

//go:embed endeavouros.txt
var endeavouros string

//go:embed fedora.txt
var fedora string

//go:embed freebsd.txt
var freebsd string

//go:embed gentoo.txt
var gentoo string

//go:embed gnu.txt
var gnu string

//go:embed guix.txt
var guix string

//go:embed ios.txt
var ios string

//go:embed macos.txt
var macos string

//go:embed manjaro-arm.txt
var manjaroarm string

//go:embed manjaro.txt
var manjaro string

//go:embed linuxmint.txt
var linuxmint string

//go:embed openbsd.txt
var openbsd string

//go:embed opensuse-leap.txt
var opensuseleap string

//go:embed opensuse-tumbleweed.txt
var opensuseweed string

//go:embed pop.txt
var pop string

//go:embed raspbian.txt
var raspbian string

//go:embed slackware.txt
var slackware string

//go:embed solus.txt
var solus string

//go:embed ubuntu.txt
var ubuntu string

//go:embed void.txt
var void string

//go:embed xerolinux.txt
var xerolinux string

func GetAsciiInternal(distroID string) string {
	switch distroID {
	case "alpine":
		return alpine
	case "arcolinux":
		return arcolinux
	case "artix":
		return artix
	case "debian":
		return debian
	case "endeavouros":
		return endeavouros
	case "fedora":
		return fedora
	case "freebsd":
		return freebsd
	case "gentoo":
		return gentoo
	case "gnu":
		return gnu
	case "guix":
		return guix
	case "ios":
		return ios
	case "linuxmint":
		return linuxmint
	case "macos":
		return macos
	case "manjaro-arm":
		return manjaroarm
	case "manjaro":
		return manjaro
	case "openbsd":
		return openbsd
	case "opensuse-leap":
		return opensuseleap
	case "opensuse-tumbleweed":
		return opensuseweed
	case "pop":
		return pop
	case "raspbian":
		return raspbian
	case "slackware":
		return slackware
	case "solus":
		return solus
	case "ubuntu":
		return ubuntu
	case "void":
		return void
	case "xerolinux":
		return xerolinux

	case "arch":
		return arch

	default:
		return unknown
	}
}
func GetAscii(distroID string) string {
	ascii := GetAsciiInternal(distroID)
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
