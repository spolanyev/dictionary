//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordInformationDto struct {
	Message string
	From    string
	IsError bool
	Data    WordInformationData
}

func (object *WordInformationDto) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["safename"] = object.Data.SafeName
	data["word"] = object.Data.Word

	typeData := make(map[string]interface{})
	for typeName, typeInfo := range object.Data.Type {
		typeMap := make(map[string]interface{})
		typeMap["transcription"] = typeInfo.Transcription
		typeMap["translation"] = typeInfo.Translation
		if typeInfo.Level != nil {
			typeMap["level"] = *typeInfo.Level
		}
		typeMap["audio"] = typeInfo.Audio
		typeData[typeName] = typeMap
	}

	data["type"] = typeData

	result := make(map[string]interface{})
	result["message"] = object.Message
	result["from"] = object.From
	result["isError"] = object.IsError
	result["data"] = data

	return result
}
