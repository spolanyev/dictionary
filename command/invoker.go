//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	stor "dictionary/storage"
	"fmt"
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

func (invoker *Invoker) Invoke(payload dto.RequestInterface) dto.ResponseInterface {
	lib.Log(lib.LogLevelDebug, "Payload:", fmt.Sprintf("%+v\n", payload))

	command, ok := invoker.commands[payload.GetCommandName()]
	if !ok {
		lib.Log(lib.LogLevelDebug, "No such command", payload.GetCommandName())
		err := &dto.ErrorMessage{Message: "Unknown command", From: "Invoke"}
		return err
	}

	result := command.Execute(payload)
	lib.Log(lib.LogLevelDebug, "Result:", result.ToMap())

	return result
}
