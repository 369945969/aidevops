# AI DevOps Pipeline 使用文档

## 项目概述

AI DevOps Pipeline 是一个多Agent协作的智能开发管道系统。用户用一句话描述需求，系统自动触发5个AI Agent依次工作：需求分析 → 架构设计 → 代码开发 → 测试验证 → 部署上线，产出完整的开发交付物。

核心技术栈：
- 后端：Go (Gin + GORM + SQLite)
- 前端：Vue 3 + Vite + Tailwind CSS
- LLM：OpenAI兼容API（支持任意兼容接口）
- 通信：REST API + SSE实时进度推送

## 快速开始

```bash
# 1. 启动后端
cd golang
go run main.go
# 服务监听 http://localhost:8080

# 2. 启动前端
cd vue
npm install
npm run dev
# 服务监听 http://localhost:5175

# 3. 配置LLM API（必须）
export LLM_API_KEY="your-api-key"
export LLM_BASE_URL="https://api.openai.com/v1"      # 或其他兼容接口
export LLM_MODEL="gpt-4o"                              # 模型名称
export LLM_MAX_TOKENS="4096"                            # 可选
export LLM_TIMEOUT_MS="120000"                          # 可选，默认120秒
```

## 环境配置

### LLM API配置

| 环境变量 | 说明 | 默认值 |
|---|---|---|
| `LLM_API_KEY` | API密钥（必须） | 无 |
| `LLM_BASE_URL` | API基础URL | `https://api.openai.com/v1` |
| `LLM_MODEL` | 模型名称 | `gpt-4o` |
| `LLM_MAX_TOKENS` | 最大生成token数 | `4096` |
| `LLM_TIMEOUT_MS` | 请求超时（毫秒） | `120000` |

支持任何OpenAI兼容API：
- OpenAI (gpt-4o, gpt-4o-mini)
- Azure OpenAI
- Anthropic (通过兼容代理)
- 本地模型 (Ollama, vLLM, LocalAI)

### 数据库配置

系统使用SQLite单文件数据库，位于 `golang/devops.db`。首次启动自动创建表结构。

## API参考

### 需求管理

#### 创建需求 + 触发管道

```
POST /api/requirements
Body: {
  "description": "用户注册需要增加手机号短信验证码功能",
  "priority": "高",          // 高/中/低
  "mode": "full"             // quick/standard/full/review
}
Response: {
  "data": {
    "id": 1,
    "pipeline_id": 1,        // 自动创建的管道ID
    "status": "进行中"
  }
}
```

#### 获取需求列表

```
GET /api/requirements?status=进行中&priority=高
Response: { "data": [...] }
```

#### 获取需求详情

```
GET /api/requirements/:id
Response: { "data": { "id": 1, "pipeline": { "stages": [...] } } }
```

### 管道管理

#### 获取管道列表

```
GET /api/pipeline
Response: { "data": [...] }
```

#### 获取管道详情（含各阶段状态）

```
GET /api/pipeline/:id
Response: {
  "data": {
    "id": 1,
    "status": "running",      // running/completed/failed/paused_for_review
    "stages": [
      { "stage_name": "requirement_analyst", "status": "completed", "duration_ms": 15000 },
      { "stage_name": "architect", "status": "running" },
      ...
    ]
  }
}
```

#### 触发管道执行

```
POST /api/pipeline/:id/run
Body: { "mode": "full" }       // 可选，覆盖创建时的mode
Response: { "data": { "id": 1, "status": "running" } }
```

#### SSE实时进度推送

```
GET /api/pipeline/:id/progress
Content-Type: text/event-stream

事件格式:
data: {"pipeline_id":1,"stage":"architect","type":"stage_done","message":"Stage architect completed","status":"completed","duration_ms":15000}

事件类型:
- pipeline_started     管道开始执行
- stage_pending        阶段注册
- stage_done           阶段完成
- stage_failed         阶段失败
- pipeline_completed   管道全部完成
- pipeline_failed      管道执行失败
- review_pause         等待人工审核
- review_approved      审核通过（继续执行）
- review_rejected      审核拒绝（终止管道）
- heartbeat            心跳（每30秒）
```

