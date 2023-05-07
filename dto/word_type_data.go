//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordTypeData struct {
	Transcription string
	Translation   []string
	Level         *[]string //is optional
	Audio         []string
}
