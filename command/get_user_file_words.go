//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	msg "dictionary/dictionary/message"
	"dictionary/dto"
	lib "dictionary/library"
	"dictionary/logger"
	stor "dictionary/storage"
	"path/filepath"
)

type GetUserFileWords struct {
	fileManipulator *lib.FileManipulator
}

func NewGetUserFileWords(fm *lib.FileManipulator) *GetUserFileWords {
	return &GetUserFileWords{fileManipulator: fm}
}

func (*GetUserFileWords) GetName() CommandName {
	return GetUserFileWordsCommand
}

func (cmd *GetUserFileWords) Execute(payload dto.RequestInterface) dto.ResponseInterface {
	params := payload.GetCommandParameters()
	commandName := string(cmd.GetName())

	fileName, ok := params["file"].(string)
	if !ok || fileName == "" {
		logger.LogMessage("file", params["file"])
		return dto.NewErrorMessage(msg.InvalidParams, commandName)
	}

	fileName = filepath.Base(fileName)
	fullPathFile := filepath.Join(cmd.fileManipulator.SourceDir, stor.PUBLIC_DIR, stor.USER_DATA_DIR, fileName)

	words, err := cmd.fileManipulator.GetLines(fullPathFile, "")
	if err != nil {
		return dto.NewErrorMessage(msg.InternalError, commandName)
	}

	return dto.NewSuccessResultMessage(commandName, words)
}
