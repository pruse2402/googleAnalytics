package dtos

// TriggerAssessmentResponse Response struct send to client
type TriggerAssessmentResponse struct {
	ID                   int64                     `json:"id"`
	QuestionOptionTypeID int64                     `json:"questionOptionTypeId"`
	Question             string                    `json:"question"`
	QuestionNo           int                       `json:"questionNo"`
	SequenceOrder        int                       `json:"sequenceOrder"`
	Options              []TriggerAssessmentOption `json:"options"`
}

type TriggerAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	SequenceOrder int     `json:"sequenceOrder"`
}
