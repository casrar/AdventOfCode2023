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

func IsNumericChar(c byte) bool {
	if (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func main() {
	raw_data, err := os.ReadFile("../test.txt")
	if err != nil {
		panic("Error reading file")
	}

	var data string = string(raw_data)
	data = strings.ReplaceAll(data, "\n", "")
	var offsets = [8]int{-141, -140, -139, -1, 1, 139, 140, 141}  


	var sum int = 0
	for i:=0;i<len(data);i++ {
		curr := string(data[i])
		if curr != "*" {
			continue
		}
		var first_index int = -1
		var last_index int = -1
		var is_continuous bool = true
		for j:=0; j<len(offsets); j++ {
			var offset int = offsets[j]
			if i + offset < 0 || i + offset > len(data) {
				continue
			} 

			fmt.Print("Offset: ")
			fmt.Print(offset)
			fmt.Println( " | curr: " + string(data[(i+offset)]))
			

			// unhandled edge cases
			if first_index == -1 && IsNumericChar(data[(i+offset)]) {
				fmt.Println("first if")
				first_index = i + offset
			} else if first_index > -1 && !IsNumericChar(data[(i+offset)]) {
				fmt.Println("second if")
				is_continuous = false
			} else if !is_continuous && IsNumericChar(data[(i+offset)]) {
				fmt.Println("third if")
				last_index = i + offset
			} else if first_index > -1 && offset == 1 && IsNumericChar(data[(i+offset)]) { 
				fmt.Println("fourth if")
				is_continuous = false
				last_index = i + offset
			}
		}
		if last_index == -1 || is_continuous {
			fmt.Println("in")
			continue
		}
		for {
			if first_index < 0 {
				break
			}
			if !IsNumericChar(data[first_index]) {
				break
			}
			first_index--
		}
		for {
			if last_index < 0 {
				break
			}
			if !IsNumericChar(data[last_index]) {
				break
			}
			last_index--
		}
		first_index++
		last_index++

		var first_num string = ""
		var second_num string = ""
		for {
			if first_index >= len(data) {
				break
			}
			if !IsNumericChar(data[first_index]) {
				break
			}
			first_num += string(data[first_index])
			first_index++
		}
		for {
			if last_index >= len(data) {
				break
			}
			if !IsNumericChar(data[last_index]) {
				break
			}
			second_num += string(data[last_index])
			last_index++
		}
		fmt.Println("First num: " + first_num + " | Second num: " + second_num)
		sum += (StringToInteger(first_num) * StringToInteger(second_num))
	}
	fmt.Println(sum)
}
