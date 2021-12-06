package main

import (
	"bufio"
	"fmt"

	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	fish []int
)

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}
	fmt.Println(partTwo(data))
	fish = data
	fmt.Println(partOne())

}

func partOne() int {
	for day := 0; day < 80; day++ {

		for i := range fish {

			if fish[i] == 0 {
				fish[i] = 7
				fish = append(fish, 8)
			}
			fish[i]--
		}

	}

	return len(fish)
}

func partTwo(data []int) int {
	var total int
	for i, f := range data {
		log.Debug().Int("fish", i).Msg("started fish")
		fi := simfish(f, 256)
		total = total + fi
		log.Debug().Int("done", i).Int("total", len(data)).Msg("finished a fish")
	}

	return total
}

func simfish(f, days int) int {
	var simfish []int
	simfish = append(simfish, f)
	for day := 0; day < days; day++ {
		log.Debug().Int("day", day).Msg("started day")
		for i := range simfish {
			if simfish[i] == 0 {
				simfish[i] = 7
				simfish = append(simfish, 8)
			}
			simfish[i]--

		}
	}
	return len(simfish)
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
		}
	}
	return data, nil
}
