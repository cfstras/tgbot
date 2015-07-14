package main

import (
	"flag"
	"fmt"
	"github.com/cfstras/tgbot"
)

var apikey string

func main() {
	flags()
	bot, err := tgbot.New(apikey)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Bot Name:", bot.Name())
}

func flags() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Errorf("Please provide an API key:\n    echobot <key>")
	}
	apikey = flag.Arg(0)
}
