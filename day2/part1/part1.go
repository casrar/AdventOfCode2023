package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

const RED_MAX = 12
const GREEN_MAX = 13
const BLUE_MAX = 14


func StringToIneger(s string) int {
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

func main() {
	
	raw_data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic("Error reading file")
	}

	data := string(raw_data)
	data = FormatData(data)

	var temp_string string = ""
	var sum int = 0
	game_id := 0
	for i:=0; i<len(data); i++ {
		if (data[i] == 10) {
			sum += game_id
			temp_string = ""
			continue
		}

		curr_string := string(data[i])
		if curr_string == ":" {
			game_id = StringToIneger(temp_string)
			temp_string = ""
			continue 
		} 

		if !(curr_string == "r" || curr_string == "g" || curr_string == "b") {
			temp_string += curr_string
			continue
		}

		value := StringToIneger(temp_string)
		if curr_string == "r" && value > RED_MAX {
			game_id = 0
		}
		if curr_string == "g" && value > GREEN_MAX {
			game_id = 0
		}
		if curr_string == "b" && value > BLUE_MAX {
			game_id = 0
		}
		temp_string = ""
	}

	fmt.Println(sum)
}