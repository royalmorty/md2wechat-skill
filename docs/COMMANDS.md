# WeChatWriter Commands Documentation

Complete command reference for the WeChatWriter plugin - a professional WeChat content creation toolkit with AI-powered writing, visual design, and publishing optimization.

## 🚀 Quick Start

```bash
# Initialize configuration
/wechatwriter:config init

# Generate content with AI
/wechatwriter:generate article "Your topic here" --style dan-koe

# Convert and optimize existing content
/wechatwriter:convert your-article.md --theme autumn-warm --publish

# Get help for any command
/wechatwriter:help [command-name]
```

## 📋 Command Overview

| Command | Description | Key Features |
|---------|-------------|--------------|
| **[config](config.md)** | Configuration management | Interactive setup, multi-environment, security |
| **[convert](convert.md)** | Smart content conversion | Auto-detection, multi-format, one-click publish |
| **[generate](generate.md)** | AI content generation | Multi-style, batch processing, quality assurance |
| **[outline](outline.md)** | Content structure planning | AI-powered, multi-domain, template-based |
| **[write](write.md)** | AI-assisted writing | Creator styles, interactive mode, optimization |
| **[humanize](humanize.md)** | AI trace removal | Multi-intensity, 24 patterns, quality scoring |
| **[draft](draft.md)** | WeChat draft management | Create, test, publish, multi-account |
| **[image](image.md)** | Image processing | Upload, download, AI generation, compression |
| **[score](score.md)** | Viral content analysis | Multi-dimensional, platform-specific, optimization |

## 🎯 Core Workflows

### 1. Content Creation Workflow

```bash
# 1. Research topic
/wechatwriter:research "artificial intelligence in education" > research.json

# 2. Generate outline
/wechatwriter:outline "AI Education Future" --based-on research.json > outline.json

# 3. Write content
/wechatwriter:write --input-type outline outline.json --style dan-koe > article.md

# 4. Generate visuals
/wechatwriter:generate cover "AI Education" --style professional > cover.jpg

# 5. Convert for WeChat
/wechatwriter:convert article.md --theme autumn-warm --auto-images > wechat.html

# 6. Create draft
/wechatwriter:draft create --content wechat.html --cover cover.jpg
```

### 2. Content Optimization Workflow

```bash
# 1. Score existing content
/wechatwriter:score --input article.md --domain technology

# 2. Humanize if needed
/wechatwriter:humanize article.md --intensity medium > humanized.md

# 3. Enhance writing
/wechatwriter:write humanized.md --input-type fragment --refine > optimized.md

# 4. Convert with optimization
/wechatwriter:convert optimized.md --enhance-writing --optimize-seo
```

### 3. Batch Processing Workflow

```bash
# Batch generate articles
for topic in "Time Management" "Goal Setting" "Habit Building"; do
  /wechatwriter:generate article "$topic" --style dan-koe > "${topic}.md"
  /wechatwriter:convert "${topic}.md" --theme spring-fresh --auto-images
done

# Batch process images
/wechatwriter:image upload images/*.jpg --compress --batch

# Batch score content
for file in articles/*.md; do
  /wechatwriter:score --input "$file" --domain lifestyle
done
```

## 🎨 Content Styles & Themes

### Writing Styles
- **dan-koe**: Deep, sharp, grounded insights
- **story**: Narrative-driven, emotional engagement
- **professional**: Analytical, data-driven, formal
- **creative**: Imaginative, artistic, innovative
- **academic**: Research-based, citation-heavy
- **viral**: Engagement-optimized, shareable

### Visual Themes
- **spring-fresh**: Clean, vibrant, youthful
- **autumn-warm**: Cozy, mature, sophisticated
- **ocean-calm**: Peaceful, professional, trustworthy
- **cyber**: Tech-focused, futuristic, bold
- **chinese**: Traditional, cultural, elegant
- **apple**: Minimalist, premium, clean
- **bytedance**: Modern, dynamic, engaging

## 🔧 Advanced Features

### AI-Powered Capabilities
- **Natural Language Understanding**: Generate content from simple descriptions
- **Context Awareness**: Maintain consistency across content pieces
- **Multi-language Support**: Create content in different languages
- **Quality Assurance**: Fact-checking, originality verification
- **Style Consistency**: Maintain brand voice across all content

