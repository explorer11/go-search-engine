package searchengine

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestSearchInitIndex(t *testing.T) {

	InitIndex(test_dir)

	assert.Check(t, cmp.Len(index.index, 13))
	assert.Check(t, cmp.Len(index.index["seventy"], 1))
	assert.Check(t, cmp.Contains(index.index["seventy"], 1))
	assert.Check(t, cmp.Len(index.index["golang"], 1))
	assert.Check(t, cmp.Contains(index.index["golang"], 2))
	assert.Check(t, cmp.Len(index.index["file1"], 1))
	assert.Check(t, cmp.Contains(index.index["file1"], 0))
	assert.Check(t, cmp.Len(index.index["file2"], 1))
	assert.Check(t, cmp.Contains(index.index["file2"], 1))
}

func TestSearchInitIndexWithSubDirectories(t *testing.T) {

	InitIndex(test_sub_dir)

	assert.Check(t, cmp.Len(index.index, 20))

	fileIds1 := index.index["chartreuse"]
	fileId1 := fileIds1[0]
	fileIds2 := index.index["parme"]
	fileId2 := fileIds2[0]
	fileIds3 := index.index["file22"]
	fileId3 := fileIds3[0]
	assert.Check(t, cmp.Len(fileIds1, 1))
	assert.Check(t, cmp.Len(fileIds2, 1))
	assert.Check(t, cmp.Len(fileIds3, 1))
	assert.Check(t, cmp.Equal(fileId1, fileId2))
	assert.Check(t, cmp.Equal(fileId1, fileId3))

	assert.Check(t, cmp.Len(index.index["lettres"], 1))
}

func TestQueryIndex(t *testing.T) {

	InitIndex(test_dir)

	scores := QueryIndex("dog")

	assert.Check(t, cmp.Len(scores, 1))
	assert.Check(t, cmp.Contains(scores, "file1.txt"))
	assert.Check(t, cmp.Equal(scores["file1.txt"], 1.0))
}

func TestQueryIndexWithFileName(t *testing.T) {

	InitIndex(test_dir)

	scores := QueryIndex("file3")

	assert.Check(t, cmp.Len(scores, 1))
	assert.Check(t, cmp.Contains(scores, "file3.txt"))
	assert.Check(t, cmp.Equal(scores["file3.txt"], 1.0))
}

func TestQueryIndexWithSeveralWords(t *testing.T) {

	InitIndex(test_dir)

	scores := QueryIndex("dog seventy")

	assert.Check(t, cmp.Len(scores, 2))
	assert.Check(t, cmp.Contains(scores, "file1.txt"))
	assert.Check(t, cmp.Equal(scores["file1.txt"], 0.5))
	assert.Check(t, cmp.Contains(scores, "file2.txt"))
	assert.Check(t, cmp.Equal(scores["file2.txt"], 0.5))
}

func TestQueryIndexWithSeveralWordsAndFileName(t *testing.T) {

	InitIndex(test_dir)

	scores := QueryIndex("dog seventy file1 file3")

	assert.Check(t, cmp.Len(scores, 3))
	assert.Check(t, cmp.Contains(scores, "file1.txt"))
	assert.Check(t, cmp.Equal(scores["file1.txt"], 0.5))
	assert.Check(t, cmp.Contains(scores, "file2.txt"))
	assert.Check(t, cmp.Equal(scores["file2.txt"], 0.25))
	assert.Check(t, cmp.Contains(scores, "file3.txt"))
	assert.Check(t, cmp.Equal(scores["file3.txt"], 0.25))
}
