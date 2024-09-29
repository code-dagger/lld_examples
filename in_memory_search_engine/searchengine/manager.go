package searchengine

import (
	"fmt"
	"strings"
)

type Manager struct {
	eng *engine
}

func NewManager() Manager {
	return Manager{
		eng: newEngine(),
	}
}

func (m *Manager) CreateDataSet(name string, ranker Ranker, splitter Splitter) {
	err := m.eng.createDataset(name, newConfig(ranker, splitter))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (m *Manager) DeleteDataSet(name string) {
	err := m.eng.deleteDataset(name)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (m *Manager) InsertDocument(dataset, docName, content string) {
	ds, err := m.eng.getDataset(dataset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = ds.insertDocument(docName, content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (m *Manager) DeleteDocument(dataset, docName, content string) {
	ds, err := m.eng.getDataset(dataset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = ds.deleteDocument(docName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (m *Manager) Search(dataset, searchTerm string) {
	ds, err := m.eng.getDataset(dataset)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	docs := ds.search(searchTerm)
	if len(docs) == 0 {
		fmt.Println("no document found that has the search term '%s'", searchTerm)
		return
	}
	docNames := []string{}
	for _, doc := range docs {
		docNames = append(docNames, doc.name)
	}
	fmt.Println("found the search term in docs: ", strings.Join(docNames, ", "))
}
