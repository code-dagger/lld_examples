package main

import (
	se "github.com/code-dagger/in-mem-search-engine/searchengine"
)

func main() {
	mgr := se.NewManager()
	mgr.CreateDataSet("apple_search", &se.RecencyRanker{}, &se.PunctuationSplitter{})
	mgr.InsertDocument("apple_search", "doc1", "apple is a fruit")
	mgr.InsertDocument("apple_search", "doc2", "apple, apple come on!")
	mgr.InsertDocument("apple_search", "doc3", "oranges are sour")
	mgr.InsertDocument("apple_search", "doc4", "apple-pie is sweet")
	mgr.Search("apple_search", "apple")
}
