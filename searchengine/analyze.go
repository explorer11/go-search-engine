package searchengine

import "strings"

func analyze(word string) string {
	return strings.ToLower(word)
}

func analyzeFileName(fileName string) string {
	result := ""
	if fileName != "" {
		lowerCase := strings.ToLower(fileName)
		tokens := strings.Split(lowerCase, "/")
		name := tokens[len(tokens)-1]
		nameWithExtension := strings.Split(name, ".")
		result = nameWithExtension[0]
	}
	return result
}
