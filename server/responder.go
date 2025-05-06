//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	msg "dictionary/dictionary/message"
	"net/http"
)

const defaultHttpStatus = http.StatusInternalServerError

type Responder struct {
	CommandStatusMap commandToStatus
	messageService   *MessageBuilder
}

func NewResponseService(httpStatusMapping commandToStatus, messageService *MessageBuilder) *Responder {
	return &Responder{
		CommandStatusMap: httpStatusMapping,
		messageService:   messageService,
	}
}

func (rs *Responder) BuildHttpResponse(commandName cmd.Name, dictionaryKey msg.Key, originalData map[string]interface{}) map[string]interface{} {
	//get HTTP status
	httpStatus, ok := rs.CommandStatusMap[commandName][dictionaryKey]
	if !ok {
		httpStatus = defaultHttpStatus
	}

	if dictionaryKey == "" {
		httpStatus = http.StatusOK
	}

	//substitute dictionary key
	message := rs.messageService.BuildMessage(commandName, dictionaryKey)

	//decorate result
	responseData := make(map[string]interface{})
	for key, value := range originalData {
		responseData[key] = value
	}
	responseData["message"] = message.Message
	responseData["httpStatusCode"] = httpStatus

	return responseData
}
