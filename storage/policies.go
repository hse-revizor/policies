package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"policies/models"
)

type DBPolicyStorage struct {
	db *sql.DB
}

func NewDBPolicyStorage(db *sql.DB) *DBPolicyStorage {
	return &DBPolicyStorage{db: db}
}

func (s *DBPolicyStorage) CreatePolicy(combination models.CheckingCombination, policy models.PolicyParams) (*models.CheckingPolicy, error) {
	id := uuid.New()

	combinationJSON, err := json.Marshal(combination)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal combination: %w", err)
	}

	policyJSON, err := json.Marshal(policy)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal policy: %w", err)
	}

	_, err = s.db.Exec(
		"INSERT INTO checking_policies (id, combination, policy) VALUES ($1, $2, $3)",
		id, combinationJSON, policyJSON,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert policy: %w", err)
	}

	return &models.CheckingPolicy{
		ID:          id,
		Combination: combination,
		Policy:      policy,
	}, nil
}
