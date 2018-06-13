package game

import (
	"fmt"
	"strconv"

	"github.com/PGo-Projects/tic-tac-toe/internal/board"
	"github.com/PGo-Projects/tic-tac-toe/internal/player"
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
	firstPlayerToken := promptUser(nil, "What token should the first player use? ")
	g.players = []player.Player{player.New("human", firstPlayerToken), player.New("human", getOtherToken(firstPlayerToken))}
	g.turn = 0

	g.DrawBoard()
	g.Play()
}

func (g *Game) DrawBoard() {
	g.board.Print()
}

func (g *Game) Play() {
	if g.board.IsOver() {
		term.Println("Game Over!")
		term.Flush()
	}

	playErr := error(nil)
	playerFirstTry := true
	for playErr != nil || playerFirstTry {
		playerFirstTry = false
		row, rowErr := strconv.Atoi(promptUser(g.getCurrentPlayer(), "Which row do you want to put your token? "))
		col, colErr := strconv.Atoi(promptUser(g.getCurrentPlayer(), "Which col do you want to put your token? "))
		if rowErr != nil {
			playErr = rowErr
		}
		if colErr != nil {
			playErr = colErr
		}

		playErr = g.playMove(row, col)
	}
	g.nextTurn()
	g.Play()
}

func (g *Game) getCurrentPlayer() player.Player {
	return g.players[g.turn]
}

func (g *Game) playMove(row int, col int) error {
	err := g.board.Put(row-1, col-1, g.getCurrentPlayer().GetToken())
	if err == nil {
		g.board.Print()
	}
	return err
}

func (g *Game) nextTurn() {
	g.turn = 1 - g.turn
}

func PrintInstructions() {
	term.Println("This is a game of tic-tac-toe!")
	term.Flush()
}

func promptUser(p player.Player, prompt string) string {
	if p != nil {
		term.Printf("For player with token %s:", p.GetToken())
		term.Println("")
	}
	term.Println(prompt)
	term.Flush()
	userResponse := ""
	for userResponse == "" {
		fmt.Scanln(&userResponse)
	}
	return userResponse
}

func getOtherToken(chosenToken string) string {
	if chosenToken == "X" {
		return "O"
	} else {
		return "X"
	}
}
