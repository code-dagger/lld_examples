package database

type ColumnType string

const (
	IntType    ColumnType = "int"
	StringType ColumnType = "string"

	IntMin int = -1024
	IntMax int = 1024

	StringMax int = 20
)

type Column struct {
	Name string
	Type ColumnType
}