#### 人工审核恢复

```
POST /api/pipeline/:id/resume
Body: {
  "stage_name": "architect",   // 当前暂停的阶段
  "decision": "approve",       // approve 或 reject
  "comment": "架构方案合理，可以继续开发"
}
Response: { "data": { ... }, "decision": "approve" }
```

#### 获取代码产物

```
GET /api/pipeline/:id/artifacts
Response: {
  "data": [
    { "name": "handler.go", "type": "code", "stage_name": "developer", "content": "..." }
  ]
}
```

### 工作流看板

```
GET /api/workflow/tasks
Response: { "data": [...] }
```

### 通知

```
GET /api/biz-notifications
PUT /api/biz-notifications/:id/read
```

## 道模式详解

| 模式 | 阶段数 | 流程 | 适用场景 |
|---|---|---|---|
| `quick` | 2 | RequirementAnalyst → Developer | 小改动、bug修复 |
| `standard` | 3 | RequirementAnalyst → Architect → Developer | 标准功能开发 |
| `full` | 5 | RequirementAnalyst → Architect → Developer → Tester → Deployer | 完整交付流程 |
| `review` | 5+ | full流程 + 人工审核节点 | 生产部署、高风险变更 |

关键/非关键阶段策略：
- 关键阶段（requirement_analyst, architect, developer）失败 → 终止管道
- 非关键阶段（tester, deployer）失败 → 降级继续，标记为skipped

```
full模式执行流程:
[需求输入] → RequirementAnalyst(结构化JSON) → Architect(架构+任务分解)
           → Developer(代码文件集) → Tester(测试计划+质量+安全)
           → Deployer(部署方案+回滚策略) → [完成]
```

## Agent架构说明

### RequirementAnalystAgent

职责：将自然语言需求转化为结构化JSON规格文档。
输入：用户原始需求文本。
输出：功能模块列表、验收标准、优先级、约束条件、技术建议。

### ArchitectAgent

职责：基于需求规格设计系统架构并分解开发任务。
输入：RequirementAnalyst的结构化输出 + 上游注入的specification。
输出：架构概述、技术栈、模块设计、API接口定义、任务分解列表。

### DeveloperAgent

职责：根据架构方案生成代码文件。
输入：Architect的架构设计 + 上游注入的specification和architecture。
输出：文件名→代码内容映射（map[string]string），每个文件都是完整可运行的代码。

### TesterAgent

职责：基于代码生成测试计划和质量评估。
输入：Developer的代码 + 上游注入的所有context。
输出：测试策略、单元测试列表、覆盖率目标、安全检查项、质量评分。

### DeployerAgent

职责：生成部署方案和回滚策略。
输入：Tester的测试结果 + 上游注入的所有context。
输出：部署步骤、环境配置、回滚方案、监控指标、健康检查配置。

## DevOpsContext数据流

所有Agent共享同一个DevOpsContext，每个Agent的输出自动注入下游：

```
DevOpsContext {
  Query:           "用户原始需求文本"
  PipelineData: {
    "specification":  RequirementAnalyst的输出  ← 第1阶段写入
    "architecture":   Architect的输出           ← 第2阶段写入
    "code_files":     Developer的输出           ← 第3阶段写入
    "test_results":   Tester的输出              ← 第4阶段写入
    "deploy_plan":    Deployer的输出            ← 第5阶段写入
  }
  Mode:             "full"
  PipelineID:       1
}
```

下游Agent的SystemPrompt中包含 `buildInjectedDataSection()` 生成的注入数据，确保：
- Architect能看到需求规格
- Developer能看到架构方案
- Tester能看到代码文件
- Deployer能看到测试结果

## 前端集成

### API调用层

所有管道相关API在 `vue/src/api/pipeline.js` 中定义：

