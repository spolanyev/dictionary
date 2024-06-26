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
	FileManipulator *lib.FileManipulator
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
	words, err := cmd.FileManipulator.GetLines(fullPathFile, "")
	if err != nil {
		return dto.NewErrorMessage(err.Error(), "GetUserFileWords", http.StatusInternalServerError)
	}
	return dto.NewSuccessResultMessage("getUserFileWords", words)
}
