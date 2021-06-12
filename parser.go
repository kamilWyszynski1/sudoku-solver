package sudoku_solver

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type SudokuParser interface {
	// Parse gets 81-digit string of sudoku arrangement, validates it and parses
	Parse(string) (*SudokuBoard, error)
}

type Parser struct{}

// e.g. 310004069000000200008005040000000005006000017807030000590700006600003050000100002

const sudokuNotationLength = 81

func (p Parser) Parse(arrangement string) (*SudokuBoard, error) {
	if len(arrangement) != sudokuNotationLength {
		return nil, errors.New("arrangement is not valid")
	}

	board := &boardType{}
	for i, v := range arrangement {
		board[i/9][i%9] = sudokuValue(v - '0')
	}
	return &SudokuBoard{board: board, log: logrus.New()}, nil
}
