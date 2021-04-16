package searchengine

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

type Files struct {
	counter    int
	filesNames map[int]string
}

func (files *Files) init() {
	files.filesNames = map[int]string{}
}

func (files *Files) fill(relativePath string, directoryName string, processFileNameFunc processFileName, processWordFunc processWord) {
	wg.Add(1)
	go files.fillRoutine(&wg, relativePath, directoryName, processFileNameFunc, processWordFunc)
	wg.Wait()
}

func (files *Files) fillRoutine(wg *sync.WaitGroup, relativePath string, directoryName string, processFileNameFunc processFileName, processWordFunc processWord) {
	defer wg.Done()

	var path string = directoryName
	if relativePath != "" {
		path = relativePath + "/" + directoryName
	}
	readFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var directoryPath string = relativePath + directoryName
	if relativePath != "" {
		directoryPath = relativePath + "/" + directoryName
	}
	for _, file := range readFiles {
		if file.IsDir() == true {
			wg.Add(1)
			go files.fillRoutine(wg, directoryPath, file.Name(), processFileNameFunc, processWordFunc)
		} else {
			var fileName string = file.Name()
			if relativePath != "" {
				fileName = directoryName + "/" + file.Name()
			}

			var fileId int
			mutex.Lock()
			fileId = files.counter
			files.filesNames[files.counter] = fileName
			files.counter++
			mutex.Unlock()

			processFileNameFunc(fileId, fileName)
			processFileWords(directoryPath, file.Name(), fileId, processWordFunc)
		}
	}
}

type processFileName func(int, string)
type processWord func(int, string)

func processFileWords(directory string, file string, fileId int, processWordFunc processWord) {
	path, err := os.Open(directory + "/" + file)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = path.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(path)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		processWordFunc(fileId, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
