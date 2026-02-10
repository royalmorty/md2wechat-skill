package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_DefaultConfig(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.DefaultTheme != "default" {
		t.Errorf("DefaultTheme = %v, want default", cfg.DefaultTheme)
	}
	if cfg.CompressImages != true {
		t.Errorf("CompressImages = %v, want true", cfg.CompressImages)
	}
	if cfg.MaxImageWidth != 1920 {
		t.Errorf("MaxImageWidth = %v, want 1920", cfg.MaxImageWidth)
	}
	if cfg.MaxImageSize != 5*1024*1024 {
		t.Errorf("MaxImageSize = %v, want 5MB", cfg.MaxImageSize)
	}
	if cfg.HTTPTimeout != 30 {
		t.Errorf("HTTPTimeout = %v, want 30", cfg.HTTPTimeout)
	}
	if cfg.ImageProvider != "openai" {
		t.Errorf("ImageProvider = %v, want openai", cfg.ImageProvider)
	}
	if cfg.ImageAPIBase != "https://api.openai.com/v1" {
		t.Errorf("ImageAPIBase = %v, want https://api.openai.com/v1", cfg.ImageAPIBase)
	}
	if cfg.ImageModel != "dall-e-3" {
		t.Errorf("ImageModel = %v, want dall-e-3", cfg.ImageModel)
	}
	if cfg.ImageSize != "1024x1024" {
		t.Errorf("ImageSize = %v, want 1024x1024", cfg.ImageSize)
	}
}

func TestLoad_YAMLConfig(t *testing.T) {
	configContent := `
wechat:
  accounts:
    - id: "test-1"
      name: "测试账号"
      appid: "wx123456"
      secret: "secret123"
      keywords:
        - "测试"
      default_style: "dan-koe"
  default: "test-1"
api:
  image_key: "test_image_key"
  image_base_url: "https://test.api.com"
  image_provider: "modelscope"
  image_model: "test-model"
  image_size: "512x512"
  default_theme: "apple"
  http_timeout: 60
image:
  compress: false
  max_width: 2560
  max_size_mb: 10
`

	tmpFile := filepath.Join(t.TempDir(), "test.yaml")
	if err := os.WriteFile(tmpFile, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}

	cfg, err := LoadWithDefaults(tmpFile)
	if err != nil {
		t.Fatalf("LoadWithDefaults() error = %v", err)
	}

	if len(cfg.WechatAccounts) != 1 {
		t.Errorf("WechatAccounts length = %v, want 1", len(cfg.WechatAccounts))
	}
	if cfg.WechatAccounts[0].ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", cfg.WechatAccounts[0].ID)
	}
	if cfg.ImageAPIKey != "test_image_key" {
		t.Errorf("ImageAPIKey = %v, want test_image_key", cfg.ImageAPIKey)
	}
	if cfg.ImageAPIBase != "https://test.api.com" {
		t.Errorf("ImageAPIBase = %v, want https://test.api.com", cfg.ImageAPIBase)
	}
	if cfg.ImageProvider != "modelscope" {
		t.Errorf("ImageProvider = %v, want modelscope", cfg.ImageProvider)
	}
	if cfg.ImageModel != "test-model" {
		t.Errorf("ImageModel = %v, want test-model", cfg.ImageModel)
	}
	if cfg.ImageSize != "512x512" {
		t.Errorf("ImageSize = %v, want 512x512", cfg.ImageSize)
	}
	if cfg.DefaultTheme != "apple" {
		t.Errorf("DefaultTheme = %v, want apple", cfg.DefaultTheme)
	}
	if cfg.HTTPTimeout != 60 {
		t.Errorf("HTTPTimeout = %v, want 60", cfg.HTTPTimeout)
	}
	if cfg.CompressImages != false {
		t.Errorf("CompressImages = %v, want false", cfg.CompressImages)
	}
	if cfg.MaxImageWidth != 2560 {
		t.Errorf("MaxImageWidth = %v, want 2560", cfg.MaxImageWidth)
	}
	if cfg.MaxImageSize != 10*1024*1024 {
		t.Errorf("MaxImageSize = %v, want 10MB", cfg.MaxImageSize)
	}
}

