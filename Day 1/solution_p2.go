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

func getNumber(numbers []int) int {
	if len(numbers) < 1 {
		return 0
	}
	firstDigit := numbers[0]
	secondDigit := numbers[len(numbers)-1]
	return firstDigit*10 + secondDigit
}

func getCalibrationSum(file *os.File) int {
	sum := 0
	numbers := make([]int, 0)
	wordToNumber := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
	window3 := ""
	window4 := ""
	window5 := ""
	for true {
		buffer := make([]byte, 1)
		_, err := file.Read(buffer)
		if err == io.EOF {
			sum += getNumber(numbers)
			break
		}
		PanicErr(err)
		intValue := int(buffer[0])
		if 47 < intValue && intValue < 58 {
			numbers = append(numbers, intValue-48)
			window3 = ""
			window4 = ""
			window5 = ""
		} else if 96 < intValue && intValue < 123 {
			window3 += string(intValue)
			window4 += string(intValue)
			window5 += string(intValue)
			if len(window3) > 3 {
				window3 = window3[1:]
			}
			if len(window4) > 4 {
				window4 = window4[1:]
			}
			if len(window5) > 5 {
				window5 = window5[1:]
			}
			number, ok := wordToNumber[window3]
			if ok {
				numbers = append(numbers, number)
			} else {
				number, ok = wordToNumber[window4]
				if ok {
					numbers = append(numbers, number)
				} else {
					number, ok = wordToNumber[window5]
					if ok {
						numbers = append(numbers, number)
					}
				}
			}
		}
		if buffer[0] == []byte{'\n'}[0] {
			sum += getNumber(numbers)
			numbers = make([]int, 0)
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 2:", getCalibrationSum(file))
}
