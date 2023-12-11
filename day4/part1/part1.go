package main

import(
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
	raw_data, err := os.ReadFile("../input.txt")
	// raw_data, err := os.ReadFile("../test.txt")

	if err != nil {
		panic("Error reading file")
	}
	var data string = string(raw_data)
	data = strings.ReplaceAll(data, "Card", "c")

	var sum, card_sum int = 0, 0
	var is_card_numbers, is_lottery_numbers bool = false, false
	var num_string string = "" 
	var is_contained map[int]bool = make(map[int]bool)
	for i:=0;i<len(data);i++ {
		var curr string = string(data[i])
		if curr == "\n" { 
			sum += card_sum
			card_sum = 0
			is_contained = make(map[int]bool)
			is_card_numbers, is_lottery_numbers = false, false
			continue
		} else if curr == ":" {
			num_string = ""
			is_card_numbers, is_lottery_numbers = true, false
			continue
		} else if curr == "|" {
			num_string = ""
			is_card_numbers, is_lottery_numbers = false, true
			continue
		}

		if !is_card_numbers && !is_lottery_numbers || !IsNumericChar(data[i]){
			continue
		}

		num_string += curr
		if (i + 1 < len(data)) && IsNumericChar(data[(i+1)]) {
			continue
		}

		var num int = StringToInteger(num_string)
		_, ok := is_contained[num]; 
		if is_card_numbers {
			is_contained[num] = true
		} else if is_lottery_numbers && ok && card_sum == 0 {
			card_sum = 1
		} else if is_lottery_numbers && ok && card_sum > 0 {
			card_sum *= 2 
		}
		num_string = ""
	}
	fmt.Println(sum)
}