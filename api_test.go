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
		t.Errorf("Connect: %s", err)
	}
	info := bot.Info()
	t.Log(info)

	if info.Id != TestId {
		t.Errorf("Test id is wrong, is %d, should be %d", info.Id, TestId)
	}
}
