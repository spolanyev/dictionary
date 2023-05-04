//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordInformationData struct {
	SafeName string
	Word     string
	Type     map[PartOfSpeech]WordTypeData
}
