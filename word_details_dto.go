//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

type WordDetailsDto struct {
	Message string
	From    string
	IsError bool
	Data    WordDetailsData
}

func (dto *WordDetailsDto) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["message"] = dto.Message
	result["from"] = dto.From
	result["isError"] = dto.IsError
	data := make(map[string]interface{})
	data["safeName"] = dto.Data.SafeName
	data["word"] = dto.Data.Word
	typeData := make(map[string]interface{})
	for k, v := range dto.Data.Type {
		typeDetails := make(map[string]interface{})
		typeDetails["transcription"] = v.Transcription
		typeDetails["translation"] = v.Translation
		if v.Level != nil {
			typeDetails["level"] = *v.Level
		}
		typeDetails["audio"] = v.Audio
		typeData[k] = typeDetails
	}
	data["type"] = typeData
	data["familyword"] = dto.Data.FamilyWord
	rawData := make(map[string]interface{})
	rawTypeData := make(map[string]interface{})
	for k, v := range dto.Data.Raw.Type {
		rawTypeDetails := make(map[string]interface{})
		rawTypeDetails["transcription"] = v.Transcription
		rawTypeDetails["translation"] = v.Translation
		rawTypeDetails["audio"] = v.Audio
		if v.Level != nil {
			rawTypeDetails["level"] = v.Level
		}
		rawTypeData[k] = rawTypeDetails
	}
	rawData["type"] = rawTypeData
	rawData["suggestedVariant"] = dto.Data.Raw.SuggestedVariant
	rawData["americanEquivalent"] = dto.Data.Raw.AmericanEquivalent
	rawData["isFamily"] = dto.Data.Raw.IsFamily
	rawData["locale"] = dto.Data.Raw.Locale
	data["raw"] = rawData
	result["data"] = data
	return result
}
