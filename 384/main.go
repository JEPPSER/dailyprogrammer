package main

import (
	"fmt"
)

var print = fmt.Println;
var list []string

func main() {
	necklaces(2, 12)
	necklaces(3, 7)
	necklaces(9, 4)
	necklaces(21, 3)
}

func necklaces(k int, length int) {
	list = list[:0]
	loop(k, length, "")

	var newList []string

	for i := 0; i < len(list); i++ {
		unique := true;
		for j := 0; j < len(newList); j++ {
			if sameNecklace(list[i], newList[j]) {
				unique = false;
				break;
			}
		}
		if unique { newList = append(newList, list[i]) }
	}

	print(len(newList))
}

func loop(k int, length int, str string) {
	if len(str) == length {
		list = append(list, str)
		return
	}

	for i := 0; i < k; i++ {
		temp := str + string(65 + i)
		loop(k, length, temp)
	}
}

func sameNecklace(one, two string) bool {
	if len(one) != len(two) { return false }
	if one == "" && two == "" { return true }

	for i := 0; i < len(one); i++ {
		if one == two { return true; }
		one = moveBead(one)
	}
	return false
}

func moveBead(str string) string {
	var temp = "";
	for i := 1; i < len(str); i++ {
		temp += string(str[i])
	}
	temp += string(str[0])
	return temp
}