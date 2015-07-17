package tgbot

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetTokenFromEnv() (token string, id Integer) {
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		fmt.Println("Please provide a bot token in the environment variable TELEGRAM_TOKEN.")
		os.Exit(1)
	}
	parts := strings.SplitN(token, ":", 2)
	if len(parts) < 2 {
		fmt.Println("Illegal Token, does not contain ID.")
		os.Exit(1)
	}
	var err error
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		fmt.Println("Illegal Token, ID", parts[0], "is not integer.")
		os.Exit(1)
	}
	id = Integer(i)
	return
}
