package userio

import (
	"fmt"
	"regexp"

	term "github.com/buger/goterm"
)

func PromptUser(addressMsg string, promptMsg string, userResponseIsValidPattern string, errMsg string) string {
	if addressMsg != "" {
		term.Println(addressMsg)
	}
	term.Println(promptMsg)
	term.MoveCursorUp(1)
	term.Flush()
	userResponse := ""
	match, err := regexp.MatchString(userResponseIsValidPattern, userResponse)
	for err != nil || !match {
		if userResponse != "" {
			term.Println(errMsg)
			term.Flush()
			userResponse = ""
		}
		for userResponse == "" {
			fmt.Scanln(&userResponse)
		}
		match, err = regexp.MatchString(userResponseIsValidPattern, userResponse)
	}
	term.Print("\n")
	term.Flush()
	return userResponse
}
