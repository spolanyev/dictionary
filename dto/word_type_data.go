//@author Stanislav Polaniev <spolanyev@gmail.com>

package dto

type WordTypeData struct {
	Transcription string
	Translation   []string
	Level         *[]string //pointer to make `Level` optional
	Audio         []string
}
