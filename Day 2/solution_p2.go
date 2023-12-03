package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getPower(line string) int {
	game := strings.Split(line, " ")
	colorMax := map[string]int{
		"r": 0,
		"g": 0,
		"b": 0,
	}
	game = game[2:]
	for idx, element := range game {
		number, err := strconv.Atoi(element)
		if err != nil {
			continue
		} else {
			color := string(game[idx+1][0])
			colorNumber, ok := colorMax[color]
			if ok && number > colorNumber {
				colorMax[color] = number
			}
		}
	}
	power := colorMax["r"] * colorMax["g"] * colorMax["b"]
	return power
}

func getPowersSum(file *os.File) int {
	sum := 0
	line := ""
	for true {
		buffer := make([]byte, 1)
		_, err := file.Read(buffer)
		if err == io.EOF {
			sum += getPower(line)
			break
		}
		line += string(buffer[0])
		PanicErr(err)
		if buffer[0] == []byte{'\n'}[0] {
			sum += getPower(line)
			line = ""
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 2:", getPowersSum(file)) // 67363
}
