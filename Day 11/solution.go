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

func expandImage(image []string) []string {
	expandedImage := make([]string, 0)
	y := 0
	for y < len(image) {
		containsGalaxy := false
		x := 0
		for x < len(image[y]) {
			containsGalaxy = containsGalaxy || image[y][x] == '#'
			x++
		}
		if !containsGalaxy {
			expandedImage = append(expandedImage, image[y], image[y])
		} else {
			expandedImage = append(expandedImage, image[y])
		}
		y++
	}
	finalExpandedImage := make([]string, 0)
	y = 0
	for y < len(expandedImage) {
		finalExpandedImage = append(finalExpandedImage, "")
		y++
	}
	x := 0
	for x < len(expandedImage[0]) {
		containsGalaxy := false
		y = 0
		for y < len(expandedImage) {
			containsGalaxy = containsGalaxy || expandedImage[y][x] == '#'
			y++
		}
		if !containsGalaxy {
			j := 0
			for j < len(expandedImage) {
				finalExpandedImage[j] += string(expandedImage[j][x]) + string(expandedImage[j][x])
				j++
			}
		} else {
			j := 0
			for j < len(expandedImage) {
				finalExpandedImage[j] += string(expandedImage[j][x])
				j++
			}
		}
		x++
	}
	return finalExpandedImage
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
	image = expandImage(image)
	galaxies := getGalaxyCoordinates(image)
	pairShortestPathSum := 0
	i, j := 0, 0
	for i < len(galaxies) {
		j = i + 1
		for j < len(galaxies) {
			i_x, i_y := galaxies[i][0], galaxies[i][1]
			j_x, j_y := galaxies[j][0], galaxies[j][1]
			pairShortestPathSum += abs(i_x-j_x) + abs(i_y-j_y)
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
	fmt.Println("PART 1:", getPairShortestPathSum(file)) // 9556896
}
