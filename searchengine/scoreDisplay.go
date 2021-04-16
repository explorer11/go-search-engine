package searchengine

func filesScores(filesNames map[int]string, scores map[int]float64) map[string]float64 {

	var filesScores = map[string]float64{}

	for fileId, score := range scores {
		fileName := filesNames[fileId]
		filesScores[fileName] = score
	}

	return filesScores
}
