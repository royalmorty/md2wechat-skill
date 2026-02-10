---
description: "AI智能大纲生成,支持多领域和多模板类型"
---

# wechatwriter智能大纲生成

基于AI技术生成高质量内容大纲，支持多领域、多风格的内容结构规划。

## 基础用法

### 生成内容大纲

```bash
/wechatwriter:outline "人工智能对未来教育的影响"
```

### 指定内容领域

```bash
/wechatwriter:outline "区块链技术的应用" --domain technology
```

支持领域：
- `technology`: 科技技术
- `business`: 商业经济  
- `lifestyle`: 生活方式
- `education`: 教育学习
- `health`: 健康养生
- `culture`: 文化艺术

## 高级功能

### 使用模板生成

```bash
/wechatwriter:outline "如何提高工作效率" --template listicle
```

可用模板：
- `listicle`: 清单体文章
- `howto`: 教程指南类
- `story`: 故事叙述类
- `analysis`: 深度分析类
- `news`: 新闻报道类
- `review`: 评测对比类

### 指定写作风格

```bash
/wechatwriter:outline "数字化转型的挑战" --style dan-koe
```

风格选项：
- `dan-koe`: 深刻犀利风格
- `professional`: 专业严谨风格
- `casual`: 轻松随意风格
- `academic`: 学术理论风格
- `viral`: 爆款传播风格

### 关键词优化

```bash
/wechatwriter:outline "新能源汽车发展趋势" --keywords "电动化,智能化,自动驾驶"
```

## 输出格式

生成的内容大纲包含：

```json
{
  "title": "主标题",
  "subtitle": "副标题", 
  "hook": "开头钩子",
  "sections": [
    {
      "title": "段落标题",
      "content": "内容要点",
      "key_points": ["要点1", "要点2"],
      "engagement": "互动元素"
    }
  ],
  "key_points": ["核心观点"],
  "call_to_action": "行动号召",
  "viral_elements": ["传播要素"],
  "seo_keywords": ["SEO关键词"]
}
```

## 实用案例

### 案例1：科技评论文章

```bash
/wechatwriter:outline "元宇宙的商业前景分析" \
  --domain technology \
  --style professional \
  --template analysis \
  --keywords "VR,AR,虚拟经济,投资机会"
```

### 案例2：生活方式清单

```bash
/wechatwriter:outline "提升生活品质的10个习惯" \
  --domain lifestyle \
  --style dan-koe \
  --template listicle
```

### 案例3：教育深度文章

```bash
/wechatwriter:outline "在线教育的未来形态" \
  --domain education \
  --style academic \
  --keywords "远程学习,个性化教育,AI教学"
```

## 工作流程集成

### 结合写作命令

```bash
# 1. 生成大纲
/wechatwriter:outline "主题" > outline.json

# 2. 基于大纲写作
/wechatwriter:write --input-type outline outline.json
```

### 结合研究命令

```bash
# 1. 研究热门话题
/wechatwriter:research "人工智能"

# 2. 生成内容大纲
/wechatwriter:outline "研究发现" --keywords "AI趋势"
```

## 参数说明

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| topic | 内容主题 | string | 是 |
| --domain | 内容领域 | string | 否 |
| --template | 文章模板 | string | 否 |
| --style | 写作风格 | string | 否 |
| --keywords | SEO关键词 | string | 否 |
| --output | 输出文件 | string | 否 |

## 最佳实践

1. **明确主题**: 提供具体清晰的主题描述
2. **选择合适领域**: 根据内容类型选择最匹配的领域
3. **使用关键词**: 添加相关关键词提升SEO效果
4. **模板匹配**: 根据内容目的选择合适模板
5. **风格一致**: 保持品牌调性的风格一致性