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

func doesExist(pipe_xy, pipe_XY rune, valid [][]rune) bool {
	exists := false
	for _, pipe := range valid[0] {
		exists = exists || pipe == pipe_xy
	}
	for _, pipe := range valid[1] {
		if pipe == pipe_XY {
			return exists && true
		}
	}
	return false
}

func isValidPath(pipe_xy, pipe_XY rune, relativeCoordinate []int) bool {
	c_x, c_y := relativeCoordinate[0], relativeCoordinate[1]
	validRight := [][]rune{
		{'-', 'L', 'F'},
		{'-', 'J', '7'},
	}
	validLeft := [][]rune{
		{'-', 'J', '7'},
		{'-', 'L', 'F'},
	}
	validUp := [][]rune{
		{'|', 'L', 'J'},
		{'|', '7', 'F'},
	}
	validDown := [][]rune{
		{'|', '7', 'F'},
		{'|', 'L', 'J'},
	}
	if c_x == 0 {
		if c_y == 1 {
			return doesExist(pipe_xy, pipe_XY, validDown)
		} else {
			return doesExist(pipe_xy, pipe_XY, validUp)
		}
	} else if c_x == 1 {
		return doesExist(pipe_xy, pipe_XY, validRight)
	}
	return doesExist(pipe_xy, pipe_XY, validLeft)
}

func getSteps(file *os.File) int {
	landscape, S_x, S_y := getLandscape(file)
	relativeCoordinates := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	pipes := []rune{'|', '-', 'L', 'J', 'F', '7'}
	placeValue := max(len(landscape), len(landscape[0]))
	steps := 0
	for _, pipe := range pipes {
		queue := [][]int{{S_x, S_y, 0}}
		visited := make(map[int]int)
		subSteps := 0
		cycle := false
		for len(queue) > 0 {
			x, y, distance := queue[0][0], queue[0][1], queue[0][2]
			queue = queue[1:]
			key := x*placeValue + y
			distanceValue, ok := visited[key]
			pipe_xy := rune(landscape[y][x])
			if pipe_xy == 'S' {
				pipe_xy = pipe
			}
			if ok {
				if distanceValue-distance == 2 || distanceValue-distance == -2 {
					continue
				} else {
					cycle = true
					break
				}
			}
			visited[key] = distance
			subSteps = max(subSteps, distance)
			for _, relativeCoordinate := range relativeCoordinates {
				X := x + relativeCoordinate[0]
				Y := y + relativeCoordinate[1]
				if Y >= 0 && Y < len(landscape) && X >= 0 && X < len(landscape[0]) {
					pipe_XY := rune(landscape[Y][X])
					if isValidPath(pipe_xy, pipe_XY, relativeCoordinate) {
						queue = append(queue, []int{X, Y, distance + 1})
					}
				}
			}
		}
		if cycle {
			steps = max(steps, subSteps)
		}
	}
	return steps
}

func getLandscape(file *os.File) ([]string, int, int) {
	line := ""
	landscape := make([]string, 0)
	i, j := 0, 0
	x, y := -1, -1
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		if !isWhitespace(buffer[0]) {
			if buffer[0] == 'S' {
				x, y = i, j
			}
			line += string(buffer[0])
			i++
		}
		if buffer[0] == 10 || err == io.EOF {
			landscape = append(landscape, line)
			line = ""
			i = 0
			j++
		}
	}
	return landscape, x, y
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 1:", getSteps(file)) // 6907
}
