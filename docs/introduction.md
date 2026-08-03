# Feishu-ChatGPT 项目介绍

## 项目概述

**Feishu-ChatGPT** 是一个将 OpenAI 的 ChatGPT 与飞书（Lark）即时通讯平台深度集成的开源机器人项目。用户可以在飞书的私聊或群聊中，直接与 ChatGPT 进行对话，享受 AI 驱动的智能助手体验。

该项目基于 **Go 语言**开发，具有高性能、低资源消耗的特点，支持多种灵活的部署方式，适合个人开发者及企业团队使用。

---

## 核心功能

| 功能 | 说明 |
|------|------|
| 🗣 **语音交流** | 支持私人语音消息，直接与机器人畅所欲言 |
| 💬 **多话题对话** | 支持私聊和群聊中的多话题独立讨论，上下文高效连贯 |
| 🖼 **文本成图** | 支持文本生成图片（DALL·E）及以图搜图 |
| 🛖 **场景预设** | 内置丰富的角色场景列表，一键切换 AI 角色 |
| 🎭 **角色扮演** | 支持自定义场景模式，增添对话趣味性与创意 |
| 🤖 **多种 AI 模式** | 内置多种 AI 模式，体验不同智慧与创意风格 |
| 🔄 **上下文保留** | 回复对话框即可继续同一话题讨论 |
| ⏰ **自动超时结束** | 超时自动结束对话，支持清除历史记录 |
| 📝 **富文本卡片** | 以飞书消息卡片形式呈现，信息展示更丰富 |
| 👍 **交互式反馈** | 即时获取机器人处理进度与结果 |
| 🎰 **余额查询** | 实时查询 OpenAI Token 消耗情况 |
| 🌐 **多 Token 负载均衡** | 支持多个 API Key 轮询，优化高频调用场景 |
| ↩️ **反向代理支持** | 为不同地区用户提供更快、更稳定的访问体验 |

---

## 技术架构

```
feishu-chatgpt/
├── code/                    # 主程序源码
│   ├── main.go              # 程序入口，初始化并启动 HTTP 服务
│   ├── handlers/            # 飞书事件与卡片消息处理器
│   │   ├── handler.go       # 事件分发入口
│   │   ├── event_msg_action.go    # 文本消息处理
│   │   ├── event_audio_action.go  # 语音消息处理
│   │   ├── event_pic_action.go    # 图片消息处理
│   │   └── card_*_action.go      # 卡片交互处理
│   ├── services/            # 外部服务封装
│   │   ├── openai/          # OpenAI API 调用（ChatGPT、DALL·E、Whisper）
│   │   └── loadbalancer/    # 多 Key 负载均衡
│   ├── initialization/      # 配置加载与客户端初始化
│   ├── config.example.yaml  # 配置文件示例
│   └── role_list.yaml       # 内置角色列表
├── docs/                    # 项目文档与截图
├── Dockerfile               # Docker 镜像构建文件
├── docker-compose.yaml      # Docker Compose 编排文件
└── s.yaml                   # Serverless 部署配置
```

### 技术栈

- **语言**：Go 1.18+
- **Web 框架**：[Gin](https://github.com/gin-gonic/gin)
- **飞书 SDK**：[oapi-sdk-go v3](https://github.com/larksuite/oapi-sdk-go) + [oapi-sdk-gin](https://github.com/larksuite/oapi-sdk-gin)
- **AI 能力**：OpenAI API（GPT-3.5-turbo / GPT-4、DALL·E、Whisper）及 Azure OpenAI
- **缓存**：[go-cache](https://github.com/patrickmn/go-cache)（内存键值对缓存）
- **配置管理**：[Viper](https://github.com/spf13/viper)

---

## 工作原理

1. **消息接收**：飞书将用户消息通过 Webhook 推送到本服务的 `/webhook/event` 接口。
2. **事件分发**：`handlers` 模块根据消息类型（文本、语音、图片）分发到对应的处理器。
3. **AI 调用**：处理器调用 `services/openai` 模块向 OpenAI API 发送请求，获取 AI 回复。
4. **结果回复**：将 AI 回复以飞书消息卡片的形式，通过 Lark SDK 发回给用户。
5. **卡片交互**：用户点击卡片上的按钮时，飞书将事件推送到 `/webhook/card` 接口进行处理。

---

## 快速开始

### 前置条件

- [飞书开放平台](https://open.feishu.cn/) 账号，已创建企业自建应用
- [OpenAI API Key](https://platform.openai.com/account/api-keys)
- Go 1.18+（本地部署）或 Docker

### 配置说明

复制 `code/config.example.yaml` 为 `code/config.yaml`，并填写以下关键配置：

```yaml
# 飞书应用凭证
APP_ID: cli_axxx               # 飞书应用 App ID
APP_SECRET: xxx                # 飞书应用 App Secret
APP_ENCRYPT_KEY: xxx           # 飞书事件加密密钥
APP_VERIFICATION_TOKEN: xxx    # 飞书事件验证 Token
BOT_NAME: chatGpt              # 机器人名称（需与飞书后台一致）

# OpenAI 配置
OPENAI_KEY: sk-xxx,sk-xxx      # 支持多个 Key，逗号分隔（负载均衡）
API_URL: https://api.openai.com
HTTP_PROXY: ""                 # HTTP 代理（可选）

# 服务器配置
HTTP_PORT: 9000
```

### 部署方式

本项目支持多种部署方式，按需选择：

| 方式 | 适合场景 |
|------|----------|
| **本地 + 内网穿透** | 开发测试，使用 cpolar 或 natapp 暴露公网地址 |
| **Docker** | 生产环境，快速启动，资源隔离 |
| **Docker Compose** | 生产环境，配置管理更便捷 |
| **Serverless（阿里云）** | 弹性扩缩容，按需付费 |
| **Railway** | 一键部署到海外云平台 |
| **二进制包** | 直接下载运行，无需编译 |

> 详细部署步骤请参阅项目根目录的 [readme.md](../readme.md)。

---

## 飞书机器人配置要点

在飞书开发者后台完成以下配置后，机器人才能正常工作：

1. **事件订阅**：将回调地址设置为 `http://<你的公网地址>/webhook/event`
2. **消息卡片**：将请求地址设置为 `http://<你的公网地址>/webhook/card`
3. **权限申请**：开启以下权限
   - `im:message`（接收与发送消息）
   - `im:message.group_at_msg`（群组 @ 消息）
   - `im:message.p2p_msg`（单聊消息）
   - `im:resource`（图片与文件资源）
   - `im:chat:readonly`（获取群组信息）

---

## Azure OpenAI 支持

本项目同时支持 Azure OpenAI 服务，在 `config.yaml` 中启用：

```yaml
AZURE_ON: true
AZURE_API_VERSION: 2023-03-15-preview
AZURE_RESOURCE_NAME: your-resource-name
AZURE_DEPLOYMENT_NAME: your-deployment-name
AZURE_OPENAI_TOKEN: your-azure-token
```

---

## 参与贡献

欢迎提交 Issue 和 Pull Request！在贡献代码前，请先阅读现有代码风格，并确保：

- 遵循 Go 语言代码规范
- 新功能附带必要的说明
- 保持向后兼容性

---

## 相关链接

- 🏠 [GitHub 仓库](https://github.com/Leizhenpeng/feishu-chatgpt)
- 📖 [飞书开放平台文档](https://open.feishu.cn/document/)
- 🤖 [OpenAI API 文档](https://platform.openai.com/docs/)
- 🌐 [企联AI 开源社区](https://github.com/ConnectAI-E)
