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

func isLetter(a byte) bool {
	return 64 < a && a < 91
}

func isWhitespace(a byte) bool {
	return a < 33
}

func getWords(line string) []string {
	word := ""
	words := make([]string, 0)
	i := 0
	for i < len(line) {
		for i < len(line) && !isLetter(line[i]) {
			i++
		}
		for i < len(line) && isLetter(line[i]) {
			word += string(line[i])
			i++
		}
		for i < len(line) && !isLetter(line[i]) {
			i++
		}
		i--
		words = append(words, word)
		word = ""
		i++
	}
	return words
}

func getInstructions(line string) []byte {
	instructions := make([]byte, 0)
	for idx := range line {
		if isLetter(line[idx]) {
			letter := line[idx]
			instructions = append(instructions, letter)
		}
	}
	return instructions
}

func generateGraph(input []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range input {
		words := getWords(line)
		graph[words[0]] = append(graph[line], words[1], words[2])
	}
	return graph
}

func getSteps(file *os.File) int {
	totalSteps := 0
	input := getInput(file)
	instructions := getInstructions(input[0])
	graph := generateGraph(input[2:])
	curr := "AAA"
	i := 0
	for curr != "ZZZ" {
		totalSteps += 1
		if instructions[i] == 'L' {
			curr = graph[curr][0]
		} else {
			curr = graph[curr][1]
		}
		i++
		i = i % len(instructions)
	}
	return totalSteps
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
	fmt.Println("PART 1:", getSteps(file)) // 21409
}
