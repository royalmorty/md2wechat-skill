package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/royalrick/wechatwriter/app/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// configCmd 配置管理命令组
func configCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "配置管理",
		Long: `管理 writer 配置

Configuration Priority:
  1. Environment variables (highest)
  2. Config file
  3. Default values (lowest)

Config file search order:
  1. ~/.config/wechatwriter/config.yaml  (global config, recommended)
  2. ~/.wechatwriter.yaml                (global config)
  3. ./wechatwriter.yaml                  (project config)`,
		Run: func(cmd *cobra.Command, args []string) {
			// 默认显示配置
			if err := showConfig(false); err != nil {
				responseError(err)
			}
		},
	}

	cmd.AddCommand(configShowCmd())
	cmd.AddCommand(configValidateCmd())
	cmd.AddCommand(configInitCmd())
	cmd.AddCommand(configDomainCmd())
	cmd.AddCommand(configStyleCmd())
	cmd.AddCommand(configAccountsCmd())

	return cmd
}

// configShowCmd 显示配置
func configShowCmd() *cobra.Command {
	var showSecret bool
	var format string

	cmd := &cobra.Command{
		Use:   "show",
		Short: "显示当前配置",
		Run: func(cmd *cobra.Command, args []string) {
			if err := showConfig(showSecret); err != nil {
				responseError(err)
			}
		},
	}

	cmd.Flags().BoolVar(&showSecret, "show-secret", false, "显示密钥值")
	cmd.Flags().StringVarP(&format, "format", "f", "json", "输出格式: json, yaml")

	return cmd
}

// configValidateCmd 验证配置
func configValidateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "验证配置",
		Run: func(cmd *cobra.Command, args []string) {
			if err := validateConfig(); err != nil {
				responseError(err)
			} else {
				responseSuccess(map[string]any{
					"valid":   true,
					"message": "Configuration is valid",
				})
			}
		},
	}
}

// configInitCmd 初始化配置文件
func configInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init [output_file]",
		Short: "创建示例配置文件",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var outputFile string
			if len(args) > 0 {
				outputFile = args[0]
			} else {
				homeDir, _ := os.UserHomeDir()
				configDir := homeDir + "/.config/wechatwriter"
				outputFile = configDir + "/config.yaml"
			}

			if err := initConfigFile(outputFile); err != nil {
				responseError(err)
				return
			}

			relPath := outputFile
			homeDir, _ := os.UserHomeDir()
			if homeDir != "" && strings.HasPrefix(outputFile, homeDir) {
				rel := strings.TrimPrefix(outputFile, homeDir)

				home, err := os.UserHomeDir()
				if err != nil {
					responseError(err)
					return
				}

				relPath = filepath.Join(home, rel)
			}

			fmt.Fprintf(os.Stderr, "\n✅ 配置文件已创建: %s\n", relPath)
			responseSuccess(map[string]any{
				"file":    relPath,
				"message": "Config file created",
			})
		},
	}
}

// configDomainCmd 领域配置管理
func configDomainCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "domain",
		Short: "管理领域配置",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listDomains()
		},
	}
}

// configStyleCmd 风格配置管理
func configStyleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "style",
		Short: "管理写作风格",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listStyles()
		},
	}
}

// Helper functions
func showConfig(showSecret bool) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("no configuration found. Create a config file with 'writer config init'")
	}

	result := map[string]any{
		"success": true,
		"config":  cfg.ToMap(!showSecret),
	}

	printJSON(result)
	return nil
}

func validateConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	log.Info("configuration validated",
		zap.String("config_file", cfg.GetConfigFile()),
		zap.String("default_theme", cfg.DefaultTheme))

	return nil
}

func initConfigFile(outputFile string) error {
	if _, err := os.Stat(outputFile); err == nil {
		return fmt.Errorf("config file already exists: %s", outputFile)
	}

	// 创建多账号配置模板
	cfg := &config.Config{
		WechatAccounts: []config.WechatAccount{
			{
				ID:           "my-account",
				Name:         "我的公众号",
				AppID:        "your_wechat_appid",
				Secret:       "your_wechat_secret",
				Keywords:     []string{"关键词1", "关键词2"},
				DefaultStyle: "dan-koe",
			},
		},
		DefaultAccount: "my-account",
		ImageAPIKey:    "your_image_api_key",
		ImageAPIBase:   "https://api.openai.com/v1",
		ImageProvider:  "openai",
		ImageModel:     "dall-e-3",
		ImageSize:      "1024x1024",
		DefaultTheme:   "default",
		CompressImages: true,
		MaxImageWidth:  1920,
		MaxImageSize:   5 * 1024 * 1024,
		HTTPTimeout:    30,
	}

	return config.SaveConfig(outputFile, cfg)
}

func listDomains() error {
	fmt.Println("可用的领域配置:")
	fmt.Println("================")

	configDir := "config/domains"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		fmt.Println("  (暂无领域配置文件)")
		fmt.Println("\n默认领域:")
		fmt.Println("  • tea - 茶文化领域（默认）")
		return nil
	}

	entries, err := os.ReadDir(configDir)
	if err != nil {
		return fmt.Errorf("读取配置目录失败: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			domain := entry.Name()[:len(entry.Name())-5]
			fmt.Printf("  • %s\n", domain)
		}
	}

	return nil
}

func listStyles() error {
	fmt.Println("可用的写作风格:")
	fmt.Println("================")

	styles := map[string]string{
		"dan-koe":        "Dan Koe 风格 - 简洁有力",
		"cultural-depth": "深度文化风格 - 引经据典",
		"casual-science": "轻松科普风格 - 通俗易懂",
	}

	for style, desc := range styles {
		fmt.Printf("  • %s - %s\n", style, desc)
	}

	return nil
}

// configAccountsCmd 列出可用的微信公众号账号
func configAccountsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts",
		Short: "列出可用的微信公众号账号",
		Long: `列出配置文件中所有可用的微信公众号账号。

当 AI 不确定使用哪个账号时，可以调用此命令获取账号列表。
输出包含账号 ID、名称、关键词等信息，帮助选择合适的账号。`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			// 获取账号列表
			accounts := cfg.WechatAccounts

			// 构建输出
			var result []map[string]any
			for _, acc := range accounts {
				result = append(result, map[string]any{
					"id":            acc.ID,
					"name":          acc.Name,
					"appid":         acc.AppID,
					"keywords":      acc.Keywords,
					"default_style": acc.DefaultStyle,
					"is_default":    acc.ID == cfg.DefaultAccount,
				})
			}

			responseSuccess(map[string]any{
				"accounts":        result,
				"default_account": cfg.DefaultAccount,
				"total":           len(accounts),
			})
		},
	}

	return cmd
}
