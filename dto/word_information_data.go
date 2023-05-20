//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordInformationData struct {
	SafeName string
	Word     string
	Type     map[PartOfSpeech]WordTypeData
}

func (info *WordInformationData) ToMap() map[string]interface{} {
	data := make(map[string]interface{})
	data["safename"] = info.SafeName
	data["word"] = info.Word
	typeData := make(map[string]interface{})
	for typeName, typeInfo := range info.Type {
		typeMap := make(map[string]interface{})
		typeMap["transcription"] = typeInfo.Transcription
		typeMap["translation"] = typeInfo.Translation
		if typeInfo.Level != nil {
			typeMap["level"] = *typeInfo.Level
		}
		typeMap["audio"] = typeInfo.Audio
		typeData[string(typeName)] = typeMap
	}
	data["type"] = typeData
	return data
}
