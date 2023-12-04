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

func getWins(line string) int {
	wins := 0
	cardNumbers, winningNumbers := getNumbers(line)
	for _, number := range cardNumbers {
		_, ok := winningNumbers[number]
		if ok {
			wins += 1
		}
	}
	return wins
}

func updateCards(cardNumber int, wins int, numberOfCards []int) []int {
	if len(numberOfCards) < cardNumber {
		numberOfCards = append(numberOfCards, 1)
	} else {
		numberOfCards[cardNumber-1] += 1
	}
	cardCopies := numberOfCards[cardNumber-1]
	for j := 0; j < cardCopies; j++ {
		for i := cardNumber + 1; i <= cardNumber+wins; i++ {
			if len(numberOfCards) < i {
				numberOfCards = append(numberOfCards, 1)
			} else {
				numberOfCards[i-1] += 1
			}
		}
	}
	return numberOfCards
}

func getTotalCards(file *os.File) int {
	totalCards := 0
	line := ""
	cardNumber := 1
	numberOfCards := make([]int, 0)
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		line += string(buffer[0])
		if buffer[0] == 10 || err == io.EOF {
			wins := getWins(line)
			numberOfCards = updateCards(cardNumber, wins, numberOfCards)
			cardNumber += 1
			line = ""
		}
	}
	for _, element := range numberOfCards {
		totalCards += element
	}
	return totalCards
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 2:", getTotalCards(file)) // 5747443
}
