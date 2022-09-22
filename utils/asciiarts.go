package utils

import (
	"strings"
)

func Getascii(name string) []string {
	none := `!!!!!!!!!!!!!!!  
!!!!!!!!!!!!!!!  
!!!!noascii!!!!  
!!!!!!!!!!!!!!!  
!!!!!!!!!!!!!!!  `
	arch := `      /\        
     /  \       
    /\   \      
   / > ω <\     
  /   __   \    
 / __|  |__-\   
/_-''    ''-_\  `

	switch name {
	case "Arch Linux":
		return strings.Split(arch, "\n")

	default:
		return strings.Split(none, "\n")

	}

}
