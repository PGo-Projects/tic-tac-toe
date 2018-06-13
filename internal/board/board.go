package board

import (
	"errors"
	"fmt"

	term "github.com/buger/goterm"
)

type Board struct {
	board [][]string
}

func New() *Board {
	return &Board{board: [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}}
}

func (b *Board) Put(x int, y int, token string) error {
	if b.board[x][y] != "" {
		return errors.New("The board already has a token there!")
	}
	b.board[x][y] = token
	return nil
}

func (b *Board) IsOver() bool {
	return b.board[0][0] != "" && b.board[0][0] == b.board[0][1] && b.board[0][0] == b.board[0][2] ||
		b.board[1][0] != "" && b.board[1][0] == b.board[1][1] && b.board[1][0] == b.board[1][2] ||
		b.board[2][0] != "" && b.board[2][0] == b.board[2][1] && b.board[2][0] == b.board[2][2] ||
		b.board[0][0] != "" && b.board[0][0] == b.board[1][0] && b.board[0][0] == b.board[2][0] ||
		b.board[0][1] != "" && b.board[0][1] == b.board[1][1] && b.board[0][1] == b.board[2][1] ||
		b.board[0][2] != "" && b.board[0][2] == b.board[1][2] && b.board[0][2] == b.board[2][2] ||
		b.board[0][0] != "" && b.board[0][0] == b.board[1][1] && b.board[0][0] == b.board[2][2] ||
		b.board[2][0] != "" && b.board[2][0] == b.board[1][1] && b.board[2][0] == b.board[0][2]
}

func (b *Board) Print() {
	term.Clear()
	term.MoveCursor(1, 1)
	term.Println(b)
	term.Flush()
}

func (b *Board) String() string {
	strBoard := ""

	for rowIndex, row := range b.board {
		for colIndex, token := range row {
			if token == "" {
				strBoard += "   "
			} else if token == "X" {
				strBoard += fmt.Sprintf(" %s ", term.Color(token, term.BLUE))
			} else if token == "O" {
				strBoard += fmt.Sprintf(" %s ", term.Color(token, term.RED))
			}
			if colIndex < 2 {
				strBoard += "|"
			}
		}
		strBoard += "\n"
		if rowIndex < 2 {
			strBoard += "-----------"
			strBoard += "\n"
		}
	}
	return strBoard
}
