//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	"dictionary/dictionary"
	msg "dictionary/dictionary/message"
	"dictionary/dto"
	"dictionary/logger"
)

type MessageBuilder struct {
	specific dictionary.CommandMessage
	common   dictionary.CommonMessage
}

func NewMessageService(specific dictionary.CommandMessage, common dictionary.CommonMessage) *MessageBuilder {
	return &MessageBuilder{
		specific: specific,
		common:   common,
	}
}

func (ms *MessageBuilder) BuildMessage(commandName cmd.Name, dictionaryKey msg.Key) *dto.Message {
	message := ""

	if cmdMessages, ok := ms.specific[commandName]; ok {
		if msgText, ok := cmdMessages[dictionaryKey]; ok {
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
