//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	stor "dictionary/storage"
	"net/http"
	"path/filepath"
)

type GetUserFileWords struct {
	fileManipulator *lib.FileManipulator
}

func NewGetUserFileWords(fileManipulator *lib.FileManipulator) *GetUserFileWords {
	return &GetUserFileWords{fileManipulator: fileManipulator}
}

func (*GetUserFileWords) GetName() CommandName {
	return GetUserFileWordsCommand
}

func (cmd *GetUserFileWords) Execute(payload dto.RequestInterface) dto.ResponseInterface {
	fileName, ok := payload.GetCommandParameters()["file"].(string)
	if !ok {
		return dto.NewErrorMessage("Invalid params", "GetUserFileWords")
	}
	fileName = filepath.Base(fileName)
	fullPathDirectory, err := lib.GetFullPathSourceDirectory(lib.NewCaller())
	if err != nil {
		return dto.NewErrorMessage(err.Error(), "GetUserFileWords", http.StatusInternalServerError)
	}
	fullPathFile := filepath.Join(fullPathDirectory, stor.PUBLIC_DIR, stor.USER_DATA_DIR, fileName)
	words, err := cmd.fileManipulator.GetLines(fullPathFile, "")
	if err != nil {
		return dto.NewErrorMessage(err.Error(), "GetUserFileWords", http.StatusInternalServerError)
	}
	return dto.NewSuccessResultMessage("getUserFileWords", words)
}
