//@author Stanislav Polaniev <spolanyev@gmail.com>

package main

import (
	cmd "dictionary/command"
	lib "dictionary/library"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

func main() {
	invoker := cmd.NewCommandInvoker()

	http.HandleFunc("/dictionary/run.php", func(writer http.ResponseWriter, request *http.Request) {

		lib.Log(lib.LogLevelDebug, lib.ColorBlue, "---------------- Triggered handler for `/dictionary/run.php` route ----------------", lib.ColorReset)

		urlStruct, err := url.Parse(request.RequestURI)
		if err != nil {
			lib.Log(lib.LogLevelDebug, "Error parsing URL:", err)
			return
		}

		urlQuery, err := url.QueryUnescape(urlStruct.RawQuery)
		if err != nil {
			lib.Log(lib.LogLevelDebug, "Error decoding URL:", err)
			return
		}

		payload := cmd.Payload{}
		err = json.Unmarshal([]byte(urlQuery), &payload)
		if err != nil {
			lib.Log(lib.LogLevelDebug, "Error parsing JSON:", err)
			return
		}

		lib.Log(lib.LogLevelDebug, fmt.Sprintf("%sStatement: name - %q, params - %q%s\n", lib.ColorCyan, payload.Name, payload.Params, lib.ColorReset))

		result := invoker.Invoke(payload.Name, payload.Params)

		lib.Log(lib.LogLevelDebug, fmt.Sprintf("%sResponse: %q%s\n", lib.ColorGreen, result, lib.ColorReset))

		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(&result)
		if err != nil {
			lib.Log(lib.LogLevelDebug, "Error encoding to JSON:", err)
			return
		}
	})

	http.HandleFunc("/dictionary/data/", func(writer http.ResponseWriter, request *http.Request) {

		lib.Log(lib.LogLevelDebug, lib.ColorBlue, "---------------- Triggered handler for `/dictionary/data/` route ----------------", lib.ColorReset)

		if strings.HasPrefix(request.URL.Path, "/dictionary/data/") {
			filePath := request.URL.Path[len("/dictionary/data/"):]
			if extension := filepath.Ext(filePath); extension == ".mp3" {

				lib.Log(lib.LogLevelDebug, fmt.Sprintf("%sFile: %q%s\n", lib.ColorYellow, filePath, lib.ColorReset))

				http.ServeFile(writer, request, "data/"+filePath)
				return
			}
		}
		http.NotFound(writer, request)
	})

	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		lib.Log(lib.LogLevelDebug, "Error starting server:", err)
		return
	}
}
