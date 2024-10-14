//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	"fmt"
)

type Invoker struct {
	commands map[CommandName]CommandInterface
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make(map[CommandName]CommandInterface),
	}
}

func (invoker *Invoker) RegisterCommand(name CommandName, command CommandInterface) {
	invoker.commands[name] = command
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
