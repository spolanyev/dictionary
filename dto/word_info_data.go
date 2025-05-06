//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordInfoData struct {
	SafeName string
	Word     string
	Type     map[PartOfSpeech]WordTypeData
}

func (info *WordInfoData) ToMap() map[string]interface{} {
	typeData := make(map[string]interface{}, len(info.Type))
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
