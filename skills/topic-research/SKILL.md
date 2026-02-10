---
name: topic-research
description: 微信公众号选题分析工具，提供热点评估、竞品研究和数据驱动的选题建议
argument-hint: "[研究主题或关键词]"
allowed-tools: Bash(./scripts/wechatwriter *), WebSearch, WebFetch, Read
---

# 微信公众号选题分析工具

专业的微信公众号选题研究工具，通过数据分析提供热点话题评估和竞品内容研究。

## 核心功能

### 1. 话题研究

分析特定主题的研究数据：

```bash
./scripts/research.sh research --topic "主题" --domain [领域]
```

### 2. 热点评分

评估话题的爆款潜力：

```bash
./scripts/research.sh score --topic "主题" --metrics [数据文件] --domain [领域]
```

### 3. 趋势分析

分析主题趋势变化：

```bash
./scripts/research.sh trends --domain [领域] --platform [平台]
```

### 4. 内容框架生成

基于研究生成内容框架：

```bash
./scripts/research.sh outline --topic "主题" --template [模板] --domain [领域]
```

**可用模板**: viral(爆款), authoritative(权威), comparison(对比), cultural(文化), practical(实用)
**可用领域**: general, tech, lifestyle, culture, business, education

## 使用示例

### 基础选题分析

```bash
# 研究话题
./scripts/research.sh research --topic "春茶选购指南" --domain lifestyle

# 话题评分（需要提供数据文件）
echo '{"read_count": 5000, "like_count": 300}' > metrics.json
./scripts/research.sh score --topic "春茶选购指南" --metrics metrics.json --domain lifestyle

# 趋势分析
./scripts/research.sh trends --domain lifestyle --platform wechat
```

### 综合选题研究

```bash
# 1. 话题研究
./scripts/research.sh research --topic "职场沟通技巧" --domain business

# 2. 生成内容框架
./scripts/research.sh outline --topic "职场沟通技巧" --template practical --domain business
```

## 输出格式

### 话题研究结果

```json
{
  "topic": "春茶选购指南",
  "domain": "lifestyle",
  "platform": "wechat",
  "hot_topics": [
    {
      "title": "春茶选购指南",
      "heat_score": 85,
      "trend": "rising",
      "keywords": ["春茶", "选购", "指南"]
    }
  ],
  "recommendations": [
    "建议从实用角度切入春茶话题",
    "结合季节性特点进行内容创作"
  ]
}
```

### 评分结果

```json
{
  "topic": "春茶选购指南",
  "score": 78,
  "level": "B",
  "dimensions": {
    "search_heat": 82,
    "social_discussion": 75,
    "timeliness": 85,
    "competition": 70
  },
  "recommendation": "中等热度话题，适合细分受众"
}
```

### 趋势分析结果

```json
{
  "domain": "lifestyle",
  "platform": "wechat",
  "trends": [
    {
      "category": "文化科普",
      "trending_topics": ["传统文化", "现代应用"],
      "engagement_rate": 0.045,
      "growth_potential": "high"
    }
  ],
  "seasonal_trends": {
    "spring": ["新品发布", "春季养生"],
    "summer": ["避暑攻略", "夏季美食"]
  }
}
```

## 数据源

- 微信指数: 搜索热度数据
- 社交媒体: 讨论趋势分析  
- 行业报告: 专业领域洞察
- 竞品监控: 内容表现追踪

## 质量标准

- **数据准确性**: 基于多源数据交叉验证
- **时效性**: 数据更新频率日级别
- **相关性**: 聚焦微信公众号生态
- **可操作性**: 提供具体执行建议

## 注意事项

- 评分结果仅供参考，需结合实际情况判断
- 竞品分析需遵守平台规则，避免过度抓取
- 建议结合多个话题进行对比分析
- 定期更新分析模型以保持准确性
