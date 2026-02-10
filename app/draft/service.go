package draft

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/royalrick/wechatwriter/app/config"
	"github.com/royalrick/wechatwriter/app/wechat"
	"github.com/silenceper/wechat/v2/officialaccount/draft"
	"go.uber.org/zap"
)

// Service 草稿服务
type Service struct {
	cfg      *config.Config
	selector *config.AccountSelector
	log      *zap.Logger
}

// NewService 创建草稿服务
func NewService(cfg *config.Config, log *zap.Logger) *Service {
	selector := config.NewAccountSelector(cfg.WechatAccounts, cfg.DefaultAccount)
	return &Service{
		cfg:      cfg,
		selector: selector,
		log:      log,
	}
}

// DraftRequest 草稿请求
type DraftRequest struct {
	Articles []Article `json:"articles"`
}

// Article 文章
type Article struct {
	Title            string `json:"title"`
	Author           string `json:"author,omitempty"`
	Digest           string `json:"digest,omitempty"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url,omitempty"`
	ThumbMediaID     string `json:"thumb_media_id,omitempty"`
	ShowCoverPic     int    `json:"show_cover_pic,omitempty"`
}

// DraftResult 草稿结果
type DraftResult struct {
	MediaID  string `json:"media_id"`
	DraftURL string `json:"draft_url,omitempty"`
}

// CreateDraftFromFile 从 JSON 文件创建草稿
func (s *Service) CreateDraftFromFile(jsonFile string) (*DraftResult, error) {
	s.log.Info("creating draft from file", zap.String("file", jsonFile))

	// 读取 JSON 文件
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	// 解析请求
	var req DraftRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, fmt.Errorf("parse json: %w", err)
	}

	// 验证
	if len(req.Articles) == 0 {
		return nil, fmt.Errorf("no articles in request")
	}

	// 使用 CreateDraft 方法（会自动选择账号）
	return s.CreateDraft(req.Articles)
}

// CreateDraft 创建草稿（使用默认账号或第一个账号）
func (s *Service) CreateDraft(articles []Article) (*DraftResult, error) {
	return s.CreateDraftWithAccount(articles, "")
}

// CreateDraftWithAccount 使用指定账号创建草稿
func (s *Service) CreateDraftWithAccount(articles []Article, accountID string) (*DraftResult, error) {
	// 选择账号
	account, err := s.selector.SelectAccount("", accountID)
	if err != nil {
		return nil, fmt.Errorf("select account: %w", err)
	}

	s.log.Info("creating draft with account",
		zap.String("account_id", account.ID),
		zap.String("account_name", account.Name))

	// 创建 WeChat Service
	ws := wechat.NewService(account, s.log)

	// 转换为 SDK 格式
	var draftArticles []*draft.Article
	for _, a := range articles {
		article := &draft.Article{
			Title:   a.Title,
			Content: a.Content,
			Digest:  a.Digest,
			Author:  a.Author,
		}

		if a.ThumbMediaID != "" {
			article.ThumbMediaID = a.ThumbMediaID
			article.ShowCoverPic = uint(a.ShowCoverPic)
		}

		if a.ContentSourceURL != "" {
			article.ContentSourceURL = a.ContentSourceURL
		}

		draftArticles = append(draftArticles, article)
	}

	// 调用微信 API
	result, err := ws.CreateDraft(draftArticles)
	if err != nil {
		return nil, err
	}

	return &DraftResult{
		MediaID:  result.MediaID,
		DraftURL: result.DraftURL,
	}, nil
}

// CreateDraftForPrompt 根据提示词选择账号并创建草稿
func (s *Service) CreateDraftForPrompt(articles []Article, prompt string) (*DraftResult, error) {
	// 选择账号
	account, err := s.selector.SelectAccount(prompt, "")
	if err != nil {
		return nil, fmt.Errorf("select account: %w", err)
	}

	s.log.Info("creating draft with account (keyword matched)",
		zap.String("account_id", account.ID),
		zap.String("account_name", account.Name),
		zap.String("prompt", prompt))

	// 创建 WeChat Service
	ws := wechat.NewService(account, s.log)

	// 转换为 SDK 格式
	var draftArticles []*draft.Article
	for _, a := range articles {
		article := &draft.Article{
			Title:   a.Title,
			Content: a.Content,
			Digest:  a.Digest,
			Author:  a.Author,
		}

		if a.ThumbMediaID != "" {
			article.ThumbMediaID = a.ThumbMediaID
			article.ShowCoverPic = uint(a.ShowCoverPic)
		}

		if a.ContentSourceURL != "" {
			article.ContentSourceURL = a.ContentSourceURL
		}

		draftArticles = append(draftArticles, article)
	}

	// 调用微信 API
	result, err := ws.CreateDraft(draftArticles)
	if err != nil {
		return nil, err
	}

	return &DraftResult{
		MediaID:  result.MediaID,
		DraftURL: result.DraftURL,
	}, nil
}

// GenerateDigestFromContent 从内容生成摘要
func GenerateDigestFromContent(content string, maxLen int) string {
	if maxLen == 0 {
		maxLen = 120
	}

	// 简化实现：去除 HTML 标签后截取
	// 实际应该使用 HTML 解析器

	// 移除 HTML 标签的简单方法
	content = stripHTML(content)

	// 截取
	if len(content) > maxLen {
		content = content[:maxLen] + "..."
	}

	return content
}

// stripHTML 去除 HTML 标签（简化版）
func stripHTML(html string) string {
	// 简化实现：移除常见标签
	// 实际应该使用 proper HTML 解析器
	result := html
	for _, tag := range []string{"</p>", "<br/>", "<br>", "</div>", "</h1>", "</h2>", "</h3>"} {
		result = strings.ReplaceAll(result, tag, "\n")
	}

	// 移除所有标签
	inTag := false
	var clean strings.Builder
	for _, r := range result {
		if r == '<' {
			inTag = true
		} else if r == '>' {
			inTag = false
		} else if !inTag {
			clean.WriteRune(r)
		}
	}

	return clean.String()
}
