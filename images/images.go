package images

import _ "embed"

//go:embed arch.png
var ArchArt []byte

//go:embed ubuntu.jpg
var UbuntuArt []byte

//go:embed saul.png
var saulArt []byte

var DistroImages = map[string][]byte{
	"arch":   ArchArt,
	"saul":   saulArt,
	"ubuntu": UbuntuArt,
}
