package vkapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//Event is container for unserialized by mapstructure eventes, where Type is string, containig event type
type Event struct {
	Type   string      `json:"type"`
	Object interface{} `json:"object"`
}

//EventsFromServer contains response from server with new events
type EventsFromServer struct {
	Ts      string  `json:"ts"`
	Updates []Event `json:"updates"`
}

//BotIsNotInitializedError is error throwing in case of calling StartPolling for uninitialized bot struct
type BotIsNotInitializedError struct {
}

func (err BotIsNotInitializedError) Error() string {
	return "Bot is not initialized. Call Bot.GetServer(token, groupID, APIver string) at first and Bot.CommandMarkerSetter(comMark string) than."
}

//StartPolling function start getting new events and handling them
func (bot *Bot) StartPolling() (err error) {
	if bot.server == "" || bot.secretKey == "" || bot.ts == "" || bot.commandMarker == "" {
		return BotIsNotInitializedError{}
	}
	go bot.startDispatcher()
	for {
		answer, err := http.Get(fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25", bot.server, bot.secretKey, bot.ts))
		if err != nil {
			log.Println(err)
			return err
		}

		var (
			answerBody bytes.Buffer
			events     EventsFromServer
		)

		io.Copy(&answerBody, answer.Body)
		json.Unmarshal(answerBody.Bytes(), &events)
		bot.ts = events.Ts
		for _, v := range events.Updates {
			if _, ok := typeSpecifier[v.Type]; ok {
				bot.eventChan <- v
				//log.Printf("New event os type %s\n", v.Type)
			}
		}
	}
}
