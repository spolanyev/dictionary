//@author Stanislav Polaniev <spolanyev@gmail.com>

package dictionary

import (
	cmd "dictionary/command"
	msg "dictionary/dictionary/message"
)

type keyToMessage map[msg.Key]string

type CommonMessage keyToMessage

type CommandMessage map[cmd.Name]keyToMessage

var CommonMap = CommonMessage{
	msg.InvalidParams: "Invalid params",
	msg.InternalError: "Internal error",
}

var CommandMap = CommandMessage{
	cmd.AddWordToFileCommand: {
		msg.InvalidValues:     "Invalid values",
		msg.WordAlreadyExists: "Word already exists",
		msg.WordAdded:         "Word added",
	},
	cmd.GetUserFileWordsCommand: {
		//no custom messages
	},
	cmd.GetWordDetailsCommand: {
		msg.InvalidWord: "Invalid word",
	},
	cmd.GetWordFromFileCommand: {
		msg.InvalidValues: "Invalid values",
		msg.IndexTooSmall: "Index too small",
		msg.IndexTooBig:   "Index too big",
	},
	cmd.GetWordInfoCommand: {
		msg.InvalidWord: "Invalid word",
	},
	cmd.SearchWordCommand: {
		//no custom messages
	},
	cmd.UpdateWordDetailsCommand: {
		msg.InvalidWord: "Invalid word",
		msg.DataSaved:   "Data saved",
	},
}
