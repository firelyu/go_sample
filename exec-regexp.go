package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	REGEXP_FORMAT = "^(edit|save|view)/([a-zA-Z0-9]*)$"
	REGEXP_QUIT   = "[qQ][uU][iI][tT]|q|Q"
)

func main() {
	validInput := regexp.MustCompile(REGEXP_FORMAT)

	// read []byte from stdin
	reader := bufio.NewReader(os.Stdin)
	input := make([]byte, 100000)
	for {
		fmt.Print("Input:")
		l, err := reader.Read(input)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		cleanInput := input[:l-1]

		if v := regexp.MustCompile(REGEXP_QUIT); v.Match(cleanInput) {
			break
		}

		fmt.Printf("The input is %s", string(cleanInput))
		m := validInput.FindStringSubmatch(string(cleanInput))
		fmt.Println(m)
	}

	fmt.Println("See U.")
}
