package database

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/code-dagger/in-mem-sql-db/utils"
)

type Table struct {
	Name         string
	Columns      []Column
	Rows         map[int]Row
	Indexes      map[string]map[string][]int
	LastInsertID int

	tableLock *sync.RWMutex
}

func (t *Table) validateValues(values []string) error {
	for idx, col := range t.Columns {
		val := values[idx]
		switch col.Type {
		case IntType:
			num, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("invalid int value for column '%s'", col.Name)
			}
			if num < -IntMin || num > IntMax {
				return fmt.Errorf("int value for column '%s' out of range (-1024 to 1024)", col.Name)
			}
		case StringType:
			if len(val) > StringMax {
				return fmt.Errorf("string value for column '%s' exceeds %d characters", col.Name, StringMax)
			}
		default:
			return fmt.Errorf("invalid column type '%s' for column '%s'", col.Type, col.Name)
		}
	}
	return nil
}

func (t *Table) InsertRow(values []string) (int, error) {
	// acquiring lock
	t.tableLock.Lock()
	defer t.tableLock.Unlock()
	// checking the values len with the schema column list len
	if len(t.Columns) != len(values) {
		return 0, errors.New("mismatched column count")
	}
	// validating the values provided
	err := t.validateValues(values)
	if err != nil {
		return 0, err
	}
	// getting the formatted rowValue
	rowValues := make(map[string]interface{})
	for idx, col := range t.Columns {
		switch col.Type {
		case IntType:
			num, _ := strconv.Atoi(values[idx])
			rowValues[col.Name] = num
		case StringType:
			rowValues[col.Name] = values[idx]
		}
	}
	// generating the row object
	row := Row{
		ID:     t.LastInsertID + 1,
		Values: rowValues,
	}
	// storing the row object
	t.Rows[row.ID] = row
	// updating the last insert id count
	t.LastInsertID = row.ID
	// updating index
	for _, col := range t.Columns {
		// update index for the column if the index exist for column
		index, indexed := t.Indexes[col.Name]
		if !indexed {
			continue
		}
		ind, valIndexed := index[fmt.Sprint(row.Values[col.Name])]
		if !valIndexed {
			ind = make([]int, 0)
		}
		ind = append(ind, row.ID)
		t.Indexes[col.Name][fmt.Sprint(row.Values[col.Name])] = ind
	}

	return t.LastInsertID, nil
}

func (t *Table) UpdateRow(rowId int, updates map[string]string) error {
	// acquiring lock
	t.tableLock.Lock()
	defer t.tableLock.Unlock()
	// checking if the row exist with ID and getting the row data
	row, exists := t.Rows[rowId]
	if !exists {
		return fmt.Errorf("row does not exist")
	}
	// generating updateValues slice of string to be validated using validateValues method
	updateValues := []string{}
	for _, col := range t.Columns {
		if newVal, ok := updates[col.Name]; ok {
			updateValues = append(updateValues, newVal)
		} else {
			updateValues = append(updateValues, fmt.Sprint(row.Values[col.Name]))
		}
	}
	// validating the updateValues
	err := t.validateValues(updateValues)
	if err != nil {
		return err
	}
	// updating the values of row
	rowValues := row.Values
	for idx, col := range t.Columns {
		oldValue := rowValues[col.Name]
		if oldValue == updateValues[idx] {
			continue
		}
		switch col.Type {
		case IntType:
			num, _ := strconv.Atoi(updateValues[idx])
			rowValues[col.Name] = num
		case StringType:
			rowValues[col.Name] = updateValues[idx]
		}

		// update index for the column if the index exist for column
		index, indexed := t.Indexes[col.Name]
		if !indexed {
			continue
		}
		ind, valIndexed := index[fmt.Sprint(row.Values[col.Name])]
		if !valIndexed {
			ind = make([]int, 0)
		}
		ind = append(ind, row.ID)
		t.Indexes[col.Name][fmt.Sprint(row.Values[col.Name])] = ind

		// remove old value index which is now changed
		v := t.Indexes[col.Name][fmt.Sprint(row.Values[col.Name])]
		t.Indexes[col.Name][fmt.Sprint(row.Values[col.Name])] = utils.PullFirstValueFromArray(v, row.ID)
	}
	return nil
}

func (t *Table) DeleteRows(deleteBy map[string]string) int {
	t.tableLock.Lock()
	defer t.tableLock.Unlock()

	deletedCount := 0

	for rowId, row := range t.Rows {
		match := true
		for col, val := range deleteBy {
			if fmt.Sprint(row.Values[col]) != val {
				match = false
				break
			}
		}
		if !match {
			continue
		}
		// delete the rows
		delete(t.Rows, rowId)
		deletedCount++
		// remove the index
		for col, val := range deleteBy {
			index, indexed := t.Indexes[col]
			if !indexed {
				continue
			}
			ind, valIndexed := index[val]
			if !valIndexed {
				continue
			}
			// remove old value index which is now changed
			t.Indexes[col][val] = utils.PullFirstValueFromArray(ind, row.ID)
		}
	}
	return deletedCount
}

func (t *Table) searchByIndex(column, value string) ([]Row, bool) {
	if len(t.Indexes) == 0 {
		return nil, false
	}
	indexMap, indexed := t.Indexes[column]
	if !indexed {
		return nil, false
	}
	rowIdList, valFound := indexMap[fmt.Sprint(value)]
	if !valFound {
		return nil, true
	}
	result := []Row{}
	for _, rowId := range rowIdList {
		result = append(result, t.Rows[rowId])
	}
	return result, true
}

func (t *Table) searchInData(column, value string) []Row {
	result := []Row{}
	for _, row := range t.Rows {
		if fmt.Sprint(row.Values[column]) != value {
			continue
		}
		result = append(result, row)
	}
	return result
}

func (t *Table) ReadRows(findBy map[string]string) []Row {
	t.tableLock.RLock()
	defer t.tableLock.Unlock()

	result := []Row{}
	for col, val := range findBy {
		rows, indexed := t.searchByIndex(col, val)
		if !indexed {
			rows = t.searchInData(col, val)
		}
		if len(rows) == 0 {
			continue
		}
		result = append(result, rows...)
	}
	return nil
}

func (t *Table) CreateIndex(column string) error {
	t.tableLock.Lock()
	defer t.tableLock.Unlock()

	colFound := false
	for _, col := range t.Columns {
		if col.Name == column {
			colFound = true
			break
		}
	}

	if !colFound {
		return fmt.Errorf("column does not exist")
	}

	index := map[string][]int{}
	for rowId, row := range t.Rows {
		val := fmt.Sprint(row.Values[column])
		index[val] = append(index[val], rowId)
	}
	t.Indexes[column] = index
	return nil
}
