package dtos

// ACBehaviourChangeResponse Response struct send to client
type ACBehaviourChangeNotification struct {
	BehaviourNotificationID      int64  `json:"behaviour_notification_id"`
	BcnCategory                  string `json:"bcn_category"`
	BcnCategoryDesc              string `json:"bcn_category_desc"`
	BcnGroup                     string `json:"bcn_group"`
	BcnGroupDescription          string `json:"bcn_group_description"`
	BcnTriggerEvent              string `json:"bcn_trigger_event"`
	AppRoute                     string `json:"app_route"`
	BcnID                        string `json:"bcn_id"`
	BcnMessage                   string `json:"bcn_message"`
	Bct1                         string `json:"bct_1"`
	Bct2                         string `json:"bct_2"`
	Bct3                         string `json:"bct_3"`
	Bct4                         string `json:"bct_4"`
	AlcochangeTheme              bool   `json:"alcochange_theme"`
	FramesFeedback               bool   `json:"frames_feedback"`
	FramesResponsibility         bool   `json:"frames_responsibility"`
	FramesAdvice                 bool   `json:"frames_advice"`
	FramesMenu                   bool   `json:"frames_menu"`
	FramesEmpathy                bool   `json:"frames_empathy"`
	FramesSupportAndSelfefficacy bool   `json:"frames_support_and_selfefficacy"`
	DevelopDiscrepancy           bool   `json:"develop_discrepancy"`
	Assessment                   bool   `json:"assessment"`
	CreatedAt                    string `json:"created_at,omitempty"`
	UpdatedAt                    string `json:"updated_at,omitempty"`
}

type ACBehaviourChangeNotificationResponse struct {
	BehaviourChangeNotification *[]ACBehaviourChangeNotification `json:"behaviour_change_notification"`
}
