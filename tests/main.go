package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main(){
	wd, _ := os.Getwd()
	fmt.Println(wd)
	b1 := filepath.Base(wd+"/file.txt")
	fmt.Println("......")
	fmt.Println(b1)
}
