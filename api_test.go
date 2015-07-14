package tgbot_test

import (
	"."
	"testing"
)

const (
	TestAPIKey = "1234:???"
	TestId     = 1234
)

func TestNew(t *testing.T) {
	bot, err := tgbot.New(TestAPIKey)
	if err != nil {
		t.Fatalf("Connect: %s", err)
	}
	info := bot.Info()
	t.Log(info)

	if info.Id != TestId {
		t.Errorf("Test id is wrong, is %d, should be %d", info.Id, TestId)
	}
}

func TestEcho(t *testing.T) {
	bot, err := tgbot.New(TestAPIKey)
	if err != nil {
		t.Fatalf("Connect: %s", err)
	}

	id := bot.Info().Id
	incoming, stop := bot.Listen(func(err error) bool {
		t.Fatal(err)
		return false
	})

	msg, err := bot.Send(id, "TestBotMessageEcho", false, nil)
	if err != nil {
		t.Fatalf("Send: %s", err)
	}

	msgGot := <-incoming
	t.Log(msg)
	t.Log(msgGot)

	stop <- true
}
