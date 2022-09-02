package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args[1:]) == 1 {
		file, _ := ioutil.ReadFile(os.Args[1])
		text := string(file)
		fmt.Print(text)
	} else if len(os.Args[1:]) > 1 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
}
