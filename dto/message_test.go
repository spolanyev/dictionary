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
			Message: "Hello",
			From:    "John",
			Data:    FieldValueWithToMap{SomeField: 42},
			IsError: false,
		}, map[string]interface{}{
			"message": "Hello",
			"from":    "John",
			"data": map[string]interface{}{
				"answer": 42,
			},
			"isError": false,
		}},
		{&Message{
			Message: "Bye",
			From:    "John",
			Data:    21,
			IsError: true,
		}, map[string]interface{}{
			"message": "Bye",
			"from":    "John",
			"data":    21,
			"isError": true,
		}},
	}

	for _, theCase := range cases {
		got := theCase.message.ToMap()
		if !reflect.DeepEqual(got, theCase.want) {
			t.Errorf("ToMap() == %v, want %v", got, theCase.want)
		}
	}
}
