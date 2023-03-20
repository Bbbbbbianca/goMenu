package main
import "fmt"
 
func main() {
	for {
		var cmd string
		fmt.Scan(&cmd)
		if cmd == "3" {
			break
		}
		switch cmd {
		case "1":
			fmt.Println("Hello world!")
		case "2":
			fmt.Println("Help command:")
		default:
			fmt.Println("Unknown command")
		}
	}
}