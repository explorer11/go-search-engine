package searchengine

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestFilesScoresNoResult(t *testing.T) {

	filesNames := map[int]string{0: "file1", 1: "file2"}
	scores := map[int]float64{}

	filesScores := filesScores(filesNames, scores)

	expectedLength := 0
	if len(filesScores) != expectedLength {
		t.Errorf("Failed ! got %v scores want %v", len(filesScores), expectedLength)
	} else {
		t.Logf("Success !")
	}
}

func TestFilesScoresOneResultForFirstFile(t *testing.T) {

	file0 := "file0"
	filesNames := map[int]string{0: file0, 1: "file1"}
	scores := map[int]float64{0: 50}

	filesScores := filesScores(filesNames, scores)

	assert.Check(t, cmp.Len(filesScores, 1))
	assert.Check(t, cmp.Equal(filesScores[file0], 50.0))
}

func TestFilesScoresOneResultForOtherFile(t *testing.T) {

	file2 := "file2"
	filesNames := map[int]string{0: "file0", 1: "file1", 2: file2}
	scores := map[int]float64{2: 30}

	filesScores := filesScores(filesNames, scores)

	assert.Check(t, cmp.Len(filesScores, 1))
	assert.Check(t, cmp.Equal(filesScores[file2], 30.0))
}

func TestFilesScoresTwoResults(t *testing.T) {

	file0 := "file0"
	file1 := "file1"
	filesNames := map[int]string{0: file0, 1: file1}
	scores := map[int]float64{0: 100, 1: 50}

	filesScores := filesScores(filesNames, scores)

	assert.Check(t, cmp.Len(filesScores, 2))
	assert.Check(t, cmp.Equal(filesScores[file0], 100.0))
	assert.Check(t, cmp.Equal(filesScores[file1], 50.0))
}

func TestFilesScoresComplexCase(t *testing.T) {

	file4 := "file4"
	file1 := "file1"
	file2 := "file2"
	file5 := "file5"
	file3 := "file3"
	filesNames := map[int]string{0: "file0", 1: file1, 2: file2, 3: file3, 4: file4, 5: file5, 6: "file6"}
	scores := map[int]float64{4: 80, 1: 50, 2: 50, 5: 20, 3: 10}

	filesScores := filesScores(filesNames, scores)

	assert.Check(t, cmp.Len(filesScores, 5))
	assert.Check(t, cmp.Equal(filesScores[file4], 80.0))
	assert.Check(t, cmp.Equal(filesScores[file1], 50.0))
	assert.Check(t, cmp.Equal(filesScores[file2], 50.0))
	assert.Check(t, cmp.Equal(filesScores[file5], 20.0))
	assert.Check(t, cmp.Equal(filesScores[file3], 10.0))
}
