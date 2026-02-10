package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// ContentRequest 内容生成请求
type ContentRequest struct {
	Topic    string   `json:"topic"`
	Template string   `json:"template"`
	Domain   string   `json:"domain"`
	Style    string   `json:"style"`
	Keywords []string `json:"keywords"`
}

// GeneratedContent 生成的内容框架
type GeneratedContent struct {
	Title         string           `json:"title"`
	Subtitle      string           `json:"subtitle"`
	Hook          string           `json:"hook"`
	Sections      []ContentSection `json:"sections"`
	KeyPoints     []string         `json:"key_points"`
	CallToAction  string           `json:"call_to_action"`
	ViralElements []string         `json:"viral_elements"`
	SEOKeywords   []string         `json:"seo_keywords"`
}

// ContentSection 内容段落
type ContentSection struct {
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	KeyPoints  []string `json:"key_points"`
	Engagement string   `json:"engagement"`
}

func outlineCmd() *cobra.Command {
	var (
		topic    string
		template string
		domain   string
		style    string
		keywords []string
		output   string
	)

	cmd := &cobra.Command{
		Use:   "outline",
		Short: "生成爆款内容框架",
		Long: `基于话题和模板生成文章框架

支持的模板：
- authoritative: 权威揭秘型
- comparison: 对比评测型
- cultural: 文化故事型
- practical: 实用干货型

支持的风格：
- dan-koe: Dan Koe 风格（简洁有力）
- cultural-depth: 深度文化风格
- casual-science: 轻松科普风格`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runGenerate(topic, template, domain, style, keywords, output)
		},
	}

	cmd.Flags().StringVarP(&topic, "topic", "t", "", "话题（必填）")
	cmd.Flags().StringVar(&template, "template", "authoritative", "模板类型")
	cmd.Flags().StringVarP(&domain, "domain", "d", "tea", "领域配置")
	cmd.Flags().StringVarP(&style, "style", "s", "dan-koe", "写作风格")
	cmd.Flags().StringSliceVarP(&keywords, "keywords", "k", []string{}, "关键词列表")
	cmd.Flags().StringVarP(&output, "output", "o", "json", "输出格式 (json, text)")

	cmd.MarkFlagRequired("topic")

	return cmd
}

func runGenerate(topic, template, domain, style string, keywords []string, outputFormat string) error {
	request := ContentRequest{
		Topic:    topic,
		Template: template,
		Domain:   domain,
		Style:    style,
		Keywords: keywords,
	}

	// 生成内容框架
	content := generateContentFramework(request)

	if outputFormat == "json" {
		output, _ := json.MarshalIndent(content, "", "  ")
		fmt.Println(string(output))
	} else {
		printTextFramework(content)
	}

	return nil
}

func generateContentFramework(req ContentRequest) GeneratedContent {
	// 根据模板生成内容框架
	template := getTemplate(req.Template)

	content := GeneratedContent{
		Title:         generateTitle(req, template),
		Subtitle:      generateSubtitle(req, template),
		Hook:          generateHook(req, template),
		Sections:      buildSections(req, template),
		KeyPoints:     generateKeyPoints(req, template),
		CallToAction:  generateCTA(req, template),
		ViralElements: getViralElements(template),
		SEOKeywords:   optimizeKeywords(req),
	}

	return content
}

func getTemplate(templateType string) map[string]interface{} {
	templates := map[string]map[string]interface{}{
		"authoritative": {
			"name": "权威揭秘型",
			"structure": []string{
				"权威背书建立",
				"常见误区揭示",
				"专业原理解析",
				"实用方法指导",
				"价值总结升华",
			},
		},
		"comparison": {
			"name": "对比评测型",
			"structure": []string{
				"对比对象介绍",
				"多维度对比分析",
				"优劣势客观评价",
				"适用场景建议",
				"选择指南总结",
			},
		},
		"cultural": {
			"name": "文化故事型",
			"structure": []string{
				"历史背景铺垫",
				"文化故事展开",
				"精神内核挖掘",
				"现代价值连接",
				"情感共鸣升华",
			},
		},
		"practical": {
			"name": "实用干货型",
			"structure": []string{
				"问题痛点提出",
				"原因原理解析",
				"具体方法步骤",
				"注意事项提醒",
				"效果预期说明",
			},
		},
	}

	if t, ok := templates[templateType]; ok {
		return t
	}
	return templates["authoritative"]
}

