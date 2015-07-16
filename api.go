package tgbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Bot struct {
	apiKey  string
	BaseURL string
	info    TGUser
	client  *http.Client
	debug   bool
}

const Timeout = 5
const DefaultBaseURL = "https://api.telegram.org/bot"

func New(apiKey string) (*Bot, error) {
	if len(apiKey) < 3 {
		return nil, ErrorInvalidArgs
	}

	b := &Bot{apiKey: apiKey, client: &http.Client{Timeout: time.Second * Timeout},
		BaseURL: DefaultBaseURL}

	err := b.connect()
	return b, err
}

func (b *Bot) Debug(enable bool) {
	b.debug = enable
}

func (b *Bot) Req(method string, receiver interface{}) error {
	url := b.BaseURL + b.apiKey + "/" + method
	httpRes, err := b.client.Get(url)

	//TODO certificate pinning
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()
	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}

	if b.debug {
		fmt.Println("response", string(body))
	}
	if httpRes.StatusCode != 200 {
		return fmt.Errorf("HTTP Error %s:\n%s",
			httpRes.Status, string(body))
	}

	var res TGResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Errorf("GET %s:\n%s", url, string(body))
		return err
	}
	if !res.Ok {
		return fmt.Errorf("Server error: %s", res.Description)
	}
	return json.Unmarshal(res.Result, receiver)
}

// Listen starts listening for updates.
// The given errHandler function will be called when an error occurs. If it
// returns false, listening is stopped.
//
// Two channels are returned: incoming, an unbuffered channel providing updates,
// and stop, which causes the listening to stop if a value is written to it.
// Stop can take up to 10 seconds to react.
func (b *Bot) Listen(errHandler func(error) bool) (incoming <-chan TGMessage, stop chan<- bool) {
	inc := make(chan TGMessage)
	st := make(chan bool, 1)
	incoming, stop = inc, st

	go b.listener(inc, st, errHandler)
	return
}

func (b *Bot) listener(inc chan<- TGMessage, st <-chan bool, errHandler func(error) bool) {
	var offset Integer = 0
	cont := true
	for cont {
		var res []TGUpdate
		err := b.Req(fmt.Sprintf("getUpdates?offset=%d&timeout=%d", offset, Timeout-2), &res)
		if err != nil {
			//TODO is this a timeout error?
			fmt.Println("timeout?", err)
			ret := errHandler(err)
			if !ret {
				break
			}
		}
		for _, u := range res {
			if u.UpdateId >= offset {
				offset = u.UpdateId + 1
			}
			inc <- u.Message
		}
		select {
		case <-st:
			cont = false
		default:
		}
	}
	close(inc)
}

func (b *Bot) connect() error {
	err := b.Req("getMe", &b.info)
	return err
}

func (b *Bot) Info() TGUser {
	return b.info
}

func (b *Bot) Send(chatId ID, text string) (TGMessage, error) {
	return b.SendAdv(chatId, text, false, nil)
}

func (b *Bot) SendAdv(chatId ID, text string, disablePreview bool,
	replyingToId *Integer) (TGMessage, error) {

	str := fmt.Sprintf("sendMessage?chat_id=%d&text=%s", chatId.ID(), url.QueryEscape(text))
	if disablePreview {
		str += "&disable_web_page_preview=true"
	}
	if replyingToId != nil {
		str += fmt.Sprintf("reply_to_message_id=%d", *replyingToId)
	}
	var msg TGMessage
	err := b.Req(str, &msg)
	return msg, err
}
