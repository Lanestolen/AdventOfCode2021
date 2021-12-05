package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
)

var (
	cordregex = regexp.MustCompile(`([\d]{1,3}),([\d]{1,3}) -> ([\d]{1,3}),([\d]{1,3})`)
	grid      [1000][1000]int
)

type vector struct {
	startx int
	starty int
	endx   int
	endy   int
}

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to read input")
	}
	res1 := partOne(data)
	res2 := partTwo(data)
	log.Info().Int("part 1", res1).Int("part 2", res2).Msg("program complete")
}

func partOne(data []vector) int {
	for _, v := range data {
		if v.startx == v.endx {
			addToGrid(v)
		}
		if v.starty == v.endy {
			addToGrid(v)
		}
	}

	return count()
}

func partTwo(data []vector) int {
	for _, v := range data {
		deltax := max(v.startx, v.endx) - min(v.startx, v.endx)
		deltay := max(v.starty, v.endy) - min(v.starty, v.endy)

		if deltax == deltay {
			addToGridP2(v, deltax)
		}
	}

	return count()
}

func addToGridP2(vec vector, delta int) {
	var xpoints []int
	var ypoints []int
	minx := min(vec.startx, vec.endx)
	if vec.startx == minx {
		for i := minx; i <= minx+delta; i++ {
			xpoints = append(xpoints, i)
		}
	} else {
		for i := minx + delta; i >= minx; i-- {
			xpoints = append(xpoints, i)
		}
	}

	miny := min(vec.starty, vec.endy)
	if vec.starty == miny {
		for i := miny; i <= miny+delta; i++ {
			ypoints = append(ypoints, i)
		}
	} else {
		for i := miny + delta; i >= miny; i-- {
			ypoints = append(ypoints, i)
		}
	}
	for i := range xpoints {

		grid[xpoints[i]][ypoints[i]]++
	}

}

func count() int {
	var count int
	for _, x := range grid {
		for _, y := range x {
			if y >= 2 {
				count++
			}
		}
	}
	return count
}

func addToGrid(vec vector) {
	if vec.startx == vec.endx {
		min := min(vec.starty, vec.endy)
		max := max(vec.starty, vec.endy)

		for i := min; i <= max; i++ {
			grid[vec.startx][i]++
		}
	}
	if vec.starty == vec.endy {
		min := min(vec.startx, vec.endx)
		max := max(vec.startx, vec.endx)
		for i := min; i <= max; i++ {
			grid[i][vec.starty]++
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func readInput() ([]vector, error) {
	var data []vector
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := cordregex.FindSubmatch(scanner.Bytes())
		startx, err := strconv.Atoi(string(matches[1]))
		if err != nil {
			return nil, err
		}
		starty, err := strconv.Atoi(string(matches[2]))
		if err != nil {
			return nil, err
		}
		endx, err := strconv.Atoi(string(matches[3]))
		if err != nil {
			return nil, err
		}
		endy, err := strconv.Atoi(string(matches[4]))
		if err != nil {
			return nil, err
		}
		vec := vector{startx: startx, starty: starty, endx: endx, endy: endy}

		data = append(data, vec)
	}
	return data, nil
}
