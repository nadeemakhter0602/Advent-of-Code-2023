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

func getNumber(number []int) int {
	firstDigit := int(number[0])
	secondDigit := int(number[len(number)-1])
	return firstDigit*10 + secondDigit
}

func getCalibrationSum(file *os.File) int {
	sum := 0
	number := make([]int, 0)
	for true {
		buffer := make([]byte, 1)
		_, err := file.Read(buffer)
		if err == io.EOF {
			sum += getNumber(number)
			break
		}
		PanicErr(err)
		intValue := int(buffer[0])
		if 47 < intValue && intValue < 58 {
			number = append(number, intValue-48)
		}
		if buffer[0] == []byte{'\n'}[0] {
			sum += getNumber(number)
			number = make([]int, 0)
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getCalibrationSum(file)) // 53921
}
