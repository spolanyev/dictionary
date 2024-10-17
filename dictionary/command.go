//@author Stanislav Polaniev <spolanyev@gmail.com>

package dictionary

import cmd "dictionary/command"

var Command = map[cmd.CommandName]map[string]string{
	cmd.AddWordToFileCommand: {
		"invalid_params":      "Invalid params",
		"invalid_values":      "Invalid values",
		"word_already_exists": "Word already exists",
		"word_added":          "Word added",
	},
	cmd.GetUserFileWordsCommand: {
		"invalid_params": "Invalid params",
	},
	cmd.GetWordDetailsCommand: {
		"invalid_params": "Invalid params",
		"invalid_word":   "Invalid word",
	},
	cmd.GetWordFromFileCommand: {
		"invalid_params": "Invalid params",
		"invalid_values": "Invalid values",
		"index_too_big":  "Index too big",
	},
	cmd.GetWordInformationCommand: {
		"invalid_params": "Invalid params",
		"invalid_word":   "Invalid word",
	},
	cmd.SearchWordCommand: {
		"invalid_params": "Invalid params",
	},
	cmd.UpdateWordDetailsCommand: {
		"invalid_params": "Invalid params",
		"invalid_word":   "Invalid word",
		"data_saved":     "Data saved",
	},
}
