package searchengine

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestTopNEmptyMap(t *testing.T) {
	topN := map[int]int{}
	result := topNValues(topN, 3)
	if len(result) != 0 {
		t.Errorf("Failed ! got %v want empty", result)
	} else {
		t.Logf("Success !")
	}
}

func TestTopNSingleValueMap(t *testing.T) {
	topN := map[int]int{2: 3}
	result := topNValues(topN, 3)
	assert.Check(t, cmp.Len(result, 1))
	assert.Check(t, cmp.Equal(result[2], 3))
}

func TestTopNThreeValuesMap(t *testing.T) {
	topN := map[int]int{2: 3, 1: 8, 4: 6}
	result := topNValues(topN, 3)
	assert.Check(t, cmp.Len(result, 3))
	assert.Check(t, cmp.Equal(result[2], 3))
	assert.Check(t, cmp.Equal(result[1], 8))
	assert.Check(t, cmp.Equal(result[4], 6))
}

func TestTopNTwoValuesFromThreeValuesMap(t *testing.T) {
	topN := map[int]int{2: 3, 1: 8, 4: 6}
	result := topNValues(topN, 2)
	assert.Check(t, cmp.Len(result, 2))
	assert.Check(t, cmp.Equal(result[1], 8))
	assert.Check(t, cmp.Equal(result[4], 6))
}

func TestTopNZeroValueFromThreeValuesMap(t *testing.T) {
	topN := map[int]int{2: 3, 1: 8, 4: 6}
	result := topNValues(topN, 0)
	if len(result) != 0 {
		t.Errorf("Failed ! got %v want empty", result)
	} else {
		t.Logf("Success !")
	}
}

func TestTopNWithNegativeArgument(t *testing.T) {
	topN := map[int]int{2: 3, 1: 8, 4: 6}
	result := topNValues(topN, -1)
	if len(result) != 0 {
		t.Errorf("Failed ! got %v want empty", result)
	} else {
		t.Logf("Success !")
	}
}

func TestTopNWithBoundEqualToSize(t *testing.T) {
	topN := map[int]int{2: 3, 1: 8, 4: 6, 7: 10}
	result := topNValues(topN, 4)
	assert.Check(t, cmp.Len(result, 4))
	assert.Check(t, cmp.Equal(result[7], 10))
	assert.Check(t, cmp.Equal(result[1], 8))
	assert.Check(t, cmp.Equal(result[4], 6))
	assert.Check(t, cmp.Equal(result[2], 3))
}
