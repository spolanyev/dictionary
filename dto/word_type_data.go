//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordTypeData struct {
	Transcription string
	Translation   []string
	Level         *[]Level //is optional
	Audio         []AudioData
}
