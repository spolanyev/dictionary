//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	msg "dictionary/dictionary/message"
	"net/http"
)

var commandToMethod = map[cmd.CommandName]string{
	cmd.GetUserFilesCommand:       http.MethodGet,
	cmd.GetUserFileWordsCommand:   http.MethodGet,
	cmd.GetLetterWordsCommand:     http.MethodGet,
	cmd.GetWordInformationCommand: http.MethodGet,
	cmd.GetWordDetailsCommand:     http.MethodGet,
	cmd.UpdateWordDetailsCommand:  http.MethodPut,
	cmd.SearchWordCommand:         http.MethodGet,
	cmd.AddWordToFileCommand:      http.MethodPost,
	cmd.GetWordFromFileCommand:    http.MethodGet,
}

type messageToStatus map[msg.Key]int
type commandToStatus map[cmd.CommandName]messageToStatus

var CommandToStatusMap = commandToStatus{
	cmd.AddWordToFileCommand: {
		msg.InvalidParams:     http.StatusBadRequest,
		msg.InvalidValues:     http.StatusBadRequest,
		msg.WordAlreadyExists: http.StatusOK,
		msg.WordAdded:         http.StatusCreated,
	},
	cmd.GetUserFileWordsCommand: {
		msg.InvalidParams: http.StatusBadRequest,
	},
	cmd.GetWordDetailsCommand: {
		msg.InvalidParams: http.StatusBadRequest,
		msg.InvalidWord:   http.StatusBadRequest,
	},
	cmd.GetWordFromFileCommand: {
		msg.InvalidParams: http.StatusBadRequest,
		msg.InvalidWord:   http.StatusBadRequest,
		msg.IndexTooBig:   http.StatusBadRequest,
	},
	cmd.GetWordInformationCommand: {
		msg.InvalidParams: http.StatusBadRequest,
		msg.InvalidWord:   http.StatusBadRequest,
	},
	cmd.SearchWordCommand: {
		msg.InvalidParams: http.StatusBadRequest,
	},
	cmd.UpdateWordDetailsCommand: {
		msg.InvalidParams: http.StatusBadRequest,
		msg.InvalidWord:   http.StatusBadRequest,
		msg.DataSaved:     http.StatusOK,
	},
}