```javascript
import { createRequirement, runPipeline, getPipeline, resumePipelineReview } from '../api/pipeline'

// 创建需求并触发管道
const res = await createRequirement({ description: '...', priority: '高', mode: 'full' })

// 查看管道进度
const pipeline = await getPipeline(pipelineId)

// SSE实时进度
import { usePipelineProgress } from '../composables/usePipelineProgress'
const { events, stages, status } = usePipelineProgress(pipelineId)
```

### SSE Composable

`usePipelineProgress(pipelineId)` 返回：
- `events` — 所有收到的事件列表
- `stages` — 各阶段最新状态映射 `{ architect: { status: 'completed' } }`
- `status` — 连接状态（connecting/connected/completed/failed）
- `error` — 错误信息
- 自动重连（5秒间隔）

### mockData迁移计划

当前以下页面仍使用mockData，计划逐步迁移到API：
- Requirements列表页 → `getRequirements()`
- Knowledge页面 → 知识库API（待实现）
- Settings各页面 → 已有完整Settings API

## 扩展指南

### 添加新Agent

1. 在 `golang/agent/agents/` 创建新文件，实现 `DevOpsAgent` 接口：
```go
type MyAgent struct {
    agent.BaseDevOpsAgent
}

func (a *MyAgent) AgentName() string { return "my_agent" }
func (a *MyAgent) SystemPrompt() string { return "..." }
func (a *MyAgent) Run(ctx agent.DevOpsContext) (*agent.AgentOutput, error) { ... }
```

2. 在 `orchestrator.go` 的 `BuildAgentChain()` 中注册：
```go
chain = append(chain, &MyAgent{BaseDevOpsAgent: agent.BaseDevOpsAgent{...}})
```

3. 在 `models/business_models.go` 中添加对应的数据模型（如需要）。

### 添加新工具

1. 定义ToolDefinition和ToolExecutor：
```go
def := agent.ToolDefinition{
    Name: "my_tool",
    Description: "描述工具功能",
    Parameters: map[string]agent.ParamSpec{
        "input": { Type: "string", Required: true },
    },
}
executor := func(params map[string]interface{}) agent.ToolResult { ... }
```

2. 注册到ToolRegistry：
```go
registry := agent.NewToolRegistry()
agent.RegisterBuiltInTools(registry, workspacePath)
registry.Register(def, executor)
```

### 添加审核节点

在 `review` 模式下，orchestrator会在指定阶段暂停并标记 `paused_for_review`。
可通过 `POST /api/pipeline/:id/resume` 审核通过或拒绝。

自定义审核节点位置：修改 `orchestrator.go` 中 `reviewStages` 配置。

## 故障排查

### LLM调用失败

症状：管道阶段报错 `LLM API call failed`。
排查：
1. 检查 `LLM_API_KEY` 是否正确
2. 检查 `LLM_BASE_URL` 是否可访问
3. 检查 `LLM_MODEL` 是否存在
4. 增加超时：`LLM_TIMEOUT_MS=300000`

### 管道超时

症状：管道状态一直为 `running`，未完成。
排查：
1. 查看具体阶段状态 `GET /api/pipeline/:id`
2. 检查是否有阶段卡在LLM调用
3. 检查 `LLM_TIMEOUT_MS` 配置
4. 管道总超时默认10分钟，可在orchestrator中调整

### SQLite锁冲突

症状：`database is locked` 错误。
排查：系统已配置 `MaxOpenConns=1` 避免并发写入冲突。如果仍有问题，检查是否有外部进程同时访问数据库文件。

### SSE连接断开

症状：前端进度推送中断。
排查：
1. SSE自动重连（5秒间隔）
2. 检查后端是否重启（重启后EventBus清空）
3. 管道完成后SSE自动断开

### Agent输出解析失败

症状：阶段状态为 `failed`，错误为 `JSON extraction failed`。
排查：
1. LLM可能返回非JSON格式内容
2. 系统有markdown JSON提取逻辑（从```json块中提取）
3. 如果持续失败，考虑使用更稳定的模型或调整SystemPrompt
