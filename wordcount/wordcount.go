package wordcount

import "strings"

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := range words {
		word := string(words[i])
		_, ok := m[word]
		if ok == false {
			m[word] = 1
		} else {
			m[word] = m[word] + 1
		}
	}
	return m
}
