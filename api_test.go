package tgbot_test

import (
	"."
	"testing"
)

const TestAPIKey = "110894750:AAFY9JYP03mc4EtofS75ogX2vkJp-67ksFc"

func TestNew(t *testing.T) {
	bot, err := tgbot.New(TestAPIKey)
	if err != nil {
		t.Error(err)
	}
	bot.Name()
}
