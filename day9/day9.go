package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type point struct {
	x int
	y int
}

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	risk, lows := partOne(data)
	res2 := partTwo(data, lows)

	log.Info().Int("part 1", risk).Int("part2", res2).Msg("program complete")
}

func partOne(data [][]int) (int, []point) {
	var points []point
	var risk int
	for q, r := range data {
		for i, p := range r {
			var up, down, left, right int
			if q == 0 {
				up = 10
			} else {
				up = data[q-1][i]
			}
			if q == len(data)-1 {
				down = 10
			} else {
				down = data[q+1][i]
			}
			if i == 0 {
				left = 10
			} else {
				left = r[i-1]
			}
			if i == len(r)-1 {
				right = 10
			} else {
				right = r[i+1]
			}
			if p < up && p < down && p < left && p < right {
				risk += (p + 1)
				points = append(points, point{x: q, y: i})
			}
		}
	}
	return risk, points
}

func partTwo(data [][]int, lows []point) int {
	var pools []int
	for _, p := range lows {
		larger := checkPoints(data, p)
		larger = append(larger, p)
		slarger := sortPoints(larger)
		pools = append(pools, len(slarger))
	}
	sort.Ints(pools)

	larg1 := pools[len(pools)-1]
	larg2 := pools[len(pools)-2]
	larg3 := pools[len(pools)-3]
	return larg1 * larg2 * larg3
}

func sortPoints(p []point) []point {
	cheat := make(map[point]int)
	var out []point
	for _, po := range p {
		cheat[po]++
	}

	for i := range cheat {
		out = append(out, i)
	}
	return out
}

func checkPoints(data [][]int, p point) []point {
	var l []point
	if p.x != 0 && data[p.x-1][p.y] > data[p.x][p.y] && data[p.x-1][p.y] != 9 {
		l = append(l, point{x: p.x - 1, y: p.y})
		out := checkPoints(data, point{x: p.x - 1, y: p.y})
		l = append(l, out...)
	}
	if p.y != 0 && data[p.x][p.y-1] > data[p.x][p.y] && data[p.x][p.y-1] != 9 {
		l = append(l, point{x: p.x, y: p.y - 1})
		out := checkPoints(data, point{x: p.x, y: p.y - 1})
		l = append(l, out...)
	}
	if p.x != len(data)-1 && data[p.x+1][p.y] > data[p.x][p.y] && data[p.x+1][p.y] != 9 {
		l = append(l, point{x: p.x + 1, y: p.y})
		out := checkPoints(data, point{x: p.x + 1, y: p.y})
		l = append(l, out...)
	}
	if p.y != len(data[0])-1 && data[p.x][p.y+1] > data[p.x][p.y] && data[p.x][p.y+1] != 9 {
		l = append(l, point{x: p.x, y: p.y + 1})
		out := checkPoints(data, point{x: p.x, y: p.y + 1})
		l = append(l, out...)
	}
	return l
}

func readInput() ([][]int, error) {
	var data [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []int
		input := strings.Split(scanner.Text(), "")
		for _, i := range input {
			point, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			row = append(row, point)
		}
		data = append(data, row)
	}
	return data, nil
}
