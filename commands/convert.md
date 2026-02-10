---
description: "Markdown转微信公众号HTML,支持多主题、智能识别和一键发布"
---

# wechatwriter智能转换

智能识别用户需求并执行最佳转换策略。支持自动识别、一键发布和高级自定义选项。

## 基础用法

### 智能识别模式

```bash
# 自动分析并转换
/wechatwriter:convert article.md

# 指定输出格式
/wechatwriter:convert article.md --to html

# 批量转换
/wechatwriter:convert articles/*.md --batch
```

自动识别内容类型并选择最优转换策略：
- **技术文档** → API快速模式
- **品牌文章** → AI精美模式
- **故事内容** → 文学主题模式
- **新闻报道** → 媒体专业模式
- **教程指南** → 步骤清晰模式

### 一键转换上传

```bash
# 完整发布流程
/wechatwriter:convert article.md --publish

# 定时发布
/wechatwriter:convert article.md --publish --schedule "2024-01-15 10:00"

# 多平台发布
/wechatwriter:convert article.md --publish --platforms "wechat,weibo,zhihu"
```

自动完成：转换 → 生成封面 → 优化内容 → 上传到草稿箱

## 高级选项

### 主题和风格选择

```bash
# 指定主题
/wechatwriter:convert article.md --theme autumn-warm

# 智能推荐主题
/wechatwriter:convert article.md --smart-theme --based-on content

# 多主题对比
/wechatwriter:convert article.md --themes "spring-fresh,autumn-warm,ocean-calm" --compare

# 自定义主题参数
/wechatwriter:convert article.md --theme custom --primary-color "#1890ff" --font "modern"
```

支持主题：
- `spring-fresh`: 春天清新
- `autumn-warm`: 秋天温暖
- `ocean-calm`: 海洋宁静
- `cyber`: 科技赛博
- `chinese`: 国风传统
- `apple`: 苹果简约
- `bytedance`: 字节风格

### 图片处理集成

```bash
# 自动生成配图
/wechatwriter:convert article.md --auto-images --image-style professional

# 使用现有封面
/wechatwriter:convert article.md --cover ./cover.jpg --cover-position top

# 智能封面生成
/wechatwriter:convert article.md --smart-cover --based-on title --style minimalist

# 配图优化
/wechatwriter:convert article.md --optimize-images --compress --max-width 1920
```

### 写作辅助功能

```bash
# AI写作优化
/wechatwriter:convert article.md --enhance-writing --style dan-koe --improve-clarity

# 移除AI痕迹
/wechatwriter:convert article.md --humanize --intensity medium --preserve-meaning

# 内容重写
/wechatwriter:convert article.md --rewrite --style professional --maintain-key-points

# 多语言转换
/wechatwriter:convert article.md --translate --target-lang en --preserve-style
```

## 智能工作流

### 内容分析与优化

```bash
# 深度内容分析
/wechatwriter:convert article.md --analyze --report detailed

# SEO优化转换
/wechatwriter:convert article.md --optimize-seo --keywords "AI,教育,未来" --meta-description auto

# 可读性优化
/wechatwriter:convert article.md --improve-readability --target-level "grade-8" --simplify-technical

# 互动元素添加
/wechatwriter:convert article.md --add-interaction --polls --questions --cta
```

### 格式兼容性处理

```bash
# 多格式输出
/wechatwriter:convert article.md --formats "html,pdf,docx" --parallel

# 微信公众号专用格式
/wechatwriter:convert article.md --target wechat --optimize-mobile --compress-images

# 邮件格式转换
/wechatwriter:convert article.md --target email --template newsletter --responsive

# 社交媒体适配
/wechatwriter:convert article.md --social "weibo,zhihu,jianshu" --auto-truncate --add-hashtags
```

## 专业模式配置

### 企业级转换

```bash
# 品牌一致性转换
/wechatwriter:convert article.md --brand-guidelines brand.json --approval-workflow --multi-review

# 批量企业文档
/wechatwriter:convert enterprise-docs/*.docx --template corporate --output-format html --style-guide strict

# 合规性检查
/wechatwriter:convert article.md --compliance-check --industry "finance" --regulations "SOX,GDPR"
```

### 技术文档处理

