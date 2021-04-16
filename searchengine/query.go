package searchengine

func (index *Index) query(words []string) map[int]float64 {
	var baseScores = map[int]int{}

	for _, word := range words {

		analyzedWord := analyze(word)
		filesIds := index.index[analyzedWord]

		var wordScore = map[int]bool{}
		for _, fileId := range filesIds {
			if _, alreadyFound := wordScore[fileId]; !alreadyFound {
				wordScore[fileId] = true
				if value, exist := baseScores[fileId]; exist {
					value = value + 1
					baseScores[fileId] = value
				} else {
					baseScores[fileId] = 1
				}
			}
		}

	}

	return finalScore(words, baseScores)
}

func finalScore(words []string, filesScores map[int]int) map[int]float64 {

	var scores = map[int]float64{}

	wordNumber := len(words)

	for fileId, fileScore := range filesScores {
		scores[fileId] = float64(fileScore) / float64(wordNumber)
	}

	return scores
}
