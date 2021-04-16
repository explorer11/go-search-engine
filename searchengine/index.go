package searchengine

import (
	"sync"
)

type Index struct {
	index map[string][]int
}

var indexMutex sync.Mutex

func (index *Index) init() {
	index.index = map[string][]int{}
}

func (index *Index) fill(fileId int, word string) {
	analyzedWord := analyze(word)
	index.fillIndex(fileId, analyzedWord)
}

func (index *Index) fillFileName(fileId int, fileName string) {
	analyzedName := analyzeFileName(fileName)
	index.fillIndex(fileId, analyzedName)
}

func (index *Index) fillIndex(fileId int, token string) {
	indexMutex.Lock()
	if value, exist := index.index[token]; exist {
		value = append(value, fileId)
		index.index[token] = value
	} else {
		var initSlice = []int{fileId}
		index.index[token] = initSlice
	}
	indexMutex.Unlock()
}
