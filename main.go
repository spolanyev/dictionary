//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	lib "dictionary/library"
	stor "dictionary/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func main() {

	http.HandleFunc("/dictionary/run.php", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println(lib.ColorBlue, "---------------- Triggered handler for `/dictionary/run.php` route ----------------", lib.ColorReset)

		urlStruct, err := url.Parse(request.RequestURI)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}

		urlQuery, err := url.QueryUnescape(urlStruct.RawQuery)
		if err != nil {
			fmt.Println("Error decoding URL:", err)
			return
		}

		payload := cmd.CommandPayload{}
		err = json.Unmarshal([]byte(urlQuery), &payload)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		command := getCommand(payload)
		if command == nil {
			fmt.Printf("%sUnknown command: %q%s\n", lib.ColorRed, payload.Name, lib.ColorReset)
			return
		}

		result := command.Execute(payload.Params)
		fmt.Printf("%sResult: %q%s\n", lib.ColorGreen, result, lib.ColorReset)

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(&result)
		if err != nil {
			fmt.Println("Error encoding to JSON:", err)
			return
		}
	})

	http.HandleFunc("/dictionary/data/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println(lib.ColorBlue, "---------------- Triggered handler for `/dictionary/data/` route ----------------", lib.ColorReset)

		if strings.HasPrefix(request.URL.Path, "/dictionary/data/") {
			filePath := request.URL.Path[len("/dictionary/data/"):]
			if extension := filepath.Ext(filePath); extension == ".mp3" {

				fmt.Printf("%sFile: %q%s\n", lib.ColorYellow, filePath, lib.ColorReset)

				http.ServeFile(writer, request, "data/"+filePath)
				return
			}
		}
		http.NotFound(writer, request)
	})

	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

func getCommand(payload cmd.CommandPayload) cmd.CommandInterface {

	fmt.Printf("%sCommand: name - %q, params - %q%s\n", lib.ColorCyan, payload.Name, payload.Params, lib.ColorReset)

	switch payload.Name {
	case "getUserFiles":
		return &cmd.GetUserFilesCommand{}
	case "getUserFileWords":
		return &cmd.GetUserFileWordsCommand{}
	case "getLetterWords":
		return &cmd.GetLetterWordsCommand{}
	case "getWordInformation":
		fileManipulator := &lib.FileManipulator{}
		storage := stor.NewWordFileStorage(fileManipulator)
		loader := stor.NewWordDataLoader(storage)
		return cmd.NewGetWordInformationCommand(loader)
	case "getWordDetails":
		fileManipulator := &lib.FileManipulator{}
		storage := stor.NewWordFileStorage(fileManipulator)
		loader := stor.NewWordDataLoader(storage)
		return cmd.NewGetWordDetailsCommand(loader)
	case "updateWordDetails":
		return &cmd.UpdateWordDetailsCommand{}
	default:
		return nil
	}
}
