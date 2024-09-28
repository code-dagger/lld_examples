package database

import (
	"errors"
	"sync"
)

type database struct {
	tables map[string]*Table
	dbLock *sync.RWMutex
}

func (d *database) createTable(name string, columnList []Column) error {
	d.dbLock.Lock()
	defer d.dbLock.Unlock()

	if _, exist := d.tables[name]; exist {
		return errors.New("table already exists")
	}
	table := &Table{
		Name:         name,
		Columns:      columnList,
		Rows:         make(map[int]Row),
		Indexes:      make(map[string]map[string][]int),
		LastInsertID: 0,
		tableLock:    d.dbLock,
	}
	d.tables[name] = table
	return nil
}

func (d *database) dropTable(name string) error {
	d.dbLock.Lock()
	defer d.dbLock.Unlock()
	if _, exist := d.tables[name]; exist {
		return errors.New("table does not exists")
	}
	delete(d.tables, name)
	return nil
}

func (d *database) getTable(name string) (*Table, error) {
	d.dbLock.Lock()
	defer d.dbLock.Unlock()

	if _, exist := d.tables[name]; exist {
		return nil, errors.New("table does not exists")
	}
	return d.tables[name], nil
}