func generateTitle(req ContentRequest, template map[string]interface{}) string {
	switch req.Template {
	case "authoritative":
		return fmt.Sprintf("专家揭秘：%s的5个关键秘密", req.Topic)
	case "comparison":
		return fmt.Sprintf("%s大对比：一张图看懂核心区别", req.Topic)
	case "cultural":
		return fmt.Sprintf("%s背后的文化密码：千年传承的现代表达", req.Topic)
	case "practical":
		return fmt.Sprintf("3分钟学会：%s的实用指南", req.Topic)
	default:
		return fmt.Sprintf("关于%s的深度解析", req.Topic)
	}
}

func generateSubtitle(req ContentRequest, template map[string]interface{}) string {
	return fmt.Sprintf("深度解析%s，助你快速掌握要点", req.Topic)
}

func generateHook(req ContentRequest, template map[string]interface{}) string {
	switch req.Template {
	case "authoritative":
		return fmt.Sprintf("你知道吗？在%s方面，90%%的人都存在认知误区。今天就来告诉你一些行业内幕...", req.Topic)
	case "comparison":
		return fmt.Sprintf("面对%s，你是不是也经常感到选择困难？今天用一张图，让你看懂它们的核心区别...", req.Topic)
	case "cultural":
		return fmt.Sprintf("每一个%s，都承载着千年的文化记忆。今天，让我们一起探寻背后的文化密码...", req.Topic)
	case "practical":
		return fmt.Sprintf("在%s上，你是不是也踩过很多坑？今天，分享一些专业方法，让你少走弯路...", req.Topic)
	default:
		return fmt.Sprintf("关于%s，你可能还有很多不了解的地方...", req.Topic)
	}
}

func buildSections(req ContentRequest, template map[string]interface{}) []ContentSection {
	structure, ok := template["structure"].([]string)
	if !ok {
		return []ContentSection{}
	}

	var sections []ContentSection
	for i, sectionType := range structure {
		section := ContentSection{
			Title:      fmt.Sprintf("%d. %s", i+1, sectionType),
			Content:    fmt.Sprintf("详细阐述%s的要点", sectionType),
			KeyPoints:  []string{"要点1", "要点2", "要点3"},
			Engagement: "💡 互动提问：你有类似的经验吗？",
		}
		sections = append(sections, section)
	}

	return sections
}

func generateKeyPoints(req ContentRequest, template map[string]interface{}) []string {
	return []string{
		"专业权威性建立",
		"实用方法指导",
		"价值总结升华",
	}
}

func generateCTA(req ContentRequest, template map[string]interface{}) string {
	return "💡 如果你觉得这些内容有价值，欢迎分享给更多需要的朋友。有什么问题，也可以在评论区交流。"
}

func getViralElements(template map[string]interface{}) []string {
	return []string{
		"权威专业背书",
		"实用价值传递",
		"情感共鸣设计",
		"互动参与引导",
	}
}

func optimizeKeywords(req ContentRequest) []string {
	keywords := []string{req.Topic}
	keywords = append(keywords, req.Keywords...)
	return keywords
}

func printTextFramework(content GeneratedContent) {
	fmt.Printf("内容框架\n")
	fmt.Printf("========\n\n")
	fmt.Printf("标题: %s\n", content.Title)
	fmt.Printf("副标题: %s\n", content.Subtitle)
	fmt.Printf("\n开头钩子:\n%s\n", content.Hook)

	fmt.Printf("\n内容结构:\n")
	for _, section := range content.Sections {
		fmt.Printf("\n%s\n", section.Title)
		fmt.Printf("  %s\n", section.Content)
	}

	fmt.Printf("\n行动号召:\n%s\n", content.CallToAction)
}
