---
description: "AI驱动的风格化写作助手,支持多种创作者风格和封面生成"
---

# wechatwriter智能写作助手

AI驱动的写作助手，支持多种创作者风格，提供智能化的内容创作体验。

## 快速开始

### 交互式写作

```bash
/wechatwriter:write
```

进入交互模式，AI会引导你完成写作过程。

### 指定风格写作

```bash
/wechatwriter:write "关于自律的思考" --style dan-koe
```

## 写作风格

### Dan Koe风格（默认）

深刻、犀利、接地气的表达方式：

```bash
/wechatwriter:write "如何在信息过载时代保持专注" --style dan-koe
```

特点：
- 直击问题本质
- 实用性强
- 语言简洁有力
- 逻辑清晰

### 故事叙述风格

```bash
/wechatwriter:write "创业路上的挑战" --style story
```

特点：
- 情节引人入胜
- 情感共鸣强
- 细节丰富生动
- 结构完整

### 专业分析风格

```bash
/wechatwriter:write "区块链技术发展趋势" --style professional
```

特点：
- 数据支撑充分
- 逻辑严谨
- 术语准确
- 结构规范

### 创意表达风格

```bash
/wechatwriter:write "未来生活方式的想象" --style creative
```

特点：
- 想象力丰富
- 语言生动活泼
- 视角独特新颖
- 感染力强

## 高级功能

### 内容优化

```bash
# 精炼现有内容
/wechatwriter:write article.md --input-type fragment --style dan-koe

# 扩展简短想法
/wechatwriter:write "AI会改变教育" --input-type idea --expand
```

输入类型：
- `fragment`: 内容片段
- `idea`: 简单想法
- `outline`: 大纲结构
- `draft`: 初稿版本

### 封面生成

```bash
/wechatwriter:write "时间管理的重要性" --style dan-koe --cover
```

自动生成与内容匹配的封面图片。

### AI痕迹去除

```bash
# 轻度处理
/wechatwriter:write "主题" --style dan-koe --humanize

# 深度处理
/wechatwriter:write "主题" --style dan-koe --humanize --humanize-intensity aggressive
```

强度选项：
- `gentle`: 温和处理
- `medium`: 中等强度（默认）
- `aggressive`: 深度处理

## 工作流程集成

### 完整创作流程

```bash
#!/bin/bash
# 1. 研究话题
/wechatwriter:research "远程办公趋势" > research.json

# 2. 生成大纲
/wechatwriter:outline "远程办公的未来" --keywords "远程办公,工作方式,企业管理" > outline.json

# 3. 基于大纲写作
/wechatwriter:write --input-type outline outline.json --style dan-koe > article.md

# 4. 去除AI痕迹
/wechatwriter:humanize article.md --intensity medium > article_humanized.md

# 5. 生成封面
/wechatwriter:generate cover "远程办公的未来趋势" > cover.jpg
```

### 批量内容生成

```bash
# 批量生成系列文章
for topic in "时间管理" "目标设定" "习惯养成"; do
  /wechatwriter:write "$topic" --style dan-koe --cover > "${topic}.md"
done
```

## 参数详解

### 基础参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| input | 输入内容或文件 | string | 否 |
| --style | 写作风格 | string | 否 |
| --input-type | 输入类型 | string | 否 |
| --cover | 生成封面 | boolean | 否 |
| --humanize | 去除AI痕迹 | boolean | 否 |
| --humanize-intensity | 处理强度 | string | 否 |
| --expand | 扩展内容 | boolean | 否 |
| --output | 输出文件 | string | 否 |

### 风格对比

| 风格 | 适用场景 | 语言特点 | 读者群体 |
|------|----------|----------|----------|
| dan-koe | 深度思考、个人成长 | 简洁有力、直击本质 | 知识工作者 |
| story | 品牌故事、案例分享 | 生动叙述、情感共鸣 | 大众读者 |
| professional | 行业分析、技术解读 | 专业严谨、数据支撑 | 专业人士 |
| creative | 未来趋势、创新想法 | 想象丰富、语言活泼 | 年轻群体 |

## 实用技巧

### 1. 明确写作目的

```bash
# 教育启发
/wechatwriter:write "如何培养批判性思维" --style dan-koe

# 娱乐消遣
/wechatwriter:write "办公室趣事分享" --style story

# 专业分享
/wechatwriter:write "数据分析方法论" --style professional
```

### 2. 内容长度控制

```bash
# 短篇快文
/wechatwriter:write "主题" --style dan-koe --length short

# 中篇深度
/wechatwriter:write "主题" --style dan-koe --length medium

# 长篇详解
/wechatwriter:write "主题" --style dan-koe --length long
```

### 3. 多轮优化

```bash
# 初稿生成
/wechatwriter:write "初稿内容" --style dan-koe > draft1.md

# 精炼优化
/wechatwriter:write draft1.md --input-type fragment --refine > draft2.md

# 最终润色
/wechatwriter:humanize draft2.md > final.md
```

## 常见问题

### Q: 写作质量如何保证？

A: 系统提供：
- 智能内容结构规划
- 语法和逻辑检查
- 风格一致性维护
- 可读性优化

### Q: 支持哪些内容类型？

A: 支持多种内容形式：
- 观点文章
- 教程指南
- 产品评测
- 故事叙述
- 新闻评论

### Q: 如何保持内容原创性？

A: 通过：
- AI痕迹去除功能
- 个性化风格调整
- 原创性检测
- 内容重构优化

## 最佳实践

1. **明确目标受众**: 根据读者群体选择合适风格
2. **提供充分背景**: 给AI足够的上下文信息
3. **多轮迭代**: 不要期望一次完美，多次优化
4. **人工审核**: AI辅助但保持人工判断
5. **持续学习**: 根据效果反馈调整使用方法