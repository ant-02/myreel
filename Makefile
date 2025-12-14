# 检查 tmux 是否存在
TMUX_EXISTS := $(shell command -v tmux)
# 当前架构
ARCH := $(shell uname -m)
PREFIX = "[Makefile]"
# 目录相关
DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
OUTPUT_PATH = $(DIR)/output
API_PATH= $(DIR)/cmd/api


# 服务名
SERVICES := gateway user chat comment follow like video
service = $(word 1, $@)

# 启动必要的环境，比如 etcd、mysql
.PHONY: env-up
env-up:
	@ docker compose -f ./docker/docker-compose.yml up -d

# 关闭必要的环境，但不清理 data（位于 docker/data 目录中）
.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

# 基于 idl 生成相关的 go 语言描述文件
.PHONY: kitex-gen-%
kitex-gen-%:
	@ kitex -module "${MODULE}" \
		-proto no_default_serdes \
		${IDL_PATH}/$*.proto
	@ go mod tidy

# 生成基于 Hertz 的脚手架
.PHONY: hz-%
hz-%:
	hz update -idl ${IDL_PATH}/api/$*.proto


# 构建指定对象，构建后在没有给 BUILD_ONLY 参的情况下会自动运行，需要熟悉 tmux 环境
# 用于本地调试
.PHONY: $(SERVICES)
$(SERVICES):
	@if [ -z "$(TMUX_EXISTS)" ]; then \
		echo "$(PREFIX) tmux is not installed. Please install tmux first."; \
		exit 1; \
	fi
	@if [ -z "$$TMUX" ]; then \
		echo "$(PREFIX) you are not in tmux, press ENTER to start tmux environment."; \
		read -r; \
		if tmux has-session -t myreel 2>/dev/null; then \
			echo "$(PREFIX) Tmux session 'myreel' already exists. Attaching to session and running command."; \
			tmux attach-session -t myreel; \
			tmux send-keys -t myreel "make $(service)" C-m; \
		else \
			echo "$(PREFIX) No tmux session found. Creating a new session."; \
			tmux new-session -s myreel "make $(service); $$SHELL"; \
		fi; \
	else \
		echo "$(PREFIX) Build $(service) target..."; \
		mkdir -p output; \
		bash $(DIR)/docker/script/build.sh $(service); \
		echo "$(PREFIX) Build $(service) target completed"; \
	fi
ifndef BUILD_ONLY
	@echo "$(PREFIX) Automatic run server"
	@if tmux list-windows -F '#{window_name}' | grep -q "^myreel-$(service)$$"; then \
		echo "$(PREFIX) Window 'myreel-$(service)' already exists. Reusing the window."; \
		tmux select-window -t "myreel-$(service)"; \
	else \
		echo "$(PREFIX) Window 'myreel-$(service)' does not exist. Creating a new window."; \
		tmux new-window -n "myreel-$(service)"; \
		tmux split-window -h ; \
		tmux select-layout -t "myreel-$(service)" even-horizontal; \
	fi
	@echo "$(PREFIX) Running $(service) service in tmux..."
	@tmux send-keys -t myreel-$(service).0 'export SERVICE=$(service) && bash ./docker/script/entrypoint.sh' C-m
	@tmux select-pane -t myreel-$(service).1
endif

# 清除所有的构建产物
.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print

# 清除所有构建产物、compose 环境和它的数据
.PHONY: clean-all
clean-all: clean
	@echo "$(PREFIX) Checking if docker-compose services are running..."
	@docker-compose -f ./docker/docker-compose.yml ps -q | grep '.' && docker-compose -f ./docker/docker-compose.yml down || echo "$(PREFIX) No services are running."
	@echo "$(PREFIX) Removing docker data..."
	rm -rf ./docker/data

.PHONY: tidy
tidy:
	go mod tidy

