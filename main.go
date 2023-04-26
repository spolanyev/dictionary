//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func main() {

	http.HandleFunc("/dictionary/run.php", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("---------------- Triggered handler for `/dictionary/run.php` route ----------------")

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

		payload := CommandPayload{}
		err = json.Unmarshal([]byte(urlQuery), &payload)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		command := getCommand(&payload)
		if command == nil {
			fmt.Printf("Unknown command name: %q\n", payload.Name)
			return
		}

		result := command.Execute(&payload.Params)
		fmt.Printf("Result: %q\n", result)

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(&result)
		if err != nil {
			fmt.Println("Error encoding to JSON:", err)
			return
		}
	})

	http.HandleFunc("/dictionary/data/", func(writer http.ResponseWriter, request *http.Request) {

		fmt.Println("---------------- Triggered handler for `/dictionary/data/` route ----------------")

		if strings.HasPrefix(request.URL.Path, "/dictionary/data/") {
			filePath := request.URL.Path[len("/dictionary/data/"):]
			if extension := filepath.Ext(filePath); extension == ".mp3" {

				fmt.Printf("File: %#q\n", filePath)

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

func getCommand(payload *CommandPayload) CommandInterface {

	fmt.Printf("Command: name - %q, params - %#q\n", payload.Name, payload.Params)

	switch payload.Name {
	case "getUserFiles":
		return &GetUserFilesCommand{}
	case "getUserFileWords":
		return &GetUserFileWordsCommand{}
	case "getLetterWords":
		return &GetLetterWordsCommand{}
	case "getWordInformation":
		return &GetWordInformationCommand{}
	case "getWordDetails":
		return &GetWordDetailsCommand{}
	case "updateWordDetails":
		return &UpdateWordDetailsCommand{}
	default:
		return nil
	}
}
