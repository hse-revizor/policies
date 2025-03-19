package handlers

import (
	"encoding/json"
	"net/http"
	"policies/models"
	"policies/storage"
)

type PolicyHandler struct {
	policyStorage *storage.DBPolicyStorage
}

func NewPolicyHandler(policyStorage *storage.DBPolicyStorage) *PolicyHandler {
	return &PolicyHandler{policyStorage: policyStorage}
}

type CreatePolicyRequest struct {
	Combination models.CheckingCombination `json:"combination"`
	Policy      models.PolicyParams        `json:"policy"`
}

func (h *PolicyHandler) CreatePolicy(w http.ResponseWriter, r *http.Request) {
	var req CreatePolicyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Policy.Type != "onCommit" && req.Policy.Type != "IntervalChecking" {
		http.Error(w, "Invalid policy type", http.StatusBadRequest)
		return
	}

	if req.Policy.Type == "IntervalChecking" && len(req.Policy.Params) == 0 {
		http.Error(w, "cronExpression is required for IntervalChecking", http.StatusBadRequest)
		return
	}

	policy, err := h.policyStorage.CreatePolicy(req.Combination, req.Policy)
	if err != nil {
		http.Error(w, "Failed to create policy", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(policy)
}
