//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	stor "dictionary/storage"
)

type Invoker struct {
	commands map[string]CommandInterface
}

func NewCommandInvoker() *Invoker {
	loader := stor.GetLoader()
	return &Invoker{
		commands: map[string]CommandInterface{
			"getUserFiles":       &GetUserFiles{},
			"getUserFileWords":   &GetUserFileWords{},
			"getLetterWords":     &GetLetterWords{},
			"getWordInformation": NewGetWordInformationCommand(loader),
			"getWordDetails":     NewGetWordDetailsCommand(loader),
			"updateWordDetails":  &UpdateWordDetails{},
			"searchWord":         &SearchWord{},
		},
	}
}

func (invoker *Invoker) Invoke(name string, params map[string]interface{}) map[string]interface{} {
	command, exists := invoker.commands[name]
	if !exists {
		return (&dto.ErrorMessage{Message: "Unknown command", From: "Invoke"}).ToMap()
	}
	return command.Execute(params).ToMap()
}
