package dtos

// ACBehaviourChangeResponse Response struct send to client
type ACBehaviourChangeResponse struct {
	BehaviourChangeID int64  `json:"behaviour_change_id"`
	BctTaxonomyID     string `json:"bct_taxonomy_id"`
	BctTaxonomy       string `json:"bct_taxonomy"`
	BctID             string `json:"bct_id"`
	BctDescription    string `json:"bct_description"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}
