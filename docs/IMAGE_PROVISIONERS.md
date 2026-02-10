# 图片生成服务配置

writer 支持多种图片生成服务，可以在 Markdown 中使用 AI 生成图片。

## 快速开始

在 Markdown 中使用以下语法生成图片：

```markdown
![图片描述](__generate:一个秋天的森林，阳光透过树叶__)
```

## 配置方式

在配置文件 `~/.config/wechatwriter/config.yaml` 中配置图片生成服务：

```yaml
api:
  # 图片服务提供者: openai, tuzi
  image_provider: "tuzi"

  # API 配置
  image_key: "your-api-key"
  image_base_url: "https://api.tu-zi.com/v1"

  # 生成参数
  image_model: "doubao-seedream-4-5-251128"
  image_size: "2048x2048"  # TuZi 要求最小 3686400 像素

image:
  compress: true
  max_width: 1920
```

## 支持的图片服务

### TuZi

TuZi 是国内可用的图片生成服务，支持多种模型。

#### 配置示例

```yaml
api:
  image_provider: "tuzi"
  image_key: "tuzi-sk-..."
  image_base_url: "https://api.tu-zi.com/v1"
  image_model: "doubao-seedream-4-5-251128"
  image_size: "1024x1024"
```

#### 支持的模型

| 模型 | 说明 |
|------|------|
| `doubao-seedream-4-5-251128` | 豆包 Seedream 4.5（推荐） |
| `gemini-3-pro-image-preview` | Gemini 3 Pro 图片预览版 |

#### 支持的尺寸

| 尺寸 | 比例 | 说明 |
|------|------|------|
| `2048x2048` | 1:1 | 正方形（默认，4.2M 像素）|
| `1920x1920` | 1:1 | 正方形（最小要求，3.7M 像素）|
| `2560x1440` | 16:9 | 横版（3.7M 像素）|
| `1440x2560` | 9:16 | 竖版（3.7M 像素）|
| `3072x2048` | 3:2 | 横版（6.3M 像素）|
| `2048x3072` | 2:3 | 竖版（6.3M 像素）|
| `3840x2160` | 16:9 | 超宽横版（8.3M 像素）|
| `2160x3840` | 9:16 | 超高竖版（8.3M 像素）|

> **注意**：TuZi 要求图片尺寸至少达到 **3,686,400 像素**

#### 获取 API Key

前往 [TuZi 控制台](https://api.tu-zi.com) 注册并获取 API Key。

---

### OpenAI

使用 OpenAI DALL-E 模型生成图片。

#### 配置示例

```yaml
api:
  image_provider: "openai"
  image_key: "sk-..."
  image_base_url: "https://api.openai.com/v1"
  image_model: "dall-e-3"
  image_size: "1024x1024"
```

#### 支持的模型

| 模型 | 说明 |
|------|------|
| `dall-e-3` | DALL-E 3（最新） |
| `dall-e-2` | DALL-E 2 |

#### 支持的尺寸

| 尺寸 | 模型 |
|------|------|
| `1024x1024` | dall-e-2, dall-e-3 |
| `1792x1024` | dall-e-3 |
| `1024x1792` | dall-e-3 |

---

## 使用示例

### 在 Markdown 中生成图片

```markdown
# 我的文章

这是一篇文章，里面有一张 AI 生成的封面图：

![封面图](__generate:温暖的秋天森林，阳光透过树叶洒在地面上__)

还可以生成插图：

![插图](__generate:赛博朋克风格的城市夜景，霓虹灯闪烁__)
```

### 命令行使用

```bash
# 转换文章（会自动生成图片并上传到微信）
writer convert article.md --draft

# 只预览（不上传）
writer convert article.md --preview
```

---

## 常见问题

### Q: 提示 "API Key 无效" 怎么办？

**A:** 请检查以下几点：
1. 配置文件中的 `api.image_key` 是否正确填写
2. API Key 是否已过期或被撤销
3. 对于 TuZi，前往控制台确认账户状态正常
4. 对于 OpenAI，确认 API Key 有效且有余额

---

### Q: 提示 "账户余额不足" 怎么办？

**A:**
- **TuZi**: 前往 [TuZi 控制台](https://api.tu-zi.com) 充值
- **OpenAI**: 前往 [OpenAI 控制台](https://platform.openai.com) 充值

---

### Q: 提示 "请求过于频繁" 怎么办？

**A:** API 服务有速率限制，请：
1. 等待一段时间后再试
2. 考虑升级服务套餐
3. 减少同时生成的图片数量

---

### Q: 提示 "参数配置有误" 怎么办？

**A:** 请检查：
1. `image_provider` 是否为 `openai` 或 `tuzi`
2. `image_model` 是否在支持的模型列表中
3. `image_size` 是否在支持的尺寸列表中

---

### Q: 生成的图片不符合预期怎么办？

**A:** 尝试优化提示词：
1. 描述更具体：`一只金色的猫坐在红色的沙发上` 比 `猫` 更好
2. 添加风格描述：`油画风格`、`照片级真实`、`卡通风格` 等
3. 指定颜色和光线：`温暖的阳光`、`冷色调`、`高对比度` 等

---

### Q: 图片生成失败但没显示具体错误？

**A:** 运行以下命令查看详细日志：

```bash
# 设置日志级别为 debug
WECHATWRITER_LOG_LEVEL=debug writer convert article.md --preview
```

---

## 环境变量配置

除了配置文件，也可以使用环境变量：

```bash
export IMAGE_PROVIDER="tuzi"
export IMAGE_API_KEY="your-api-key"
export IMAGE_API_BASE="https://api.tu-zi.com/v1"
export IMAGE_MODEL="doubao-seedream-4-5-251128"
export IMAGE_SIZE="2048x2048"  # TuZi 要求最小 3686400 像素
```

环境变量优先级高于配置文件。

---

## 错误代码速查

| 错误代码 | 说明 | 解决方案 |
|----------|------|----------|
| `unauthorized` | API Key 无效 | 检查 API Key 配置 |
| `payment_required` | 余额不足 | 前往控制台充值 |
| `rate_limit` | 请求过于频繁 | 等待后重试 |
| `bad_request` | 参数错误 | 检查模型和尺寸配置 |
| `network_error` | 网络错误 | 检查网络连接和 API 地址 |
| `no_image` | 未生成图片 | 检查提示词是否符合内容政策 |
