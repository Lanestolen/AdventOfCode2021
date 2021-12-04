package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var (
	input []int
)

type board struct {
	won  bool
	rows []map[int]bool
	col  []map[int]bool
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

func partOne(boards []board) int {
	for _, n := range input {
		for _, board := range boards {
			for _, r := range board.rows {
				if _, ok := r[n]; ok {
					r[n] = true
				}
				continue
			}
			for _, c := range board.col {
				if _, ok := c[n]; ok {
					c[n] = true
				}
				continue
			}
			if checkBoard(board) {
				sum := sumUnmarked(board)
				return n * sum
			}
		}
	}
	return 0
}

func partTwo(boards []board) int {
	boards = boards
	for _, n := range input {

		var newboards []board
		for _, board := range boards {
			for _, r := range board.rows {
				if _, ok := r[n]; ok {
					r[n] = true
				}
				continue
			}
			for _, c := range board.col {
				if _, ok := c[n]; ok {
					c[n] = true
				}
				continue
			}
			if !checkBoard(board) {
				newboards = append(newboards, board)
			}
			if len(boards) == 1 && checkBoard(board) {
				sum := sumUnmarked(boards[0])
				return sum * n
			}
		}
		boards = newboards

	}
	return 0
}

func sumUnmarked(in board) int {
	var sum int

	for _, r := range in.rows {
		for v, m := range r {
			if !m {
				sum = sum + v
			}
		}
	}

	return sum
}

func checkBoard(in board) bool {

	for _, col := range in.col {
		var wcount int
		for _, val := range col {
			if val {
				wcount++
			}
		}
		if wcount == 5 {
			return true
		}
	}

	for _, row := range in.rows {
		var wcount int
		for _, val := range row {
			if val {
				wcount++
			}
		}
		if wcount == 5 {
			return true
		}
	}

	return false
}

func readInput() ([]board, error) {
	var rows [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputRow := true
		if inputRow {
			in := strings.Split(scanner.Text(), ",")
			for _, v := range in {
				amount, err := strconv.Atoi(v)
				if err != nil {
					continue
				}
				input = append(input, amount)
				inputRow = false
			}

		}

		if scanner.Text() == "" {

			continue
		}

		var row []int
		in := strings.Split(scanner.Text(), " ")
		for _, v := range in {
			amount, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			row = append(row, amount)
		}
		if len(row) == 0 {
			continue
		}
		rows = append(rows, row)

	}

	return boardsFromArray(rows), nil
}

func boardsFromArray(input [][]int) []board {
	var boards []board
	for i := 0; i < len(input); i += 5 {
		batch := input[i:min(i+5, len(input))]
		var board board
		for j, v := range batch {
			rowMap := make(map[int]bool)
			colMap := make(map[int]bool)
			for _, j := range v {
				rowMap[j] = false
			}
			for i := 0; i < 5; i++ {
				colMap[batch[i][j]] = false
			}
			board.rows = append(board.rows, rowMap)
			board.col = append(board.col, colMap)
		}

		boards = append(boards, board)

	}

	return boards
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
