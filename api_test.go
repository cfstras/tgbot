package tgbot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	TestAPIKey string
	TestId     Integer
)

func init() {
	TestAPIKey, TestId = GetTokenFromEnv()
}

func TestNew(t *testing.T) {
	bot, err := New(TestAPIKey)
	if err != nil {
		t.Fatalf("Connect: %s", err)
	}
	info := bot.Info()
	t.Log(info)

	if info.Id != TestId {
		t.Errorf("Test id is wrong, is %d, should be %d", info.Id, TestId)
	}
}

func TestSend(t *testing.T) {
	bot, err := New(TestAPIKey)
	if err != nil {
		t.Fatalf("Connect: %s", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok":true,"result":{"message_id":37}}`)
	}))
	defer server.Close()
	bot.BaseURL = server.URL + "/"

	_, err = bot.Send(Integer(1234), "Testing tgbot")
	if err != nil {
		t.Error(err)
	}
}
