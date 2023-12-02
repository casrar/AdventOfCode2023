package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsNumericChar(c byte) bool {
	if (c >= '0' && c <= '9') {
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
	numbers := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	data_raw, err := os.ReadFile("../input.txt")
	if err != nil {
		panic("Error reading file")
	}
	var data string = string(data_raw) 

	for i:=0; i<len(numbers); i++ {
		data = strings.ReplaceAll(data, numbers[i], strconv.Itoa(i)) // this is replacing some of the letters we need, eightwo becomes eigh2
	}

	var sum int = 0
	var first string
	var second string 
	var new_line = true
	var temp string = ""
	for i:=0; i<len(data); i++{
		temp += string(data[i])
		if !IsNumericChar(data[i]) && data[i] != 10{
			continue
		}
		if (data[i] == 10) {
			sum += ConvertStringToInt(first + second)
			// fmt.Println(temp + "|" + first + second)
			new_line = true
			temp = ""
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