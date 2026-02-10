package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/royalrick/wechatwriter/app/draft"
	"github.com/spf13/cobra"
)

// draftCmd 草稿管理命令组
func draftCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "draft",
		Short: "草稿管理（创建、测试、发布）",
		Long: `草稿管理命令组

支持的操作：
  create   - 从 JSON 文件创建草稿
  test     - 测试草稿 HTML
  publish  - 创建并发布草稿`,
	}

	cmd.AddCommand(draftCreateCmd())
	cmd.AddCommand(draftTestCmd())
	cmd.AddCommand(draftPublishCmd())

	return cmd
}

// draftCreateCmd 从 JSON 创建草稿
func draftCreateCmd() *cobra.Command {
	var accountID string

	cmd := &cobra.Command{
		Use:   "create <json_file>",
		Short: "从 JSON 文件创建微信草稿",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			jsonFile := args[0]
			svc := draft.NewService(cfg, log)

			// 读取 JSON 文件
			data, err := os.ReadFile(jsonFile)
			if err != nil {
				responseError(fmt.Errorf("read file: %w", err))
				return
			}

			// 解析请求
			var req struct {
				Articles []draft.Article `json:"articles"`
			}
			if err := json.Unmarshal(data, &req); err != nil {
				responseError(fmt.Errorf("parse json: %w", err))
				return
			}

			// 验证
			if len(req.Articles) == 0 {
				responseError(fmt.Errorf("no articles in request"))
				return
			}

			// 使用指定账号创建草稿
			var result *draft.DraftResult
			if accountID != "" {
				result, err = svc.CreateDraftWithAccount(req.Articles, accountID)
			} else {
				result, err = svc.CreateDraft(req.Articles)
			}

			if err != nil {
				responseError(err)
				return
			}
			responseSuccess(result)
		},
	}

	cmd.Flags().StringVarP(&accountID, "account", "a", "", "指定微信公众号账号ID（可选，不指定则自动选择）")

	return cmd
}

// draftTestCmd 测试草稿
func draftTestCmd() *cobra.Command {
	var accountID string

	cmd := &cobra.Command{
		Use:   "test <html_file> <cover_image>",
		Short: "测试从 HTML 文件创建微信草稿",
		Args:  cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			htmlFile := args[0]
			coverImage := args[1]

			// 读取 HTML
			html, err := os.ReadFile(htmlFile)
			if err != nil {
				responseError(fmt.Errorf("read HTML file: %w", err))
				return
			}

			// 上传封面图片
			coverMediaID, err := uploadCoverImage(coverImage)
			if err != nil {
				responseError(fmt.Errorf("upload cover: %w", err))
				return
			}

			// 创建草稿
			svc := draft.NewService(cfg, log)
			articles := []draft.Article{
				{
					Title:        "AI生成测试文章",
					Content:      string(html),
					Digest:       "这是AI生成的微信公众号文章测试",
					ThumbMediaID: coverMediaID,
					ShowCoverPic: 1,
				},
			}

			var result *draft.DraftResult
			if accountID != "" {
				result, err = svc.CreateDraftWithAccount(articles, accountID)
			} else {
				result, err = svc.CreateDraft(articles)
			}

			if err != nil {
				responseError(fmt.Errorf("create draft: %w", err))
				return
			}

			responseSuccess(map[string]any{
				"success":   true,
				"media_id":  result.MediaID,
				"draft_url": result.DraftURL,
				"message":   "Draft created successfully!",
			})
		},
	}

	cmd.Flags().StringVarP(&accountID, "account", "a", "", "指定微信公众号账号ID（可选，不指定则自动选择）")

	return cmd
}

// draftPublishCmd 创建并发布草稿
func draftPublishCmd() *cobra.Command {
	var (
		title     string
		content   string
		author    string
		digest    string
		coverID   string
		outputDir string
	)

	cmd := &cobra.Command{
		Use:   "publish",
		Short: "创建草稿并上传到微信公众号",
		Long: `将文章内容上传到微信公众号草稿箱

需要在配置文件中设置微信公众号账号信息。
使用 'writer config init' 创建配置文件。
支持多账号配置，可通过 --account 参数指定账号。`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPublish(title, content, author, digest, coverID, outputDir)
		},
	}

	cmd.Flags().StringVarP(&title, "title", "t", "", "文章标题（必填）")
	cmd.Flags().StringVarP(&content, "content", "c", "", "文章内容（HTML格式）")
	cmd.Flags().StringVarP(&author, "author", "a", "", "作者名称")
	cmd.Flags().StringVar(&digest, "digest", "", "文章摘要")
	cmd.Flags().StringVar(&coverID, "cover", "", "封面图片Media ID")
	cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "输出目录（保存草稿JSON）")

	cmd.MarkFlagRequired("title")
	cmd.MarkFlagRequired("content")

	return cmd
}

func runPublish(title, content, author, digest, coverID, outputDir string) error {
	// 创建草稿JSON
	draftData := map[string]interface{}{
		"articles": []map[string]interface{}{
			{
				"title":   title,
				"content": content,
				"author":  author,
				"digest":  digest,
			},
		},
	}

	if coverID != "" {
		draftData["articles"].([]map[string]interface{})[0]["thumb_media_id"] = coverID
		draftData["articles"].([]map[string]interface{})[0]["show_cover_pic"] = 1
	}

	// 保存到文件
	draftFile := fmt.Sprintf("%s/draft.json", outputDir)
	data, _ := json.MarshalIndent(draftData, "", "  ")
	if err := os.WriteFile(draftFile, data, 0644); err != nil {
		return fmt.Errorf("无法保存草稿文件: %w", err)
	}

	fmt.Printf("✅ 草稿JSON已保存到: %s\n", draftFile)
	fmt.Printf("\n使用以下命令上传草稿:\n")
	fmt.Printf("  writer draft create %s\n", draftFile)

	return nil
}
