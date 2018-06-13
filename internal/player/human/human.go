package human

import (
	"fmt"
	"strconv"

	"github.com/PGo-Projects/tic-tac-toe/internal/board"
	"github.com/PGo-Projects/tic-tac-toe/internal/userio"
)

const TYPE = "human"

type Human struct {
	token string
}

func New(token string) *Human {
	return &Human{token: token}
}

func (h *Human) GetToken() string {
	return h.token
}

func (h *Human) GetType() string {
	return TYPE
}

func (h *Human) PlayMove(b *board.Board) error {
	row, rowErr := strconv.Atoi(userio.PromptUser(fmt.Sprintf("To the player with token %s:", h.token), "Which row do you want to put your token? ", "[123]", "Not a valid row, please try again!"))
	col, colErr := strconv.Atoi(userio.PromptUser(fmt.Sprintf("To the player with token %s:", h.token), "Which column do you want to put your token? ", "[123]", "Not a valid column, please try again!"))
	if rowErr == nil && colErr == nil {
		playErr := b.Put(row-1, col-1, h.token)
		return playErr
	}
	panic("This should not happen!")
}
