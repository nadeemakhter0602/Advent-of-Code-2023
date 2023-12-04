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

func getNumbers(line string) ([]string, map[string]bool) {
	winningNumbers := make(map[string]bool)
	cardNumbers := make([]string, 0)
	i := 0
	for line[i] != ':' {
		i++
	}
	number := ""
	for line[i] != '|' {
		if isDigit(line[i]) {
			number += string(line[i])
		} else if len(number) > 0 {
			cardNumbers = append(cardNumbers, number)
			number = ""
		}
		i++
	}
	cardNumbers = append(cardNumbers, number)
	number = ""
	for i < len(line) {
		if isDigit(line[i]) {
			number += string(line[i])
		} else if len(number) > 0 {
			winningNumbers[number] = true
			number = ""
		}
		i++
	}
	return cardNumbers, winningNumbers
}

func getPoint(line string) int {
	point := 1
	cardNumbers, winningNumbers := getNumbers(line)
	for _, number := range cardNumbers {
		_, ok := winningNumbers[number]
		if ok {
			point <<= 1
		}
	}
	return point >> 1
}

func getTotalPoints(file *os.File) int {
	totalPoints := 0
	line := ""
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		line += string(buffer[0])
		if buffer[0] == 10 || err == io.EOF {
			totalPoints += getPoint(line)
			line = ""
		}
	}
	return totalPoints
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getTotalPoints(file)) // 22674
}
