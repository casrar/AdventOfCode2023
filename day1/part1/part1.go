package main

import (
	"fmt"
	"os"
	"strconv"
)

func IsNumericChar(c byte) bool {
	
	if (c >= 0 && c <= '9') {
		return true
	}
	return false
}

func ConvertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Error converting string to int")
	}
	return i
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Error reading file")
	}

	var sum int = 0
	var first string
	var second string 
	var new_line = true
	for i:=0; i<len(data); i++{
		if !IsNumericChar(data[i]){
			continue
		}
		if (data[i] == 10) {
			sum += ConvertStringToInt(first + second)
			new_line = true
			continue
		}

		curr := string(data[i])
		if new_line == true {
			first = curr
			second = first
			new_line = false
		} else {
			second = curr
		}

	}
	fmt.Println(sum)
}