//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordInformationData struct {
	SafeName string
	Word     string
	Type     map[PartOfSpeech]WordTypeData
}

func (info *WordInformationData) ToMap() map[string]interface{} {
	typeData := make(map[string]interface{})
	for typeName, typeInfo := range info.Type {
		typeMap := map[string]interface{}{
			"transcription": typeInfo.Transcription,
			"translation":   typeInfo.Translation,
			"audio":         typeInfo.Audio,
		}

		if typeInfo.Level != nil {
			typeMap["level"] = *typeInfo.Level
		}
		typeData[string(typeName)] = typeMap
	}

	return map[string]interface{}{
		"safename": info.SafeName,
		"word":     info.Word,
		"type":     typeData,
	}
}
