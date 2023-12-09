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

func isSign(a byte) bool {
	return a == 45
}

func isDigit(a byte) bool {
	return 47 < a && a < 58
}

func isDigitOrSign(a byte) bool {
	return isDigit(a) || isSign(a)
}

func isWhitespace(a byte) bool {
	return a < 33
}

func getNumbers(line string) []int {
	i := 0
	number := 0
	sign := 1
	numbers := make([]int, 0)
	for i < len(line) {
		for i < len(line) && !isDigitOrSign(line[i]) {
			i++
		}
		for i < len(line) && isSign(line[i]) {
			sign = -1
			i++
		}
		for i < len(line) && isDigit(line[i]) {
			number = number*10 + (int(line[i]) - '0')
			i++
		}
		for i < len(line) && !isDigitOrSign(line[i]) {
			i++
		}
		i--
		number = sign * number
		numbers = append(numbers, number)
		sign = 1
		number = 0
		i++
	}
	return numbers
}

func predictNext(history []int) int {
	nextPrediction := history[len(history)-1]
	allZeroes := false
	for !allZeroes {
		idx := 0
		temp := make([]int, len(history))
		copy(temp, history)
		history = make([]int, 0)
		allZeroes = true
		for idx < len(temp)-1 {
			difference := temp[idx+1] - temp[idx]
			if difference != 0 {
				allZeroes = false
			}
			history = append(history, difference)
			idx++
		}
		nextPrediction += history[len(history)-1]
	}
	return nextPrediction
}

func getPredictionSum(file *os.File) int {
	predictionSum := 0
	histories := getOASISReport(file)
	for _, history := range histories {
		predictionSum += predictNext(history)
	}
	return predictionSum
}

func getOASISReport(file *os.File) [][]int {
	line := ""
	histories := make([][]int, 0)
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		line += string(buffer[0])
		if buffer[0] == 10 || err == io.EOF {
			histories = append(histories, getNumbers(line))
			line = ""
		}
	}
	return histories
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getPredictionSum(file)) // 1696140818
}