```bash
# API文档转换
/wechatwriter:convert api-spec.yaml --to developer-guide --include-examples --code-highlighting

# 代码文档生成
/wechatwriter:convert source-code/ --generate-docs --format markdown --include-readme

# 技术博客优化
/wechatwriter:convert tech-article.md --add-diagrams --code-snippets --performance-charts
```

## 工作流程集成

### 完整内容生产流程

```bash
#!/bin/bash
# 智能内容转换发布流程

# 1. 内容预处理
/wechatwriter:convert raw-content.md --preprocess --normalize-format > prepared.md

# 2. 智能主题选择
THEME=$(/wechatwriter:convert prepared.md --smart-theme --return-suggestion)
echo "推荐主题: $THEME"

# 3. 内容转换优化
/wechatwriter:convert prepared.md --theme "$THEME" --enhance-writing --auto-images > optimized.html

# 4. 质量检查
SCORE=$(/wechatwriter:score --input optimized.html --domain technology | jq -r '.viral_score')
echo "内容评分: $SCORE"

# 5. 生成配套内容
/wechatwriter:generate cover "文章标题" --style "$THEME" > cover.jpg
/wechatwriter:generate summary --style concise optimized.html > summary.txt

# 6. 多平台适配
/wechatwriter:convert optimized.html --target wechat --mobile-optimize > wechat_version.html
/wechatwriter:convert optimized.html --target weibo --truncate 2000 > weibo_version.txt

# 7. 发布到草稿箱
/wechatwriter:draft create --content wechat_version.html --cover cover.jpg --summary summary.txt
```

### 自动化转换管道

```bash
# 监控文件夹自动转换
inotifywait -m ./source/ -e create -e moved_to |
while read path action file; do
  if [[ "$file" =~ \.md$ ]]; then
    echo "检测到新文件: $path$file"
    
    # 自动转换处理
    /wechatwriter:convert "$path$file" \
      --smart-theme \
      --auto-images \
      --enhance-writing \
      --publish
      
    echo "转换完成: $file"
  fi
done
```

## 参数详解

### 基础转换参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| file | 输入文件 | string | 是 |
| --to | 目标格式 | string | 否 |
| --theme | 主题风格 | string | 否 |
| --publish | 直接发布 | boolean | 否 |
| --batch | 批量模式 | boolean | 否 |

### 智能优化参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| --smart-theme | 智能主题推荐 | boolean | 否 |
| --auto-images | 自动生成配图 | boolean | 否 |
| --enhance-writing | AI写作优化 | boolean | 否 |
| --humanize | 移除AI痕迹 | boolean | 否 |
| --optimize-seo | SEO优化 | boolean | 否 |

### 高级功能参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| --target | 目标平台 | string | 否 |
| --formats | 输出格式 | string | 否 |
| --brand-guidelines | 品牌规范 | string | 否 |
| --compliance-check | 合规检查 | boolean | 否 |
| --approval-workflow | 审批流程 | boolean | 否 |

## 转换质量保障

### 内容质量保证
- 智能语义分析确保转换准确性
- 格式兼容性检查避免信息丢失
- 多轮优化迭代提升内容质量
- 人工审核接口支持最终确认

### 技术可靠性
- 异常处理机制确保稳定性
- 版本回滚功能保障安全性
- 增量转换支持提高效率
- 缓存机制优化性能表现

### 平台适配性
- 自动检测目标平台要求
- 智能调整内容格式规范
- 优化显示效果提升体验
- 兼容性测试确保正常发布

## 常见问题

### Q: 转换后的格式不符合预期怎么办？

A: 系统提供多种调整选项：
- 手动指定目标格式参数
- 使用自定义主题配置
- 启用详细模式查看处理过程
- 回滚到之前版本重新调整

### Q: 如何处理特殊内容元素？

A: 支持丰富的内容元素：
- 复杂表格智能转换
- 代码块语法高亮保持
- 数学公式正确渲染
- 多媒体内容适配处理

### Q: 批量转换的性能如何？

A: 优化的批量处理机制：
- 并发处理提升效率
- 内存管理避免溢出
- 进度显示便于监控
- 错误隔离确保完成度

## 最佳实践

1. **预处理准备**: 转换前先标准化原始内容格式
2. **主题一致性**: 保持项目内主题风格统一
3. **质量检查**: 转换后进行必要的审核和优化
4. **版本管理**: 建立内容版本追踪机制
5. **自动化集成**: 将转换流程集成到工作流中