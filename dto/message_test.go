//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

import (
	"reflect"
	"testing"
)

type FieldValueWithToMap struct {
	SomeField int
}

func (result FieldValueWithToMap) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"answer": result.SomeField,
	}
}

func TestMessageToMap(t *testing.T) {
	cases := []struct {
		message *Message
		want    map[string]interface{}
	}{
		{&Message{
			Message:        "Hello",
			From:           "John",
			Data:           FieldValueWithToMap{SomeField: 42},
			IsError:        false,
			HttpStatusCode: 200,
		}, map[string]interface{}{
			"message": "Hello",
			"from":    "John",
			"data": map[string]interface{}{
				"answer": 42,
			},
			"isError":        false,
			"httpStatusCode": 200,
		}},
		{&Message{
			Message:        "Bye",
			From:           "John",
			Data:           21,
			IsError:        true,
			HttpStatusCode: 400,
		}, map[string]interface{}{
			"message":        "Bye",
			"from":           "John",
			"data":           21,
			"isError":        true,
			"httpStatusCode": 400,
		}},
	}

	for _, theCase := range cases {
		got := theCase.message.ToMap()
		if !reflect.DeepEqual(got, theCase.want) {
			t.Errorf("ToMap() == %v, want %v", got, theCase.want)
		}
	}
}
