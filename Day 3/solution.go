package main

import (
	"fmt"
	"io"
	"os"
)

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func isDigit(a rune) bool {
	return 47 < a && a < 58
}

func isWhitespace(a rune) bool {
	return a < 33
}

func isSymbol(a rune) bool {
	return a != '.' && !isDigit(a) && !isWhitespace(a)
}

func getPartNumbersSum(file *os.File) int {
	schematic := generateEngineSchematic(file)
	sum := 0
	for j, line := range schematic {
		number := 0
		isPartNumber := false
		for i, element := range line {
			if isDigit(element) {
				number = 10*number + (int(element) - '0')
				right := '\x00'
				left := '\x00'
				up := '\x00'
				down := '\x00'
				up_left := '\x00'
				up_right := '\x00'
				down_left := '\x00'
				down_right := '\x00'
				if i+1 < len(line) {
					right = rune(schematic[j][i+1])
				}
				if i-1 > 0 {
					left = rune(schematic[j][i-1])
				}
				if j+1 < len(schematic) {
					down = rune(schematic[j+1][i])
				}
				if j-1 > 0 {
					up = rune(schematic[j-1][i])
				}
				if i+1 < len(line) && j+1 < len(schematic) {
					down_right = rune(schematic[j+1][i+1])
				}
				if i-1 > 0 && j+1 < len(schematic) {
					down_left = rune(schematic[j+1][i-1])
				}
				if i-1 > 0 && j-1 > 0 {
					up_left = rune(schematic[j-1][i-1])
				}
				if i+1 < len(line) && j-1 > 0 {
					up_right = rune(schematic[j-1][i+1])
				}
				isPartNumber = isPartNumber || isSymbol(right)
				isPartNumber = isPartNumber || isSymbol(left)
				isPartNumber = isPartNumber || isSymbol(up)
				isPartNumber = isPartNumber || isSymbol(down)
				isPartNumber = isPartNumber || isSymbol(up_left)
				isPartNumber = isPartNumber || isSymbol(up_right)
				isPartNumber = isPartNumber || isSymbol(down_left)
				isPartNumber = isPartNumber || isSymbol(down_right)
			} else {
				if isPartNumber {
					sum += number
				}
				number = 0
				isPartNumber = false
			}
		}
	}
	return sum
}

func generateEngineSchematic(file *os.File) []string {
	schematic := make([]string, 0)
	line := ""
	for true {
		buffer := make([]byte, 1)
		_, err := file.Read(buffer)
		if err == io.EOF {
			schematic = append(schematic, line)
			break
		}
		line += string(buffer[0])
		PanicErr(err)
		if buffer[0] == []byte{'\n'}[0] {
			schematic = append(schematic, line)
			line = ""
		}
	}
	return schematic
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getPartNumbersSum(file))
}
