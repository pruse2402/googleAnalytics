package dtos

// ReasonAssessmentResponse Response struct send to client
type ReasonAssessmentResponse struct {
	ID                   int64                    `json:"id"`
	QuestionOptionTypeID int64                    `json:"questionOptionTypeId"`
	Question             string                   `json:"question"`
	QuestionNo           int                      `json:"questionNo"`
	SequenceOrder        int                      `json:"sequenceOrder"`
	HeaderNote           string                   `json:"headerNote"`
	Options              []ReasonAssessmentOption `json:"options"`
}

type ReasonAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	MaxPoints     float64 `json:"maxPoints"`
	SequenceOrder int     `json:"sequenceOrder"`
}
