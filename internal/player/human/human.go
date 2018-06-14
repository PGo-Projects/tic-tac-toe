package human

import (
	"fmt"
	"strconv"

	"github.com/PGo-Projects/tic-tac-toe/internal/board"
	"github.com/PGo-Projects/tic-tac-toe/internal/userio"
)

const (
	TYPE                    = "human"
	MOVE_VALIDATION_PATTERN = "[123]"
)

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
	addressMsg := fmt.Sprintf("To the player with token %s:", h.token)
	rowPrompt := "Which row do you want to put your token? "
	rowErrMsg := "Not a valid row, please try again!"
	colPrompt := "Which column do you want to put your token?"
	colErrMsg := "Not a valid column, please try again!"

	row, rowErr := strconv.Atoi(userio.PromptUser(&userio.PromptUserInfo{
		AddressMsg:                 addressMsg,
		PromptMsg:                  rowPrompt,
		UserResponseIsValidPattern: MOVE_VALIDATION_PATTERN,
		ErrMsg: rowErrMsg,
	}))
	col, colErr := strconv.Atoi(userio.PromptUser(&userio.PromptUserInfo{
		AddressMsg:                 addressMsg,
		PromptMsg:                  colPrompt,
		UserResponseIsValidPattern: MOVE_VALIDATION_PATTERN,
		ErrMsg: colErrMsg,
	}))
	if rowErr == nil && colErr == nil {
		playErr := b.Put(row-1, col-1, h.token)
		return playErr
	}
	panic("This should not happen!")
}
