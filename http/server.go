package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"searchengine/searchengine"
	"strconv"
)

func Run(directory string) {

	searchengine.InitIndex(directory)
	fmt.Printf("Server started\n")

	handler := func(responseWriter http.ResponseWriter, request *http.Request) {

		query := request.FormValue("q")
		filesScores := searchengine.QueryIndex(query)

		for file, score := range filesScores {
			io.WriteString(responseWriter, file+" : ")
			io.WriteString(responseWriter, strconv.FormatFloat(score, 'f', -1, 64))
			io.WriteString(responseWriter, "\n")
		}
	}

	http.HandleFunc("/search", handler)

	log.Fatal(http.ListenAndServe(":10000", nil))
}
