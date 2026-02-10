package config

import (
	"fmt"
	"strings"
)

// AccountSelector handles WeChat account selection logic
type AccountSelector struct {
	accounts       []WechatAccount
	defaultAccount string
}

// NewAccountSelector creates a new AccountSelector
func NewAccountSelector(accounts []WechatAccount, defaultAccount string) *AccountSelector {
	return &AccountSelector{
		accounts:       accounts,
		defaultAccount: defaultAccount,
	}
}

// SelectAccount selects a WeChat account based on prompt and explicit account ID
// Selection priority:
// 1. If only 1 account → use it
// 2. Explicit account ID → find by ID
// 3. Keyword match from prompt → return matched account
// 4. Default account → use default
// 5. First account → fallback
func (s *AccountSelector) SelectAccount(prompt, accountID string) (*WechatAccount, error) {
	if len(s.accounts) == 0 {
		return nil, fmt.Errorf("no WeChat accounts configured")
	}

	// Priority 1: Single account
	if len(s.accounts) == 1 {
		return &s.accounts[0], nil
	}

	// Priority 2: Explicit account (by ID or Name)
	if accountID != "" {
		// Try by ID first
		for i := range s.accounts {
			if s.accounts[i].ID == accountID {
				return &s.accounts[i], nil
			}
		}
		// Try by Name
		for i := range s.accounts {
			if s.accounts[i].Name == accountID {
				return &s.accounts[i], nil
			}
		}
		return nil, fmt.Errorf("account not found: %s", accountID)
	}

	// Priority 3: Keyword matching from prompt
	if prompt != "" {
		if account := s.matchByKeywords(prompt); account != nil {
			return account, nil
		}
	}

	// Priority 4: Default account
	if s.defaultAccount != "" {
		for i := range s.accounts {
			if s.accounts[i].ID == s.defaultAccount {
				return &s.accounts[i], nil
			}
		}
	}

	// Priority 5: First account as fallback
	return &s.accounts[0], nil
}

// matchByKeywords matches an account based on keywords in the prompt
func (s *AccountSelector) matchByKeywords(prompt string) *WechatAccount {
	promptLower := strings.ToLower(prompt)

	// Find the account with the most keyword matches
	var bestMatch *WechatAccount
	maxMatches := 0

	for i := range s.accounts {
		matches := 0
		for _, keyword := range s.accounts[i].Keywords {
			if strings.Contains(promptLower, strings.ToLower(keyword)) {
				matches++
			}
		}
		if matches > maxMatches {
			maxMatches = matches
			bestMatch = &s.accounts[i]
		}
	}

	return bestMatch
}

// GetAllAccounts returns all configured accounts
func (s *AccountSelector) GetAllAccounts() []WechatAccount {
	return s.accounts
}

// GetAccountByID returns an account by its ID
func (s *AccountSelector) GetAccountByID(id string) (*WechatAccount, error) {
	for i := range s.accounts {
		if s.accounts[i].ID == id {
			return &s.accounts[i], nil
		}
	}
	return nil, fmt.Errorf("account not found: %s", id)
}

// GetAccountByName returns an account by its display name
func (s *AccountSelector) GetAccountByName(name string) (*WechatAccount, error) {
	for i := range s.accounts {
		if s.accounts[i].Name == name {
			return &s.accounts[i], nil
		}
	}
	return nil, fmt.Errorf("account not found: %s", name)
}
