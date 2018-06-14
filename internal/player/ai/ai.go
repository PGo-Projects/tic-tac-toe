package ai

import (
	"errors"

	"github.com/PGo-Projects/tic-tac-toe/internal/board"
	"github.com/PGo-Projects/tic-tac-toe/internal/utils"
)

const (
	TYPE                   = "ai"
	OUT_OF_BOUND_ERROR_MSG = "Out of bound!"
)

type Ai struct {
	token string
}

func New(token string) *Ai {
	return &Ai{token: token}
}

func (a *Ai) GetToken() string {
	return a.token
}

func (a *Ai) GetType() string {
	return TYPE
}

func (a *Ai) PlayMove(b *board.Board) error {
	bestMove, _ := a.minimax(b, 0)
	return b.Put(bestMove[0], bestMove[1], a.token)
}

func (a *Ai) minimax(b *board.Board, turn int) ([]int, int) {
	if b.SomeoneWon() {
		if turn == 0 {
			return nil, -10
		} else {
			return nil, 10
		}
	}
	if b.IsOver() {
		return nil, 0
	}
	availableMoves := b.AvailableMoves()
	moves := make([][]int, 0)
	scores := make([]int, 0)
	for _, move := range availableMoves {
		row := move[0]
		col := move[1]
		if turn == 0 {
			b.Put(row, col, a.token)
		} else {
			b.Put(row, col, utils.GetOtherToken(a.token))
		}
		moves = append(moves, move)
		_, score := a.minimax(b, 1-turn)
		scores = append(scores, score)
		b.Remove(row, col)
	}
	bestScore, index := getBestScore(scores, turn)
	return moves[index], bestScore
}

func getBestScore(scores []int, turn int) (int, int) {
	if turn == 0 {
		scoreIndex, err := getMaxIndexOfSlice(scores)
		if err == nil {
			return scores[scoreIndex], scoreIndex
		}
	} else {
		scoreIndex, err := getMinIndexOfSlice(scores)
		if err == nil {
			return scores[scoreIndex], scoreIndex
		}
	}
	return -1, -1
}

func getMinIndexOfSlice(slice []int) (minIndex int, err error) {
	if len(slice) == 0 {
		return -1, errors.New(OUT_OF_BOUND_ERROR_MSG)
	}
	currentMin := slice[0]
	for index, num := range slice {
		if num < currentMin {
			currentMin = num
			minIndex = index
		}
	}
	return minIndex, nil
}

func getMaxIndexOfSlice(slice []int) (maxIndex int, err error) {
	if len(slice) == 0 {
		return -1, errors.New(OUT_OF_BOUND_ERROR_MSG)
	}
	currentMax := slice[0]
	for index, num := range slice {
		if num > currentMax {
			currentMax = num
			maxIndex = index
		}
	}
	return maxIndex, nil
}
