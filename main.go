//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	dic "dictionary/dictionary"
	lib "dictionary/library"
	serv "dictionary/server"
	stor "dictionary/storage"
	"os"
)

func main() {
	commandInvoker := cmd.NewInvoker()

	rootDir, err := lib.GetFullPathSourceDir(lib.NewCaller())
	if err != nil {
		os.Exit(3)
	}
	fileManipulator := lib.NewFileManipulator(rootDir)

	wordStorage, err := stor.NewWordFileStorage(fileManipulator)
	if err != nil {
		os.Exit(2)
	}

	wordLoader := stor.NewWordDataLoader(wordStorage)

	commandInvoker.RegisterCommand(cmd.NewGetUserFiles(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetUserFileWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetLetterWords(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetWordInfo(wordLoader))
	commandInvoker.RegisterCommand(cmd.NewGetWordDetails(wordLoader))
	commandInvoker.RegisterCommand(cmd.NewUpdateWordDetails(wordStorage))
	commandInvoker.RegisterCommand(cmd.NewSearchWord(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewAddWordToFile(fileManipulator))
	commandInvoker.RegisterCommand(cmd.NewGetWordFromFile(fileManipulator))

	server := serv.NewServer(commandInvoker, serv.CommandToStatusMap, dic.CommandMap, dic.CommonMap)
	if err := server.Start(); err != nil {
		os.Exit(1)
	}
}

//TODO
//Lie – lie – lay – lain
//Lie – lie – lied – lied

//add not found word
