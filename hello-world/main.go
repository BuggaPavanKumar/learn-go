package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// Topic: Using function
/* func main() {
	fmt.Println("Hello, world!")
} */

// Topic: function calling
/* func main() {
	sayHelloWorld("Hello World")
}
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
} */

// Topic: Variables
/* func main() {
	//var whatToSay string
	//whatToSay = "Hello world, again!"
	whatToSay := "Hello world, again!"
	sayHelloWorld(whatToSay)
}

func sayHelloWorld(wayToSay string) {
	fmt.Println(wayToSay)
} */

// Create a module: `go mod init myapp`
// Topic: import modules, read i/p, respond
func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // Windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // unix / macos
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}
