package dtos

// DrinkHabitAssessmentResponse Response struct send to client
type DrinkHabitAssessmentResponse struct {
	DrinkProfiles       []DrinkProfiles       `json:"drinkProfiles"`
	DrinkHabitQuestions []DrinkHabitQuestions `json:"drinkHabitQuestions"`
}

type DrinkProfiles struct {
	DrinkID       int64  `json:"drinkId"`
	Name          string `json:"name"`
	SequenceOrder int    `json:"sequenceOrder"`
}

type DrinkHabitQuestions struct {
	ID                   int64                        `json:"id"`
	QuestionOptionTypeID int64                        `json:"questionOptionTypeId"`
	Question             string                       `json:"question"`
	QuestionNo           int                          `json:"questionNo"`
	SequenceOrder        int                          `json:"sequenceOrder"`
	Options              []DrinkHabitAssessmentOption `json:"options"`
}

type DrinkHabitAssessmentOption struct {
	ID            int64   `json:"id"`
	QuestionID    int64   `json:"questionId"`
	Name          string  `json:"name"`
	Points        float64 `json:"points"`
	MaxPoints     int     `json:"maxPoints"`
	SequenceOrder int     `json:"sequenceOrder"`
}
