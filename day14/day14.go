package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	data, lettermap, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	fmt.Println(data, lettermap)
	partOne(data, lettermap, 10)
}

func partOne(data map[string]int, lettermap map[string][]string, days int) {
	finalL := data
	for day := 0; day < days; day++ {
		fmt.Println(day)
		next := make(map[string]int)
		for st, i := range finalL {
			val := lettermap[st]
			for j := 0; j < i; j++ {
				for _, str := range val {
					next[str]++
				}
			}
		}
		finalL = next
	}
	lettercount := make(map[string]int)
	for i, amount := range finalL {
		for j := 0; j < amount; j++ {
			letters := strings.Split(i, "")
			lettercount[letters[0]]++
			lettercount[letters[1]]++
		}
	}
	var countArr []int
	for letter, i := range lettercount {
		lettercount[letter] = i/2 + i%2
		countArr = append(countArr, i/2+i%2)
	}
	sort.Ints(countArr)
	fmt.Println(countArr[len(countArr)-1] - countArr[0])
}

func readInput() (map[string]int, map[string][]string, error) {
	letterMap := make(map[string][]string)
	data := make(map[string]int)
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			if !strings.Contains(scanner.Text(), "->") {
				text := scanner.Text()
				for i := 0; i < len(text)-1; i++ {
					data[text[i:i+2]]++
				}
				continue
			}
			input := strings.Split(scanner.Text(), "->")
			ol := strings.Trim(input[1], " ")
			il := strings.Trim(input[0], " ")
			letters := strings.Split(il, "")

			letterMap[il] = []string{letters[0] + ol, ol + letters[1]}
		}
	}
	return data, letterMap, nil
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
