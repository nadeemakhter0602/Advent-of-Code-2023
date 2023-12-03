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

func checkGame(line string) int {
	game := strings.Split(line, " ")
	gameIDString := game[1]
	gameIDString = gameIDString[:len(gameIDString)-1]
	gameID, err := strconv.Atoi(gameIDString)
	colorNumbers := map[string]int{
		"r": 12,
		"g": 13,
		"b": 14,
	}
	PanicErr(err)
	game = game[2:]
	for idx, element := range game {
		number, err := strconv.Atoi(element)
		if err != nil {
			continue
		} else {
			color := string(game[idx+1][0])
			colorNumber, ok := colorNumbers[color]
			if ok && number > colorNumber {
				return 0
			}
		}
	}
	return gameID
}

func getPossibleGamesSum(file *os.File) int {
	sum := 0
	line := ""
	for true {
		buffer := make([]byte, 1)
		_, err := file.Read(buffer)
		if err == io.EOF {
			sum += checkGame(line)
			break
		}
		line += string(buffer[0])
		PanicErr(err)
		if buffer[0] == []byte{'\n'}[0] {
			sum += checkGame(line)
			line = ""
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getPossibleGamesSum(file))
}
