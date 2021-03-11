package dtos

// ACPatientEngagementReminder Response struct from client to server
type ACPatientEngagementReminder struct {
	ID                    int64                   `json:"id"`
	UserID                int64                   `json:"user_id"`
	UserUuid              string                  `json:"user_uuid"`
	ScheduledIntervention []ScheduledIntervention `json:"scheduledInterventions"`
	EngagedIntervention   []EngagedIntervention   `json:"engagedInterventions"`
}

type ScheduledIntervention struct {
	InterventionTypeID    int64  `json:"intervention_type_id"`
	NotificationID        int64  `json:"notification_id"`
	PatientEngagementTime string `json:"patient_engagement_time"`
	MessageShown          string `json:"message_shown"`
	UserAction            string `json:"user_action"`
	CreatedAt             string `json:"created_at,omitempty"`
	UpdatedAt             string `json:"updated_at,omitempty"`
}

type EngagedIntervention struct {
	InterventionTypeID    int64  `json:"intervention_type_id"`
	NotificationID        int64  `json:"notification_id"`
	PatientEngagementTime string `json:"patient_engagement_time"`
	MessageShown          string `json:"message_shown"`
	UserAction            string `json:"user_action"`
	CreatedAt             string `json:"created_at,omitempty"`
	UpdatedAt             string `json:"updated_at,omitempty"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}
