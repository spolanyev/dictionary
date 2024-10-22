//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	dic "dictionary/dictionary"
	lib "dictionary/library"
	"dictionary/logger"
	serv "dictionary/server"
	stor "dictionary/storage"
	"os"
)

func main() {
	commandInvoker := cmd.NewInvoker()
	fileManipulator := lib.NewFileManipulator()
	wordStorage, err := stor.NewWordFileStorage(fileManipulator)
	if err != nil {
		os.Exit(2)
	}
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

	if err := serv.NewServer(commandInvoker, serv.DictionaryKeyHttpStatusMapping, dic.Command).Start(); err != nil {
		logger.LogError(err, "Start")
		os.Exit(1)
	}
}
