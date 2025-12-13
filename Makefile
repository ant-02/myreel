SHELL := /bin/bash

# 获取 Go 路径（兼容 goenv）
GO := $(shell command -v goenv >/dev/null 2>&1 && goenv which go || echo go)

# 输出目录和日志目录
OUTPUT := output/bin
LOGS := logs

# Go 服务列表
SERVICES := gateway user video like comment follow chat
BINS := $(addprefix $(OUTPUT)/, $(SERVICES))

# tmux 会话名
TMUX_SESSION := go-apps

# Docker Compose 配置
DOCKER_COMPOSE_DIR := docker
DOCKER_COMPOSE_FILE := $(DOCKER_COMPOSE_DIR)/docker-compose.yml

.PHONY: all build clean run restart up down up-and-run stop-tmux

#======== Docker ========
up:
	@echo "🐳 Starting Docker containers..."
	@cd $(DOCKER_COMPOSE_DIR) && docker compose up -d

down:
	@echo "🛑 Stopping Docker containers..."
	@cd $(DOCKER_COMPOSE_DIR) && docker compose down

#======== Go Build（每次都构建）========
build:
	@echo "🔧 Building all Go services..."
	@mkdir -p $(OUTPUT)
	@for srv in $(SERVICES); do \
		echo "Building $$srv..."; \
		$(GO) build -o $(OUTPUT)/$$srv ./cmd/$$srv; \
	done

#======== Run (tmux，构建 + 启动) ========
run: build
	@echo "🧹 Killing old tmux session..."
	-@tmux kill-session -t $(TMUX_SESSION) 2>/dev/null || true

	@echo "🚀 Starting Go services in tmux..."
	@mkdir -p $(LOGS)

	# gateway
	tmux new-session -d -s $(TMUX_SESSION) -n gateway \
		"$(OUTPUT)/gateway 2>&1 | tee -a $(LOGS)/gateway.log"

	sleep 0.3

	# 其他服务
	for srv in user video like comment follow chat; do \
		echo "Starting $$srv..."; \
		tmux new-window -t $(TMUX_SESSION) -n $$srv \
			"$(OUTPUT)/$$srv 2>&1 | tee -a $(LOGS)/$$srv.log"; \
		sleep 0.3; \
	done

	# 选择第一个窗口
	tmux select-window -t $(TMUX_SESSION):gateway

	@echo "✅ All services started. Attaching to tmux session..."
	tmux attach -t $(TMUX_SESSION)

#======== Restart (不构建，直接重开) ========
restart:
	@echo "🧹 Killing old tmux session..."
	-@tmux kill-session -t $(TMUX_SESSION) 2>/dev/null || true

	@echo "🚀 Restarting Go services in tmux..."
	@mkdir -p $(LOGS)

	# gateway
	tmux new-session -d -s $(TMUX_SESSION) -n gateway \
		"$(OUTPUT)/gateway 2>&1 | tee -a $(LOGS)/gateway.log"

	sleep 0.3

	# 其他服务
	for srv in user video like comment; do \
		echo "Restarting $$srv..."; \
		tmux new-window -t $(TMUX_SESSION) -n $$srv \
			"$(OUTPUT)/$$srv 2>&1 | tee -a $(LOGS)/$$srv.log"; \
		sleep 0.3; \
	done

	# 选择第一个窗口
	tmux select-window -t $(TMUX_SESSION):gateway

	@echo "✅ All services restarted. Attaching to tmux session..."
	tmux attach -t $(TMUX_SESSION)

#======== All-in-one: Docker + Go ========
up-and-run: up
	@echo "⏳ Waiting for dependent services..."
	sleep 2
	@$(MAKE) run

#======== Clean ========
clean:
	rm -rf $(OUTPUT) $(LOGS)

#======== Stop tmux + kill Go 服务 ========
stop-tmux:
	@echo "🛑 Killing tmux session and all Go services..."
	-@tmux list-panes -s -F "#{session_name}:#{window_index}:#{pane_pid}" | grep $(TMUX_SESSION) | \
	while read line; do \
		pid=$$(echo $$line | cut -d: -f3); \
		kill -9 $$pid 2>/dev/null || true; \
	done
	-@tmux kill-session -t $(TMUX_SESSION) 2>/dev/null || true
