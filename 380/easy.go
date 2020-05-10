package main

import (
	"fmt"
)

var print = fmt.Println;

var letters = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func main() {
	print(smorse("sos"))
	print(smorse("daily"))
	print(smorse("programmer"))
	print(smorse("bits"))
	print(smorse("three"))
}

func smorse(str string) string {
	result := ""

	for i := 0; i < len(str); i++ {
		index := str[i] - 97
		result += string(letters[index])
	}

	return result
}