//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	lib "dictionary/library"
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func main() {
	invoker := cmd.NewCommandInvoker()

	http.HandleFunc("/dictionary/run.php", func(writer http.ResponseWriter, request *http.Request) {
		lib.Log(lib.LogLevelDebug, "\n---------------- Triggered handler for `/dictionary/run.php` route, method "+request.Method+" ----------------\n")

		writer.Header().Set("Access-Control-Allow-Origin", "*")

		if request.Method == http.MethodOptions {
			writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			writer.WriteHeader(http.StatusOK)
			return
		}

		payload := cmd.Payload{}

		if request.Method == http.MethodGet {
			lib.Log(lib.LogLevelDebug, "RequestURI:", request.RequestURI)
			urlStruct, err := url.Parse(request.RequestURI)
			if err != nil {
				lib.Log(lib.LogLevelDebug, "Error Parse:", err)
				http.Error(writer, "Bad Request", http.StatusBadRequest)
				return
			}

			lib.Log(lib.LogLevelDebug, "RawQuery:", urlStruct.RawQuery)
			urlQuery, err := url.QueryUnescape(urlStruct.RawQuery)
			if err != nil {
				lib.Log(lib.LogLevelDebug, "Error QueryUnescape:", err)
				http.Error(writer, "Bad Request", http.StatusBadRequest)
				return
			}

			err = json.Unmarshal([]byte(urlQuery), &payload)
			if err != nil {
				lib.Log(lib.LogLevelDebug, "Error Unmarshal:", err)
				http.Error(writer, "Bad Request", http.StatusBadRequest)
				return
			}
		} else if request.Method == http.MethodPost || request.Method == http.MethodPut {
			lib.Log(lib.LogLevelDebug, "Body:", request.Body)
			err := json.NewDecoder(request.Body).Decode(&payload)
			if err != nil {
				lib.Log(lib.LogLevelDebug, "Error Decode:", err)
				http.Error(writer, "Bad Request", http.StatusBadRequest)
				return
			}
		} else {
			lib.Log(lib.LogLevelDebug, "Unhandled HTTP method")
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		commandName := cmd.CommandName(payload.GetCommandName())
		requiredMethod, ok := cmd.CommandMethods[commandName]
		if !ok {
			lib.Log(lib.LogLevelDebug, "No such command:", commandName)
			http.Error(writer, "Command not found", http.StatusNotFound)
			return
		}

		if request.Method != requiredMethod {
			lib.Log(lib.LogLevelDebug, "Unexpected method for a command")
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		result := invoker.Invoke(&payload).ToMap()
		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			lib.Log(lib.LogLevelDebug, "Error Encode:", err)
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/dictionary/data/", func(writer http.ResponseWriter, request *http.Request) {
		lib.Log(lib.LogLevelDebug, "\n---------------- Triggered handler for `/dictionary/data/` route ----------------\n")

		lib.Log(lib.LogLevelDebug, "RequestURI:", request.RequestURI)

		if strings.HasPrefix(request.URL.Path, "/dictionary/data/") {
			filePath := request.URL.Path[len("/dictionary/"):]
			if extension := filepath.Ext(filePath); extension == ".mp3" {
				lib.Log(lib.LogLevelDebug, "MP3 file path:", filePath)
				http.ServeFile(writer, request, filePath)
				return
			}
		}
		http.NotFound(writer, request)
	})

	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		lib.Log(lib.LogLevelDebug, "Error ListenAndServe:", err)
		return
	}
}
