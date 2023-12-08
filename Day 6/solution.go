package main

import (
	"fmt"
	"io"
	"os"
)

func PanicErr(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

func isDigit(a byte) bool {
	return 47 < a && a < 58
}

func isWhitespace(a byte) bool {
	return a < 33
}

func getNumbers(line string) []int {
	i := 0
	number := 0
	numbers := make([]int, 0)
	for i < len(line) {
		for i < len(line) && !isDigit(line[i]) {
			i++
		}
		for i < len(line) && isDigit(line[i]) {
			number = number*10 + (int(line[i]) - '0')
			i++
		}
		for i < len(line) && !isDigit(line[i]) {
			i++
		}
		i--
		numbers = append(numbers, number)
		number = 0
		i++
	}
	return numbers
}

func waysToWinProduct(file *os.File) int {
	input := getInput(file)
	times := getNumbers(input[0])
	distances := getNumbers(input[1])
	product := 1
	for idx, time := range times {
		distance := distances[idx]
		waysToWin := 0
		for hold := 0; hold <= time; hold++ {
			possibleDistance := hold * (time - hold)
			if possibleDistance > distance {
				waysToWin++
			}
		}
		product *= waysToWin
	}
	return product
}

func getInput(file *os.File) []string {
	line := ""
	input := make([]string, 0)
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		line += string(buffer[0])
		if buffer[0] == 10 || err == io.EOF {
			input = append(input, line)
			line = ""
		}
	}
	return input
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", waysToWinProduct(file)) // 316800
}
