package dtos

type PatientAccessCodeReq struct {
	SolutionType string `json:"solutionType"`
	TimeZone     string `json:"timeZone"`
	AccessCode   string `json:"accessCode"`
}

type PatientAccessCodeResponse struct {
	Message string `json:"message"`
}
