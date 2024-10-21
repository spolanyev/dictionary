//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	serv "dictionary/service"
	"net/http"
)

type ResponseService struct {
	httpStatusMapping map[cmd.CommandName]map[string]int
	messageService    *serv.MessageService
}

func NewResponseService(httpStatusMapping map[cmd.CommandName]map[string]int, messageService *serv.MessageService) *ResponseService {
	return &ResponseService{
		httpStatusMapping: httpStatusMapping,
		messageService:    messageService,
	}
}

func (resp *ResponseService) BuildResponse(commandName cmd.CommandName, dictionaryKey string, originalData map[string]interface{}) map[string]interface{} {
	//get HTTP status
	httpStatus, ok := resp.httpStatusMapping[commandName][dictionaryKey]
	if !ok {
		httpStatus = http.StatusInternalServerError // default 500
	}

	if dictionaryKey == "" {
		httpStatus = http.StatusOK
	}

	//substitute dictionary key
	message := resp.messageService.BuildMessage(commandName, dictionaryKey)

	//decorate result
	responseData := make(map[string]interface{})
	for key, value := range originalData {
		responseData[key] = value
	}
	responseData["message"] = message.Message
	responseData["httpStatusCode"] = httpStatus

	return responseData
}
