---
name: content-writing
description: 微信公众号文章写作工具，支持内容框架生成、风格化写作、AI去痕和格式转换
argument-hint: "[写作任务或内容需求]"
allowed-tools: Bash(./scripts/wechatwriter *), Read, Write
---

# 微信公众号内容写作工具

专业的微信公众号文章写作工具，提供内容框架生成、多风格写作、AI去痕和格式转换功能。

## 核心功能

### 1. 内容框架生成

生成文章框架结构，支持多种模板：

```bash
./scripts/write.sh outline --topic "主题" --template [模板]
```

**可用模板**: authoritative(权威型), comparison(对比型), cultural(文化型), practical(实用型)

### 2. 风格化写作

基于指定风格进行文章创作：

```bash
./scripts/write.sh write --topic "主题" --style [风格]
```

**可用风格**:

- dan-koe: 简洁有力，金句频出
- cultural-depth: 引经据典，文化底蕴
- casual-science: 通俗易懂，趣味性强

### 3. AI去痕优化

去除AI生成痕迹，提升内容自然度：

```bash
./scripts/write.sh humanize [文件] --intensity [强度]
```

**处理强度**: gentle(温和), medium(中等), aggressive(深度)

### 4. 格式转换

Markdown转微信HTML格式：

```bash
./scripts/write.sh convert [markdown文件] --theme [主题]
```

**可用主题**: default, autumn-warm, spring-fresh, ocean-calm, custom

## 使用示例

### 完整写作流程

```bash
# 1. 生成框架
./scripts/write.sh outline --topic "人工智能对职场的影响" --template authoritative

# 2. 基于框架写作
./scripts/write.sh write --topic "人工智能对职场的影响" --style dan-koe

# 3. 去痕处理（可选）
./scripts/write.sh humanize draft.md --intensity medium

# 4. 格式转换
./scripts/write.sh convert humanized.md --theme elegant
```

### 快速写作

```bash
# 直接写作指定主题
./scripts/write.sh write --topic "如何提升工作效率" --style dan-koe

# 转换现有Markdown
./scripts/write.sh convert existing-content.md --theme default
```

## 输出格式

### 框架输出

```json
{
  "title": "文章标题",
  "hook": "开头引入",
  "sections": [
    {
      "heading": "段落标题",
      "key_points": ["要点1", "要点2"]
    }
  ],
  "conclusion": "总结"
}
```

### 文章输出

```markdown
# 文章标题
## 副标题

[开头段落]

## 第一部分
[详细内容]

## 总结
[核心观点]
```

## 质量标准

- **原创性**: 确保内容原创，避免抄袭
- **可读性**: 语言流畅，结构清晰
- **合规性**: 符合微信公众号规范
- **实用性**: 提供实际价值和可操作建议

## 注意事项

- 使用正版素材，避免版权问题
- 根据目标受众选择合适风格
- 保持内容真实性和专业性
- 定期更新写作策略和模板
