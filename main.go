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

		payload := cmd.Payload{}
		err = json.Unmarshal([]byte(urlQuery), &payload)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		fmt.Printf("%sStatement: name - %q, params - %q%s\n", lib.ColorCyan, payload.Name, payload.Params, lib.ColorReset)

		result := invoker.Invoke(payload.Name, payload.Params)

		fmt.Printf("%sResponse: %q%s\n", lib.ColorGreen, result, lib.ColorReset)

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
