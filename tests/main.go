package main

import (
	"fmt"
	

	"github.com/itsllyaz/gommit/utils"

)


func main(){
	fmt.Println("............") 
	fmt.Println("............")
	
	jsonContent := utils.HandleGemini()
	message := utils.ExtractContent(jsonContent)
	fmt.Println(message)
}
