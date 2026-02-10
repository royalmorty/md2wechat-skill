---
description: "多维度爆款潜力评分工具,支持多平台评估和优化建议"
---

# wechatwriter爆款潜力评分

基于多维度数据分析，智能评估内容的爆款潜力，提供科学的传播预测。

## 快速评分

### 基础评分

```bash
/wechatwriter:score --read-count 10000 --like-count 500 --share-count 200
```

### 从数据文件评分

```bash
/wechatwriter:score --input metrics.json
```

## 评分维度

### 基础指标

**互动数据**
- 阅读量 (read_count): 内容曝光度
- 点赞量 (like_count): 用户认可度
- 分享量 (share_count): 传播意愿
- 评论量 (comment_count): 互动深度
- 收藏量 (collect_count): 价值认知

**衍生指标**
- 互动率 = (点赞+分享+评论+收藏) / 阅读量
- 分享率 = 分享量 / 阅读量
- 点赞率 = 点赞量 / 阅读量
- 评论率 = 评论量 / 阅读量

### 平台适配

```bash
# 微信公众号评分
/wechatwriter:score --platform wechat --read-count 50000 --like-count 2000

# 抖音评分
/wechatwriter:score --platform douyin --play-count 100000 --like-count 8000

# 小红书评分
/wechatwriter:score --platform xiaohongshu --read-count 20000 --like-count 3000
```

## 高级分析

### 领域专业评分

```bash
/wechatwriter:score \
  --domain technology \
  --topic "人工智能发展趋势" \
  --read-count 15000 \
  --like-count 800 \
  --share-count 400 \
  --comment-count 200
```

支持领域：
- `technology`: 科技数码
- `business`: 商业财经
- `lifestyle`: 生活方式
- `education`: 教育学习
- `health`: 健康养生
- `entertainment`: 娱乐八卦
- `fashion`: 时尚美妆
- `food`: 美食料理
- `travel`: 旅游出行

### 竞品对比分析

```bash
# 分析竞品数据
/wechatwriter:score --benchmark competitors.json --input my_article.json
```

### 趋势预测

```bash
# 基于历史数据预测
/wechatwriter:score --trend-analysis --history history_data.json
```

## 评分输出

### 详细评分报告

```json
{
  "viral_score": 85.5,
  "viral_level": "high_potential",
  "metrics": {
    "engagement_rate": 0.15,
    "share_rate": 0.08,
    "like_rate": 0.04,
    "comment_rate": 0.02
  },
  "analysis": {
    "strengths": ["高分享率", "良好互动"],
    "weaknesses": ["评论率偏低"],
    "opportunities": ["增加互动引导"],
    "threats": ["同类内容竞争激烈"]
  },
  "recommendations": [
    "增加评论区互动",
    "优化标题吸引力",
    "选择合适发布时间"
  ]
}
```

### 评分等级

| 分数范围 | 等级 | 说明 |
|----------|------|------|
| 90-100 | 爆款 | 极高传播潜力 |
| 80-89 | 优质 | 良好传播效果 |
| 70-79 | 普通 | 中等传播表现 |
| 60-69 | 一般 | 较低传播效果 |
| 0-59 | 待优化 | 需要改进 |

## 优化建议

### 内容优化

```bash
# 获取具体优化建议
/wechatwriter:score --optimize-suggestions --input article.json
```

常见建议：
- **标题优化**: 增加吸引力词汇
- **内容结构**: 调整段落安排
- **互动引导**: 增加用户参与
- **发布时间**: 选择最佳时段

### A/B测试支持

```bash
# 版本对比
/wechatwriter:score --ab-test version_a.json version_b.json
```

## 工作流程集成

### 内容生产流程

```bash
#!/bin/bash
# 1. 内容创作
/wechatwriter:write "人工智能对教育的影响" --style dan-koe > article.md

# 2. 质量评分
/wechatwriter:score --input article.md --domain technology > score.json

# 3. 优化建议
if [ $(jq -r '.viral_score' score.json) -lt 70 ]; then
  echo "需要优化，查看建议:"
  jq -r '.recommendations[]' score.json
fi

# 4. 应用优化
/wechatwriter:optimize --suggestions score.json --input article.md > optimized.md

# 5. 重新评分
/wechatwriter:score --input optimized.md --domain technology
```

### 批量内容评估

```bash
# 评估内容库
for file in articles/*.md; do
  score=$(/wechatwriter:score --input "$file" --domain technology | jq -r '.viral_score')
  echo "$file: $score"
done | sort -k2 -nr
```

## 参数详解

### 基础参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| --read-count | 阅读量 | number | 是 |
| --like-count | 点赞量 | number | 是 |
| --share-count | 分享量 | number | 是 |
| --comment-count | 评论量 | number | 否 |
| --collect-count | 收藏量 | number | 否 |
| --platform | 平台类型 | string | 否 |
| --domain | 内容领域 | string | 否 |
| --topic | 内容主题 | string | 否 |

### 高级参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| --input | 数据文件 | string | 否 |
| --benchmark | 竞品数据 | string | 否 |
| --history | 历史数据 | string | 否 |
| --optimize-suggestions | 优化建议 | boolean | 否 |
| --trend-analysis | 趋势分析 | boolean | 否 |
| --ab-test | A/B测试 | string | 否 |

## 平台基准数据

### 微信公众号

| 指标 | 优秀 | 良好 | 一般 |
|------|------|------|------|
| 阅读量 | 10W+ | 5W-10W | 1W-5W |
| 点赞率 | >5% | 3-5% | 1-3% |
| 分享率 | >3% | 1-3% | 0.5-1% |
| 评论率 | >1% | 0.5-1% | 0.1-0.5% |

### 抖音短视频

| 指标 | 优秀 | 良好 | 一般 |
|------|------|------|------|
| 播放量 | 100W+ | 50W-100W | 10W-50W |
| 点赞率 | >8% | 5-8% | 2-5% |
| 分享率 | >2% | 1-2% | 0.5-1% |
| 评论率 | >2% | 1-2% | 0.5-1% |

## 常见问题

### Q: 评分准确性如何？

A: 基于：
- 大数据分析模型
- 多维度综合评估
- 行业基准对比
- 机器学习优化

### Q: 如何提高评分？

A: 系统提供：
- 具体优化建议
- 竞品对比分析
- 内容改进方向
- 发布策略指导

### Q: 支持哪些平台？

A: 支持主流平台：
- 微信公众号
- 抖音短视频
- 小红书
- 今日头条
- 新浪微博

## 最佳实践

1. **定期评估**: 建立内容评分档案
2. **对比分析**: 与竞品持续对比
3. **数据积累**: 收集历史表现数据
4. **策略调整**: 根据评分优化策略
5. **效果跟踪**: 验证优化效果