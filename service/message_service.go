//@author Stanislav Polaniev <spolanyev@gmail.com>

package service

import (
	cmd "dictionary/command"
	"dictionary/dto"
	"dictionary/logger"
)

type MessageService struct {
	messageMapping map[cmd.CommandName]map[string]string
}

func NewMessageService(messageMapping map[cmd.CommandName]map[string]string) *MessageService {
	return &MessageService{
		messageMapping: messageMapping,
	}
}

func (ms *MessageService) BuildMessage(commandName cmd.CommandName, dictionaryKey string) *dto.Message {
	message, ok := ms.messageMapping[commandName][dictionaryKey]
	if !ok {
		logger.LogMessage("commandName", commandName)
		logger.LogMessage("dictionaryKey", dictionaryKey)
		message = dictionaryKey //default message is the key itself
	}

	return dto.NewMessage(message, "server", nil, false)
}
