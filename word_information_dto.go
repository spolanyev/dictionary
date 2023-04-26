//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

type WordInformationDto struct {
	Message string
	From    string
	IsError bool
	Data    WordInformationData
}

func (dto *WordInformationDto) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["safename"] = dto.Data.SafeName
	data["word"] = dto.Data.Word

	typeData := make(map[string]interface{})
	for typeName, typeInfo := range dto.Data.Type {
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
	result["message"] = dto.Message
	result["from"] = dto.From
	result["isError"] = dto.IsError
	result["data"] = data

	return result
}
