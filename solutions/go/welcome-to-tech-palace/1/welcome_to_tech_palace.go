package techpalace


import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %v", strings.ToUpper(customer))
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf("%s\n%s\n%s",strings.Repeat("*", numStarsPerLine),welcomeMsg,strings.Repeat("*", numStarsPerLine))
}


func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg,"*",""))
}

