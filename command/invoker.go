//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	"dictionary/logger"
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

func (invoker *Invoker) RegisterCommand(command CommandInterface) {
	invoker.commands[command.GetName()] = command
}

func (invoker *Invoker) Invoke(payload dto.RequestInterface) dto.ResponseInterface {
	logger.LogMessage("Payload", fmt.Sprintf("%+v", payload))

	command, ok := invoker.commands[CommandName(payload.GetCommandName())]
	if !ok {
		logger.LogMessage("No such command", payload.GetCommandName())
		err := dto.NewErrorMessage("Unknown command", "Invoke")
		return err
	}

	result := command.Execute(payload)
	logger.LogMessage("Result", result.ToMap())
	return result
}
