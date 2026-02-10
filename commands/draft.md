---
description: "微信公众号草稿管理,支持创建、预览和发布到草稿箱"
---

# wechatwriter草稿管理

专业的微信公众号草稿管理工具，支持创建、测试、发布全流程操作。

## 快速开始

### 创建草稿

```bash
/wechatwriter:draft create article.json
```

### 测试草稿HTML

```bash
/wechatwriter:draft test article.json
```

### 发布草稿

```bash
/wechatwriter:draft publish article.json
```

## 详细功能

### 从JSON创建草稿

```bash
/wechatwriter:draft create /path/to/article.json --account official
```

JSON文件格式：

```json
{
  "title": "文章标题",
  "author": "作者名称",
  "content": "HTML内容",
  "digest": "文章摘要",
  "thumb_media_id": "封面图片媒体ID",
  "show_cover_in_text": true,
  "content_source_url": "原文链接",
  "need_open_comment": 1,
  "only_fans_can_comment": 0
}
```

### 测试草稿渲染

```bash
/wechatwriter:draft test article.json --output preview.html
```

生成预览文件检查：
- 格式渲染效果
- 图片显示状态  
- 链接跳转功能
- 移动端适配

### 发布到草稿箱

```bash
/wechatwriter:draft publish article.json \
  --account official \
  --schedule "2024-01-15 10:00"
```

## 高级用法

### 多账号管理

```bash
# 查看可用账号
/wechatwriter:draft list-accounts

# 指定特定账号
/wechatwriter:draft create article.json --account tech-official
```

### 批量操作

```bash
# 批量创建草稿
for file in articles/*.json; do
  /wechatwriter:draft create "$file"
done
```

### 错误处理

```bash
# 详细错误信息
/wechatwriter:draft create article.json --verbose

# 忽略警告继续执行
/wechatwriter:draft publish article.json --force
```

## 工作流程集成

### 完整发布流程

```bash
#!/bin/bash
# 1. 转换文章
/wechatwriter:convert article.md --theme autumn-warm > article.json

# 2. 生成封面
/wechatwriter:generate cover "文章标题" > cover.jpg

# 3. 上传封面
COVER_ID=$(/wechatwriter:image upload cover.jpg | jq -r '.media_id')

# 4. 更新JSON中的封面ID
jq --arg id "$COVER_ID" '.thumb_media_id = $id' article.json > article_final.json

# 5. 创建草稿
/wechatwriter:draft create article_final.json

# 6. 测试预览
/wechatwriter:draft test article_final.json

# 7. 发布草稿
/wechatwriter:draft publish article_final.json
```

### 自动化发布

```bash
# 定时发布脚本
/wechatwriter:draft publish article.json --schedule "$(date -d 'tomorrow 09:00' '+%Y-%m-%d %H:%M')"
```

## 参数说明

### create命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| json_file | JSON文件路径 | string | 是 |
| --account | 目标账号 | string | 否 |
| --verbose | 详细输出 | boolean | 否 |

### test命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| json_file | JSON文件路径 | string | 是 |
| --output | 输出HTML文件 | string | 否 |

### publish命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| json_file | JSON文件路径 | string | 是 |
| --account | 目标账号 | string | 否 |
| --schedule | 定时发布 | string | 否 |
| --force | 强制发布 | boolean | 否 |

## 常见问题

### Q: 草稿创建失败怎么办？

A: 检查以下几点：
- JSON格式是否正确
- 必填字段是否完整
- 媒体文件是否已上传
- API权限是否正常

### Q: 如何获取thumb_media_id？

A: 使用图片上传命令：
```bash
/wechatwriter:image upload cover.jpg
```

### Q: 支持哪些内容格式？

A: 支持标准微信公众号内容格式：
- HTML标签（div, p, h1-h6, img等）
- 内联样式
- 图片引用
- 超链接

## 最佳实践

1. **内容准备**: 确保JSON文件格式正确，内容完整
2. **预览测试**: 发布前务必进行预览测试
3. **账号选择**: 根据内容类型选择合适的发布账号
4. **定时发布**: 选择最佳发布时间提高阅读量
5. **错误处理**: 建立完善的错误处理机制