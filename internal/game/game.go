package game

import (
	"math/rand"
	"regexp"
	"time"

	"github.com/PGo-Projects/tic-tac-toe/internal/board"
	"github.com/PGo-Projects/tic-tac-toe/internal/player"
	"github.com/PGo-Projects/tic-tac-toe/internal/userio"
	term "github.com/buger/goterm"
)

type Game struct {
	board   *board.Board
	players []player.Player
	turn    int
}

func New() *Game {
	return &Game{board: board.New()}
}

func (g *Game) Start() {
	term.Clear()
	term.MoveCursor(1, 1)
	PrintInstructions()

	playAgainstAI := userio.PromptUser("", "Want to play against an AI? ", "([yY][eE][sS])|([nN][oO])", "Please answer yes or no.")
	if match, err := regexp.MatchString("[yY][eE][sS]", playAgainstAI); err != nil && match {
		playerGoesFirst := userio.PromptUser("", "Do you want to go first? ", "([yY][eE][sS])|([nN][oO])", "Please answer yes or no.")
		if match, err := regexp.MatchString("[yY][eE][sS]", playerGoesFirst); err != nil && match {
			playerTokenErrMsg := "Token must be a X or a O, please try again!"
			playerToken := userio.PromptUser("", "What token should the first player use? ", "X|O", playerTokenErrMsg)
			g.players = []player.Player{player.New("human", playerToken), player.New("ai", getOtherToken(playerToken))}
		} else {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(100) < 50 {
				g.players = []player.Player{player.New("ai", "X"), player.New("human", getOtherToken("X"))}
			} else {
				g.players = []player.Player{player.New("ai", "O"), player.New("human", getOtherToken("O"))}
			}
		}
	} else {
		firstPlayerTokenErrMsg := "Token must be a X or a O, please try again!"
		firstPlayerToken := userio.PromptUser("", "What token should the first player use? ", "X|O", firstPlayerTokenErrMsg)
		g.players = []player.Player{player.New("human", firstPlayerToken), player.New("human", getOtherToken(firstPlayerToken))}
	}

	g.turn = 0
	g.drawBoard()
	g.play()
}

func (g *Game) drawBoard() {
	g.board.Print()
}

func (g *Game) play() {
	if g.board.IsOver() {
		term.Println("Game Over!")
		term.Flush()
		return
	}

	playErr := g.getCurrentPlayer().PlayMove(g.board)
	g.drawBoard()
	if playErr != nil {
		term.Println(playErr)
		term.Flush()
	} else {
		g.nextTurn()
	}

	g.play()
}

func (g *Game) getCurrentPlayer() player.Player {
	return g.players[g.turn]
}

func (g *Game) nextTurn() {
	g.turn = 1 - g.turn
}

func PrintInstructions() {
	term.Println("This is a game of tic-tac-toe!")
	term.Flush()
}

func getOtherToken(chosenToken string) string {
	if chosenToken == "X" {
		return "O"
	} else {
		return "X"
	}
}