func TestLoad_JSONConfig(t *testing.T) {
	configContent := `{
  "wechat": {
    "accounts": [
      {
        "id": "test-1",
        "name": "测试账号",
        "appid": "wx123456",
        "secret": "secret123",
        "keywords": ["测试"],
        "default_style": "dan-koe"
      }
    ],
    "default": "test-1"
  },
  "api": {
    "image_key": "test_image_key",
    "image_base_url": "https://test.api.com"
  }
}`

	tmpFile := filepath.Join(t.TempDir(), "test.json")
	if err := os.WriteFile(tmpFile, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}

	cfg, err := LoadWithDefaults(tmpFile)
	if err != nil {
		t.Fatalf("LoadWithDefaults() error = %v", err)
	}

	if len(cfg.WechatAccounts) != 1 {
		t.Errorf("WechatAccounts length = %v, want 1", len(cfg.WechatAccounts))
	}
	if cfg.WechatAccounts[0].ID != "test-1" {
		t.Errorf("Account ID = %v, want test-1", cfg.WechatAccounts[0].ID)
	}
	if cfg.ImageAPIKey != "test_image_key" {
		t.Errorf("ImageAPIKey = %v, want test_image_key", cfg.ImageAPIKey)
	}
}

func TestLoad_EnvOverride(t *testing.T) {
	configContent := `
wechat:
  accounts:
    - id: "test-1"
      name: "测试账号"
      appid: "wx123456"
      secret: "secret123"
  default: "test-1"
`

	tmpFile := filepath.Join(t.TempDir(), "test.yaml")
	if err := os.WriteFile(tmpFile, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to create temp config file: %v", err)
	}

	os.Setenv("IMAGE_API_KEY", "env_image_key")
	os.Setenv("IMAGE_API_BASE", "https://env.api.com")
	os.Setenv("IMAGE_PROVIDER", "env_provider")
	os.Setenv("IMAGE_MODEL", "env_model")
	os.Setenv("IMAGE_SIZE", "env_size")
	os.Setenv("DEFAULT_THEME", "env_theme")
	os.Setenv("COMPRESS_IMAGES", "false")
	os.Setenv("MAX_IMAGE_WIDTH", "2048")
	os.Setenv("MAX_IMAGE_SIZE", "8")
	os.Setenv("HTTP_TIMEOUT", "45")
	defer func() {
		os.Unsetenv("IMAGE_API_KEY")
		os.Unsetenv("IMAGE_API_BASE")
		os.Unsetenv("IMAGE_PROVIDER")
		os.Unsetenv("IMAGE_MODEL")
		os.Unsetenv("IMAGE_SIZE")
		os.Unsetenv("DEFAULT_THEME")
		os.Unsetenv("COMPRESS_IMAGES")
		os.Unsetenv("MAX_IMAGE_WIDTH")
		os.Unsetenv("MAX_IMAGE_SIZE")
		os.Unsetenv("HTTP_TIMEOUT")
	}()

	cfg, err := LoadWithDefaults(tmpFile)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.ImageAPIKey != "env_image_key" {
		t.Errorf("ImageAPIKey = %v, want env_image_key", cfg.ImageAPIKey)
	}
	if cfg.ImageAPIBase != "https://env.api.com" {
		t.Errorf("ImageAPIBase = %v, want https://env.api.com", cfg.ImageAPIBase)
	}
	if cfg.ImageProvider != "env_provider" {
		t.Errorf("ImageProvider = %v, want env_provider", cfg.ImageProvider)
	}
	if cfg.ImageModel != "env_model" {
		t.Errorf("ImageModel = %v, want env_model", cfg.ImageModel)
	}
	if cfg.ImageSize != "env_size" {
		t.Errorf("ImageSize = %v, want env_size", cfg.ImageSize)
	}
	if cfg.DefaultTheme != "env_theme" {
		t.Errorf("DefaultTheme = %v, want env_theme", cfg.DefaultTheme)
	}
	if cfg.CompressImages != false {
		t.Errorf("CompressImages = %v, want false", cfg.CompressImages)
	}
	if cfg.MaxImageWidth != 2048 {
		t.Errorf("MaxImageWidth = %v, want 2048", cfg.MaxImageWidth)
	}
	if cfg.MaxImageSize != 8*1024*1024 {
		t.Errorf("MaxImageSize = %v, want 8MB", cfg.MaxImageSize)
	}
	if cfg.HTTPTimeout != 45 {
		t.Errorf("HTTPTimeout = %v, want 45", cfg.HTTPTimeout)
	}
}

