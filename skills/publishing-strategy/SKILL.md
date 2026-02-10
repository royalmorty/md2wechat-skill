---
name: publishing-strategy
description: 微信公众号发布策略，制定发布计划、管理草稿、优化发布时间、分析数据表现，集成草稿管理工具，提供数据驱动的发布策略建议
allowed-tools: Bash(./scripts/publish.sh *), WebSearch, WebFetch, Read, Write
argument-hint: "[发布策略需求]"
---

# 微信公众号发布策略专家

当用户需要制定发布计划、管理草稿、优化发布时间、分析数据表现等发布相关问题时，自动激活此技能提供专业的发布策略建议。支持手动调用 `/publishing-strategy [发布策略需求]`。

## 核心功能

### 1. 草稿管理工具集成 (`publish.sh draft`)
- 创建、保存、管理和版本控制草稿
- 支持草稿的分类、标签和搜索功能
- 提供草稿的版本历史和协作编辑
- 支持草稿的定时发布和自动备份

### 2. 发布计划制定
- 基于数据制定最优的发布时间表
- 考虑受众活跃时间和平台算法规律
- 平衡内容质量和发布频率
- 提供发布日历和提醒功能

### 3. 发布时间优化
- 分析历史数据找出最佳发布时间
- 考虑不同内容类型的最佳发布时段
- 适配目标受众的在线习惯
- 提供A/B测试和持续优化建议

### 4. 数据表现分析
- 跟踪和分析发布后的数据表现
- 评估不同发布策略的效果
- 提供数据驱动的优化建议
- 建立发布效果评估体系

### 5. 内容排期管理
- 制定长期的内容发布计划
- 协调不同主题和类型的内容搭配
- 管理内容发布节奏和频率
- 提供内容日历和进度跟踪

## 草稿管理系统

### 草稿操作命令
```bash
# 创建新草稿
./scripts/publish.sh draft create "文章标题" --content content.md

# 保存草稿
./scripts/publish.sh draft save draft-id --content updated-content.md

# 列出所有草稿
./scripts/publish.sh draft list --status all --sort updated

# 获取草稿详情
./scripts/publish.sh draft get draft-id --include-history

# 删除草稿
./scripts/publish.sh draft delete draft-id --confirm
```

### 草稿状态管理
- **草稿（draft）**：正在编辑中的内容
- **待发布（scheduled）**：已排期等待发布的内容
- **已发布（published）**：已发布的内容
- **已归档（archived）**：历史内容归档

### 草稿标签和分类
```bash
# 添加标签
./scripts/publish.sh draft tag draft-id --add "茶文化,春茶,选购指南"

# 设置分类
./scripts/publish.sh draft categorize draft-id --category "文化科普"

# 搜索草稿
./scripts/publish.sh draft search --tags "茶文化" --category "文化科普"
```

## 发布计划制定

### 发布时间表优化
```bash
# 分析最佳发布时间
./scripts/publish.sh schedule analyze --content-type "文化科普" --audience "茶文化爱好者"

# 制定发布计划
./scripts/publish.sh schedule create --content-id draft-id --time "2024-03-15 14:00" --timezone "Asia/Shanghai"
```

### 发布频率建议
- **高频发布**：每日1-2次，适合新闻资讯类
- **中频发布**：每周3-4次，适合知识分享类
- **低频发布**：每周1-2次，适合深度内容类
- **混合发布**：根据内容类型调整频率

### 发布时机考虑因素
- **受众活跃时间**：分析目标受众的在线习惯
- **平台算法规律**：了解微信推荐机制
- **内容时效性**：考虑内容的时效要求
- **竞争环境**：避开竞品密集发布时段

## 数据表现分析

### 关键指标跟踪
- **阅读量**：文章的阅读次数和完读率
- **点赞数**：用户对内容的认可度
- **分享数**：内容的传播力和影响力
- **评论数**：用户的参与度和互动性
- **收藏数**：内容的实用价值和保存率

### 发布效果评估
```bash
# 分析发布数据
./scripts/publish.sh analytics draft-id --metrics "views,likes,shares,comments" --period 7d

# 对比不同发布时间的效果
./scripts/publish.sh analytics compare --content-ids "draft1,draft2" --metrics "views,engagement"
```

### 优化建议生成
基于数据分析提供以下优化建议：
- **时间优化**：调整发布时间以获得更好的曝光
- **内容优化**：改进内容质量和形式
- **标题优化**：优化标题以提高点击率
- **互动优化**：增强与读者的互动

