package vkapi

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

//SendRequest function calls API method wpecified in method arg provided with arguments specified in arguments arg.
//All arguments names must be same as in official vkontakte api documentation
func (bot *Bot) SendRequest(method string, arguments map[string]string) (answer []byte, err error) {
	address := fmt.Sprintf("https://api.vk.com/method/%s?access_token=%s&v=%s", method, bot.accessToken, bot.apiVer)
	for key, value := range arguments { //TODO without sprintf. Only stream
		address = fmt.Sprintf("%s&%s=%s", address, key, url.QueryEscape(value))
	}
	fmt.Println(address)
	ans, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	//var result interface{}
	var bodybytes bytes.Buffer
	io.Copy(&bodybytes, ans.Body)

	return bodybytes.Bytes(), nil
}
