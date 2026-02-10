# Writer CLI Makefile
# 微信公众号写作工具统一构建

.PHONY: all build clean test install help lint fmt vet release sync

# 默认目标
all: build

# 构建所有平台的二进制文件（发布到 scripts/ 目录）
release:
	@echo "🔨 构建 wechatwriter 所有平台版本..."
	@echo ""
	@echo "📦 Building for Linux amd64..."
	@GOOS=linux GOARCH=amd64 go build -o scripts/wechatwriter-linux-amd64 ./app
	@echo "✓ Linux amd64"
	@echo "📦 Building for Linux arm64..."
	@GOOS=linux GOARCH=arm64 go build -o scripts/wechatwriter-linux-arm64 ./app
	@echo "✓ Linux arm64"
	@echo "📦 Building for macOS amd64 (Intel)..."
	@GOOS=darwin GOARCH=amd64 go build -o scripts/wechatwriter-darwin-amd64 ./app
	@echo "✓ macOS amd64"
	@echo "📦 Building for macOS arm64 (Apple Silicon)..."
	@GOOS=darwin GOARCH=arm64 go build -o scripts/wechatwriter-darwin-arm64 ./app
	@echo "✓ macOS arm64"
	@echo "📦 Building for Windows amd64..."
	@GOOS=windows GOARCH=amd64 go build -o scripts/wechatwriter-windows-amd64.exe ./app
	@echo "✓ Windows amd64"
	@echo ""
	@chmod +x scripts/*-linux* scripts/*-darwin* 2>/dev/null || true
	@echo "✅ 构建完成！二进制文件在 scripts/ 目录"
	@echo ""
	@ls -lh scripts/

# 构建当前平台
build:
	@echo "🔨 构建当前平台..."
	@go build -o scripts/wechatwriter ./app/*.go
	@echo "✅ 构建完成: ./scripts/wechatwriter"

# 快速构建（仅当前平台，用于开发）
fast:
	@go build -o wechatwriter ./app

# 清理
clean:
	@echo "🧹 清理..."
	@rm -f scripts/wechatwriter
	@rm -rf dist/ release/
	@rm -f *.log

# 运行测试
test:
	@echo "🧪 运行测试..."
	@go test -v ./...

# 代码检查
lint:
	@echo "🔍 代码检查..."
	@golangci-lint run ./... 2>/dev/null || echo "  (需要安装 golangci-lint)"

# 格式化代码
fmt:
	@echo "🎨 格式化代码..."
	@go fmt ./...
	@gofmt -w .

# 静态分析
vet:
	@echo "🔬 静态分析..."
	@go vet ./...

# 安装到 GOPATH/bin
install:
	@echo "📦 安装到 $(GOPATH)/bin..."
	@go install ./app

# 下载依赖
deps:
	@echo "📥 下载依赖..."
	@go mod download
	@go mod tidy

# 同步 Skill 目录
sync:
	@echo "🔄 同步 Skill 目录..."
	@bash scripts/sync.sh

# 帮助
help:
	@echo "Writer CLI Makefile 命令:"
	@echo ""
	@echo "开发命令:"
	@echo "  make build       - 构建当前平台二进制"
	@echo "  make fast        - 快速构建（开发用）"
	@echo "  make release     - 构建所有平台二进制到 bin/"
	@echo "  make clean       - 清理构建文件"
	@echo ""
	@echo "代码质量:"
	@echo "  make fmt         - 格式化代码"
	@echo "  make vet         - 静态分析"
	@echo "  make test        - 运行测试"
	@echo ""
	@echo "依赖管理:"
	@echo "  make deps        - 下载依赖"
	@echo "  make install     - 安装到 GOPATH/bin"
	@echo ""
	@echo "文档同步:"
	@echo "  make sync        - 同步 Skill 目录到插件目录"
	@echo ""
	@echo "用户快速安装:"
	@echo "  go install github.com/royalrick/wechatwriter/app/cmd/writer@latest"
