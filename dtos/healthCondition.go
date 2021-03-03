package dtos

import "time"

// HealthConditionAssessmentResponse Response struct send to client
type HealthConditionAssessmentResponse struct {
	ID                   int64                             `json:"id"`
	QuestionOptionTypeID int64                             `json:"questionOptionTypeId"`
	Question             string                            `json:"question"`
	QuestionNo           int                               `json:"questionNo"`
	SequenceOrder        int                               `json:"sequenceOrder"`
	Options              []HealthConditionAssessmentOption `json:"options"`
	LatestHistory        LatestHistory                     `json:"latestHistory"`
}

type HealthConditionAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	MaxPoints     int     `json:"maxPoints"`
	SequenceOrder int     `json:"sequenceOrder"`
}

type LatestHistory struct {
	UserID      int64     `json:"userId"`
	DateCreated time.Time `json:"dateCreated"`
	LastUpdated time.Time `json:"lastUpdated"`
	Answers     []Answers `json:"answers"`
}

type Answers struct {
	ID         int64 `json:"id"`
	QuestionID int64 `json:"questionId"`
	OptionID   int64 `json:"optionId"`
	Points     int   `json:"points"`
}
