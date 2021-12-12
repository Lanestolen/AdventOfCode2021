package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/rs/zerolog/log"
)

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	res1 := partOne(data)
	res2 := partTwo(data)
	log.Info().Int("part 1", res1).Int("part2", res2).Msg("program complete")
}

func partOne(data map[string][]string) int {
	visits := make(map[string]int)

	return findPaths("start", data, visits)
}

func partTwo(data map[string][]string) int {
	visits := make(map[string]int)

	return sightseeing("start", data, visits, true)
}

func sightseeing(cave string, data map[string][]string, visits map[string]int, revisit bool) int {
	var paths int

	if cave == "end" {
		return 1
	}
	visits[cave]++

	for _, cons := range data[cave] {
		if isSmall(cons) {
			if visits[cons] == 0 {
				paths += sightseeing(cons, data, visits, revisit)
			}
			if visits[cons] == 1 {
				if revisit && cons != "start" {
					paths += sightseeing(cons, data, visits, false)
				}
				continue
			}
			continue
		}
		paths += sightseeing(cons, data, visits, revisit)
	}
	visits[cave]--

	return paths
}

func findPaths(cave string, data map[string][]string, visits map[string]int) int {
	var paths int

	if cave == "end" {
		return 1
	}
	visits[cave]++

	for _, cons := range data[cave] {
		if isSmall(cons) {
			if visits[cons] == 0 {
				paths += findPaths(cons, data, visits)
			}
			continue
		}
		paths += findPaths(cons, data, visits)
	}
	visits[cave]--

	return paths
}

func isSmall(cave string) bool {
	return unicode.IsLower(rune(cave[0]))
}

func readInput() (map[string][]string, error) {
	data := make(map[string][]string)
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "-")
		data[input[0]] = append(data[input[0]], input[1])
		data[input[1]] = append(data[input[1]], input[0])
	}
	return data, nil
}
