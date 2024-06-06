//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type Message struct {
	Message        string
	From           string
	Data           interface{}
	IsError        bool
	HttpStatusCode int
}

func NewMessage(message, from string, data interface{}, isError bool, httpStatusCode int) *Message {
	return &Message{
		Message:        message,
		From:           from,
		Data:           data,
		IsError:        isError,
		HttpStatusCode: httpStatusCode,
	}
}

func (message *Message) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["message"] = message.Message
	result["from"] = message.From
	if message.Data != nil {
		if response, ok := message.Data.(ResponseInterface); ok {
			result["data"] = response.ToMap()
		} else {
			result["data"] = message.Data
		}
	}
	result["isError"] = message.IsError
	result["httpStatusCode"] = message.HttpStatusCode
	return result
}
