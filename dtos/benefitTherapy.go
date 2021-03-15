package dtos

type BenefitTherapy struct {
	ID        int64  `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type BenefitTherapyResponse struct {
	Title          string            `json:"title"`
	Header         string            `json:"header"`
	ButtonText     string            `json:"button_text"`
	BenefitTherapy *[]BenefitTherapy `json:"benefits_list"`
}
