package main

import (
    "fmt"
)

var print = fmt.Println;

func main() {
	print(sameNecklace("nicole", "icolen"))
	print(sameNecklace("nicole", "lenico"))
	print(sameNecklace("nicole", "coneli"))
	print(sameNecklace("aabaaaaabaab", "aabaabaabaaa"))
	print(sameNecklace("abc", "cba"))
	print(sameNecklace("xxyyy", "xxxyy"))
	print(sameNecklace("xyxxz", "xxyxz"))
	print(sameNecklace("x", "x"))
	print(sameNecklace("x", "xx"))
	print(sameNecklace("x", ""))
	print(sameNecklace("", ""))
}

func sameNecklace(one, two string) bool {
	if (len(one) != len(two)) { return false }
	if (one == "" && two == "") { return true }

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