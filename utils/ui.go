package utils

import (
	"fmt"
)
var Logo string  =  `
██████╗  ██████╗ ███╗  ███╗███╗  ███╗██╗███╗   ██╗
██╔════╝ ██╔═══██╗████╗ ████║████╗ ████║██║████╗ ██║
██║  ███╗██║   ██║██╔██╗╚███║██╔██╗╚███║██║██╔██╗██║
██║   ██║██║   ██║██║╚██╗██║██║╚██╗██║██║██║╚████║
╚██████╔╝╚██████╔╝██║ ╚████║██║ ╚████║██║██║ ╚███║
╚═════╝  ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═══╝╚═╝╚═╝  ╚══╝
`
var Tagline string = "Your personal Git assistant"
func main() {

    //fmt.Println(logo)

    //Tagline := "Your personal Git assistant'"
}

func PrintHeader(commandTitle string){
	fmt.Println(Logo)
	fmt.Printf("     ***GOMMIT YOUR PERSONAL GIT ASSISTANT***\n\n")

	fmt.Printf("%v \n\n", commandTitle)
}
