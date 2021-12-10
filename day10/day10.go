package main

import (
	"bufio"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	scores  = map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	scores2 = map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	open    = map[string]bool{
		"(": true,
		"<": true,
		"[": true,
		"{": true,
	}
	opp = map[string]string{
		"(": ")",
		"<": ">",
		"[": "]",
		"{": "}",
		")": "(",
		">": "<",
		"]": "[",
		"}": "{",
	}
)

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	res1, godlines := partOne(data)
	res2 := partTwo(godlines)
	log.Info().Int("Part 1", res1).Int("Part 2", res2).Msg("program complete")
}

func partOne(data []string) (int, []string) {
	var score int
	var godLines []string
	for _, l := range data {
		chars := strings.Split(l, "")
		var stack []string
		var valid bool
		for _, c := range chars {
			valid = true
			if open[c] {
				stack = append(stack, c)
				continue
			}
			tail := stack[len(stack)-1]
			if tail != opp[c] {
				score += scores[c]
				valid = false
				break
			}
			stack = stack[:len(stack)-1]
		}
		if valid {
			godLines = append(godLines, l)
		}

	}
	return score, godLines
}

func partTwo(data []string) int {
	var scores []int
	for _, l := range data {
		var score int
		var stack []string
		chars := strings.Split(l, "")
		for _, c := range chars {
			if open[c] {
				stack = append(stack, c)
				continue
			}

			stack = stack[:len(stack)-1]

		}
		for i := range stack {
			s := stack[len(stack)-1-i]
			score *= 5
			score += scores2[opp[s]]
		}
		scores = append(scores, score)

	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func readInput() ([]string, error) {
	var data []string
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}
