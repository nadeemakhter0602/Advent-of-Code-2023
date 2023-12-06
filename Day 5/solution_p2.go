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

func splitRange(start, end int, mapping []int) ([][]int, [][]int) {
	dest, src, length := mapping[0], mapping[1], mapping[2]
	mapped := make([][]int, 0)
	unmapped := make([][]int, 0)
	if start >= src && start < src+length {
		if end < src+length {
			mapped = append(mapped, []int{dest + (start - src), dest + (end - src)})
		} else if end >= src+length {
			mapped = append(mapped, []int{dest + (start - src), dest + length - 1})
			unmapped = append(unmapped, []int{src + length, end})
		}
	} else if start < src {
		if end < src {
			unmapped = append(unmapped, []int{start, end})
		} else if end >= src && end < src+length {
			unmapped = append(unmapped, []int{start, src - 1})
			mapped = append(mapped, []int{dest, dest + (end - src)})
		} else {
			unmapped = append(unmapped, []int{start, src - 1})
			mapped = append(mapped, []int{dest, dest + length - 1})
			unmapped = append(unmapped, []int{src + length, end})
		}
	} else {
		unmapped = append(unmapped, []int{start, end})
	}
	return mapped, unmapped
}

func transformSeeds(ranges [][]int, mapping [][]int) [][]int {
	transformed := make([][]int, 0)
	for _, interval := range ranges {
		mapped := make([][]int, 0)
		unmapped := [][]int{interval}
		for _, element := range mapping {
			temp := make([][]int, len(unmapped))
			copy(temp, unmapped)
			for _, r := range temp {
				mapped, unmapped = splitRange(r[0], r[1], element)
				transformed = append(transformed, mapped...)
			}
		}
		transformed = append(transformed, unmapped...)
	}
	return transformed
}

func mapSeedsToLocations(input []string) [][]int {
	seedRanges := getNumbers(input[0])
	mapping := make([][]int, 0)
	ranges := make([][]int, 0)
	j := 0
	for j < len(seedRanges) {
		start := seedRanges[j]
		end := seedRanges[j+1] + start - 1
		ranges = append(ranges, []int{start, end})
		j += 2
	}
	i := 2
	for i < len(input) {
		if isDigit(input[i][0]) {
			mapping = append(mapping, getNumbers(input[i]))
		} else if isWhitespace(input[i][0]) {
			ranges = transformSeeds(ranges, mapping)
			mapping = make([][]int, 0)
		}
		i++
	}
	ranges = transformSeeds(ranges, mapping)
	return ranges
}

func getLowestLocation(file *os.File) int {
	input := getInput(file)
	locations := mapSeedsToLocations(input)
	lowestLocation := math.Inf(1)
	for _, location := range locations {
		lowestLocation = min(lowestLocation, float64(location[0]))
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
	fmt.Println("PART 2:", getLowestLocation(file)) // 1240035
}
