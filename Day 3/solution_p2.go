package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
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
	gearToPartNumber := make(map[string][]int)
	relativeCoordinates := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}
	for j, line := range schematic {
		number := 0
		gearCoordinates := make([]string, 0)
		for i, element := range line {
			if isDigit(element) {
				number = 10*number + (int(element) - '0')
				for _, relativeCoordinate := range relativeCoordinates {
					y := j + relativeCoordinate[0]
					x := i + relativeCoordinate[1]
					if y >= 0 && y < len(schematic) && x >= 0 && x < len(line) {
						elementXY := rune(schematic[y][x])
						if elementXY == '*' {
							gearCoordinate := strconv.Itoa(y) + "," + strconv.Itoa(x)
							gearCoordinatesLen := len(gearCoordinates)
							if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
								gearCoordinates = append(gearCoordinates, gearCoordinate)
							}
						}
					}
				}
			} else {
				for _, gearCoordinate := range gearCoordinates {
					gearToPartNumber[gearCoordinate] = append(gearToPartNumber[gearCoordinate], number)
				}
				number = 0
				gearCoordinates = make([]string, 0)
			}
		}
	}
	for _, numbers := range gearToPartNumber {
		if len(numbers) == 2 {
			gearPartNumber1 := numbers[0]
			gearPartNumber2 := numbers[1]
			gearRatio := gearPartNumber1 * gearPartNumber2
			sum += gearRatio
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
	fmt.Println("PART 2:", getPartNumbersSum(file)) // 76314915
}
