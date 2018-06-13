package player

import "github.com/PGo-Projects/tic-tac-toe/internal/player/human"

var (
	YOUR_TOKEN     string = "X"
	OPPONENT_TOKEN string = "O"
)

type Player interface {
	GetToken() string
}

func New(playerType string, token string) Player {
	if playerType == "human" {
		return human.New(token)
	}
	return nil
}
