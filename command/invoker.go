//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	msg "dictionary/dictionary/message"
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
	logger.LogMessage("payload", fmt.Sprintf("%+v", payload))

	payload.SanitizeParameters()

	commandName := payload.GetCommandName()
	command, ok := invoker.commands[CommandName(commandName)]
	if !ok {
		logger.LogMessage("Unknown command:", commandName)

		return dto.NewErrorMessage(msg.InternalError, "Invoke")
	}

	result := command.Execute(payload)
	logger.LogMessage("result", result.ToMap())

	return result
}
