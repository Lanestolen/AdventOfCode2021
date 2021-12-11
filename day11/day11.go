package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type octupus struct {
	Energy     int
	HasFlashed bool
}

func main() {
	data, err := readInput()
	if err != nil {
		log.Error().Err(err).Msg("failed to load data")
	}

	res1 := partOne(data)
	res2 := partTwo(data)

	log.Info().Int("part 1", res1).Int("part2", res2).Msg("program complete")
}

func partTwo(data [][]octupus) int {
	var allFlashed bool
	var day int
	for !allFlashed {
		var flashesday int
		for i, r := range data {
			for j := range r {
				data[i][j].Energy++
			}
		}
		checkflashes(data)

		for i, r := range data {
			for j, o := range r {
				if o.HasFlashed {
					flashesday++
					data[i][j].Energy = 0
					data[i][j].HasFlashed = false
				}
			}
		}
		if flashesday == 100 {
			allFlashed = true
		}
		day++
		log.Debug().Int("flashes", flashesday).Int("day", day+100).Msg("flashes on day")
	}
	return day + 100
}

func partOne(data [][]octupus) int {
	var totalFlashes int
	for day := 0; day < 100; day++ {
		var flashesday int
		for i, r := range data {
			for j := range r {
				data[i][j].Energy++
			}
		}
		checkflashes(data)

		for i, r := range data {
			for j, o := range r {
				if o.HasFlashed {
					flashesday++
					data[i][j].Energy = 0
					data[i][j].HasFlashed = false
				}
			}
		}
		log.Debug().Int("flashes", flashesday).Int("day", day).Msg("flashes on day")
		totalFlashes += flashesday
	}
	return totalFlashes
}

func checkflashes(data [][]octupus) {
	for i, r := range data {
		for j, o := range r {
			if o.Energy > 9 {
				if o.HasFlashed {
					continue
				}
				data[i][j].HasFlashed = true
				handleFlash(data, i, j)
			}
		}
	}
}

func handleFlash(data [][]octupus, x, y int) {
	data[x][y].HasFlashed = true
	//handle topleft
	if x != 0 && y != 0 {
		data[x-1][y-1].Energy++
		if data[x-1][y-1].Energy > 9 && !data[x-1][y-1].HasFlashed {
			handleFlash(data, x-1, y-1)
		}
	}
	// Top middle
	if y != 0 {
		data[x][y-1].Energy++
		if data[x][y-1].Energy > 9 && !data[x][y-1].HasFlashed {
			handleFlash(data, x, y-1)
		}
	}
	// Top right
	if x != len(data)-1 && y != 0 {
		data[x+1][y-1].Energy++
		if data[x+1][y-1].Energy > 9 && !data[x+1][y-1].HasFlashed {
			handleFlash(data, x+1, y-1)
		}
	}
	// Middle left
	if x != 0 {
		data[x-1][y].Energy++
		if data[x-1][y].Energy > 9 && !data[x-1][y].HasFlashed {
			handleFlash(data, x-1, y)
		}
	}
	// Middle Right
	if x != len(data)-1 {
		data[x+1][y].Energy++
		if data[x+1][y].Energy > 9 && !data[x+1][y].HasFlashed {
			handleFlash(data, x+1, y)
		}
	}
	// bottom left
	if x != 0 && y != len(data[x])-1 {
		data[x-1][y+1].Energy++
		if data[x-1][y+1].Energy > 9 && !data[x-1][y+1].HasFlashed {
			handleFlash(data, x-1, y+1)
		}
	}
	// bottom middle
	if y != len(data[x])-1 {
		data[x][y+1].Energy++
		if data[x][y+1].Energy > 9 && !data[x][y+1].HasFlashed {
			handleFlash(data, x, y+1)
		}
	}
	// bottom Right
	if x != len(data)-1 && y != len(data[x])-1 {
		data[x+1][y+1].Energy++
		if data[x+1][y+1].Energy > 9 && !data[x+1][y+1].HasFlashed {
			handleFlash(data, x+1, y+1)
		}
	}
}

func readInput() ([][]octupus, error) {
	var data [][]octupus
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var row []octupus
		input := strings.Split(scanner.Text(), "")
		for _, i := range input {
			point, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			row = append(row, octupus{Energy: point})
		}
		data = append(data, row)
	}
	return data, nil
}
