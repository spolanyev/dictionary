//@author Stanislav Polaniev <spolanyev@gmail.com>

package command

import (
	"dictionary/dto"
	lib "dictionary/library"
	stor "dictionary/storage"
	"path/filepath"
)

type GetUserFileWords struct {
	FileManipulator *lib.FileManipulator
}

func (cmd *GetUserFileWords) Execute(params map[string]interface{}) dto.ResponseInterface {
	fileName, ok := params["file"].(string)
	if !ok {
		return &dto.ErrorMessage{Message: "Invalid params", From: "GetUserFileWords"}
	}
	fileName = filepath.Base(fileName)
	fullPathDirectory, err := lib.GetFullPathSourceDirectory()
	if err != nil {
		return &dto.ErrorMessage{Message: err.Error(), From: "GetUserFileWords"}
	}
	fullPathFile := filepath.Join(fullPathDirectory, stor.PUBLIC_DIR, stor.USER_DATA_DIR, fileName)
	words, err := cmd.FileManipulator.GetSlice(fullPathFile)
	if err != nil {
		return &dto.ErrorMessage{Message: err.Error(), From: "GetUserFileWords"}
	}
	return &dto.Message{Message: "", From: "getUserFileWords", Data: words, IsError: false}
}
