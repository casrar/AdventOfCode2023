package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringToInteger(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Error converting string to int: \"" + s +"\"")
	}
	return i
}

func IsValidSymbol(c byte) bool {
	if (c == '@' || c == '#' || c == '$' || c == '%' || c == '&' ||
	c == '*' || c == '-' ||c == '=' || c == '+' ||c == '/' ) {
		return true
	}
	return false
}


func main() {
	raw_data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic("Error reading file")
	}
	var data string = string(raw_data)
	data = strings.ReplaceAll(data, "\n", "")
	// every line is 140 char long
	// 	-141	-140	-139
	//	-1		here	+1
	//  +139	+140	+141
	var offsets = [8]int{-141, -140, -139, -1, 1, 139, 140, 141}  
	var is_valid_number bool = false
	var number string = ""
	var sum int = 0
	for i:=0; i<len(data); i++ {
		var curr byte = data[i]
		if (curr < '0' || curr > '9') && is_valid_number {
			sum += StringToInteger(number)
			number = ""
			is_valid_number = false
			continue
		}
		if (curr < '0' || curr > '9') && !is_valid_number {
			number = ""
			continue
		}

		number += string(curr)
		for j:=0; j<len(offsets); j++ {
			var offset int = offsets[j]
			if offset < 0 && ((i + offset) >= 0) && IsValidSymbol(data[(i+offset)]) {
				is_valid_number = true
			}
			if offset > 0 && ((i + offset) < len(data)) && IsValidSymbol(data[(i+offset)]) {
				is_valid_number = true
			}
		}

	}
	fmt.Println(sum)

}