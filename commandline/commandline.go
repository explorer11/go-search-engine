package commandline

import (
	"bufio"
	"fmt"
	"os"
	"searchengine/searchengine"
)

func Run(directory string) {

	searchengine.InitIndex(directory)

	var exit = "#exit"
	fmt.Printf("Index loaded. %s to exit\n", exit)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		query := scanner.Text()
		if query == exit {
			os.Exit(0)
		}

		filesScores := searchengine.QueryIndex(query)

		fmt.Println(filesScores)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
