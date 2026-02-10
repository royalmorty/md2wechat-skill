# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**Writer for WeChat** (wechatwriter) is a Go CLI tool that transforms Markdown articles into WeChat-formatted HTML with professional styling, AI-powered writing assistance, humanization features, and direct publishing to WeChat draft box.

- **Language**: Go 1.24.0
- **CLI Framework**: Cobra
- **Logging**: Zap (structured logging)
- **WeChat SDK**: silenceper/wechat/v2
- **Architecture**: Modular CLI with separate concerns for conversion, image processing, draft management, and writing assistance

## Build & Test Commands

### Building

```bash
# Quick build for current platform (development)
make fast

# Build for current platform (outputs to scripts/wechatwriter)
make build

# Build for all platforms (release)
make release

# Build via go directly
go build -o wechatwriter ./app
```

### Testing

```bash
# Run all tests
make test
# or
go test -v ./...

# Run specific package tests
go test -v ./app/config
go test -v ./app/image

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -v -run TestConfig_Validate ./app/config
```

### Code Quality

```bash
# Format code
make fmt

# Static analysis
make vet

# Lint (requires golangci-lint)
make lint

# Download/update dependencies
make deps
```

### Installation

```bash
# Install to GOPATH/bin
make install
# or
go install ./app
```

## Architecture

### Core Module Structure

The project follows a layered architecture with clear separation of concerns:

```
app/
├── main.go                 # CLI entry point, command routing
├── {command}.go            # Individual command implementations (convert, write, humanize, etc.)
│
├── config/                 # Configuration management
│   ├── config.go          # Multi-account config with YAML/JSON/env support
│   └── account_selector.go # Smart account selection based on keywords
│
├── converter/             # Markdown → WeChat HTML conversion
│   ├── converter.go       # Core conversion interface & orchestration
│   ├── ai.go             # AI mode implementation (Claude-based)
│   ├── image.go          # Image reference extraction & placeholder handling
│   ├── prompt.go         # AI prompt building with theme support
│   └── theme.go          # Theme management system
│
├── writer/               # Styled writing assistance
│   ├── assistant.go      # Writing style orchestration
│   ├── generator.go      # Content generation
│   ├── cover_generator.go # Cover image prompt generation
│   ├── style.go          # Style definition loading (YAML-based)
│   └── types.go          # Data structures
│
├── humanizer/            # AI writing trace removal
│   ├── humanizer.go      # Detection & removal of AI patterns
│   ├── prompt.go         # Humanization prompts
│   └── result.go         # Quality scoring (5 dimensions)
│
├── image/                # Image processing & generation
│   ├── processor.go      # Unified image handling (upload, compress, generate)
│   ├── compress.go       # Image compression (max 1920px, preserves aspect ratio)
│   ├── provider.go       # Provider interface
│   ├── openai.go         # OpenAI DALL-E provider
│   ├── modelscope.go     # ModelScope provider (async task-based)
│   └── tuzi.go           # Tuzi provider
│
├── draft/                # WeChat draft management
│   └── service.go        # Draft creation & publishing
│
└── wechat/               # WeChat API wrapper
    └── service.go        # Material upload, access token management
```

### Configuration System

**Multi-Account Support**: The config system supports multiple WeChat accounts with automatic keyword-based selection. Each account can have:
- Unique credentials (AppID, Secret)
- Domain-specific keywords for auto-selection
- Per-account default writing styles

**Configuration Sources** (priority order):
1. Environment variables (highest priority)
2. Config file (`~/.config/wechatwriter/config.yaml` or `.json`)
3. Default values

**Key Environment Variables**:
```bash
IMAGE_API_KEY          # Image generation API key
IMAGE_API_BASE         # API base URL (default: https://api.openai.com/v1)
IMAGE_PROVIDER         # Provider: openai|modelscope|tuzi
IMAGE_MODEL            # Model name (e.g., dall-e-3, Tongyi-MAI/Z-Image-Turbo)
IMAGE_SIZE             # Image size (e.g., 1024x1024, 2560x1440)
COMPRESS_IMAGES        # Auto-compress images (default: true)
MAX_IMAGE_WIDTH        # Max width in pixels (default: 1920)
```

### Conversion Flow

The converter module orchestrates a multi-step process:

1. **Image Extraction**: Parse Markdown for image references (local/online/AI-generated)
2. **Markdown → HTML**: Generate WeChat-compatible HTML with theme styling
   - **AI Mode**: Uses Claude with theme-specific prompts (autumn-warm, spring-fresh, ocean-calm, custom)
   - All CSS must be inlined (no external stylesheets)
   - Safe HTML tags only (no script, iframe, form elements)
3. **Image Placeholders**: Replace image references with `<!-- IMG:0 -->` format
4. **Image Processing** (if enabled):
   - Local: compress → upload to WeChat CDN
   - Online: download → compress → upload
   - AI: generate → download → compress → upload
5. **Placeholder Replacement**: Replace placeholders with WeChat CDN URLs
6. **Draft Publishing** (optional): Create draft in WeChat backend

### Image Generation Providers

**ModelScope** (recommended for free tier):
- Async API (task_id + polling every 5 seconds)
- Typical generation time: 10-30 seconds
- Max polling timeout: 120 seconds
- Default model: `Tongyi-MAI/Z-Image-Turbo`

**OpenAI**:
- Synchronous API
- Models: dall-e-2, dall-e-3
- Faster but requires paid API key

**Tuzi**:
- Chinese alternative with stable performance

