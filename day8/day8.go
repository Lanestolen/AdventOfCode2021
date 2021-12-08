package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	data, setup, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	res1 := partOne(data)
	res2 := partTwo(data, setup)

	log.Info().Int("part 1", res1).Int("part 2", res2).Msg("program complete")
}

func partOne(data [][]string) int {
	var counter int
	for _, v := range data {
		for _, d := range v {
			if len(d) == 2 || len(d) == 4 || len(d) == 7 || len(d) == 3 {
				counter++
			}
		}
	}
	return counter
}

func partTwo(data [][]string, setup [][]string) int {
	var count int
	for i, v := range data {
		maps := findNumbers(setup[i])
		var nrs []int
		for _, k := range v {
			if k == "" {
				continue
			}

			for t, q := range maps {
				if len(t) == len(k) {
					if checkifContains(k, t) {
						nrs = append(nrs, q)
					}
				}
			}

		}
		count += sliceToInt(nrs)
	}
	return count
}

func findNumbers(setup []string) map[string]int {
	nrmap := make(map[string]int)
	otherway := make(map[int]string)
	for _, v := range setup {
		if len(v) == 2 {
			nrmap[v] = 1
			otherway[1] = v
		}
		if len(v) == 3 {
			nrmap[v] = 7
			otherway[7] = v
		}
		if len(v) == 4 {
			nrmap[v] = 4
			otherway[4] = v
		}
		if len(v) == 7 {
			nrmap[v] = 8
			otherway[8] = v
		}
	}

	for _, v := range setup {
		for i, n := range nrmap {
			if checkifContains(v, i) {
				if len(v) == 5 && n == 1 {
					nrmap[v] = 3
					otherway[3] = v
				}
				if len(v) == 6 && n == 4 {
					nrmap[v] = 9
					otherway[9] = v
					continue
				}
			}
			if len(v) == 6 && n == 4 {
				if checkifContains(v, otherway[1]) {
					nrmap[v] = 0
					otherway[0] = v
				}
			}

		}
	}

	for _, v := range setup {
		if len(v) == 6 && !checkifContains(v, otherway[9]) && !checkifContains(v, otherway[0]) {
			nrmap[v] = 6
			otherway[6] = v
		}
	}
	for _, v := range setup {
		if len(v) == 5 && checkifContains(otherway[6], v) {
			nrmap[v] = 5
			otherway[5] = v
		}
		if v != "" {
			if _, ok := nrmap[v]; !ok {
				nrmap[v] = 2
			}
		}
	}

	return nrmap
}

func checkifContains(in, contains string) bool {
	var amount int
	for _, l := range contains {
		if strings.Contains(in, string(l)) {
			amount++
		}
	}
	if amount == len(contains) {
		return true
	}
	return false
}

func readInput() ([][]string, [][]string, error) {
	var data [][]string
	var setup [][]string
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "|")
		digits := strings.Split(input[1], " ")
		sdigits := strings.Split(input[0], " ")
		data = append(data, digits)
		setup = append(setup, sdigits)
	}
	return data, setup, nil
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
