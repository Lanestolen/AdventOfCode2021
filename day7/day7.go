package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	minp int
	maxp int
)

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	log.Debug().Int("part 1", partOne(data)).Int("part 2", partTwo(data)).Msg("program complete")
}

func partOne(data []int) int {
	fuel := make(map[int]int)
	minfuel := 999999999999999999
	for i := minp; i <= maxp; i++ {
		for _, v := range data {
			fuel[i] += max(i, v) - min(i, v)
		}
	}
	for _, v := range fuel {
		minfuel = min(minfuel, v)
	}

	return minfuel
}

func partTwo(data []int) int {
	fuel := make(map[int]int)
	minfuel := 999999999999999999
	for i := minp; i <= maxp; i++ {
		for _, v := range data {
			fuel[i] += calcfuel(max(i, v) - min(i, v))
		}
	}
	for _, v := range fuel {
		minfuel = min(minfuel, v)
	}

	return minfuel
}

func calcfuel(in int) int {
	return in * (in + 1) / 2
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func readInput() ([]int, error) {
	var data []int
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		for _, i := range input {
			point, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			data = append(data, point)
			minp = min(minp, point)
			maxp = max(maxp, point)
		}
	}
	return data, nil
}
