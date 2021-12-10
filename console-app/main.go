package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-->")

		unserInput, _ := reader.ReadString('\n')
		unserInput = strings.Replace(unserInput, "\n", "", -1)

	}
}
