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

	commandInvoker.RegisterCommand(cmd.GetUserFilesCommand, cmd.NewGetUserFiles(fileManipulator))
	commandInvoker.RegisterCommand(cmd.GetUserFileWordsCommand, cmd.NewGetUserFileWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.GetLetterWordsCommand, cmd.NewGetLetterWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.GetWordInformationCommand, cmd.NewGetWordInformation(wordLoader))
	commandInvoker.RegisterCommand(cmd.GetWordDetailsCommand, cmd.NewGetWordDetails(wordLoader))
	commandInvoker.RegisterCommand(cmd.UpdateWordDetailsCommand, cmd.NewUpdateWordDetails(wordStorage))
	commandInvoker.RegisterCommand(cmd.SearchWordCommand, cmd.NewSearchWord(fileManipulator))
	commandInvoker.RegisterCommand(cmd.AddWordToFileCommand, cmd.NewAddWordToFile(fileManipulator))
	commandInvoker.RegisterCommand(cmd.GetWordFromFileCommand, cmd.NewGetWordFromFile(fileManipulator))

	if err := server.NewServer(commandInvoker).Start(); err != nil {
		os.Exit(1)
	}
}