func TestConfig_Validate_Success(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:     "test-1",
				Name:   "测试账号",
				AppID:  "wx123456",
				Secret: "secret123",
			},
		},
		DefaultAccount: "test-1",
		MaxImageWidth:  1920,
		MaxImageSize:   5 * 1024 * 1024,
		HTTPTimeout:    30,
	}

	err := cfg.Validate()
	if err != nil {
		t.Errorf("Validate() error = %v, want nil", err)
	}
}

func TestConfig_Validate_MissingAccount(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{},
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("Validate() should return error for missing accounts")
	}

	configErr, ok := err.(*ConfigError)
	if !ok {
		t.Fatalf("Error type = %T, want *ConfigError", err)
	}

	if configErr.Field != "WechatAccounts" {
		t.Errorf("Error field = %v, want WechatAccounts", configErr.Field)
	}
}

func TestConfig_Validate_MissingAccountID(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				Name:   "测试账号",
				AppID:  "wx123456",
				Secret: "secret123",
			},
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("Validate() should return error for missing account ID")
	}

	configErr, ok := err.(*ConfigError)
	if !ok {
		t.Fatalf("Error type = %T, want *ConfigError", err)
	}

	if configErr.Field != "WechatAccounts[0].ID" {
		t.Errorf("Error field = %v, want WechatAccounts[0].ID", configErr.Field)
	}
}

func TestConfig_Validate_MissingAppID(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:     "test-1",
				Name:   "测试账号",
				Secret: "secret123",
			},
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("Validate() should return error for missing AppID")
	}

	configErr, ok := err.(*ConfigError)
	if !ok {
		t.Fatalf("Error type = %T, want *ConfigError", err)
	}

	if configErr.Field != "WechatAccounts[0].AppID" {
		t.Errorf("Error field = %v, want WechatAccounts[0].AppID", configErr.Field)
	}
}

func TestConfig_Validate_MissingSecret(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:    "test-1",
				Name:  "测试账号",
				AppID: "wx123456",
			},
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("Validate() should return error for missing Secret")
	}

	configErr, ok := err.(*ConfigError)
	if !ok {
		t.Fatalf("Error type = %T, want *ConfigError", err)
	}

	if configErr.Field != "WechatAccounts[0].Secret" {
		t.Errorf("Error field = %v, want WechatAccounts[0].Secret", configErr.Field)
	}
}

