package searchengine

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

const test_dir string = "./test_dir/"
const test_sub_dir string = "./test_sub_dir/"

func TestFill(t *testing.T) {

	var files Files
	files.init()
	files.fill("", test_dir, doNothingOnFileName, doNothing)

	assert.Check(t, cmp.Len(files.filesNames, 3))
	assert.Check(t, cmp.Equal(files.filesNames[0], "file1.txt"))
	assert.Check(t, cmp.Equal(files.filesNames[1], "file2.txt"))
	assert.Check(t, cmp.Equal(files.filesNames[2], "file3.txt"))
}

func TestFillAndProcessFileName(t *testing.T) {

	testResult = ""

	var files Files
	files.init()
	files.fill("", test_dir, copyFileName, doNothing)

	assert.Check(t, cmp.Len(files.filesNames, 3))

	expectedResult := "file1.txt file2.txt file3.txt "
	if testResult != expectedResult {
		t.Errorf("Failed ! got %v want %v", testResult, expectedResult)
	} else {
		t.Logf("Success !")
	}
}

func TestFillFromSubDirectories(t *testing.T) {

	var files Files
	files.init()
	files.fill("", test_sub_dir, doNothingOnFileName, doNothing)

	if len(files.filesNames) != 5 {
		t.Errorf("Failed ! got %v want %v", len(files.filesNames), 5)
	} else {
		t.Logf("Success !")
	}

	var keys = make([]int, 0, len(files.filesNames))
	var values = make([]string, 0, len(files.filesNames))
	for key, value := range files.filesNames {
		keys = append(keys, key)
		values = append(values, value)
	}

	assert.Check(t, cmp.Contains(keys, 0))
	assert.Check(t, cmp.Contains(keys, 1))
	assert.Check(t, cmp.Contains(keys, 2))
	assert.Check(t, cmp.Contains(keys, 3))
	assert.Check(t, cmp.Contains(keys, 4))

	assert.Check(t, cmp.Contains(values, "file11.txt"))
	assert.Check(t, cmp.Contains(values, "sub_dir1/file21.txt"))
	assert.Check(t, cmp.Contains(values, "sub_dir1/file22.txt"))
	assert.Check(t, cmp.Contains(values, "sub_dir2/file31.txt"))
	assert.Check(t, cmp.Contains(values, "sub_dir2/file32.txt"))
}

func TestFillAndProcessFileNameFromSubDirectories(t *testing.T) {

	testResult = ""

	var files Files
	files.init()
	files.fill("", test_sub_dir, copyFileName, doNothing)

	if len(files.filesNames) != 5 {
		t.Errorf("Failed ! got %v want %v", len(files.filesNames), 5)
	} else {
		t.Logf("Success !")
	}

	assert.Check(t, cmp.Contains(testResult, "file11.txt"))
	assert.Check(t, cmp.Contains(testResult, "sub_dir1/file21.txt"))
	assert.Check(t, cmp.Contains(testResult, "sub_dir1/file22.txt"))
	assert.Check(t, cmp.Contains(testResult, "sub_dir2/file31.txt"))
	assert.Check(t, cmp.Contains(testResult, "sub_dir2/file32.txt"))

	replacedResult := strings.Replace(testResult, "file11.txt", "", 1)
	replacedResult = strings.Replace(replacedResult, "sub_dir1/file21.txt", "", 1)
	replacedResult = strings.Replace(replacedResult, "sub_dir1/file22.txt", "", 1)
	replacedResult = strings.Replace(replacedResult, "sub_dir2/file31.txt", "", 1)
	replacedResult = strings.Replace(replacedResult, "sub_dir2/file32.txt", "", 1)
	replacedResult = strings.Trim(replacedResult, " ")
	if replacedResult != "" {
		t.Errorf("Failed ! got %v want empty", replacedResult)
	} else {
		t.Logf("Success !")
	}
}

func doNothingOnFileName(fileId int, fileName string) {
}

func doNothing(fileId int, word string) {
}

func TestProcessFileWords(t *testing.T) {
	processFileWords(test_dir, "file1.txt", 0, doNothing)
}

func TestProcessFileWordsInSubDirectory(t *testing.T) {
	processFileWords(test_sub_dir, "sub_dir1/file21.txt", 0, doNothing)
}

var testResult string

func copyFileName(fileId int, fileName string) {
	testResult = testResult + fileName + " "
}

func concatenateFile(fileId int, word string) {
	testResult += word
}

func TestGetFileWordsWithProcessor(t *testing.T) {

	testResult = ""
	processFileWords(test_dir, "file1.txt", 0, concatenateFile)

	expected := "Thereisadog"
	if testResult != expected {
		t.Errorf("Failed ! got %s want %s", testResult, expected)
	} else {
		t.Logf("Success !")
	}
}

func TestGetFileWordsInSubDirectoriesWithProcessor(t *testing.T) {

	testResult = ""
	processFileWords(test_sub_dir, "sub_dir1/file22.txt", 0, concatenateFile)
	expected := "LachartreusedeParme"
	if testResult != expected {
		t.Errorf("Failed ! want %s containing %s", testResult, expected)
	} else {
		t.Logf("Success !")
	}
}
