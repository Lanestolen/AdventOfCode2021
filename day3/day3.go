package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func readInput() ([][]int, error) {
	var data [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		var point []int
		for _, v := range input {
			amount, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			point = append(point, amount)
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

func partOne(data [][]int) int {
	one := make(map[int]int)
	zero := make(map[int]int)
	out1 := make(map[int]int)
	out2 := make(map[int]int)
	for _, v := range data {
		for i, j := range v {
			if j == 1 {
				one[i]++
				continue
			}
			zero[i]++
		}
	}

	for x := range one {
		if one[x] > zero[x] {
			out1[x] = 1
			out2[x] = 0
			continue
		}
		out1[x] = 0
		out2[x] = 1
	}

	return map2int(out1) * map2int(out2)
}

func map2int(in map[int]int) int {
	bstr := ""
	for i := 0; i <= 11; i++ {
		bstr = bstr + fmt.Sprint(in[i])
	}

	out, _ := strconv.ParseInt(bstr, 2, 64)
	return int(out)
}

func arr2int(in []int) int {
	bstr := ""
	for i := 0; i <= 11; i++ {
		bstr = bstr + fmt.Sprint(in[i])
	}

	out, _ := strconv.ParseInt(bstr, 2, 64)
	return int(out)
}

func partTwo(data [][]int) int {
	cpdata := data
	for i := 0; i <= 11; i++ {
		var newdata [][]int
		sig := findSig(data, i)
		for _, v := range data {
			if v[i] == sig {
				newdata = append(newdata, v)
			}
		}
		data = newdata
	}

	for i := 0; i <= 11; i++ {
		var newdata [][]int
		if len(cpdata) == 1 {
			break
		}
		least := findLeast(cpdata, i)
		for _, v := range cpdata {
			if v[i] == least {
				newdata = append(newdata, v)
			}
		}
		cpdata = newdata
	}

	return arr2int(data[0]) * arr2int(cpdata[0])
}

func findSig(data [][]int, slot int) int {
	var ones int
	var zeroes int
	for _, v := range data {
		if v[slot] == 1 {
			ones++
			continue
		}
		zeroes++
	}
	if ones >= zeroes {
		return 1
	}
	return 0
}

func findLeast(data [][]int, slot int) int {
	var ones int
	var zeroes int
	for _, v := range data {
		if v[slot] == 1 {
			ones++
			continue
		}
		zeroes++
	}

	if ones < zeroes {
		return 1
	}
	return 0
}
