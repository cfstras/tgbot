package tgbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bot struct {
	apiKey string
	info   TGUser
}

const BaseURL = "https://api.telegram.org/bot"

func New(apiKey string) (*Bot, error) {
	if len(apiKey) < 3 {
		return nil, ErrorInvalidArgs
	}
	b := &Bot{apiKey: apiKey}
	err := b.connect()
	return b, err
}

func (b *Bot) Req(method string, receiver interface{}) error {
	httpRes, err := http.Get(BaseURL + b.apiKey + "/" + method)
	//TODO certificate pinning
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	var res TGResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	if !res.Ok {
		return fmt.Errorf("API error: %s", res.Description)
	}
	return json.Unmarshal(res.Result, receiver)
}

func (b *Bot) connect() error {
	err := b.Req("getMe", &b.info)
	return err
}

func (b *Bot) Info() TGUser {
	return b.info
}
