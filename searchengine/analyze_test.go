package searchengine

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestEmptyWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze(""), ""))
}

func TestLowerCaseWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze("token"), "token"))
}

func TestFirstCharUpperCaseWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze("Token"), "token"))
}

func TestUpperCaseWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze("TOKEN"), "token"))
}

func TestMixUpperCaseWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze("tOKen"), "token"))
}

func TestSpecialCharsWord(t *testing.T) {
	assert.Check(t, cmp.Equal(analyze("Caché"), "caché"))
}

func TestAnalyzeEmptyFileName(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName(""), ""))
}

func TestAnalyzeSimpleFileName(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("file"), "file"))
}

func TestAnalyzeFileNameWithFirstCharUpperCase(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("File"), "file"))
}

func TestAnalyzeFileNameUpperCase(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("FILE"), "file"))
}

func TestAnalyzeFileNameMixUpperCase(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("fILe"), "file"))
}

func TestAnalyzeSpecialCharsFileName(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("filé"), "filé"))
}

func TestAnalyzeFileNameInDirectory(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("dir/file"), "file"))
}

func TestAnalyzeFileNameWithExtension(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("file.txt"), "file"))
}

func TestAnalyzeFileNameComplexCase(t *testing.T) {
	assert.Check(t, cmp.Equal(analyzeFileName("dir/SubDir/File.go"), "file"))
}
