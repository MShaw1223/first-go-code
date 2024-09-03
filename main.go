package main

import "fmt"

func main(){
	hello_world()
	take_cli_input()
}

func hello_world(){
	fmt.Println("Hello, World")
}

func take_cli_input(){
	var input string

	fmt.Print("Enter your name: ")
	_,err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:",err)
		return
	}
    // instead of println as pln adds space between args; Print does not add a space between args
	fmt.Print("Your name is ",input,", hi there ",input,"!\n")
}
