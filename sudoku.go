package sudoku_solver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type sudokuValue int8

func (s sudokuValue) isValid() bool {
	return s >= 0 && s <= 9
}

const sudokuSize = 9

type SudokuBoard struct {
	board *boardType
	log   *logrus.Logger
}

type boardType [sudokuSize][sudokuSize]sudokuValue

func (s boardType) IsValid() bool {
	var rows [9][10]int
	var col [9][10]int
	var grid [3][3][10]int

	for i := 0; i < sudokuSize; i++ {
		for j := 0; j < sudokuSize; j++ {
			num := s[i][j]
			if !num.isValid() {
				return false
			}

			if rows[i][num] < 1 {
				if num != 0 {
					rows[i][num] = rows[i][num] + 1

				}
			} else {
				//s.log.Info("rows error")
				return false
			}
			if col[j][num] < 1 {
				if num != 0 {
					col[j][num] = col[j][num] + 1
				}
			} else {
				//s.log.Info("col error")
				return false
			}

			if grid[i/3][j/3][num] < 1 {
				if num != 0 {
					grid[i/3][j/3][num] = grid[i/3][j/3][num] + 1
				}
			} else {
				//s.log.Info("grid error")
				return false
			}
		}
	}
	return true
}

func (s *SudokuBoard) Solve() error {
	t1 := time.Now()
	newBoard, ok := backtrackingSolve(*s.board, 0, 0)
	t2 := time.Now()
	if !ok {
		return errors.New("failed to backtrackingSolve sudoku board")
	}
	s.log.Info(fmt.Sprintf("took[ms]: %d", t2.Sub(t1).Milliseconds()))
	s.board = &newBoard
	return nil
}

func backtrackingSolve(board boardType, i, j int) (boardType, bool) {
	for ; i < sudokuSize; i++ {
		if j == sudokuSize {
			j = 0
		}
		for ; j < sudokuSize; j++ {
			if board[i][j] == 0 {
				// find proper value
				for v := 1; v < 10; v++ {
					board[i][j] = sudokuValue(v)

					if board.IsValid() {
						if solvedBoard, ok := backtrackingSolve(board, i, j); ok {
							return solvedBoard, true
						}
					}
				}
				return board, false
			}
		}
	}

	return board, true
}

func (s SudokuBoard) String() string {
	line := ".------.-------.------.\n"
	lastLine := "'------'-------'------'\n"
	lineBetween := ":------ ------- ------:\n"

	printedBoard := strings.Builder{}

	printedBoard.WriteString(line)
	for j, row := range s.board {
		for i, value := range row {
			if i == 3 || i == 6 {
				printedBoard.WriteString("| ")
			} else if i == 0 {
				printedBoard.WriteString("|")
			}

			if value != 0 {
				if i == 8 {
					printedBoard.WriteString(strconv.Itoa(int(value)))
				} else {
					printedBoard.WriteString(strconv.Itoa(int(value)) + " ")
				}
			} else {
				if i == 8 {
					printedBoard.WriteString(".")
				} else {
					printedBoard.WriteString(". ")
				}
			}

			if i == 8 {
				printedBoard.WriteString("|\n")
			}
		}
		if j == 8 {
			printedBoard.WriteString(lastLine)
		} else if (j+1)%3 == 0 {
			printedBoard.WriteString(lineBetween)
		}
	}
	return printedBoard.String()
}
