package service

import (
	"fmt"
	"sync"
)

type ratingService struct {
	surveys map[int]*survey
	mu      sync.Mutex
}

func (r *ratingService) CreateSurvey(title string) *survey {
	r.mu.Lock()
	defer r.mu.Unlock()

	survey := newSurvey(len(r.surveys)+1, title)
	r.surveys[survey.id] = survey
	fmt.Println("Draft survey created")
	return survey
}

func (r *ratingService) PublishSurvey(id int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	survey, exists := r.surveys[id]
	if !exists {
		fmt.Printf("survey with id '%d' not found\n", id)
	}
	survey.publish()
	fmt.Printf("survey published with id '%d'", id)
}

func (r *ratingService) AddQuestion(surveyId int, questionText string, options []option, rateable bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	survey, exists := r.surveys[surveyId]
	if !exists {
		fmt.Printf("survey not found '%d'", surveyId)
		return
	}
	survey.addQuestion(len(survey.questions)+1, questionText, options, rateable)
}

func NewRatingService() *ratingService {
	return &ratingService{
		surveys: make(map[int]*survey),
	}
}
