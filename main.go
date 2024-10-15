//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	lib "dictionary/library"
	"dictionary/server"
	stor "dictionary/storage"
	"os"
)

func main() {
	commandInvoker := cmd.NewInvoker()
	fileManipulator := lib.NewFileManipulator()
	wordStorage := stor.NewWordFileStorage(fileManipulator)
	wordLoader := stor.NewWordDataLoader(wordStorage)

	commandInvoker.RegisterCommand(cmd.NewGetUserFiles(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetUserFileWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetLetterWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetWordInformation(wordLoader))
	commandInvoker.RegisterCommand(cmd.NewGetWordDetails(wordLoader))
	commandInvoker.RegisterCommand(cmd.NewUpdateWordDetails(wordStorage))
	commandInvoker.RegisterCommand(cmd.NewSearchWord(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewAddWordToFile(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetWordFromFile(fileManipulator))

	if err := server.NewServer(commandInvoker).Start(); err != nil {
		os.Exit(1)
	}
}
