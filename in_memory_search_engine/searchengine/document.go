package searchengine

import (
	"strings"
	"time"
)

type document struct {
	name    string
	content string
	created time.Time
	config  config
}

func (d *document) occurenceAllCheck(searchTerm string) bool {
	// splitting the document content based on the splitter
	docWords := d.config.splitter.split(d.content)
	// storing all the words in the wordSet
	wordSet := make(map[string]bool)
	for _, word := range docWords {
		wordSet[strings.ToLower(word)] = true
	}
	// splitting the search terms
	terms := strings.Split(searchTerm, " ")
	// checking for all terms
	for _, term := range terms {
		if _, exists := wordSet[strings.ToLower(term)]; !exists {
			return false
		}
	}
	return true // all terms are found
}

func (d *document) occurenceCount(searchTerm string) int {
	count := 0
	// splitting the document content based on the splitter
	docWords := d.config.splitter.split(strings.ToLower(d.content))
	// splitting the search terms
	terms := strings.Split(strings.ToLower(searchTerm), " ")

	for _, term := range terms {
		for _, word := range docWords {
			if word != term {
				continue
			}
			count++
		}
	}
	return count
}

func newDocument(name, content string, config config) *document {
	return &document{
		name:    name,
		content: content,
		created: time.Now(),
		config:  config,
	}
}
