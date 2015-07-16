package main

import (
	"flag"
	"fmt"
	"github.com/cfstras/tgbot"
	"os"
	"os/signal"
)

var apikey string

func main() {
	if !flags() {
		return
	}
	fmt.Println("Starting echoBot on key", apikey)
	bot, err := tgbot.New(apikey)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Bot Name:", bot.Info().Username)

	incoming, stop := bot.Listen(func(err error) bool {
		fmt.Println(err)
		return false
	})

	// handle ctrl+c
	sigs := make(chan os.Signal)
	signal.Notify(sigs)
	go func() {
		for _ = range sigs {
			stop <- true
		}
	}()

	for v := range incoming {
		fmt.Println("New Message:", v)
		_, err = bot.Send(v.From, v.String())
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func flags() bool {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Please provide an API key:\n    echobot <key>")
		return false
	}
	apikey = flag.Arg(0)
	return true
}
