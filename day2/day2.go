package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type movement struct {
	direction string
	amount    int
}

func readInput() ([]movement, error) {
	var data []movement
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		var point movement
		amount, err := strconv.Atoi(input[1])
		if err != nil {
			return nil, err
		}
		point.direction = input[0]
		point.amount = amount

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

func partOne(data []movement) int {
	var hor, depth int

	for _, k := range data {
		if k.direction == "forward" {
			hor = hor + k.amount
			continue
		}
		if k.direction == "down" {
			depth = depth + k.amount
			continue
		}
		if k.direction == "up" {
			depth = depth - k.amount
			continue
		}

	}
	return hor * depth
}

func partTwo(data []movement) int {
	var hor, depth, aim int

	for _, k := range data {
		if k.direction == "forward" {
			hor = hor + k.amount
			depth = depth + k.amount*aim
			continue
		}
		if k.direction == "down" {
			aim = aim + k.amount
			continue
		}
		if k.direction == "up" {
			aim = aim - k.amount
			continue
		}

	}
	return hor * depth
}
