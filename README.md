# myreel

## 技术栈

| 组件     | 技术  |
| -------- | ----- |
| WEB框架  | Hertz |
| RPC框架  | Kitex |
| 数据库   | MySQL |
| 缓存     | Redis |
| 服务注册 | etcd  |

## 项目结构

```bash
.
├── app	# 业务具体实现
│   ├── chat	# 聊天服务
│   │   ├── controller	
│   │   │   └── rpc # 聊天接口处理
│   │   │       └── pack	# 接口数据转化工具
│   │   ├── domain # 领域层
│   │   │   ├── model # 内部传输模型
│   │   │   ├── repository	# 数据存储接口
│   │   │   └── service	# 业务实现
│   │   ├── infrastructure # 数据存储实现
│   │   │   ├── cache	# 缓存数据
│   │   │   │   └── pack # 数据转化工具
│   │   │   ├── mysql # 持久化数据
│   │   │   └── rpc	# 微服务调用
│   │   └── usecase # 用例层 业务排版
│   ├── comment # 评论服务
│   │   ├── controller 
│   │   │   └── rpc
│   │   │       └── pack 
│   │   ├── domain 
│   │   │   ├── model 
│   │   │   ├── repository 
│   │   │   └── service
│   │   ├── infrastructure
│   │   │   ├── mysql
│   │   │   └── rpc
│   │   └── usecase
│   ├── follow # 好友服务
│   │   ├── controller
│   │   │   └── rpc
│   │   │       └── pack
│   │   ├── domain
│   │   │   ├── model
│   │   │   ├── repository
│   │   │   └── service
│   │   ├── infrastructure
│   │   │   ├── mysql
│   │   │   └── rpc
│   │   └── usecase
│   ├── gateway # http网关
│   │   ├── handler # 接口数据处理
│   │   │   └── api
│   │   │       ├── chat # 聊天
│   │   │       │   └── pack
│   │   │       ├── comment # 评论
│   │   │       ├── follow # 好友
│   │   │       ├── like # 点赞
│   │   │       ├── user # 用户
│   │   │       └── video # 视频
│   │   ├── model # hertz生成数据传输模型
│   │   │   ├── api
│   │   │   │   ├── chat
│   │   │   │   ├── comment
│   │   │   │   ├── follow
│   │   │   │   ├── like
│   │   │   │   ├── user
│   │   │   │   └── video
│   │   │   └── model
│   │   ├── mv # 中间件 身份认证
│   │   ├── pack # 数据转化，统一返回数据格式
│   │   ├── router # hertz生成路由层
│   │   │   ├── chat
│   │   │   ├── comment
│   │   │   ├── follow
│   │   │   ├── like
│   │   │   ├── user
│   │   │   └── video
│   │   └── rpc # 微服务调用层
│   ├── like # 点赞服务
│   │   ├── controller
│   │   │   └── rpc
│   │   │       └── pack
│   │   ├── domain
│   │   │   ├── model
│   │   │   ├── repository
│   │   │   └── service
│   │   ├── infrastructure
│   │   │   ├── cache
│   │   │   ├── mysql
│   │   │   └── rpc
│   │   └── usecase
│   ├── user # 用户服务
│   │   ├── controller
│   │   │   └── rpc
│   │   │       └── pack
│   │   ├── domain
│   │   │   ├── model
│   │   │   ├── repository
│   │   │   └── service
│   │   ├── infrastructure
│   │   │   ├── cache
│   │   │   └── mysql
│   │   └── usecase
│   └── video # 视频服务
│       ├── controller
│       │   └── rpc
│       │       └── pack
│       ├── domain
│       │   ├── model
│       │   ├── repository
│       │   └── service
│       ├── infrastructure
│       │   ├── cache
│       │   │   └── pack
│       │   ├── mysql
│       │   └── rpc
│       └── usecase
├── cmd # 各服务函数入口
│   ├── chat
│   ├── comment
│   ├── follow
│   ├── gateway
│   ├── like
│   ├── user
│   └── video
├── config # 配置信息，容器配置卷映射
│   ├── elasticsearch # es配置
│   │   ├── config
│   │   └── plugins
│   │       └── ik
│   │           └── config
│   ├── filebeat # filebeat配置
│   ├── kibana # kibana配置
│   └── mysql # mysql配置
├── docker # docker服务
│   ├── data	# 卷映射容器数据存储
│   │   ├── elasticsearch
│   │   │   ├── snapshot_cache
│   │   │   └── _state
│   │   ├── etcd
│   │   ├── filebeat
│   │   │   └── registry
│   │   │       └── filebeat
│   │   ├── kibana
│   │   ├── mysql
│   │   │   ├── #innodb_redo
│   │   │   ├── #innodb_temp
│   │   │   ├── myreel
│   │   │   ├── mysql
│   │   │   ├── performance_schema
│   │   │   └── sys
│   │   └── redis
│   ├── env # 容器启动环境配置
│   └── script # makefile 启动服务调用的脚本
├── idl	# hertz和kitex的proto文件
│   └── api # hertz
├── kitex_gen # kitex生成
│   ├── chat
│   │   └── chatservice
│   ├── comment
│   │   └── commentservice
│   ├── follow
│   │   └── followservice
│   ├── like
│   │   └── likeservice
│   ├── model
│   ├── user
│   │   └── userservice
│   └── video
│       └── videoservice
├── output # 代码编译执行文件目录
│   ├── chat
│   ├── comment
│   ├── follow
│   ├── gateway
│   ├── like
│   ├── log # 日志存储
│   │   └── gateway
│   │       └── 2025-12-16
│   ├── user 
│   └── video
├── pkg # 公共工具目录
│   ├── base # 基础工具
│   │   ├── client # 统一客户端：redis、mysql、rpcClient
│   │   └── context # 上下文统一
│   ├── constants # 统一常量
│   ├── errno # 错误工具
│   ├── logger # 日志工具
│   ├── upyun # 有拍云工具，上传图片、视频
│   └── util # 工具
└── script # hertz生成脚本
```