## 完整发布工作流

### 标准发布流程
```bash
# 1. 创建草稿
./scripts/publish.sh draft create "春茶选购指南" --content content.md --tags "茶文化,春茶" --category "文化科普"

# 2. 分析最佳发布时间
./scripts/publish.sh schedule analyze --content-type "文化科普" --audience "茶文化爱好者"

# 3. 制定发布计划
./scripts/publish.sh schedule create --content-id draft-123 --time "2024-03-15 14:00"

# 4. 跟踪发布效果
./scripts/publish.sh analytics draft-123 --metrics "views,likes,shares" --period 7d

# 5. 优化后续策略
./scripts/publish.sh analytics compare --content-ids "draft-123,draft-124" --metrics "engagement"
```

### 批量发布管理
```bash
# 批量创建发布计划
for draft in $(./scripts/publish.sh draft list --status draft --format ids); do
    ./scripts/publish.sh schedule create --content-id $draft --time "2024-03-$(date +%d) 14:00"
done

# 批量分析发布效果
./scripts/publish.sh analytics batch --content-ids "$(./scripts/publish.sh draft list --status published --format ids)" --metrics "views,engagement"
```

### 内容排期优化
```bash
# 制定月度内容计划
./scripts/publish.sh schedule monthly --month 2024-03 --content-types "文化科普:2,产品评测:1,行业分析:1"

# 调整发布频率
./scripts/publish.sh schedule frequency --frequency "每周3次" --distribution "周二四六"
```

## 输出格式

### 发布计划报告
```markdown
# 发布计划报告

## 发布时间表
- **3月15日 14:00**：春茶选购指南（文化科普类）
- **3月18日 10:00**：AI在茶行业的应用（技术分析类）
- **3月22日 16:00**：茶文化的现代传承（文化深度类）

## 发布策略
- **目标受众**：茶文化爱好者，年龄25-45岁
- **发布频率**：每周3次，周二四六
- **最佳时间**：下午2-4点，晚上7-9点

## 预期效果
- **预计阅读量**：5000-8000/篇
- **预计互动率**：3-5%
- **预计分享率**：1-2%
```

### 草稿管理报告
```markdown
# 草稿管理报告

## 草稿统计
- **总草稿数**：45篇
- **待发布**：8篇
- **已发布**：37篇
- **本月新增**：12篇

## 草稿分类
- **文化科普**：20篇（44%）
- **产品评测**：15篇（33%）
- **行业分析**：10篇（23%）

## 草稿状态
- **高质量草稿**：15篇（33%）
- **需要优化**：20篇（44%）
- **待完善**：10篇（23%）
```

### 数据表现分析
```markdown
# 发布效果分析报告

## 整体表现（最近30天）
- **总阅读量**：156,789次
- **平均阅读量**：4,893次/篇
- **总互动数**：8,456次（点赞+评论+分享）
- **平均互动率**：5.4%

## 最佳表现内容
1. **春茶选购指南**（3月15日）
   - 阅读量：12,456
   - 互动率：8.2%
   - 分享数：234

2. **茶文化的现代传承**（3月22日）
   - 阅读量：9,876
   - 互动率：6.8%
   - 分享数：187

## 优化建议
- **发布时间**：下午2-4点效果最佳
- **内容类型**：文化科普类表现优于技术分析类
- **标题优化**：包含"指南"、"如何"等关键词的文章表现更好
```

## 高级功能

### A/B测试
- 测试不同发布时间的效果
- 对比不同内容类型的表现
- 评估不同标题策略的影响
- 分析不同互动方式的效果

### 自动化发布
- 支持定时自动发布功能
- 基于数据自动优化发布时间
- 智能推荐发布频率
- 自动跟踪和分析发布效果

### 预测分析
- 基于历史数据预测发布效果
- 识别最佳发布时机
- 预测内容生命周期
- 提供风险预警和建议

## 注意事项
- 确保草稿内容的完整性和准确性
- 合理安排发布时间，避免过度集中
- 持续监控发布效果，及时调整策略
- 重视用户反馈，优化内容质量
- 遵守平台规则，避免违规操作
- 建立长期的内容规划和发布节奏

通过集成草稿管理工具和数据分析功能，我可以为你提供科学、系统的发布策略服务，确保每个内容都能在最佳时机获得最大的曝光和互动效果。