package config

import (
	"testing"
)

func TestAccountSelector_SelectDefault(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}
}

func TestAccountSelector_SelectByKeyword(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:       "test-1",
			Name:     "测试账号1",
			AppID:    "wx123456",
			Secret:   "secret123",
			Keywords: []string{"测试", "demo"},
		},
		{
			ID:       "test-2",
			Name:     "测试账号2",
			AppID:    "wx789012",
			Secret:   "secret456",
			Keywords: []string{"生产", "prod"},
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("这是一篇关于测试的文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}

	account, err = selector.SelectAccount("这是生产环境的文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-2" {
		t.Errorf("Account ID = %v, want test-2", account.ID)
	}
}

func TestAccountSelector_NoMatch(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:       "test-1",
			Name:     "测试账号1",
			AppID:    "wx123456",
			Secret:   "secret123",
			Keywords: []string{"测试"},
		},
		{
			ID:       "test-2",
			Name:     "测试账号2",
			AppID:    "wx789012",
			Secret:   "secret456",
			Keywords: []string{"生产"},
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("这是一篇普通文章，没有关键词", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1 (default)", account.ID)
	}
}

func TestAccountSelector_MultipleAccounts(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:       "test-1",
			Name:     "测试账号1",
			AppID:    "wx123456",
			Secret:   "secret123",
			Keywords: []string{"测试", "demo"},
		},
		{
			ID:       "test-2",
			Name:     "测试账号2",
			AppID:    "wx789012",
			Secret:   "secret456",
			Keywords: []string{"生产", "prod"},
		},
		{
			ID:       "test-3",
			Name:     "测试账号3",
			AppID:    "wx345678",
			Secret:   "secret789",
			Keywords: []string{"开发", "dev"},
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("测试demo文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}
}

func TestAccountSelector_SingleAccount(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
	}

	selector := NewAccountSelector(accounts, "")

	account, err := selector.SelectAccount("任何文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}
}

func TestAccountSelector_SelectByID(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("", "test-2")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-2" {
		t.Errorf("Account ID = %v, want test-2", account.ID)
	}
}

func TestAccountSelector_SelectByName(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("", "测试账号2")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-2" {
		t.Errorf("Account ID = %v, want test-2", account.ID)
	}
}

func TestAccountSelector_SelectByInvalidID(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	_, err := selector.SelectAccount("", "invalid-id")
	if err == nil {
		t.Error("SelectAccount() should return error for invalid ID")
	}
}

func TestAccountSelector_NoAccounts(t *testing.T) {
	accounts := []WechatAccount{}

	selector := NewAccountSelector(accounts, "")

	_, err := selector.SelectAccount("", "")
	if err == nil {
		t.Error("SelectAccount() should return error for no accounts")
	}
}

func TestAccountSelector_GetAllAccounts(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	allAccounts := selector.GetAllAccounts()

	if len(allAccounts) != 2 {
		t.Errorf("GetAllAccounts() length = %v, want 2", len(allAccounts))
	}
}

func TestAccountSelector_GetAccountByID(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.GetAccountByID("test-2")
	if err != nil {
		t.Fatalf("GetAccountByID() error = %v", err)
	}

	if account.ID != "test-2" {
		t.Errorf("Account ID = %v, want test-2", account.ID)
	}

	_, err = selector.GetAccountByID("invalid-id")
	if err == nil {
		t.Error("GetAccountByID() should return error for invalid ID")
	}
}

func TestAccountSelector_GetAccountByName(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:     "test-1",
			Name:   "测试账号1",
			AppID:  "wx123456",
			Secret: "secret123",
		},
		{
			ID:     "test-2",
			Name:   "测试账号2",
			AppID:  "wx789012",
			Secret: "secret456",
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.GetAccountByName("测试账号2")
	if err != nil {
		t.Fatalf("GetAccountByName() error = %v", err)
	}

	if account.ID != "test-2" {
		t.Errorf("Account ID = %v, want test-2", account.ID)
	}

	_, err = selector.GetAccountByName("invalid-name")
	if err == nil {
		t.Error("GetAccountByName() should return error for invalid name")
	}
}

func TestAccountSelector_KeywordMatchingCaseInsensitive(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:       "test-1",
			Name:     "测试账号1",
			AppID:    "wx123456",
			Secret:   "secret123",
			Keywords: []string{"测试", "DEMO"},
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("这是Demo文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}

	account, err = selector.SelectAccount("这是测试文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", account.ID)
	}
}

func TestAccountSelector_KeywordMatchingMultipleMatches(t *testing.T) {
	accounts := []WechatAccount{
		{
			ID:       "test-1",
			Name:     "测试账号1",
			AppID:    "wx123456",
			Secret:   "secret123",
			Keywords: []string{"测试", "demo"},
		},
		{
			ID:       "test-2",
			Name:     "测试账号2",
			AppID:    "wx789012",
			Secret:   "secret456",
			Keywords: []string{"测试", "prod"},
		},
	}

	selector := NewAccountSelector(accounts, "test-1")

	account, err := selector.SelectAccount("这是测试demo文章", "")
	if err != nil {
		t.Fatalf("SelectAccount() error = %v", err)
	}

	if account.ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1 (more keyword matches)", account.ID)
	}
}
