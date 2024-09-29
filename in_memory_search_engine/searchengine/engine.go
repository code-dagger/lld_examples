package searchengine

import "fmt"

type engine struct {
	datasets map[string]*dataset
}

func (e *engine) createDataset(name string, conf config) error {
	if _, exists := e.datasets[name]; exists {
		return fmt.Errorf("dataset already exist the name '%s'", name)
	}
	e.datasets[name] = newDataSet(name, conf)
	return nil
}

func (e *engine) deleteDataset(name string) error {
	if _, exists := e.datasets[name]; exists {
		return fmt.Errorf("dataset already exist the name '%s'", name)
	}
	delete(e.datasets, name)
	return nil
}

func (e *engine) getDataset(name string) (*dataset, error) {
	ds, ok := e.datasets[name]
	if !ok {
		return nil, fmt.Errorf("no dataset found with the name '%s'", name)
	}
	return ds, nil
}

func newEngine() *engine {
	return &engine{
		datasets: make(map[string]*dataset),
	}
}
