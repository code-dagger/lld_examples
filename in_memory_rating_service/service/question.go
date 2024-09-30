package service

type option struct {
	text   string
	weight int
}

func NewOption(text string, weight int) option {
	return option{
		text:   text,
		weight: weight,
	}
}

type question struct {
	id       int
	text     string
	options  []option
	rateable bool
	// todo: add support for conditional question
}

func newQuestion(id int, text string, options []option, rateable bool) *question {
	return &question{
		id:       id,
		text:     text,
		options:  options,
		rateable: rateable,
	}
}
