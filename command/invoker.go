//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import stor "dictionary/storage"

type Invoker struct {
	commands map[string]Interface
}

func NewCommandInvoker() *Invoker {
	loader := stor.GetLoader()
	return &Invoker{
		commands: map[string]Interface{
			"getUserFiles":       &GetUserFiles{},
			"getUserFileWords":   &GetUserFileWords{},
			"getLetterWords":     &GetLetterWords{},
			"getWordInformation": NewGetWordInformationCommand(loader),
			"getWordDetails":     NewGetWordDetailsCommand(loader),
			"updateWordDetails":  &UpdateWordDetails{},
		},
	}
}

func (invoker *Invoker) Invoke(name string, params map[string]interface{}) map[string]interface{} {
	command, exists := invoker.commands[name]
	if !exists {
		return map[string]interface{}{
			"isError": true,
			"from":    "invoker",
			"message": "Unknown command",
		}
	}
	return command.Execute(params)
}
