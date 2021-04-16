package searchengine

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestFillIndex(t *testing.T) {

	var fileId int = 1
	var word string = "word"

	var index Index
	index.init()
	index.fill(fileId, word)

	assert.Check(t, cmp.Len(index.index[word], 1))
	assert.Check(t, cmp.Contains(index.index[word], fileId))
}

func TestFillIndexWithSpecialChars(t *testing.T) {

	var fileId int = 1
	var word string = "Caraïbes"

	var index Index
	index.init()
	index.fill(fileId, word)

	var expectedWord string = "caraïbes"
	assert.Check(t, cmp.Len(index.index[word], 0))
	assert.Check(t, cmp.Len(index.index[expectedWord], 1))
	assert.Check(t, cmp.Contains(index.index[expectedWord], fileId))
}

func TestFillIndexWithOtherWordAndSameFile(t *testing.T) {

	var fileId int = 1
	var firstWord string = "firstWord"

	var index Index
	index.init()
	index.index[firstWord] = []int{fileId}

	var secondWord string = "secondWord"
	index.fill(fileId, secondWord)

	assert.Check(t, cmp.Len(index.index[firstWord], 1))
	assert.Check(t, cmp.Contains(index.index[firstWord], fileId))

	var expectedWord string = "secondword"
	assert.Check(t, cmp.Len(index.index[secondWord], 0))
	assert.Check(t, cmp.Len(index.index[expectedWord], 1))
	assert.Check(t, cmp.Contains(index.index[expectedWord], fileId))
}

func TestFillIndexWithSameWordAndOtherFile(t *testing.T) {

	var firstFile int = 1
	var word string = "word"

	var index Index
	index.init()
	index.index[word] = []int{firstFile}

	var secondFile int = 2
	index.fill(secondFile, word)

	assert.Check(t, cmp.Len(index.index[word], 2))
	assert.Check(t, cmp.Contains(index.index[word], firstFile))
	assert.Check(t, cmp.Contains(index.index[word], secondFile))
}

func TestFillIndexWithSameWordMultipleTimes(t *testing.T) {

	var firstFile int = 1
	var word string = "word"

	var index Index
	index.init()
	index.index[word] = []int{firstFile, firstFile}

	index.fill(firstFile, word)

	assert.Check(t, cmp.Len(index.index[word], 3))
	assert.Check(t, cmp.Contains(index.index[word], firstFile))
}

func TestFillFileName(t *testing.T) {

	var fileId int = 1
	var fileName string = "file-name"

	var index Index
	index.init()
	index.fillFileName(fileId, fileName)

	assert.Check(t, cmp.Len(index.index[fileName], 1))
	assert.Check(t, cmp.Contains(index.index[fileName], fileId))
}

func TestFillFileNameWithUpperCase(t *testing.T) {

	var fileId int = 1
	var fileName string = "fileName"

	var index Index
	index.init()
	index.fillFileName(fileId, fileName)

	var expectedFileName string = "filename"
	assert.Check(t, cmp.Len(index.index[fileName], 0))
	assert.Check(t, cmp.Len(index.index[expectedFileName], 1))
	assert.Check(t, cmp.Contains(index.index[expectedFileName], fileId))
}

func TestFillFileNameWithSeparator(t *testing.T) {

	var fileId int = 1
	var fileName string = "dir/File.txt"

	var index Index
	index.init()
	index.fillFileName(fileId, fileName)

	var expectedToken string = "file"
	assert.Check(t, cmp.Len(index.index, 1))
	assert.Check(t, cmp.Len(index.index[expectedToken], 1))
	assert.Check(t, cmp.Contains(index.index[expectedToken], fileId))
}

func TestFillFileNameWithSameNameAndOtherFile(t *testing.T) {

	var firstFile int = 1

	var index Index
	index.init()

	var expectedToken string = "file"
	index.index[expectedToken] = []int{firstFile}

	var secondFile int = 1
	index.fillFileName(secondFile, "dir/File.txt")

	assert.Check(t, cmp.Len(index.index, 1))
	assert.Check(t, cmp.Len(index.index[expectedToken], 2))
	assert.Check(t, cmp.Contains(index.index[expectedToken], firstFile))
	assert.Check(t, cmp.Contains(index.index[expectedToken], secondFile))
}
