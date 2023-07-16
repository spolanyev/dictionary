package dto

type Message struct {
	Message string
	From    string
	Data    interface{}
	IsError bool
}

func (message *Message) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["message"] = message.Message
	result["from"] = message.From
	if message.Data != nil {
		if response, ok := message.Data.(ResponseInterface); ok {
			result["data"] = response.ToMap()
		} else if message.Data != nil {
			result["data"] = message.Data
		}
	}
	result["isError"] = message.IsError
	return result
}
