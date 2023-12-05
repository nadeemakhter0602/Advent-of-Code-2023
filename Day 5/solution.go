package main

import (
	"fmt"
	"io"
	"math"
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

func transformSeeds(seeds []int, mapping [][]int) []int {
	for idx, seed := range seeds {
		for _, element := range mapping {
			dest, src, length := element[0], element[1], element[2]
			if seed >= src && seed < src+length {
				seeds[idx] = dest + (seed - src)
			}
		}
	}
	return seeds
}

func mapSeedsToLocations(input []string) []int {
	seeds := getNumbers(input[0])
	mapping := make([][]int, 0)
	i := 2
	for i < len(input) {
		if isDigit(input[i][0]) {
			mapping = append(mapping, getNumbers(input[i]))
		} else if isWhitespace(input[i][0]) {
			seeds = transformSeeds(seeds, mapping)
			mapping = make([][]int, 0)
		}
		i++
	}
	seeds = transformSeeds(seeds, mapping)
	return seeds
}

func getLowestLocation(file *os.File) int {
	input := getInput(file)
	locations := mapSeedsToLocations(input)
	lowestLocation := math.Inf(1)
	for _, location := range locations {
		lowestLocation = min(lowestLocation, float64(location))
	}
	return int(lowestLocation)
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
	fmt.Println("PART 1:", getLowestLocation(file)) // 650599855
}
