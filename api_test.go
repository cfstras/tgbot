package tgbot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
)

var (
	TestAPIKey string
	TestId     Integer
)

func init() {
	TestAPIKey = os.Getenv("TELEGRAM_TOKEN")
	if TestAPIKey == "" {
		fmt.Println("Please provide a bot token in the environment variable TELEGRAM_TOKEN.")
		os.Exit(1)
	}
	parts := strings.SplitN(TestAPIKey, ":", 2)
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
	TestId = Integer(i)
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
