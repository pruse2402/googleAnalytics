package dtos

// ACBehaviourChangeResponse Response struct send to client
type ACBehaviourChange struct {
	BehaviourChangeID int64  `json:"behaviour_change_id"`
	BctTaxonomyID     string `json:"bct_taxonomy_id"`
	BctTaxonomy       string `json:"bct_taxonomy"`
	BctID             string `json:"bct_id"`
	BctDescription    string `json:"bct_description"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
}

type ACBehaviourChangeResponse struct {
	BehaviourChange *[]ACBehaviourChange `json:"behaviour_change"`
}
