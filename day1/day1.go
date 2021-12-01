package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

func readInput() ([]int, error) {
	var data []int
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		point, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, point)
	}
	return data, nil
}

func main() {
	data, err := readInput()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load input")
	}
	res1 := partOne(data)
	res2 := partTwo(data)
	log.Info().Int("part 1", res1).Int("part 2", res2).Msg("completed the program")
}

func partOne(data []int) int {
	var result int
	for i, k := range data {
		if i == 0 {
			continue
		}
		prev := data[i-1]
		if k > prev {
			result++
		}
	}
	return result
}

func partTwo(data []int) int {
	var result int

	for i, k := range data {
		if i < 3 {
			continue
		}
		prev := data[i-3] + data[i-2] + data[i-1]
		curr := data[i-2] + data[i-1] + k
		if curr > prev {
			result++
		}
	}

	return result
}
