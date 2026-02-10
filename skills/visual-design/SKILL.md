---
name: visual-design
description: 微信公众号视觉设计工具，提供配图建议、封面设计和版式优化服务
argument-hint: "[图片处理命令或设计需求]"
allowed-tools: Bash(./scripts/wechatwriter *), Read, Write
---

# 微信公众号视觉设计工具

专业的微信公众号视觉设计工具，提供配图方案推荐、封面设计指导和版式布局优化。

## 核心功能

### 1. 图片处理

处理文章配图，支持多种方式：

```bash
# 上传本地图片到微信素材库
./scripts/image.sh upload --input ./photo.jpg

# 下载并上传在线图片
./scripts/image.sh download --input https://example.com/image.jpg

# AI生成配图
./scripts/image.sh generate --prompt "茶园清晨薄雾"
```

### 2. 图片压缩优化

压缩图片大小，提升加载速度：

```bash
./scripts/image.sh compress --input image.jpg --quality 85 --width 1920
```

### 3. 封面图设计

创建专业的封面图：

```bash
./scripts/image.sh create-cover --title "文章标题" --theme elegant
```

**设计要素**:

- 视觉冲击力: 色彩对比和构图平衡
- 主题突出: 一眼识别文章核心内容
- 文字清晰: 标题文字可读性强
- 品牌一致: 符合整体视觉风格

## 使用示例

### 完整视觉设计流程

```bash
# 1. 生成配图
./scripts/image.sh generate --prompt "茶园清晨薄雾，阳光透过茶树" --style realistic

# 2. 压缩优化
./scripts/image.sh compress --input generated.jpg --quality 85 --width 1024

# 3. 创建封面
./scripts/image.sh create-cover --title "春茶品鉴指南" --theme elegant

# 4. 上传到微信
./scripts/image.sh upload --input cover.jpg
```

### 快速图片处理

```bash
# 直接生成配图
./scripts/image.sh generate --prompt "工作效率提升主题插画" --style illustration

# 批量压缩图片
for img in images/*.jpg; do
    ./scripts/image.sh compress --input "$img" --quality 80
done
```

## 输出格式

### 图片生成结果

```json
{
  "success": true,
  "data": {
    "prompt": "茶园清晨薄雾",
    "style": "realistic",
    "image_url": "https://generated-image.example.com/ai-image-123.jpg",
    "local_path": "./generated-images/ai-image-123.jpg",
    "dimensions": {
      "width": 1024,
      "height": 1024
    },
    "file_size": "2.3MB"
  }
}
```

### 封面图设计方案

```json
{
  "success": true,
  "data": {
    "title": "春茶品鉴指南",
    "theme": "elegant",
    "cover_design": {
      "background": "gradient_elegant",
      "typography": "elegant_serif",
      "color_palette": ["#2c3e50", "#ecf0f1", "#e74c3c"],
      "layout": "centered_title_with_decorative_elements",
      "dimensions": {
        "width": 900,
        "height": 500
      }
    },
    "generated_files": [
      "./covers/春茶品鉴指南_cover_900x500.png"
    ]
  }
}
```

## 设计标准

- **分辨率**: 封面图建议 900x500 像素
- **文件大小**: 单张图片不超过 2MB
- **格式支持**: JPG、PNG、GIF
- **色彩模式**: RGB，避免CMYK

## 可用命令

- **upload**: 上传图片到微信素材库
- **download**: 下载在线图片并上传
- **generate**: AI生成配图
- **compress**: 压缩优化图片
- **create-cover**: 创建专业封面图

## 注意事项

- 确保图片版权合规，优先使用AI生成
- 考虑不同设备的显示效果，特别是移动端
- 保持视觉风格的一致性
- 避免过度压缩导致画质损失
- 定期清理生成的临时文件
