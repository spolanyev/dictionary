//@author Stanislav Polaniev <spolanyev@gmail.com>

package service

import (
	cmd "dictionary/command"
	dic "dictionary/dictionary"
	msg "dictionary/dictionary/message"
	"dictionary/dto"
	"dictionary/logger"
)

type MessageService struct {
	specific dic.CommandMessage
	common   dic.CommonMessage
}

func NewMessageService(specific dic.CommandMessage, common dic.CommonMessage) *MessageService {

	return &MessageService{
		specific: specific,
		common:   common,
	}
}

func (ms *MessageService) BuildMessage(commandName cmd.CommandName, dictionaryKey msg.Key) *dto.Message {
	message := ""

	if cmdMessages, ok := ms.specific[commandName]; ok {
		if msgText, found := cmdMessages[dictionaryKey]; found {
			message = msgText
		}
	}

	if message == "" {
		if commonMsg, ok := ms.common[dictionaryKey]; ok {
			message = commonMsg
		}
	}

	if message == "" {
		message = string(dictionaryKey)
		if dictionaryKey != "" {
			logger.LogMessage("commandName", commandName)
			logger.LogMessage("dictionaryKey", dictionaryKey)
		}
	}

	return dto.NewMessage(message, "server", nil, false)
}
