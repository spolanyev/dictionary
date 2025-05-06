//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	msg "dictionary/dictionary/message"
	"dictionary/dto"
	"dictionary/logger"
	"fmt"
)

type Invoker struct {
	commands map[Name]Command
}

func NewInvoker() *Invoker {
	return &Invoker{
		commands: make(map[Name]Command),
	}
}

func (invoker *Invoker) RegisterCommand(command Command) {
	invoker.commands[command.GetName()] = command
}

func (invoker *Invoker) Invoke(payload dto.Request) dto.Response {
	logger.LogMessage("payload", fmt.Sprintf("%+v", payload))

	payload.SanitizeParams()

	commandName := payload.GetCommandName()
	command, ok := invoker.commands[Name(commandName)]
	if !ok {
		logger.LogMessage("Unknown command:", commandName)

		return dto.NewErrorMessage(msg.InternalError, "Invoke")
	}

	result := command.Execute(payload)
	logger.LogMessage("result", result.ToMap())

	return result
}
