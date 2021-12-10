package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	userInput := readString("What is your name ?")
	age := readInt("How old are you ?")
	//fmt.Println("Your name is", userInput, ". You are", age, "years old")
	//fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userInput, age))
	fmt.Printf("Your name is %s. You are %d years old \n", userInput, age)
}

func promt() {
	fmt.Print("-->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		promt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // Windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // unix / macos

		if userInput == "" {
			fmt.Println("Please enter value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		promt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1) // Windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // unix / macos

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}

	}
}
