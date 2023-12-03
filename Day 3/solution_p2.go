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
	for j, line := range schematic {
		number := 0
		gearCoordinates := make([]string, 0)
		for i, element := range line {
			if isDigit(element) {
				number = 10*number + (int(element) - '0')
				if i+1 < len(line) {
					right := rune(schematic[j][i+1])
					if right == '*' {
						gearCoordinate := strconv.Itoa(j) + "," + strconv.Itoa(i+1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if i-1 > 0 {
					left := rune(schematic[j][i-1])
					if left == '*' {
						gearCoordinate := strconv.Itoa(j) + "," + strconv.Itoa(i-1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if j+1 < len(schematic) {
					down := rune(schematic[j+1][i])
					if down == '*' {
						gearCoordinate := strconv.Itoa(j+1) + "," + strconv.Itoa(i)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if j-1 > 0 {
					up := rune(schematic[j-1][i])
					if up == '*' {
						gearCoordinate := strconv.Itoa(j-1) + "," + strconv.Itoa(i)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if i+1 < len(line) && j+1 < len(schematic) {
					down_right := rune(schematic[j+1][i+1])
					if down_right == '*' {
						gearCoordinate := strconv.Itoa(j+1) + "," + strconv.Itoa(i+1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if i-1 > 0 && j+1 < len(schematic) {
					down_left := rune(schematic[j+1][i-1])
					if down_left == '*' {
						gearCoordinate := strconv.Itoa(j+1) + "," + strconv.Itoa(i-1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if i-1 > 0 && j-1 > 0 {
					up_left := rune(schematic[j-1][i-1])
					if up_left == '*' {
						gearCoordinate := strconv.Itoa(j-1) + "," + strconv.Itoa(i-1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
						}
					}
				}
				if i+1 < len(line) && j-1 > 0 {
					up_right := rune(schematic[j-1][i+1])
					if up_right == '*' {
						gearCoordinate := strconv.Itoa(j-1) + "," + strconv.Itoa(i+1)
						gearCoordinatesLen := len(gearCoordinates)
						if gearCoordinatesLen == 0 || gearCoordinates[gearCoordinatesLen-1] != gearCoordinate {
							gearCoordinates = append(gearCoordinates, gearCoordinate)
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
	fmt.Println("PART 2:", getPartNumbersSum(file))
}
