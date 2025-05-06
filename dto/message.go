//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

import msg "dictionary/dictionary/message"

type Message struct {
	Message string
	From    string
	Data    interface{}
	IsError bool
}

func NewMessage(message, from string, data interface{}, isError bool) *Message {
	return &Message{
		Message: message,
		From:    from,
		Data:    data,
		IsError: isError,
	}
}

func NewErrorMessage(message msg.Key, from string) *Message {

	return NewMessage(string(message), from, nil, true)
}

func NewSuccessResultMessage(from string, data interface{}) *Message {

	return NewMessage("", from, data, false)
}

func (message *Message) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["message"] = message.Message
	result["from"] = message.From
	if message.Data != nil {
		if response, ok := message.Data.(Response); ok {
			result["data"] = response.ToMap()
		} else {
			result["data"] = message.Data
		}
	}
	result["isError"] = message.IsError

	return result
}
