SHELL := /bin/bash

# 获取 goenv 的 go 路径，如果 goenv 不可用则回退到 go
GO := $(shell command -v goenv >/dev/null 2>&1 && goenv which go || echo go)

OUTPUT_DIR := output
GATEWAY_DIR := $(OUTPUT_DIR)/gateway
USER_DIR := $(OUTPUT_DIR)/user
TMUX_SESSION := go-apps

.PHONY: build build-gateway build-user run-tmux clean

build: build-gateway build-user

build-gateway:
	go build -o $(GATEWAY_DIR) ./cmd/gateway

build-user:
	go build -o $(USER_DIR) ./cmd/user

run-tmux:
	@mkdir -p $(OUTPUT_DIR)
	@echo "🔧 Building services..."
	@$(MAKE) build

	@if [ ! -f "$(USER_DIR)" ]; then echo "❌ $(USER_DIR) not built!"; exit 1; fi
	@if [ ! -f "$(GATEWAY_DIR)" ]; then echo "❌ $(GATEWAY_DIR) not built!"; exit 1; fi

	@echo "🧹 Killing old session..."
	-tmux kill-session -t $(TMUX_SESSION) 2>/dev/null

	@echo "🚀 Starting gateway and user in tmux..."
	tmux new-session -d -s $(TMUX_SESSION) "$(GATEWAY_DIR)"
	sleep 0.2
	tmux split-window -h -t $(TMUX_SESSION) "$(USER_DIR)"

	@echo "✅ Attaching to tmux session: $(TMUX_SESSION)"
	tmux -CC attach -t $(TMUX_SESSION)

clean:
	rm -rf $(OUTPUT_DIR)