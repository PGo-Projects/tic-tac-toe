package main

import (
	"github.com/PGo-Projects/tic-tac-toe/internal/game"
	"github.com/spf13/cobra"
)

func main() {
	var cmdPlay = &cobra.Command{
		Use:   "play",
		Short: "Start a tic-tac-toe game",
		Long:  "Play a tic-tac-toe game against the computer or another player",
		Args:  cobra.ExactArgs(0),
		Run:   startGame,
	}

	var rootCmd = &cobra.Command{Use: "tictactoe"}
	rootCmd.AddCommand(cmdPlay)
	rootCmd.Execute()
}

func startGame(cmd *cobra.Command, args []string) {
	tictactoe := game.New()
	tictactoe.Start()
}
