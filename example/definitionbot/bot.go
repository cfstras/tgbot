package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/cfstras/tgbot"
)

const BaseURL = "https://%s.wiktionary.org/w/api.php?action=query&format=json&" +
	"prop=extracts&exchars=300&explaintext&titles=%s"
const InfoURL = "https://%s.wiktionary.org/wiki/%s"

var prefixes = map[string]string{
	"definiere ": "de",
	"define ":    "en",
}

var apikey string

func main() {
	apikey, _ = tgbot.GetTokenFromEnv()
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
		if v.Text == nil {
			continue
		}
		qual, lang, query := qualifies(*v.Text)
		if !qual {
			continue
		}
		fmt.Println("New Message:", v)
		res := getWikiDefinition(query, lang)
		if res == "" {
			_, err = bot.Send(v.From, "No definition on Wiktionary "+lang+" for "+query)
		} else {
			fmt.Println("answer:\n", res)
			_, err = bot.Send(v.From, res)
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func buildString(lang, query string) string {
	return fmt.Sprintf(BaseURL, lang, url.QueryEscape(query))
}

func qualifies(text string) (result bool, lang string, query string) {
	for p, l := range prefixes {
		does := strings.HasPrefix(text, p)
		if !does {
			continue
		}
		query := strings.TrimPrefix(text, p)
		return true, l, query
	}
	return false, "", ""
}

func getWikiDefinition(query, lang string) string {
	wikiURL := buildString(lang, query)
	fmt.Println(wikiURL)
	httpRes, err := http.Get(wikiURL)

	if err != nil {
		return err.Error()
	}
	defer httpRes.Body.Close()
	body, err := ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err.Error()
	}

	if httpRes.StatusCode != 200 {
		return fmt.Sprintf("HTTP Error %s:\n%s",
			httpRes.Status, string(body))
	}

	var res WikiResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return fmt.Sprintf("GET %s:\n%s", wikiURL, string(body))
	}

	for _, v := range res.Query.Pages {
		if v.Missing != nil {
			continue
		}
		ex, oldex := v.Extract, ""
		for oldex != ex {
			oldex = ex
			ex = strings.Replace(ex, "\n\n", "\n", -1)
		}

		return ex + "\n" +
			fmt.Sprintf(InfoURL, lang, url.QueryEscape(v.Title))
	}
	return ""
}

type WikiResponse struct {
	Query struct {
		Pages map[string]struct {
			Title   string  `json:"title"`
			Extract string  `json:"extract"`
			Missing *string `json:"missing"`
		} `json:"pages"`
	} `json:"query"`
}
