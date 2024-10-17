//@author Stanislav Polaniev <spolanyev@gmail.com>

package server

import (
	cmd "dictionary/command"
	"net/http"
)

type CommandMethodPair struct {
	CommandName cmd.CommandName
	HttpMethod  string
}

var commandMethodMapping = []CommandMethodPair{
	{cmd.GetUserFilesCommand, http.MethodGet},
	{cmd.GetUserFileWordsCommand, http.MethodGet},
	{cmd.GetLetterWordsCommand, http.MethodGet},
	{cmd.GetWordInformationCommand, http.MethodGet},
	{cmd.GetWordDetailsCommand, http.MethodGet},
	{cmd.UpdateWordDetailsCommand, http.MethodPut},
	{cmd.SearchWordCommand, http.MethodGet},
	{cmd.AddWordToFileCommand, http.MethodPost},
	{cmd.GetWordFromFileCommand, http.MethodGet},
}

var DictionaryKeyHttpStatusMapping = map[cmd.CommandName]map[string]int{
	cmd.AddWordToFileCommand: {
		"invalid_params":      http.StatusBadRequest,
		"invalid_values":      http.StatusBadRequest,
		"word_already_exists": http.StatusOK,
		"word_added":          http.StatusCreated,
	},
	cmd.GetUserFileWordsCommand: {
		"invalid_params": http.StatusBadRequest,
	},
	cmd.GetWordDetailsCommand: {
		"invalid_params": http.StatusBadRequest,
		"invalid_word":   http.StatusBadRequest,
	},
	cmd.GetWordFromFileCommand: {
		"invalid_params": http.StatusBadRequest,
		"invalid_word":   http.StatusBadRequest,
		"index_too_big":  http.StatusBadRequest,
	},
	cmd.GetWordInformationCommand: {
		"invalid_params": http.StatusBadRequest,
		"invalid_word":   http.StatusBadRequest,
	},
	cmd.SearchWordCommand: {
		"invalid_params": http.StatusBadRequest,
	},
	cmd.UpdateWordDetailsCommand: {
		"invalid_params": http.StatusBadRequest,
		"invalid_word":   http.StatusBadRequest,
		"data_saved":     http.StatusOK,
	},
}
