package dtos

import "time"

type SaveAssessmentRequest struct {
	HealthConditionAssessmentAnswer HealthConditionAssessmentAnswer `json:"healthConditionAssessmentAnswer"`
	AuditAssessmentAnswer           AuditAssessmentAnswer           `json:"auditAssessmentAnswer"`
	GoalSettingAssessmentAnswer     GoalSettingAssessmentAnswer     `json:"goalSettingAssessmentAnswer"`
	SupportiveContact               SupportiveContact               `json:"supportiveContact"`
}

type HealthConditionAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type AuditAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type GoalSettingAssessmentAnswer struct {
	UserAnswer  []UserAnswer `json:"userAnswer"`
	CreatedDate time.Time    `json:"createdDate"`
}

type UserAnswer struct {
	QuestionID int64   `json:"questionId"`
	OptionID   int64   `json:"optionId"`
	Points     float64 `json:"points"`
	MaxPoints  int     `json:"maxPoints"`
}

type SupportiveContact struct {
	Contacts    []Contacts `json:"contacts"`
	CreatedDate time.Time  `json:"createdDate"`
}

type Contacts struct {
	Name           string `json:"name"`
	ContactNumber  string `json:"contactNumber"`
	RelationShipID int64  `json:"relationShipId"`
}
