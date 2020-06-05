package vkapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type groupsGetLongPollServerAnswer struct {
	Key    string `json:"key"`
	Ts     string `json:"ts"`
	Server string `json:"server"`
}

//GetLongPollServerError is returnrd in case of some of values being not returned by server while getting secretKey, ts and server address
type GetLongPollServerError struct {
}

func (g GetLongPollServerError) Error() string {
	return "Invalid creds"
}

//GetServer func request api.vk.com for required information to get new events
func (bot *Bot) GetServer(token, groupID, APIver string) (err error) {
	response, err := http.Get(fmt.Sprintf("https://api.vk.com/method/groups.getLongPollServer?group_id=%s&access_token=%s&v=%s", groupID, token, APIver))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var (
		answer         struct{ Response groupsGetLongPollServerAnswer }
		responseStruct = &answer.Response
		answerBody     bytes.Buffer
	)
	io.Copy(&answerBody, response.Body)

	json.Unmarshal(answerBody.Bytes(), &answer)
	if responseStruct.Key == "" || responseStruct.Ts == "" || responseStruct.Server == "" {
		return GetLongPollServerError{}
	}
	bot.ts = responseStruct.Ts
	bot.server = responseStruct.Server
	bot.secretKey = responseStruct.Key
	bot.apiVer = APIver
	bot.accessToken = token
	return nil
}
