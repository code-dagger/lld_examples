package searchengine

import (
	"regexp"
	"strings"
)

type Splitter interface {
	split(text string) []string
}

type WhiteSpaceSplitter struct{}

func (w *WhiteSpaceSplitter) split(text string) []string {
	return strings.Fields(text)
}

type PunctuationSplitter struct{}

func (p *PunctuationSplitter) split(text string) []string {
	return regexp.MustCompile(`\W+`).Split(text, -1)
}
