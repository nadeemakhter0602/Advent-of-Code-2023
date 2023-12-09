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

func isAlphaNumeric(a byte) bool {
	return (64 < a && a < 91) || isDigit(a)
}

func isWhitespace(a byte) bool {
	return a < 33
}

func getWords(line string) []string {
	word := ""
	words := make([]string, 0)
	i := 0
	for i < len(line) {
		for i < len(line) && !isAlphaNumeric(line[i]) {
			i++
		}
		for i < len(line) && isAlphaNumeric(line[i]) {
			word += string(line[i])
			i++
		}
		for i < len(line) && !isAlphaNumeric(line[i]) {
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
		if isAlphaNumeric(line[idx]) {
			letter := line[idx]
			instructions = append(instructions, letter)
		}
	}
	return instructions
}

func generateGraph(input []string) (map[string][]string, []string) {
	graph := make(map[string][]string)
	starts := make([]string, 0)
	for _, line := range input {
		words := getWords(line)
		if words[0][len(words)-1] == 'A' {
			starts = append(starts, words[0])
		}
		graph[words[0]] = append(graph[line], words[1], words[2])
	}
	return graph, starts
}

func getGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func getLCM(a, b int, integers ...int) int {
	result := a * b / getGCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = getLCM(result, integers[i])
	}
	return result
}

func getSteps(file *os.File) int {
	input := getInput(file)
	instructions := getInstructions(input[0])
	graph, starts := generateGraph(input[2:])
	subSteps := make([]int, 0)
	for _, curr := range starts {
		steps := 0
		i := 0
		for curr[len(curr)-1] != 'Z' {
			steps += 1
			if instructions[i] == 'L' {
				curr = graph[curr][0]
			} else {
				curr = graph[curr][1]
			}
			i++
			i = i % len(instructions)
		}
		subSteps = append(subSteps, steps)
	}
	return getLCM(1, 1, subSteps...)
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
	fmt.Println("PART 2:", getSteps(file)) // 21165830176709
}
