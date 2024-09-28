package database

type Manager struct {
	db *database
}

func NewManager() *Manager {
	return &Manager{db: newDatabase()}
}

func (m *Manager) CreateTable(table string, columnMap map[string]string) error {
	colList := []Column{}
	for colName, colType := range columnMap {
		colList = append(colList, Column{Name: colName, Type: ColumnType(colType)})
	}
	return m.db.createTable(table, colList)
}

func (m *Manager) dropTable() error {
	return nil
}

func (m *Manager) createIndex(table, column string) error {
	return nil
}

func (m *Manager) insertRow(table string, values []string) (int, error) {
	return 0, nil
}

func (m *Manager) updateRow(table string, rowId string, values []string) error {
	return nil
}

func (m *Manager) deleteRows(table string, filter map[string]string) (int, error) {
	return 0, nil
}

func (m *Manager) getRows(table string, filter map[string]string) []Row {
	return nil
}
