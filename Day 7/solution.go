package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
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

func getNumbers(line string) int {
	i := 0
	number := 0
	for i < len(line) {
		if isDigit(line[i]) {
			number = number*10 + (int(line[i]) - '0')
		}
		i++
	}
	return number
}

func getStrength(hand string) (int, int) {
	cardToStrength := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	numberToType := map[int]int{
		23:    4,
		113:   3,
		122:   2,
		1112:  1,
		11111: 0,
	}
	countCards := make([]int, 14)
	handStrength := 0
	for _, card := range hand {
		strength := cardToStrength[card]
		countCards[strength] += 1
		handStrength = 100*handStrength + strength
	}
	slices.Sort(countCards)
	if countCards[len(countCards)-1] > 3 {
		return handStrength, countCards[len(countCards)-1] + 1
	}
	number := 0
	for _, value := range countCards {
		if value != 0 {
			number = number*10 + value
		}
	}
	return handStrength, numberToType[number]
}

func getTotalWinnings(file *os.File) int {
	totalWinnings := 0
	handsAndBids := getInput(file)
	sort.Slice(handsAndBids, func(i, j int) bool {
		hand1 := handsAndBids[i][0]
		hand2 := handsAndBids[j][0]
		handStrength1, handType1 := getStrength(hand1)
		handStrength2, handType2 := getStrength(hand2)
		if handType1 != handType2 {
			return handType1 < handType2
		}
		return handStrength1 < handStrength2
	})
	for rank, handAndBid := range handsAndBids {
		bid := getNumbers(handAndBid[1])
		totalWinnings += (rank + 1) * bid
	}
	return totalWinnings
}

func getInput(file *os.File) [][]string {
	handsAndBids := make([][]string, 0)
	line := ""
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		line += string(buffer[0])
		if buffer[0] == 10 || err == io.EOF {
			handAndBid := []string{line[:5], line[5:]}
			handsAndBids = append(handsAndBids, handAndBid)
			line = ""
		}
	}
	return handsAndBids
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getTotalWinnings(file)) // 249390788
}