### Writing Styles

Located in `config/styles/*.yaml`, each style defines:
- **Core Traits**: Distinctive voice characteristics
- **Structure Patterns**: Preferred content organization
- **Language Usage**: Vocabulary, sentence rhythm, formatting preferences
- **Domain Knowledge**: Specialized knowledge to incorporate

Built-in styles:
- `dan-koe`: Concise, punchy, philosophical depth with practical insights
- `cultural-depth`: Rich cultural references, literary style
- `casual-science`: Accessible explanations with engaging examples

### AI Humanization

Detects and removes 5 categories of AI patterns:
1. **Content Patterns**: Over-emphasis, vague attribution, promotional language
2. **Language Patterns**: AI vocabulary, negative parallelism, three-part structures
3. **Style Patterns**: Excessive dashes, bold overuse, emoji patterns
4. **Filler Patterns**: Filler phrases, over-qualification, generic conclusions
5. **Collaboration Traces**: Dialogue-style fillers, knowledge cutoff disclaimers

Intensity levels: `gentle`, `medium`, `aggressive`

Quality scoring (5 dimensions, 10 points each):
- Directness: Gets to the point quickly
- Rhythm: Varied sentence length
- Trust: Respects reader intelligence
- Authenticity: Sounds human-written
- Precision: No redundant content

## Development Patterns

### Adding New Commands

1. Create `app/{command}.go` with cobra command definition
2. Add command to `rootCmd` in `main.go`
3. Use `initConfig()` for lazy config loading (allows --help without config)
4. Return JSON responses via `responseSuccess()` or `responseError()`

### Adding New Image Providers

1. Implement `ImageProvider` interface in `app/image/{provider}.go`:
   ```go
   type ImageProvider interface {
       GenerateImage(prompt string, options GenerateOptions) (string, error)
   }
   ```
2. Register in `app/image/provider.go` factory
3. Add tests following `modelscope_test.go` pattern (use httptest for mocking)

### Adding New Themes

1. Create YAML file in `config/styles/{name}.yaml`
2. Theme system auto-loads from YAML with hot-reload support
3. Theme structure includes: core_traits, structure_patterns, language_usage, domain_knowledge

### Writing Tests

Follow the established patterns in `app/config/config_test.go`:
- Use table-driven tests for multiple scenarios
- Create temp files/dirs with `t.TempDir()`
- Mock HTTP calls with `httptest.NewServer`
- Test both success and error paths
- Use custom error types for better error handling

Example test structure:
```go
func TestFeature(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid case", "input", "expected", false},
        {"error case", "bad", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

### Error Handling

- Use custom error types (e.g., `ConfigError` with `Field` and `Message`)
- Log errors with zap structured logging: `log.Error("msg", zap.Error(err), zap.String("field", value))`
- Return errors via JSON response in CLI commands
- Mask sensitive values in logs (see `maskMediaID()` pattern)

### WeChat API Integration

**Access Token Management**:
- Automatically cached and refreshed by wechat SDK
- No manual token management needed

**Material Upload**:
- Images must be < 10MB (enforced by compression)
- Returns media_id and CDN URL
- Retry logic handles transient failures

**Draft Creation**:
- Content size limit: < 20,000 characters or 1MB
- API mode generates larger HTML due to inline CSS (use sparingly for long articles)
- HTML must use safe tags only

## Important Constraints

1. **WeChat HTML Requirements**:
   - All CSS must be inlined (style attributes)
   - No external resources (fonts, scripts, stylesheets)
   - Safe tags only: section, p, span, strong, em, h1-h6, ul, ol, li, blockquote, pre, code, table, img, br, hr
   - No: script, iframe, form, input, style, link

2. **Image Processing**:
   - Max width: 1920px (configurable)
   - Max size: 10MB for WeChat upload
   - Compression preserves aspect ratio
   - Supported formats: JPG, PNG

3. **Configuration**:
   - Sensitive data (AppID, Secret, API keys) can be environment variables
   - Config files support both YAML and JSON
   - Account selection is automatic based on keywords in input

4. **AI Generation**:
   - Prompts must be in Chinese for better results with Chinese content
   - Theme prompts define complete styling system (colors, typography, spacing)
   - Image generation prompts should be descriptive but concise

## CLI Commands Overview

- `writer config init` - Create config file with guided setup
- `writer convert <file>` - Convert Markdown to WeChat HTML
- `writer write` - Style-based writing assistance
- `writer humanize <file>` - Remove AI writing traces
- `writer score <file>` - Evaluate article quality (viral potential scoring)
- `writer outline` - Generate article outline
- `writer draft` - Manage WeChat drafts
- `writer image generate <prompt>` - Generate AI images

## Skills Integration

The project includes Claude Code skills in `skills/` directory:
- `content-writing` - Article writing workflow
- `visual-design` - Image and theme management
- `topic-research` - Research and scoring
- `seo-optimization` - SEO best practices
- `publishing-strategy` - Publishing workflows
- `content-analysis` - Content quality analysis

Skills are auto-loaded when working in this repository or via plugin marketplace.

## Notes for Development

- The codebase uses Go 1.24 features - ensure compatibility when adding new code
- Log structured data with zap, not fmt.Printf
- JSON responses should use `printJSON()` helper for consistent formatting
- When modifying config structure, update both YAML/JSON loading and validation
- Test both environment variable and file-based configuration paths
- Image processing is performance-sensitive - consider adding concurrency for batch operations
- WeChat API calls have rate limits - implement backoff/retry where needed
