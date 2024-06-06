//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	stor "dictionary/storage"
	"fmt"
)

type Invoker struct {
	commands map[CommandName]CommandInterface
}

func NewCommandInvoker() *Invoker {
	storage := stor.GetStorage()
	loader := stor.GetLoader(storage)
	return &Invoker{
		commands: map[CommandName]CommandInterface{
			GetUserFilesCommand:       &GetUserFiles{},
			GetUserFileWordsCommand:   &GetUserFileWords{},
			GetLetterWordsCommand:     &GetLetterWords{},
			GetWordInformationCommand: NewGetWordInformationCommand(loader),
			GetWordDetailsCommand:     NewGetWordDetailsCommand(loader),
			UpdateWordDetailsCommand:  NewUpdateWordDetailsCommand(storage),
			SearchWordCommand:         &SearchWord{},
			AddWordToFileCommand:      &AddWordToFile{},
			GetWordFromFileCommand:    &GetWordFromFile{},
		},
	}
}

func (invoker *Invoker) Invoke(payload dto.RequestInterface) dto.ResponseInterface {
	lib.Log(lib.LogLevelDebug, "Payload:", fmt.Sprintf("%+v\n", payload))

	command, ok := invoker.commands[CommandName(payload.GetCommandName())]
	if !ok {
		lib.Log(lib.LogLevelDebug, "No such command:", payload.GetCommandName())
		err := dto.NewErrorMessage("Unknown command", "Invoke")
		return err
	}

	result := command.Execute(payload)
	lib.Log(lib.LogLevelDebug, "Result:", result.ToMap())
	return result
}
