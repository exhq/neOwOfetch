package images

import _ "embed"

//go:embed arch.png
var ArchArt []byte

//go:embed ubuntu.jpg
var UbuntuArt []byte

var DistroImages = map[string][]byte{
	"arch":   ArchArt,
	"ubuntu": UbuntuArt,
}
