//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	msg "dictionary/dictionary/message"
	hdlr "dictionary/service"
	"net/http"
)

const defaultHttpStatus = http.StatusInternalServerError

type ResponseService struct {
	CommandStatusMap commandToStatus
	messageService   *hdlr.MessageService
}

func NewResponseService(httpStatusMapping commandToStatus, messageService *hdlr.MessageService) *ResponseService {

	return &ResponseService{
		CommandStatusMap: httpStatusMapping,
		messageService:   messageService,
	}
}

func (rs *ResponseService) BuildHttpResponse(commandName cmd.CommandName, dictionaryKey msg.Key, originalData map[string]interface{}) map[string]interface{} {
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
