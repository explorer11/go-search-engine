package searchengine

import (
	"fmt"
	"math"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestQueryIndexWithOneWordAndOneFile(t *testing.T) {

	var file1 int = 1
	var file2 int = 2
	var word string = "word"
	var otherWord string = "otherWord"

	var index Index
	index.init()
	index.index[word] = []int{file1}
	index.index[otherWord] = []int{file2}

	words := []string{word}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 1))
	assert.Check(t, cmp.Equal(filesScores[file1], 1.0))
}

func TestQueryIndexWithOneWordAndSpecialChars(t *testing.T) {

	var file1 int = 1
	var file2 int = 2
	var word string = "codé"
	var otherWord string = "otherword"

	var index Index
	index.init()
	index.index[word] = []int{file1}
	index.index[otherWord] = []int{file2}

	words := []string{"Codé"}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 1))
	assert.Check(t, cmp.Equal(filesScores[file1], 1.0))
}

func TestQueryIndexWithOneWordAndTwoFiles(t *testing.T) {

	var file1 int = 1
	var file2 int = 2
	var word string = "word"

	var index Index
	index.init()
	index.index[word] = []int{file1, file2}

	words := []string{word}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 2))
	assert.Check(t, cmp.Equal(filesScores[file1], 1.0))
	assert.Check(t, cmp.Equal(filesScores[file2], 1.0))
}

func TestQueryIndexWithTwoWords(t *testing.T) {

	var file1 int = 1
	var file2 int = 2
	var word string = "word"
	var otherWord string = "otherword"

	var index Index
	index.init()
	index.index[word] = []int{file1, file2}
	index.index[otherWord] = []int{file2}

	words := []string{word, otherWord}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 2))
	assert.Check(t, cmp.Equal(filesScores[file1], 0.5))
	assert.Check(t, cmp.Equal(filesScores[file2], 1.0))
}

func TestQueryIndexWithSameWordFoundMultipleTimes(t *testing.T) {

	var file1 int = 1
	var word string = "word"

	var index Index
	index.init()
	index.index[word] = []int{file1, file1}

	words := []string{word}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 1))
	assert.Check(t, cmp.Equal(filesScores[file1], 1.0))
}

func TestQueryIndexComplexCase(t *testing.T) {

	var file1 int = 1
	var file2 int = 2
	var file3 int = 3
	var file4 int = 4
	var word1 string = "word1"
	var word2 string = "word2"
	var word3 string = "word3"
	var word4 string = "word4"
	var word5 string = "word5"

	var index Index
	index.init()
	index.index[word1] = []int{file1, file1, file2}
	index.index[word2] = []int{file2, file3, file4}
	index.index[word3] = []int{file2, file3}
	index.index[word4] = []int{file2}

	words := []string{word2, word3, word5}
	filesScores := index.query(words)

	assert.Check(t, cmp.Len(filesScores, 3))
	assert.Check(t, math.Abs(filesScores[file2]-0.66) < 0.01, fmt.Sprintf("actual value %f", filesScores[file2]))
	assert.Check(t, math.Abs(filesScores[file3]-0.66) < 0.01, fmt.Sprintf("actual value %f", filesScores[file3]))
	assert.Check(t, math.Abs(filesScores[file4]-0.33) < 0.01, fmt.Sprintf("actual value %f", filesScores[file4]))
}

func TestFinalScoreOneWordNoResult(t *testing.T) {

	words := []string{"word"}
	scores := finalScore(words, map[int]int{})

	expectedLength := 0
	if len(scores) != expectedLength {
		t.Errorf("Failed ! got %v scores want %v", len(scores), expectedLength)
	} else {
		t.Logf("Success !")
	}
}

func TestFinalScoreOneWordOneResult(t *testing.T) {

	words := []string{"word"}

	wordsCount := map[int]int{}
	file1 := 1
	wordsCount[file1] = 1

	scores := finalScore(words, wordsCount)

	assert.Check(t, cmp.Len(scores, 1))
	assert.Check(t, cmp.Equal(scores[file1], 1.0))
}

func TestFinalScoreOneWordTwoResults(t *testing.T) {

	words := []string{"word"}

	wordsCount := map[int]int{}
	file1 := 1
	wordsCount[file1] = 1
	file2 := 2
	wordsCount[file2] = 1

	scores := finalScore(words, wordsCount)

	assert.Check(t, cmp.Len(scores, 2))
	assert.Check(t, cmp.Equal(scores[file1], 1.0))
	assert.Check(t, cmp.Equal(scores[file2], 1.0))
}

func TestFinalScoreTwoWords(t *testing.T) {

	words := []string{"word", "otherWord"}

	wordsCount := map[int]int{}
	file1 := 1
	wordsCount[file1] = 1
	file2 := 2
	wordsCount[file2] = 2

	scores := finalScore(words, wordsCount)

	assert.Check(t, cmp.Len(scores, 2))
	assert.Check(t, cmp.Equal(scores[file1], 0.5))
	assert.Check(t, cmp.Equal(scores[file2], 1.0))
}
