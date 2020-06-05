package vkapi

import (
	"log"
)

//Bot - struct for new bot
type Bot struct {
	server, ts, secretKey string
	commandMarker         string
	eventChan             chan Event
	commandHandlers       map[string]func(*MessageNew) error
	extendedPipeline      bool
	eventHandlers         map[string]func(*interface{}) error
	apiVer                string
	accessToken           string
}

//NewBot creates new bot object with initialized chan
func NewBot() Bot {
	bot := Bot{
		commandHandlers: make(map[string]func(*MessageNew) error),
		eventChan:       make(chan Event, 1000),
		eventHandlers:   make(map[string]func(*interface{}) error),
		commandMarker:   "",
	}
	bot.commandHandlers["default"] = func(mess *MessageNew) error {
		log.Printf("Unhandled command: %s\n", mess.Message.Text)
		return nil
	}
	bot.eventHandlers["default"] = func(smth *interface{}) error {
		return nil
	}
	bot.extendedPipeline = false
	return bot
}

//CommandMarkerSetter function sets command prefix for chatbots
func (bot *Bot) CommandMarkerSetter(comMark string) {
	bot.commandMarker = comMark
}

//AddNewCommandHandler function adds new handler function for command
func (bot *Bot) AddNewCommandHandler(command string, handler func(*MessageNew) error) error {
	bot.commandHandlers[command] = handler
	return nil
}

func init() {
	typeSpecifier["message_typing_state"] = MessageTypingState{}
	typeSpecifier["message_new"] = MessageNew{}
	typeSpecifier["message_reply"] = MessageReply{}
	typeSpecifier["message_edit"] = MessageEdit{}
	typeSpecifier["message_allow"] = MessageAllow{}
	typeSpecifier["message_deny"] = MessageDeny{}
	typeSpecifier["photo_new"] = PhotoNew{}
	typeSpecifier["photo_comment_new"] = PhotoCommentNew{}
	typeSpecifier["photo_comment_edit"] = PhotoCommentEdit{}
	typeSpecifier["photo_comment_delete"] = PhotoCommentDelete{}
	typeSpecifier["photo_comment_restore"] = PhotoCommentRestore{}
}
