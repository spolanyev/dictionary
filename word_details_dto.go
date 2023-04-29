//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordDetailsDto struct {
	Message string
	From    string
	IsError bool
	Data    WordDetailsData
}

func (object *WordDetailsDto) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["message"] = object.Message
	result["from"] = object.From
	result["isError"] = object.IsError
	data := make(map[string]interface{})
	data["safeName"] = object.Data.SafeName
	data["word"] = object.Data.Word
	typeData := make(map[string]interface{})
	for k, v := range object.Data.Type {
		typeDetails := make(map[string]interface{})
		typeDetails["transcription"] = v.Transcription
		typeDetails["translation"] = v.Translation
		typeDetails["audio"] = v.Audio
		typeData[k] = typeDetails
	}
	data["type"] = typeData
	data["familyword"] = object.Data.FamilyWord
	rawData := make(map[string]interface{})
	rawTypeData := make(map[string]interface{})
	for k, v := range object.Data.Raw.Type {
		rawTypeDetails := make(map[string]interface{})
		rawTypeDetails["transcription"] = v.Transcription
		rawTypeDetails["translation"] = v.Translation
		rawTypeDetails["audio"] = v.Audio
		rawTypeDetails["level"] = v.Level
		rawTypeData[k] = rawTypeDetails
	}
	rawData["type"] = rawTypeData
	rawData["suggestedVariant"] = object.Data.Raw.SuggestedVariant
	rawData["americanEquivalent"] = object.Data.Raw.AmericanEquivalent
	rawData["isFamily"] = object.Data.Raw.IsFamily
	rawData["locale"] = object.Data.Raw.Locale
	data["raw"] = rawData
	result["data"] = data
	return result
}
