---
description: "图片全流程处理工具,支持上传、AI生成、压缩和批量操作"
---

# wechatwriter图片处理工具

专业的图片处理解决方案，支持上传、下载、生成、压缩等全流程图片管理。

## 快速开始

### 上传本地图片

```bash
/wechatwriter:image upload ./photo.jpg
```

### 下载网络图片

```bash
/wechatwriter:image download "https://example.com/image.jpg"
```

### AI生成图片

```bash
/wechatwriter:image generate "科技感的未来城市" --style cyberpunk
```

## 核心功能

### 图片上传

```bash
# 基础上传
/wechatwriter:image upload ./local-image.jpg

# 指定压缩
/wechatwriter:image upload ./large-image.jpg --compress --max-width 1920

# 批量上传
/wechatwriter:image upload ./images/*.jpg --batch
```

上传成功后返回：
```json
{
  "success": true,
  "media_id": "media_id_123456",
  "url": "http://mmbiz.qpic.cn/...",
  "size": {
    "width": 1920,
    "height": 1080,
    "file_size": 524288
  }
}
```

### 网络图片处理

```bash
# 下载并上传
/wechatwriter:image download "https://example.com/image.jpg"

# 带压缩处理
/wechatwriter:image download "https://example.com/large-image.jpg" --compress --quality 85

# 批量下载
/wechatwriter:image download --batch urls.txt
```

### AI图片生成

```bash
# 基础生成
/wechatwriter:image generate "极简主义的产品展示图"

# 指定风格
/wechatwriter:image generate "未来科技概念图" --style cyberpunk --size 16:9

# 高质量生成
/wechatwriter:image generate "商业会议场景" --quality hd --aspect-ratio 4:3
```

支持风格：
- `realistic`: 写实风格
- `cyberpunk`: 赛博朋克
- `minimalist`: 极简主义
- `cartoon`: 卡通风格
- `professional`: 商务专业
- `creative`: 创意艺术

## 高级功能

### 智能压缩

```bash
# 自动压缩（大于1920px的图片）
/wechatwriter:image upload ./photo.jpg --auto-compress

# 自定义压缩参数
/wechatwriter:image upload ./photo.jpg --compress --quality 90 --max-width 1600
```

压缩选项：
- `--quality`: 压缩质量 (70-95)
- `--max-width`: 最大宽度
- `--max-height`: 最大高度
- `--format`: 输出格式 (jpg, png, webp)

### 图片预处理

```bash
# 自动调整尺寸
/wechatwriter:image upload ./photo.jpg --resize --width 800 --height 600

# 智能裁剪
/wechatwriter:image upload ./photo.jpg --smart-crop --aspect-ratio 1:1

# 格式转换
/wechatwriter:image upload ./photo.png --format jpg
```

### 批量处理

```bash
# 批量上传文件夹
/wechatwriter:image upload ./assets/images/ --recursive

# 批量压缩
/wechatwriter:image compress ./images/*.jpg --output ./compressed/

# 批量格式转换
/wechatwriter:image convert ./images/*.png --format jpg
```

## 工作流程集成

### 文章配图完整流程

```bash
#!/bin/bash
# 1. 生成封面
/wechatwriter:image generate "人工智能教育" --style professional --size 16:9 > cover.jpg

# 2. 压缩优化
/wechatwriter:image compress cover.jpg --quality 85 --max-width 1920 > cover_optimized.jpg

# 3. 上传到微信
COVER_RESULT=$(/wechatwriter:image upload cover_optimized.jpg)
COVER_MEDIA_ID=$(echo $COVER_RESULT | jq -r '.media_id')

# 4. 生成配图
/wechatwriter:image generate "AI学习场景" --style realistic > illustration.jpg

# 5. 处理并上传配图
/wechatwriter:image compress illustration.jpg --max-width 1200 > illustration_optimized.jpg
ILLUSTRATION_RESULT=$(/wechatwriter:image upload illustration_optimized.jpg)
ILLUSTRATION_MEDIA_ID=$(echo $ILLUSTRATION_RESULT | jq -r '.media_id')

echo "封面ID: $COVER_MEDIA_ID"
echo "配图ID: $ILLUSTRATION_MEDIA_ID"
```

### 自动化图片处理

```bash
# 监控文件夹自动处理
inotifywait -m ./uploads/ -e create -e moved_to |
while read path action file; do
  if [[ "$file" =~ \.(jpg|jpeg|png)$ ]]; then
    echo "处理新图片: $path$file"
    /wechatwriter:image compress "$path$file" --auto-compress
    /wechatwriter:image upload "$path$file"
  fi
done
```

### 内容管理系统集成

```bash
# 批量处理内容图片
for article in articles/*.md; do
  # 提取图片链接
  grep -oE '!\[.*\]\(.*\)' "$article" | while read -r line; do
    image_url=$(echo "$line" | sed -n 's/.*](\([^)]*\)).*/\1/p')
    if [[ "$image_url" =~ ^https?:// ]]; then
      echo "处理图片: $image_url"
      /wechatwriter:image download "$image_url" --compress
    fi
  done
done
```

## 参数详解

### 上传命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| file | 图片文件路径 | string | 是 |
| --compress | 启用压缩 | boolean | 否 |
| --quality | 压缩质量 | number | 否 |
| --max-width | 最大宽度 | number | 否 |
| --max-height | 最大高度 | number | 否 |
| --format | 输出格式 | string | 否 |
| --batch | 批量模式 | boolean | 否 |

### 下载命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| url | 图片URL | string | 是 |
| --compress | 启用压缩 | boolean | 否 |
| --output | 输出路径 | string | 否 |
| --timeout | 超时时间 | number | 否 |

### 生成命令参数

| 参数 | 说明 | 类型 | 必需 |
|------|------|------|------|
| prompt | 生成提示词 | string | 是 |
| --style | 图片风格 | string | 否 |
| --size | 图片尺寸 | string | 否 |
| --aspect-ratio | 宽高比 | string | 否 |
| --quality | 生成质量 | string | 否 |

## 图片规格要求

### 微信公众号

| 类型 | 尺寸建议 | 大小限制 | 格式 |
|------|----------|----------|------|
| 封面图 | 900x500px | ≤5MB | JPG, PNG |
| 正文配图 | 无限制 | ≤5MB | JPG, PNG, GIF |
| 缩略图 | 200x200px | ≤1MB | JPG, PNG |

### 通用优化建议

1. **尺寸优化**: 宽度不超过1920px
2. **压缩质量**: 85-90%平衡质量与大小
3. **格式选择**: 照片用JPG，图标用PNG
4. **文件命名**: 使用有意义的文件名

## 常见问题

### Q: 上传失败怎么办？

A: 检查以下问题：
- 文件大小是否超限（5MB）
- 图片格式是否支持
- 网络连接是否正常
- API权限是否配置

### Q: AI生成图片质量如何？

A: 支持多种质量级别：
- `standard`: 标准质量，快速生成
- `hd`: 高清质量，细节丰富
- `ultra`: 超高质量，专业级别

### Q: 如何处理大批量图片？

A: 系统支持：
- 批量上传/下载
- 并发处理优化
- 断点续传功能
- 进度状态显示

## 最佳实践

1. **预处理优化**: 上传前先压缩和裁剪
2. **格式统一**: 保持项目中图片格式一致
3. **命名规范**: 使用清晰、有意义的文件名
4. **备份策略**: 重要图片保留原始文件
5. **质量平衡**: 在视觉效果和加载速度间找平衡