func TestConfig_Validate_InvalidImageWidth(t *testing.T) {
	tests := []struct {
		name    string
		width   int
		wantErr bool
	}{
		{"too small", 50, true},
		{"too large", 15000, true},
		{"valid", 1920, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				WechatAccounts: []WechatAccount{
					{
						ID:     "test-1",
						Name:   "测试账号",
						AppID:  "wx123456",
						Secret: "secret123",
					},
				},
				MaxImageWidth: tt.width,
				MaxImageSize:  5 * 1024 * 1024,
				HTTPTimeout:   30,
			}

			err := cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_Validate_InvalidTimeout(t *testing.T) {
	tests := []struct {
		name    string
		timeout int
		wantErr bool
	}{
		{"too small", 0, true},
		{"too large", 500, true},
		{"valid", 30, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				WechatAccounts: []WechatAccount{
					{
						ID:     "test-1",
						Name:   "测试账号",
						AppID:  "wx123456",
						Secret: "secret123",
					},
				},
				MaxImageWidth: 1920,
				MaxImageSize:  5 * 1024 * 1024,
				HTTPTimeout:   tt.timeout,
			}

			err := cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_MultiAccount(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
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
		},
		DefaultAccount: "test-1",
		MaxImageWidth:  1920,
		MaxImageSize:   5 * 1024 * 1024,
		HTTPTimeout:    30,
	}

	err := cfg.Validate()
	if err != nil {
		t.Fatalf("Validate() error = %v", err)
	}

	if len(cfg.WechatAccounts) != 2 {
		t.Errorf("WechatAccounts length = %v, want 2", len(cfg.WechatAccounts))
	}
	if cfg.DefaultAccount != "test-1" {
		t.Errorf("DefaultAccount = %v, want test-1", cfg.DefaultAccount)
	}
}

func TestSaveConfig_YAML(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:     "test-1",
				Name:   "测试账号",
				AppID:  "wx123456",
				Secret: "secret123",
			},
		},
		DefaultAccount: "test-1",
		ImageAPIKey:    "test_key",
		MaxImageWidth:  1920,
		MaxImageSize:   5 * 1024 * 1024,
		HTTPTimeout:    30,
	}

	tmpFile := filepath.Join(t.TempDir(), "test.yaml")
	err := SaveConfig(tmpFile, cfg)
	if err != nil {
		t.Fatalf("SaveConfig() error = %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	content := string(data)
	if !contains(content, "test-1") {
		t.Error("Saved config should contain account ID")
	}
	if !contains(content, "test_key") {
		t.Error("Saved config should contain image API key")
	}
}

func TestSaveConfig_JSON(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:     "test-1",
				Name:   "测试账号",
				AppID:  "wx123456",
				Secret: "secret123",
			},
		},
		DefaultAccount: "test-1",
		ImageAPIKey:    "test_key",
		MaxImageWidth:  1920,
		MaxImageSize:   5 * 1024 * 1024,
		HTTPTimeout:    30,
	}

	tmpFile := filepath.Join(t.TempDir(), "test.json")
	err := SaveConfig(tmpFile, cfg)
	if err != nil {
		t.Fatalf("SaveConfig() error = %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	content := string(data)
	if !contains(content, "test-1") {
		t.Error("Saved config should contain account ID")
	}
	if !contains(content, "test_key") {
		t.Error("Saved config should contain image API key")
	}
}

func TestConfig_ToMap(t *testing.T) {
	cfg := &Config{
		WechatAccounts: []WechatAccount{
			{
				ID:     "test-1",
				Name:   "测试账号",
				AppID:  "wx123456",
				Secret: "secret1234567890",
			},
		},
		DefaultAccount: "test-1",
		ImageAPIKey:    "test_image_key_123456",
	}

	m := cfg.ToMap(true)

	accounts, ok := m["wechat_accounts"].([]map[string]any)
	if !ok {
		t.Fatal("wechat_accounts should be a slice of maps")
	}

	if len(accounts) != 1 {
		t.Errorf("accounts length = %v, want 1", len(accounts))
	}

	if accounts[0]["secret"] != "se***90" {
		t.Errorf("secret should be masked, got %v", accounts[0]["secret"])
	}

	if m["image_api_key"] != "te***56" {
		t.Errorf("image_api_key should be masked, got %v", m["image_api_key"])
	}

	m2 := cfg.ToMap(false)
	accounts2, ok := m2["wechat_accounts"].([]map[string]any)
	if !ok {
		t.Fatal("wechat_accounts should be a slice of maps")
	}

	if accounts2[0]["secret"] != "secret1234567890" {
		t.Errorf("secret should not be masked, got %v", accounts2[0]["secret"])
	}
}

func TestConfig_ValidateForImageGeneration(t *testing.T) {
	cfg := &Config{
		ImageAPIKey: "",
	}

	err := cfg.ValidateForImageGeneration()
	if err == nil {
		t.Error("ValidateForImageGeneration() should return error for missing API key")
	}

	cfg.ImageAPIKey = "test_key"
	err = cfg.ValidateForImageGeneration()
	if err != nil {
		t.Errorf("ValidateForImageGeneration() error = %v, want nil", err)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
