package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	gzb "github.com/ifo/gozulipbot"
	"io/ioutil"
	"log"
	"time"
)

func SendMd(token string, content string) error {
	bot := gzb.Bot{}
	bot.APIKey = token
	bot.APIURL = "https://zulip.v2chengdu.club/api/v1/"
	bot.Email = "smart-bot@zulip.v2chengdu.club"
	bot.Backoff = 2 * time.Second
	bot.Init()
	m := gzb.Message{
		Stream:  "LeetCode",
		Topic:   "Daily Challenge",
		Content: content,
	}
	resp, err := bot.Message(m)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var toPrint bytes.Buffer

	err = json.Indent(&toPrint, body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(toPrint.String())
	return nil
}
