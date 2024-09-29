package docservice

type AccessType string

const (
	AccessRead  AccessType = "read"
	AccessWrite AccessType = "write"
)

type document struct {
	name    string
	content string
	owner   string
	access  map[string]AccessType
}

func newDocument(name, content, user string) *document {
	return &document{
		name:    name,
		content: content,
		owner:   user,
		access:  make(map[string]AccessType),
	}
}
