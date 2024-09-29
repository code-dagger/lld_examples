package searchengine

import (
	"fmt"
)

type dataset struct {
	name      string
	documents map[string]*document
	config    config
}

func (d *dataset) insertDocument(name, content string) error {
	// checking if the document already exist with the name
	if _, exists := d.documents[name]; exists {
		return fmt.Errorf("document already exist with the provided name '%s'", name)
	}
	// generate new document
	doc := newDocument(name, content, d.config)
	d.documents[name] = doc
	return nil
}

func (d *dataset) deleteDocument(name string) error {
	// checking if the doc with name exist or not
	if _, exists := d.documents[name]; !exists {
		return fmt.Errorf("document does not exist")
	}
	// deleting the document
	delete(d.documents, name)
	return nil
}

func (d *dataset) getDocumentCount() int {
	return len(d.documents)
}

func (d *dataset) search(searchTerm string) []*document {
	results := []*document{}
	for _, doc := range d.documents {
		if !doc.occurenceAllCheck(searchTerm) {
			continue
		}
		results = append(results, doc)
	}
	return d.config.ranker.rank(results, searchTerm)
}

func newDataSet(name string, conf config) *dataset {
	return &dataset{
		name:      name,
		documents: map[string]*document{},
		config:    conf,
	}
}
