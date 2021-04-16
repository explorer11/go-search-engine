package searchengine

import (
	"strings"
)

var filesNames map[int]string

var files Files
var index Index

func InitIndex(directory string) {
	files.init()
	index.init()
	files.fill("", directory, index.fillFileName, index.fill)
}

func QueryIndex(query string) map[string]float64 {
	words := strings.Split(query, " ")
	scores := index.query(words)
	return filesScores(files.filesNames, scores)
}
