SHELL := /bin/bash

# 获取 go 路径（兼容 goenv）
GO := $(shell command -v goenv >/dev/null 2>&1 && goenv which go || echo go)

OUTPUT_DIR := output
GATEWAY_DIR := $(OUTPUT_DIR)/gateway
USER_DIR := $(OUTPUT_DIR)/user
VIDEO_DIR := $(OUTPUT_DIR)/video
TMUX_SESSION := go-apps
DOCKER_COMPOSE_DIR := docker
DOCKER_COMPOSE_FILE := $(DOCKER_COMPOSE_DIR)/docker-compose.yml

.PHONY: build build-gateway build-user run-tmux clean up down up-and-run

# 启动 Docker 容器
up:
	@echo "🐳 Starting Docker containers..."
	@cd $(DOCKER_COMPOSE_DIR) && docker compose up -d

# 停止 Docker 容器
down:
	@echo "🛑 Stopping Docker containers..."
	@cd $(DOCKER_COMPOSE_DIR) && docker compose down

# 构建 Go 服务
build: build-gateway build-user build-video

build-gateway:
	$(GO) build -o $(GATEWAY_DIR) ./cmd/gateway

build-user:
	$(GO) build -o $(USER_DIR) ./cmd/user

build-video:
	$(GO) build -o $(VIDEO_DIR) ./cmd/video

run:
	@echo "🔧 Building Go services..."
	@$(MAKE) build

	@if [ ! -f "$(USER_DIR)" ]; then echo "❌ $(USER_DIR) not built!"; exit 1; fi
	@if [ -f "$(GATEWAY_DIR)" ]; then echo "✅ Gateway built"; else echo "❌ $(GATEWAY_DIR) not built!"; exit 1; fi
	@if [ -f "$(VIDEO_DIR)" ]; then echo "✅ Video built"; else echo "❌ $(VIDEO_DIR) not built!"; exit 1; fi

	@echo "🧹 Killing old tmux session..."
	-tmux kill-session -t $(TMUX_SESSION) 2>/dev/null

	@echo "🚀 Starting Go services in tmux..."
	tmux new-session -d -s $(TMUX_SESSION) "$(GATEWAY_DIR)"
	sleep 0.2
	tmux split-window -h -t $(TMUX_SESSION) "$(USER_DIR)"
	sleep 0.2 
	tmux split-window -h -t $(TMUX_SESSION) "$(VIDEO_DIR)"

	@echo "✅ Attaching to tmux session: $(TMUX_SESSION)"
	tmux attach -t $(TMUX_SESSION)

# 启动 Docker + Go 服务（一体化）
up-and-run: up 
	@echo "⏳ Waiting for services to be ready..."
	# 可选：等待 MySQL/Redis 就绪（简单 sleep）
	sleep 5
	
	@$(MAKE) run

# 清理 Go 构建产物
clean:
	rm -rf $(OUTPUT_DIR)

# 完整清理：停容器 + 清构建
up-and-run-clean: down clean