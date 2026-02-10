# Context

Filename: 2026-02-06-comprehensive-test-plan.md
Created On: 2026-02-06
Created By: AI Assistant
Associated Protocol: RIPER-5 + Multidimensional + Agent Protocol

# Task Description

增加详尽的测试计划，测试所有相关功能，保证整体的生产可运行！

# Project Overview

Writer for WeChat - 微信公众号写作工具 CLI 应用

- 语言：Go 1.24.0
- 主要功能：Markdown 转换、风格化写作、AI 去痕、图片处理、草稿管理
- 核心依赖：cobra (CLI)、zap (logging)、yaml.v3 (配置)、wechat.v2 (微信 API)
- 现有测试：仅 `app/image/modelscope_test.go` 一个测试文件

---

*The following sections are maintained by the AI during protocol execution*
---

# Analysis (Populated by RESEARCH mode)

## 项目结构分析

### 核心模块

1. **app/config/** - 配置管理
   - config.go (539 行) - 配置加载、验证、保存
   - account_selector.go - 多账号选择器

2. **app/converter/** - Markdown 转 HTML
   - converter.go - 核心转换逻辑
   - ai.go - AI 模式处理
   - image.go - 图片引用处理
   - prompt.go - AI 提示词生成
   - theme.go - 主题系统

3. **app/writer/** - 写作风格助手
   - assistant.go - 风格助手主逻辑
   - generator.go - 文章生成
   - cover_generator.go - 封面生成
   - style.go - 风格管理
   - types.go - 类型定义

4. **app/humanizer/** - AI 去痕
   - humanizer.go - 去痕核心逻辑
   - prompt.go - 提示词生成
   - result.go - 结果解析

5. **app/image/** - 图片处理
   - processor.go - 图片处理器
   - compress.go - 图片压缩
   - provider.go - Provider 接口
   - openai.go - OpenAI provider
   - modelscope.go - ModelScope provider（已有测试）
   - tuzi.go - Tuzi provider

6. **app/draft/** - 草稿管理
   - service.go - 草稿服务

7. **app/wechat/** - 微信 API
   - service.go - 微信服务

### CLI 命令

- main.go - 主入口
- convert.go - convert 命令
- write.go - write 命令
- humanize.go - humanize 命令
- score.go - score 命令
- outline.go - outline 命令
- draft.go - draft 命令
- image.go - image 命令

### 现有测试覆盖

- `app/image/modelscope_test.go` (340 行)
  - 测试 ModelScope provider 的创建、任务创建、状态轮询、图片生成
  - 测试错误处理（未授权、限流、任务失败）
  - 测试默认值和配置

### 测试工具

- Go 标准测试框架
- httptest - HTTP mock server
- Makefile 提供 `make test` 命令

## 关键依赖和约束

### 外部依赖

- 微信公众号 API（需要真实账号才能测试）
- OpenAI API（需要 API key）
- ModelScope API（需要 API key）
- Tuzi API（需要 API key）

### 测试约束

1. 微信 API 调用需要 mock 或集成测试环境
2. AI API 调用需要 mock 或真实 API key
3. 图片生成需要 mock 或真实 API
4. 文件系统操作需要临时目录
5. 配置加载需要测试多种场景

### 生产可运行性要求

- 所有核心功能必须有测试覆盖
- 错误处理必须有测试
- 边界条件必须有测试
- 集成测试覆盖关键用户流程

# Proposed Solution (Populated by INNOVATE mode)

## 测试策略方案

### 方案 1：单元测试 + 集成测试（推荐）

**优点**：

- 测试粒度细，定位问题快
- 执行速度快，适合 CI/CD
- 可以 mock 外部依赖，测试稳定

**缺点**：

- 需要编写大量 mock 代码
- 集成测试需要额外配置

**适用场景**：

- 代码逻辑复杂，需要细致测试
- 外部依赖不稳定，需要 mock
- CI/CD 环境需要快速反馈

### 方案 2：端到端测试（E2E）

**优点**：

- 测试真实用户场景
- 发现集成问题
- 测试覆盖面广

**缺点**：

- 执行速度慢
- 需要真实环境配置
- 定位问题困难

**适用场景**：

- 关键用户流程验证
- 发布前回归测试
- 需要真实环境测试

### 方案 3：混合测试策略（最佳实践）

**优点**：

- 结合单元测试和集成测试的优点
- 平衡测试速度和覆盖面
- 灵活适应不同测试需求

**缺点**：

- 需要维护多种测试类型
- 测试策略复杂

**适用场景**：

- 大型项目
- 需要快速迭代
- 对质量要求高

## 推荐方案：混合测试策略

### 测试金字塔

```
        /\
       /  \      E2E 测试 (5%)
      /____\     关键用户流程
     /      \
    /        \   集成测试 (25%)
   /__________\  模块间交互
  /            \
 /______________\ 单元测试 (70%)
                  函数级别
```

### 测试分层

**第 1 层：单元测试（70%）**

- 测试单个函数和方法的正确性
- 使用 mock 隔离外部依赖
- 覆盖正常流程、边界条件、错误处理

**第 2 层：集成测试（25%）**

- 测试模块间的交互
- 测试关键业务流程
- 使用真实或 mock 的外部服务

**第 3 层：端到端测试（5%）**

- 测试完整的用户场景
- 使用真实环境或接近真实的环境
- 验证生产可运行性

## 测试优先级

### P0 - 核心功能（必须测试）

1. 配置加载和验证
2. Markdown 转 HTML（converter）
3. 图片上传和处理
4. 草稿创建和管理
5. CLI 命令基本功能

### P1 - 重要功能（应该测试）

1. 写作风格助手
2. AI 去痕
3. 图片生成（多个 provider）
4. 热点评分
5. 内容框架生成

### P2 - 辅助功能（可以测试）

1. 封面生成
2. 主题系统
3. 账号选择器
4. 错误处理细节

## 测试工具选择

### 单元测试

- Go 标准测试框架
- testify/assert - 断言库
- httptest - HTTP mock
- gomock - 接口 mock（可选）

### 集成测试

- testcontainers - 容器化测试环境（可选）
- docker-compose - 测试环境编排

### E2E 测试

- 自定义测试脚本
- 真实环境配置

### 测试覆盖率

- go test -cover
- gocov - 覆盖率报告
- codecov - 覆盖率可视化（可选）

# Implementation Plan (Generated by PLAN mode)

## 测试计划详细规格

### 阶段 1：基础设施和配置测试（P0）

#### 1.1 配置模块测试

**文件**: `app/config/config_test.go`

**测试内容**:

- 测试默认配置加载
- 测试 YAML 配置文件加载
- 测试 JSON 配置文件加载
- 测试环境变量覆盖
- 测试配置验证（正常、边界、错误）
- 测试多账号配置
- 测试配置保存

**测试函数**:

```go
func TestLoad_DefaultConfig(t *testing.T)
func TestLoad_YAMLConfig(t *testing.T)
func TestLoad_JSONConfig(t *testing.T)
func TestLoad_EnvOverride(t *testing.T)
func TestConfig_Validate_Success(t *testing.T)
func TestConfig_Validate_MissingAccount(t *testing.T)
func TestConfig_Validate_InvalidImageWidth(t *testing.T)
func TestConfig_Validate_InvalidTimeout(t *testing.T)
func TestConfig_MultiAccount(t *testing.T)
func TestSaveConfig_YAML(t *testing.T)
func TestSaveConfig_JSON(t *testing.T)
func TestFindConfigFile_UserDir(t *testing.T)
func TestFindConfigFile_CurrentDir(t *testing.T)
```

**Mock 需求**:

- 文件系统操作（临时文件）
- 环境变量

#### 1.2 账号选择器测试

**文件**: `app/config/account_selector_test.go`

**测试内容**:

- 测试默认账号选择
- 测试关键词匹配
- 测试无匹配时的错误处理
- 测试多账号场景

**测试函数**:

```go
func TestAccountSelector_SelectDefault(t *testing.T)
func TestAccountSelector_SelectByKeyword(t *testing.T)
func TestAccountSelector_NoMatch(t *testing.T)
func TestAccountSelector_MultipleAccounts(t *testing.T)
```

### 阶段 2：转换器测试（P0）

#### 2.1 转换器核心测试

**文件**: `app/converter/converter_test.go`

**测试内容**:

- 测试基本 Markdown 转 HTML
- 测试图片引用识别（本地、在线、AI）
- 测试主题应用
- 测试 AI 模式请求生成
- 测试图片占位符替换

**测试函数**:

```go
func TestConverter_BasicMarkdown(t *testing.T)
func TestConverter_WithImages(t *testing.T)
func TestConverter_ThemeApplication(t *testing.T)
func TestConverter_AIMode(t *testing.T)
func TestConverter_ImagePlaceholderReplacement(t *testing.T)
func TestConverter_EmptyInput(t *testing.T)
func TestConverter_InvalidMarkdown(t *testing.T)
```

**Mock 需求**:

- AI 服务（httptest）
- 文件系统（临时文件）

#### 2.2 AI 模式测试

**文件**: `app/converter/ai_test.go`

**测试内容**:

- 测试 AI 请求信息提取
- 测试提示词生成
- 测试 AI 响应解析

**测试函数**:

```go
func TestGetAIRequestInfo(t *testing.T)
func TestBuildAIPrompt(t *testing.T)
func TestParseAIResponse(t *testing.T)
```

#### 2.3 主题系统测试

**文件**: `app/converter/theme_test.go`

**测试内容**:

- 测试主题加载
- 测试主题应用
- 测试主题验证

**测试函数**:

```go
func TestTheme_Load(t *testing.T)
func TestTheme_Apply(t *testing.T)
func TestTheme_Validation(t *testing.T)
```

### 阶段 3：图片处理测试（P0）

#### 3.1 图片处理器测试

**文件**: `app/image/processor_test.go`

**测试内容**:

- 测试本地图片上传
- 测试在线图片下载和上传
- 测试图片压缩
- 测试图片生成和上传
- 测试错误处理

**测试函数**:

```go
func TestProcessor_UploadLocalImage(t *testing.T)
func TestProcessor_DownloadAndUpload(t *testing.T)
func TestProcessor_GenerateAndUpload(t *testing.T)
func TestProcessor_CompressImage(t *testing.T)
func TestProcessor_InvalidImage(t *testing.T)
func TestProcessor_UploadError(t *testing.T)
```

**Mock 需求**:

- 微信 API（httptest）
- 图片生成 API（httptest）
- 文件系统（临时文件）

#### 3.2 图片压缩测试

**文件**: `app/image/compress_test.go`

**测试内容**:

- 测试图片压缩（不同尺寸）
- 测试图片质量调整
- 测试格式转换

**测试函数**:

```go
func TestCompressImage_Resize(t *testing.T)
func TestCompressImage_Quality(t *testing.T)
func TestCompressImage_Format(t *testing.T)
```

#### 3.3 OpenAI Provider 测试

**文件**: `app/image/openai_test.go`

**测试内容**:

- 测试 OpenAI 图片生成
- 测试错误处理
- 测试配置

**测试函数**:

```go
func TestOpenAIProvider_Generate(t *testing.T)
func TestOpenAIProvider_ErrorHandling(t *testing.T)
func TestOpenAIProvider_Configuration(t *testing.T)
```

**Mock 需求**:

- OpenAI API（httptest）

#### 3.4 Tuzi Provider 测试

**文件**: `app/image/tuzi_test.go`

**测试内容**:

- 测试 Tuzi 图片生成
- 测试错误处理
- 测试配置

**测试函数**:

```go
func TestTuziProvider_Generate(t *testing.T)
func TestTuziProvider_ErrorHandling(t *testing.T)
func TestTuziProvider_Configuration(t *testing.T)
```

**Mock 需求**:

- Tuzi API（httptest）

### 阶段 4：草稿管理测试（P0）

#### 4.1 草稿服务测试

**文件**: `app/draft/service_test.go`

**测试内容**:

- 测试草稿创建
- 测试草稿摘要生成
- 测试多文章草稿
- 测试错误处理

**测试函数**:

```go
func TestDraftService_CreateDraft(t *testing.T)
func TestDraftService_GenerateDigest(t *testing.T)
func TestDraftService_MultipleArticles(t *testing.T)
func TestDraftService_ErrorHandling(t *testing.T)
```

**Mock 需求**:

- 微信 API（httptest）

### 阶段 5：写作助手测试（P1）

#### 5.1 写作助手测试

**文件**: `app/writer/assistant_test.go`

**测试内容**:

- 测试风格加载
- 测试文章生成
- 测试风格列表
- 测试风格详情

**测试函数**:

```go
func TestAssistant_ListStyles(t *testing.T)
func TestAssistant_GetStyleInfo(t *testing.T)
func TestAssistant_Write(t *testing.T)
func TestAssistant_InvalidStyle(t *testing.T)
```

#### 5.2 风格管理测试

**文件**: `app/writer/style_test.go`

**测试内容**:

- 测试风格加载
- 测试风格验证
- 测试风格解析

**测试函数**:

```go
func TestStyleManager_LoadStyles(t *testing.T)
func TestStyleManager_ValidateStyle(t *testing.T)
func TestStyleManager_ParseStyleInput(t *testing.T)
```

#### 5.3 封面生成器测试

**文件**: `app/writer/cover_generator_test.go`

**测试内容**:

- 测试封面提示词生成
- 测试不同风格的封面

**测试函数**:

```go
func TestCoverGenerator_GeneratePrompt(t *testing.T)
func TestCoverGenerator_DifferentStyles(t *testing.T)
```

### 阶段 6：AI 去痕测试（P1）

#### 6.1 Humanizer 测试

**文件**: `app/humanizer/humanizer_test.go`

**测试内容**:

- 测试去痕提示词生成
- 测试 AI 响应解析
- 测试不同强度
- 测试质量评分

**测试函数**:

```go
func TestHumanizer_BuildAIRequest(t *testing.T)
func TestHumanizer_ParseAIResponse(t *testing.T)
func TestHumanizer_IntensityLevels(t *testing.T)
func TestHumanizer_QualityScore(t *testing.T)
```

### 阶段 7：微信服务测试（P1）

#### 7.1 微信服务测试

**文件**: `app/wechat/service_test.go`

**测试内容**:

- 测试素材上传
- 测试草稿创建
- 测试错误处理

**测试函数**:

```go
func TestWechatService_UploadMaterial(t *testing.T)
func TestWechatService_CreateDraft(t *testing.T)
func TestWechatService_ErrorHandling(t *testing.T)
```

**Mock 需求**:

- 微信 API（httptest）

### 阶段 8：CLI 命令测试（P0）

#### 8.1 Convert 命令测试

**文件**: `app/convert_test.go`

**测试内容**:

- 测试基本转换
- 测试不同主题
- 测试图片上传
- 测试草稿创建
- 测试错误处理

**测试函数**:

```go
func TestConvertCmd_Basic(t *testing.T)
func TestConvertCmd_WithTheme(t *testing.T)
func TestConvertCmd_WithUpload(t *testing.T)
func TestConvertCmd_WithDraft(t *testing.T)
func TestConvertCmd_ErrorHandling(t *testing.T)
```

#### 8.2 Write 命令测试

**文件**: `app/write_test.go`

**测试内容**:

- 测试基本写作
- 测试风格选择
- 测试交互模式
- 测试封面生成
- 测试 AI 去痕集成

**测试函数**:

```go
func TestWriteCmd_Basic(t *testing.T)
func TestWriteCmd_StyleSelection(t *testing.T)
func TestWriteCmd_Interactive(t *testing.T)
func TestWriteCmd_WithCover(t *testing.T)
func TestWriteCmd_WithHumanize(t *testing.T)
```

#### 8.3 Humanize 命令测试

**文件**: `app/humanize_test.go`

**测试内容**:

- 测试基本去痕
- 测试不同强度
- 测试输出选项

**测试函数**:

```go
func TestHumanizeCmd_Basic(t *testing.T)
func TestHumanizeCmd_Intensity(t *testing.T)
func TestHumanizeCmd_Output(t *testing.T)
```

#### 8.4 Draft 命令测试

**文件**: `app/draft_test.go`

**测试内容**:

- 测试草稿创建
- 测试草稿测试
- 测试草稿发布

**测试函数**:

```go
func TestDraftCmd_Create(t *testing.T)
func TestDraftCmd_Test(t *testing.T)
func TestDraftCmd_Publish(t *testing.T)
```

#### 8.5 Image 命令测试

**文件**: `app/image_test.go`

**测试内容**:

- 测试图片上传
- 测试图片下载
- 测试图片生成

**测试函数**:

```go
func TestImageCmd_Upload(t *testing.T)
func TestImageCmd_Download(t *testing.T)
func TestImageCmd_Generate(t *testing.T)
```

#### 8.6 Score 命令测试

**文件**: `app/score_test.go`

**测试内容**:

- 测试评分计算
- 测试不同指标
- 测试输出格式

**测试函数**:

```go
func TestScoreCmd_Calculation(t *testing.T)
func TestScoreCmd_DifferentMetrics(t *testing.T)
func TestScoreCmd_OutputFormat(t *testing.T)
```

#### 8.7 Outline 命令测试

**文件**: `app/outline_test.go`

**测试内容**:

- 测试内容框架生成
- 测试不同模板
- 测试输出格式

**测试函数**:

```go
func TestOutlineCmd_Generation(t *testing.T)
func TestOutlineCmd_DifferentTemplates(t *testing.T)
func TestOutlineCmd_OutputFormat(t *testing.T)
```

### 阶段 9：集成测试（P0）

#### 9.1 完整流程测试

**文件**: `tests/integration_test.go`

**测试内容**:

- 测试 Markdown 转换 -> 图片上传 -> 草稿创建完整流程
- 测试写作 -> 去痕 -> 草稿创建完整流程
- 测试配置加载 -> 命令执行完整流程

**测试函数**:

```go
func TestIntegration_ConvertToDraft(t *testing.T)
func TestIntegration_WriteToDraft(t *testing.T)
func TestIntegration_ConfigToExecution(t *testing.T)
```

### 阶段 10：E2E 测试（P0）

#### 10.1 E2E 测试脚本

**文件**: `tests/e2e_test.sh`

**测试内容**:

- 测试真实环境下的完整用户流程
- 测试配置初始化
- 测试所有主要命令

**测试场景**:

```bash
# 场景 1: 配置初始化
writer config init

# 场景 2: Markdown 转换
writer convert article.md --theme default --preview

# 场景 3: 图片上传
writer image upload cover.jpg

# 场景 4: 写作
writer write --style dan-koe

# 场景 5: AI 去痕
writer humanize article.md --intensity medium

# 场景 6: 草稿创建
writer draft create draft.json
```

## 测试覆盖率目标

### 总体目标

- 语句覆盖率：≥ 80%
- 分支覆盖率：≥ 75%
- 函数覆盖率：≥ 90%

### 模块覆盖率目标

- config：≥ 90%
- converter：≥ 85%
- writer：≥ 80%
- humanizer：≥ 80%
- image：≥ 85%
- draft：≥ 85%
- wechat：≥ 80%
- CLI 命令：≥ 75%

## 测试执行计划

### 开发阶段

1. 编写测试代码
2. 运行单元测试：`go test ./...`
3. 检查覆盖率：`go test -cover ./...`
4. 修复失败的测试

### CI/CD 阶段

1. 自动运行所有测试
2. 生成覆盖率报告
3. 检查覆盖率是否达标
4. 失败时阻止合并

### 发布阶段

1. 运行完整测试套件
2. 运行集成测试
3. 运行 E2E 测试（可选）
4. 生成测试报告

## 测试数据准备

### 测试文件

- 测试 Markdown 文件：`tests/fixtures/sample.md`
- 测试图片文件：`tests/fixtures/sample.jpg`
- 测试配置文件：`tests/fixtures/config.yaml`
- 测试 JSON 文件：`tests/fixtures/draft.json`

### Mock 数据

- 微信 API 响应：`tests/mocks/wechat_responses.json`
- OpenAI API 响应：`tests/mocks/openai_responses.json`
- ModelScope API 响应：`tests/mocks/modelscope_responses.json`

## 测试工具配置

### Makefile 更新

```makefile
# 运行测试
test:
 @echo "🧪 运行测试..."
 @go test -v -race -coverprofile=coverage.out ./...

# 运行测试并生成覆盖率报告
test-coverage:
 @echo "🧪 运行测试并生成覆盖率报告..."
 @go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
 @go tool cover -html=coverage.out -o coverage.html
 @echo "✅ 覆盖率报告已生成: coverage.html"

# 运行单元测试
test-unit:
 @echo "🧪 运行单元测试..."
 @go test -v -short ./...

# 运行集成测试
test-integration:
 @echo "🧪 运行集成测试..."
 @go test -v -tags=integration ./...

# 运行 E2E 测试
test-e2e:
 @echo "🧪 运行 E2E 测试..."
 @bash tests/e2e_test.sh

# 检查覆盖率
test-coverage-check:
 @echo "🧪 检查覆盖率..."
 @go test -coverprofile=coverage.out ./...
 @go tool cover -func=coverage.out | grep total
```

### .gitignore 更新

```
coverage.out
coverage.html
*.test
```

## 测试文档

### 测试指南

创建 `tests/README.md`，包含：

- 测试结构说明
- 如何运行测试
- 如何编写新测试
- Mock 使用指南
- 测试数据准备

### 测试报告模板

创建 `tests/REPORT_TEMPLATE.md`，包含：

- 测试执行摘要
- 覆盖率统计
- 失败测试列表
- 问题分析
- 改进建议

Implementation Checklist:

1. 创建测试目录结构
2. 创建测试数据文件
3. 编写配置模块测试（config_test.go, account_selector_test.go）
4. 编写转换器测试（converter_test.go, ai_test.go, theme_test.go）
5. 编写图片处理测试（processor_test.go, compress_test.go, openai_test.go, tuzi_test.go）
6. 编写草稿管理测试（service_test.go）
7. 编写写作助手测试（assistant_test.go, style_test.go, cover_generator_test.go）
8. 编写 AI 去痕测试（humanizer_test.go）
9. 编写微信服务测试（service_test.go）
10. 编写 CLI 命令测试（convert_test.go, write_test.go, humanize_test.go, draft_test.go, image_test.go, score_test.go, outline_test.go）
11. 编写集成测试（integration_test.go）
12. 编写 E2E 测试脚本（e2e_test.sh）
13. 更新 Makefile 添加测试命令
14. 更新 .gitignore
15. 创建测试文档（README.md, REPORT_TEMPLATE.md）
16. 运行所有测试并修复失败
17. 生成覆盖率报告并验证达标
18. 编写测试执行报告

# Current Execution Step (Updated by EXECUTE mode when starting a step)
>
> Currently executing: "No step started yet"

# Task Progress (Appended by EXECUTE mode after each step completion)

* [DateTime]
  - Step: [Checklist item number and description]
  - Modifications: [List of file and code changes, including reported minor deviation corrections]
  - Change Summary: [Brief summary of this change]
  - Reason: [Executing plan step [X]]
  - Blockers: [Any issues encountered, or None]
  - User Confirmation Status: [Success / Success with minor issues / Failure]
- [DateTime]
  - Step: ...

# Final Review (Populated by REVIEW mode)

[Summary of implementation compliance assessment against the final plan, whether unreported deviations were found]
