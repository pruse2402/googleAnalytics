package dtos

// AuditAssessmentResponse Response struct send to client
type AboutPrivacyPolicy struct {
	AboutPrivacyPolicyID int64  `json:"about_privacy_policy_id"`
	VersionCode          int64  `json:"version_code"`
	VersionName          string `json:"version_name"`
	UpdatedAt            string `json:"updated_at"`
}

type AboutPrivacyPolicyResponse struct {
	AboutPrivacyPolicyID   int64                     `json:"about_privacy_policy_id"`
	VersionCode            int64                     `json:"version_code"`
	VersionName            string                    `json:"version_name"`
	UpdatedAt              string                    `json:"updated_at"`
	AboutPrivacyPolicyInfo *[]AboutPrivacyPolicyInfo `json:"privacy_policy_info"`
}

type AboutPrivacyPolicyInfo struct {
	ID          int64  `json:"id"`
	VersionCode int64  `json:"version_code"`
	VersionName string `json:"version_name"`
	SequenceNO  int64  `json:"sequence_no"`
	MessageInfo string `json:"message_info"`
	ContentType string `json:"content_type"`
}
