package dtos

// AuditAssessmentResponse Response struct send to client
type AuditAssessmentResponse struct {
	ID                   int64                   `json:"id"`
	QuestionOptionTypeID int64                   `json:"questionOptionTypeId"`
	Question             string                  `json:"question"`
	QuestionNo           int                     `json:"questionNo"`
	SequenceOrder        int                     `json:"sequenceOrder"`
	Options              []AuditAssessmentOption `json:"options"`
}

type AuditAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	MaxPoints     int     `json:"maxPoints"`
	SequenceOrder int     `json:"sequenceOrder"`
}
