package utils

func Getascii(name string) string {
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
	case "arch":
		return arch

	default:
		return none
	}

}
