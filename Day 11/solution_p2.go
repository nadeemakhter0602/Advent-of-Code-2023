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

func isWhitespace(a byte) bool {
	return a < 33
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func getEmptyRows(image []string) []int {
	emptyRows := make([]int, 0)
	y := 0
	for y < len(image) {
		containsGalaxy := false
		x := 0
		for x < len(image[y]) {
			containsGalaxy = containsGalaxy || image[y][x] == '#'
			x++
		}
		if !containsGalaxy {
			emptyRows = append(emptyRows, y)
		}
		y++
	}
	return emptyRows
}

func getEmptyColumns(image []string) []int {
	emptyColumns := make([]int, 0)
	x := 0
	for x < len(image[0]) {
		containsGalaxy := false
		y := 0
		for y < len(image) {
			containsGalaxy = containsGalaxy || image[y][x] == '#'
			y++
		}
		if !containsGalaxy {
			emptyColumns = append(emptyColumns, x)
		}
		x++
	}
	return emptyColumns
}

func getGalaxyCoordinates(image []string) [][]int {
	galaxies := make([][]int, 0)
	y := 0
	for y < len(image) {
		x := 0
		for x < len(image[y]) {
			if image[y][x] == '#' {
				galaxies = append(galaxies, []int{x, y})
			}
			x++
		}
		y++
	}
	return galaxies
}

func getPairShortestPathSum(file *os.File) int {
	image := getImage(file)
	emptyColumns := getEmptyColumns(image)
	emptyRows := getEmptyRows(image)
	galaxies := getGalaxyCoordinates(image)
	expansion := 1000000
	pairShortestPathSum := 0
	i, j := 0, 0
	for i < len(galaxies) {
		j = i + 1
		for j < len(galaxies) {
			i_x, i_y := galaxies[i][0], galaxies[i][1]
			j_x, j_y := galaxies[j][0], galaxies[j][1]
			numEmptyRows, numEmptyColumns := 0, 0
			for _, emptyRow := range emptyRows {
				if i_y < j_y && i_y < emptyRow && j_y > emptyRow {
					numEmptyRows++
				} else if i_y > j_y && j_y < emptyRow && i_y > emptyRow {
					numEmptyRows++
				}
			}
			for _, emptyColumn := range emptyColumns {
				if i_x < j_x && i_x < emptyColumn && j_x > emptyColumn {
					numEmptyColumns++
				} else if i_x > j_x && j_x < emptyColumn && i_x > emptyColumn {
					numEmptyColumns++
				}
			}
			manhattanDistance := abs(i_x-j_x) + abs(i_y-j_y)
			expandedDistance := manhattanDistance + (expansion-1)*(numEmptyRows+numEmptyColumns)
			pairShortestPathSum += expandedDistance
			j++
		}
		i++
	}
	return pairShortestPathSum
}

func getImage(file *os.File) []string {
	line := ""
	image := make([]string, 0)
	err := error(nil)
	for err != io.EOF {
		buffer := make([]byte, 1)
		_, err = file.Read(buffer)
		PanicErr(err)
		if !isWhitespace(buffer[0]) {
			line += string(buffer[0])
		}
		if buffer[0] == 10 || err == io.EOF {
			image = append(image, line)
			line = ""
		}
	}
	return image
}

func main() {
	file, err := os.Open("input.txt")
	PanicErr(err)
	defer file.Close()
	fmt.Println("PART 2:", getPairShortestPathSum(file)) // 685038186836
}
