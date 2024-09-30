package service

import "fmt"

type SurveyStatus string

const (
	Draft     SurveyStatus = "draft"
	Published SurveyStatus = "published"
)

type survey struct {
	id        int
	title     string
	status    SurveyStatus
	questions []*question
	responses map[int]bool
}

func (s *survey) addQuestion(id int, text string, options []option, rateable bool) {
	s.questions = append(s.questions, newQuestion(id, text, options, rateable))
}

func (s *survey) publish() error {
	if s.status != Draft {
		return fmt.Errorf("only draft survey can be published")
	}
	s.status = Published
	return nil
}

func newSurvey(id int, title string) *survey {
	return &survey{
		id:        id,
		title:     title,
		status:    Draft,
		questions: make([]*question, 0),
	}
}
