package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var print = fmt.Println

var letters = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
var list []string

func main() {
	dat, _ := ioutil.ReadFile("data.txt")
	str := strings.ReplaceAll(string(dat), "\r", "")
	parts := strings.Split(str, "\n")
	for i := 0; i < len(parts); i++ {
		fmt.Printf(smalpha(parts[i]) + "    %d\n", i)
	}
}

func smalpha(str string) string {
	list = list[:0]
	perm(str, "", 0)
	if len(list) > 0 { return list[0] }
	return ""
}

func perm(str string, result string, index int) {
	if index == len(str) {
		list = append(list, result)
	}

	if len(list) > 0 {
		return
	}

	for i := 0; i < 26; i++ {
		target := letters[i]

		if contains(result, byte(i + 97)) { continue }

		size := index + len(target)
		temp := ""
		for j := index; j < size; j++ {
			temp += string(str[j])
		}

		if temp == target {
			newResult := result + string(i + 97)
			perm(str, newResult, size)
		}
	}
}

func smorse(str string) string {
	result := ""

	for i := 0; i < len(str); i++ {
		index := str[i] - 97
		result += string(letters[index])
	}

	return result
}

func contains(str string, letter byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == letter {
			return true
		}
	}
	return false
}

func indexOf(slice []string, val string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return i
		}
	}
	return -1
}