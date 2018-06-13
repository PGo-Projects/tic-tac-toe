package ai

import "github.com/PGo-Projects/tic-tac-toe/internal/board"

const TYPE = "ai"

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
	return nil
}
