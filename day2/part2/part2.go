package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func StringToInteger(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Error converting string to int: \"" + s +"\"")
	}
	return i
}

func FormatData(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "Game", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ";", "")
	s = strings.ReplaceAll(s, "red", "r")
	s = strings.ReplaceAll(s, "green", "g")
	s = strings.ReplaceAll(s, "blue", "b")
	return s
}

func SetHighest(count int, highest_map 	map[string]int, color string) map[string]int {
	if color == "r" && count > highest_map[color] {
		highest_map[color] = count 
	}
	if color == "g" && count > highest_map[color] {
		highest_map[color] = count 
	}
	if color == "b" && count > highest_map[color] {
		highest_map[color] = count 
	}
	return highest_map
} 

func main() {
	raw_data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic("Error reading file")
	}
	var highest_map = map[string]int{"r": 0, "g": 0, "b": 0}

	data := string(raw_data)
	data = FormatData(data)

	var temp_string string = ""
	var sum int = 0
	var game_number_skipped bool = false
	for i:=0; i<len(data); i++ {
		if (data[i] == 10) {
			sum += (highest_map["r"] * highest_map["g"] * highest_map["b"])
			highest_map["r"] = 0
			highest_map["g"] = 0
			highest_map["b"] = 0
			temp_string = ""
			game_number_skipped = false
			continue
		}

		curr_string := string(data[i])
		if curr_string == ":" { 
			game_number_skipped = true
			temp_string = "" 
			continue 
		} 

		if !game_number_skipped {
			continue
		}

		if !(curr_string == "r" || curr_string == "g" || curr_string == "b") {
			temp_string += curr_string
			continue
		}

		count := StringToInteger(temp_string)
		highest_map = SetHighest(count, highest_map, curr_string)
		temp_string = ""
	}

	fmt.Println(sum)
}