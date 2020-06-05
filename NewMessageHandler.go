package vkapi

import (
	"strings"
)

//NewMessageHandler function handle event of "message_new" type
func (bot *Bot) NewMessageHandler(event MessageNew) {
	if strings.HasPrefix(event.Message.Text, bot.commandMarker) {
		var message string
		if strings.IndexRune(event.Message.Text, rune(' ')) != -1 {
			message = event.Message.Text[1:strings.IndexRune(event.Message.Text, rune(' '))]
		} else {
			message = event.Message.Text[1:]
		}
		if function, ok := bot.commandHandlers[message]; ok {
			function(&event)
		} else {
			bot.commandHandlers["default"](&event)
		}
	}
}
