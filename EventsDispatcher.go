package vkapi

import (
	"github.com/mitchellh/mapstructure"
)

func (bot *Bot) startDispatcher() {
	for {
		i := <-bot.eventChan
		Type := i.Type
		switch Type {
		case "message_new":
			var event MessageNew
			mapstructure.Decode(i.Object, &event)
			go bot.NewMessageHandler(event)
		}
		if bot.extendedPipeline || Type != "message_new" {
			if function, ok := bot.eventHandlers[Type]; ok {
				function(&i.Object)
			} else {
				bot.eventHandlers["default"](&i.Object)
			}
		}
	}
}