### Professional Tools
- **Multi-account Management**: Handle multiple WeChat accounts
- **Team Collaboration**: Shared templates, version control
- **Performance Analytics**: Track content performance metrics
- **Automated Workflows**: CI/CD integration, scheduled publishing
- **Security Features**: Encrypted storage, access control

## 📊 Platform Support

### Content Platforms
- **WeChat Official Accounts**: Full API integration
- **Weibo**: Optimized formatting and features
- **Zhihu**: Professional Q&A formatting
- **Jianshu**: Literary community features
- **Custom Platforms**: Flexible output formatting

### AI Service Providers
- **OpenAI**: GPT models for text generation
- **Claude**: Advanced reasoning and analysis
- **Gemini**: Multimodal content creation
- **Local Models**: Self-hosted AI solutions

## 🛠️ Configuration Management

### Environment Setup
```bash
# Development
/wechatwriter:config init --env development

# Staging
/wechatwriter:config init --env staging

# Production
/wechatwriter:config init --env production --encrypt
```

### Team Collaboration
```bash
# Export team configuration
/wechatwriter:config export --team-shared > team-config.yaml

# Create personal template
/wechatwriter:config create-template "personal-${USER}" --based-on team-config.yaml

# Version control
/wechatwriter:config version --save --message "Team config v2.0"
```

## 📈 Performance & Optimization

### Content Scoring
- **Viral Potential**: Multi-dimensional engagement analysis
- **SEO Optimization**: Keyword integration and metadata
- **Readability**: Grade-level assessment and improvement
- **Platform-specific**: Tailored scoring for each platform

### Image Optimization
- **Smart Compression**: Balance quality and file size
- **Format Selection**: Optimal format for each use case
- **Size Optimization**: Platform-specific dimensions
- **Batch Processing**: Efficient bulk operations

## 🔒 Security & Compliance

### Data Protection
- **Encrypted Storage**: Sensitive information protection
- **API Security**: Secure key management
- **Access Control**: Role-based permissions
- **Audit Logging**: Complete activity tracking

### Content Compliance
- **Fact Checking**: Source verification and citation
- **Copyright Compliance**: Originality detection and attribution
- **Platform Guidelines**: Automatic compliance checking
- **Industry Standards**: Regulatory requirement support

## 🆘 Troubleshooting

### Common Issues

**Configuration Problems**
```bash
# Check configuration status
/wechatwriter:config check --verbose

# Reset to defaults
/wechatwriter:config init --force

# Test API connections
/wechatwriter:config test --all-services
```

**Content Generation Issues**
```bash
# Check AI service status
/wechatwriter:config check --module ai

# Verify API keys
/wechatwriter:config list --category ai

# Test with simple prompt
/wechatwriter:generate article "Test topic" --debug
```

**Conversion Errors**
```bash
# Validate input format
/wechatwriter:convert input.md --dry-run

# Check theme compatibility
/wechatwriter:convert input.md --theme autumn-warm --validate

# Enable detailed logging
/wechatwriter:convert input.md --verbose --log-level debug
```

### Getting Help

```bash
# General help
/wechatwriter:help

# Command-specific help
/wechatwriter:help generate
/wechatwriter:help convert

# Examples and tutorials
/wechatwriter:examples
/wechatwriter:tutorials
```

## 📚 Additional Resources

### Documentation Files
- [Configuration Guide](config.md) - Complete configuration reference
- [Content Generation](generate.md) - AI content creation details
- [Conversion Guide](convert.md) - Content transformation workflows
- [Writing Assistant](write.md) - AI writing optimization
- [Image Processing](image.md) - Visual content management
- [Draft Management](draft.md) - WeChat publishing workflow
- [Content Scoring](score.md) - Performance analysis tools
- [AI Humanization](humanize.md) - Content naturalization
- [Outline Generation](outline.md) - Content structure planning

### External Resources
- [Official Plugin Documentation](https://code.claude.com/docs/en/plugins)
- [WeChat API Documentation](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)
- [GitHub Repository](https://github.com/royalrick/wechatwriter)
- [Community Support](https://github.com/royalrick/wechatwriter/discussions)

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Commands
```bash
# Run tests
make test

# Build plugin
make build

# Generate documentation
make docs

# Check code quality
make lint
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Happy Writing!** 🚀

For support, please visit our [GitHub Issues](https://github.com/royalrick/wechatwriter/issues) or start a [Discussion](https://github.com/royalrick/wechatwriter/discussions).