package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const boardSize = 5
const dataPath = "day_4/data.txt"

type Case struct {
	Number int64
	Marked bool
}

type Board [boardSize][boardSize]Case
type Boards []Board

func (b *Board) Mark(n int64) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j].Number == n {
				b[i][j].Marked = true
			}
		}
	}
}

func (b *Board) hasWonByRow() bool {
	for i := 0; i < boardSize; i++ {
		full := true
		for j := 0; j < boardSize; j++ {
			if !b[i][j].Marked {
				full = false
			}
		}
		if full {
			return true
		}
	}
	return false
}

func (b *Board) hasWonByCol() bool {
	for i := 0; i < boardSize; i++ {
		full := true
		for j := 0; j < boardSize; j++ {
			if !b[j][i].Marked {
				full = false
			}
		}
		if full {
			return true
		}
	}
	return false
}

func (b *Board) HasWon() bool {
	return b.hasWonByRow() || b.hasWonByCol()
}

func (b *Board) UnmarkedAmount() int64 {
	var s int64
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b[i][j].Marked {
				s += b[i][j].Number
			}
		}
	}
	return s
}

func (b *Board) Score(lastNumber int64) int64 {
	return lastNumber * b.UnmarkedAmount()
}

func (b Boards) Mark(n int64) Boards {
	for i := 0; i < len(b); i++ {
		b[i].Mark(n)
	}
	return b
}

// HasAWinner returns wether a board has won, and if yes, the id of the winner board.
func (b Boards) HasAWinner() (bool, []int) {
	var winners []int
	for i, board := range b {
		if board.HasWon() {
			winners = append(winners, i)
		}
	}
	return len(winners) > 0, winners
}

func readDraws(input string) ([]int64, error) {
	var err error
	inputs := strings.Split(input, ",")
	draws := make([]int64, len(inputs))
	for i, s := range inputs {
		draws[i], err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return draws, nil
}

func readBoards(input string) (Boards, error) {
	var err error
	boardStr := strings.Split(input, "\r\n\r\n")
	boards := make([]Board, len(boardStr))
	for i, s := range boardStr {
		lines := strings.Split(s, "\r\n")
		lines = dropStrings(lines, "")
		for i2, line := range lines {
			numbers := strings.Split(line, " ")
			numbers = dropStrings(numbers, "")
			if len(numbers) != boardSize {
				return nil, fmt.Errorf("wrong length input: %s", line)
			}
			for i3, number := range numbers {
				boards[i][i2][i3].Number, err = strconv.ParseInt(number, 10, 64)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return boards, err
}

func dropStrings(stringSlice []string, toDrop string) []string {
	var result []string
	for _, s := range stringSlice {
		if s != toDrop {
			result = append(result, s)
		}
	}
	return result
}

func readData(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.SplitN(string(content), "\r\n", 2), nil
}

// main for the first part of the puzzle
//func main() {
//	input, err := readData(dataPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	drawn, err := readDraws(input[0])
//	if err != nil {
//		log.Fatal(err)
//	}
//	boards, err := readBoards(input[1])
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for i, x := range drawn {
//		boards = boards.Mark(x)
//		won, winnersId := boards.HasAWinner()
//		if won {
//			for _, winnerId := range winnersId {
//				fmt.Printf("Board %d has won during round %d with a score of %d.\n", winnerId, i, boards[winnerId].Score(x))
//			}
//			break
//		}
//	}
//}

func DropBoards(boards Boards, ids []int) Boards {
	var newBoards Boards
	for i, board := range boards {
		toDrop := false
		for _, id := range ids {
			if id == i {
				toDrop = true
			}
		}
		if !toDrop {
			newBoards = append(newBoards, board)
		}
	}
	return newBoards
}

// main for the second part of the puzzle
func main() {
	input, err := readData(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	drawn, err := readDraws(input[0])
	if err != nil {
		log.Fatal(err)
	}
	boards, err := readBoards(input[1])
	if err != nil {
		log.Fatal(err)
	}

	for i, x := range drawn {
		boards = boards.Mark(x)
		won, winnersId := boards.HasAWinner()
		if won {
			for _, winnerId := range winnersId {
				fmt.Printf("Board %d has won during round %d with a score of %d.\n", winnerId, i, boards[winnerId].Score(x))
			}
			boards = DropBoards(boards, winnersId)
		}
	}
}
