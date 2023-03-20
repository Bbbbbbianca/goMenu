package main

import "fmt"

func main() {
	for {
		var cmd string
		fmt.Print("MENU: \n 1. Welcome. \n 2. Help. \n 3. Quit.\n")
		fmt.Scanln(&cmd)
		if cmd == "3" {
			break
		}
		switch cmd {
		case "1":
			fmt.Println("Hello world!")
		case "2":
			fmt.Println("Help command")
		default:
			fmt.Println("Unknown command")
		}
	}
}
