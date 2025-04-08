package wordcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	input1 := "hello world hello"
	input2 := "I am learning Go!"
	input3 := "The quick brown fox jumped over the lazy dog."
	input4 := "I ate a donut. Then I ate another donut."
	input5 := "A man a plan a canal panama."

	expected1 := map[string]int{"hello": 2, "world": 1}
	expected2 := map[string]int{"Go!": 1, "I": 1, "am": 1, "learning": 1}
	expected3 := map[string]int{"The": 1, "brown": 1, "dog.": 1, "fox": 1, "jumped": 1, "lazy": 1, "over": 1, "quick": 1, "the": 1}
	expected4 := map[string]int{"I": 2, "Then": 1, "a": 1, "another": 1, "ate": 2, "donut.": 2}
	expected5 := map[string]int{"A": 1, "a": 2, "canal": 1, "man": 1, "panama.": 1, "plan": 1}

	result1 := WordCount(input1)
	result2 := WordCount(input2)
	result3 := WordCount(input3)
	result4 := WordCount(input4)
	result5 := WordCount(input5)

	assert.Equal(t, expected1, result1)
	assert.Equal(t, expected2, result2)
	assert.Equal(t, expected3, result3)
	assert.Equal(t, expected4, result4)
	assert.Equal(t, expected5, result5)
}
