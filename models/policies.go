package models

import "github.com/google/uuid"

type CheckingCombination struct {
	Name        string   `json:"name,omitempty"`
	ProjectsIDs []string `json:"projects_ids"`
	RulesIDs    []string `json:"rules_ids"`
}

type PolicyParams struct {
	Type   string        `json:"type"` // "onCommit" или "IntervalChecking"
	Params []interface{} `json:"params"`
}

type CheckingPolicy struct {
	ID          uuid.UUID           `json:"id"`
	Combination CheckingCombination `json:"combination"`
	Policy      PolicyParams        `json:"policy"`
}
