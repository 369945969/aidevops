export const navItems = [
  { label: '工作流看板', icon: 'lucide:layout-grid', route: '/workflow/dashboard', key: 'workflow' },
  { label: '需求管理', icon: 'lucide:file-text', route: '/requirements', key: 'requirements' },
  { label: '知识库', icon: 'lucide:library', route: '/knowledge', key: 'knowledge' },
  { label: '系统设置', icon: 'lucide:settings', route: '/settings', key: 'settings' },
]

export const knowledgeCategories = [
  { label: '全部知识条目', icon: 'lucide:library', route: '/knowledge', key: 'all' },
  { label: '需求文档', icon: 'lucide:file-text', route: '/knowledge/requirement-docs', key: 'requirement-docs' },
  { label: '架构方案', icon: 'lucide:blocks', route: '/knowledge/architecture', key: 'architecture' },
  { label: '代码模式', icon: 'lucide:code-2', route: '/knowledge/code-patterns', key: 'code-patterns' },
  { label: '问题解决方案', icon: 'lucide:puzzle', route: '/knowledge/solutions', key: 'solutions' },
  { label: '部署配置', icon: 'lucide:server', route: '/knowledge/deploy-config', key: 'deploy-config' },
  { label: '我的收藏', icon: 'lucide:star', route: '/knowledge/favorites', key: 'favorites' },
]

export const settingsSections = [
  { label: '自动化配置', icon: 'lucide:cpu', route: '/settings', key: 'automation' },
  { label: '审核规则', icon: 'lucide:shield-check', route: '/settings/review-rules', key: 'review-rules' },
  { label: '开发环境', icon: 'lucide:terminal', route: '/settings/dev-environment', key: 'dev-environment' },
  { label: '模型管理', icon: 'lucide:brain', route: '/settings', key: 'model-management' },
  { label: 'AI Agent配置', icon: 'lucide:bot', route: '/settings/ai-agent', key: 'ai-agent' },
  { label: '团队成员', icon: 'lucide:users', route: '/settings/team-members', key: 'team-members' },
  { label: '通知设置', icon: 'lucide:bell', route: '/settings/notifications', key: 'notifications' },
  { label: '集成配置', icon: 'lucide:plug', route: '/settings/integrations', key: 'integrations' },
  { label: '安全与权限', icon: 'lucide:lock', route: '/settings/security', key: 'security' },
  { label: '账单与订阅', icon: 'lucide:credit-card', route: '/settings', key: 'billing' },
]

export const requirements = [
  { id: 'REQ-0012', title: '用户注册手机号验证', status: '需求分析', priority: '高', description: '用户注册流程中需要增加手机号短信验证码验证功能，确保用户身份真实性，减少虚假账号注册。支持中国大陆手机号格式...', author: 'Sarah', authorAvatar: 'sarah', createdAt: '2026-06-04', updatedAt: '2小时前', taskCount: 0 },
  { id: 'REQ-0008', title: '多语言支持模块', status: '架构设计', priority: '高', description: '系统需要支持中英文双语界面切换，需要设计可扩展的国际化架构，支持动态语言包加载和翻译管理...', author: 'Alex', authorAvatar: 'alex', createdAt: '2026-06-02', updatedAt: '6小时前', taskCount: 3 },
  { id: 'REQ-0007', title: '短信验证码功能', status: '开发中', priority: '高', description: '实现短信验证码发送、验证和过期管理功能，集成第三方短信服务商API，使用Redis存储验证码并设置5分钟过期时间...', author: 'Sarah', authorAvatar: 'sarah', createdAt: '2026-05-30', updatedAt: '1天前', taskCount: 4 },
  { id: 'REQ-0005', title: '用户权限管理优化', status: '代码审查', priority: '高', description: '优化现有权限管理系统，支持RBAC模型，增加角色继承和权限组功能，提供更灵活的权限配置方案...', author: 'Alex', authorAvatar: 'alex', createdAt: '2026-05-28', updatedAt: '3天前', taskCount: 2 },
  { id: 'REQ-0003', title: '支付流程重构', status: '已完成', priority: '高', description: '重构支付流程，将原有的单一支付入口改为多渠道支付，支持微信支付、支付宝和银行卡支付三种方式...', author: 'Alex', authorAvatar: 'alex', createdAt: '2026-05-20', updatedAt: null, taskCount: 5 },
]

export const statusConfig = {
  '需求分析': { bg: '#E8F0FE', text: '#3367D6' },
  '架构设计': { bg: '#FEF6E8', text: '#F59D0D' },
  '开发中': { bg: '#E8F7F0', text: '#0F8B5D' },
  '代码审查': { bg: '#FEF6E8', text: '#F59D0D' },
  '测试': { bg: '#E8F0FE', text: '#3367D6' },
  '待审核': { bg: '#FEF6E8', text: '#F59D0D' },
  '需人工审核': { bg: '#FDECEA', text: '#D93025' },
  '已完成': { bg: '#E8F7F0', text: '#0F8B5D' },
  '已归档': { bg: 'hsla(200,15%,95%,1)', text: '#4A5259' },
}

export const priorityConfig = {
  '高': { bg: '#FDECEA', text: '#D93025' },
  '中': { bg: '#FEF6E8', text: '#F59D0D' },
  '低': { bg: '#E8F7F0', text: '#0F8B5D' },
}

export const knowledgeTypeConfig = {
  '架构方案': { bg: '#E8F0FE', text: '#3367D6', icon: 'lucide:blocks' },
  '代码模式': { bg: '#E8F7F0', text: '#0F8B5D', icon: 'lucide:code-2' },
  '问题解决方案': { bg: '#FEF6E8', text: '#F59D0D', icon: 'lucide:puzzle' },
  '部署配置': { bg: '#E8F0FE', text: '#3367D6', icon: 'lucide:server' },
  '需求文档': { bg: 'hsla(200,15%,95%,1)', text: '#4A5259', icon: 'lucide:file-text' },
}

export const knowledgeEntries = [
  { id: 1, title: '短信服务架构设计模式', type: '架构方案', description: '基于Redis缓存的验证码存储架构,包含发送频率限制、过期管理、加密存储等核心设计要点,适用于短信验证码类功能开发。', tags: ['Redis', 'Node.js', '短信API'], citations: 28, quality: 4.5, source: 'REQ-0007', date: '2026-06-01', aiRecommended: true, recommendationReason: '与当前需求REQ-0012高度相关' },
  { id: 2, title: 'React表单验证Hook模式', type: '代码模式', description: '自定义React Hook用于手机号格式校验和验证码输入,包含防抖提交、错误状态管理、重发倒计时等通用逻辑封装。', tags: ['React', 'TypeScript'], citations: 15, quality: 4.2, source: 'REQ-0007', date: '2026-05-31', aiRecommended: false },
  { id: 3, title: 'Redis连接池超时问题', type: '问题解决方案', description: '高并发场景下Redis连接池耗尽导致验证码存储超时,解决方案: 增加连接池上限、实现连接复用、添加降级策略。', tags: ['Redis', '性能'], citations: 42, quality: 4.8, source: 'REQ-0005', date: '2026-05-25', aiRecommended: false },
  { id: 4, title: 'Node.js生产环境部署模板', type: '部署配置', description: '标准化的Node.js服务部署配置模板,包含Dockerfile、K8s部署YAML、健康检查配置、日志收集和监控设置。', tags: ['DevOps', 'Docker'], citations: 67, quality: 4.9, source: 'REQ-0003', date: '2026-05-20', aiRecommended: false },
  { id: 5, title: '支付流程重构需求分析', type: '需求文档', description: '多渠道支付重构的完整需求分析文档,涵盖业务目标、用户场景、功能范围和验收标准,作为后续类似支付需求的参考。', tags: ['支付', '业务'], citations: 33, quality: 4.3, source: 'REQ-0003', date: '2026-05-18', aiRecommended: false },
  { id: 6, title: 'RBAC权限管理实现模式', type: '代码模式', description: '基于角色的访问控制(RBAC)完整实现,包含角色继承、权限组、动态权限加载等核心代码模式,适用于权限管理功能开发。', tags: ['安全', 'Node.js'], citations: 19, quality: 4.1, source: 'REQ-0005', date: '2026-05-28', aiRecommended: false },
]

export const kanbanColumns = [
  { key: 'requirement-analysis', label: '需求分析', icon: 'lucide:file-text', iconColor: '#3367D6', count: 2, tasks: [
    { id: 'REQ-0012', title: '用户注册手机号验证', agent: '需求分析师', priority: '高', progress: 60, estimate: '预计2h后完成' },
    { id: 'REQ-0013', title: '订单导出Excel功能', agent: '需求分析师', priority: '中', progress: 30, estimate: '预计4h后完成' },
  ]},
  { key: 'architecture-design', label: '架构设计', icon: 'lucide:blocks', iconColor: '#F59D0D', count: 1, warning: true, tasks: [
    { id: 'REQ-0008', title: '多语言支持模块', agent: '架构师', priority: '高', progress: 100, status: '待审核', reviewRoute: '/review/architecture' },
  ]},
  { key: 'development', label: '开发中', icon: 'lucide:code-2', iconColor: '#0F8B5D', count: 4, tasks: [
    { id: 'REQ-0007-A', title: '短信验证码API', agent: '后端开发', priority: '高', progress: 75, estimate: '预计1h后完成' },
    { id: 'REQ-0007-B', title: '手机号验证前端表单', agent: '前端开发', priority: '中', progress: 55, estimate: '预计2h后完成' },
    { id: 'REQ-0007-C', title: 'Redis验证码存储', agent: '后端开发', priority: '高', progress: 80 },
    { id: 'REQ-0006-A', title: '数据报表页面', agent: '前端开发', priority: '低', progress: 40 },
  ]},
  { key: 'code-review', label: '代码审查', icon: 'lucide:git-pull-request', iconColor: '#3367D6', count: 2, tasks: [
    { id: 'REQ-0005-A', title: '用户权限管理API', agent: '代码审查', priority: '高', progress: 90, reviewRoute: '/review/code' },
    { id: 'REQ-0004-A', title: '邮件通知模块', agent: '代码审查', priority: '中', progress: 70 },
  ]},
  { key: 'testing', label: '测试', icon: 'lucide:test-tubes', iconColor: '#3367D6', count: 1, tasks: [
    { id: 'REQ-0003-A', title: '支付流程单元测试', agent: '测试Agent', priority: '高', progress: 50, extra: '覆盖率 72%' },
  ]},
  { key: 'pending-review', label: '待审核', icon: 'lucide:shield-check', iconColor: '#F59D0D', bg: '#FEF6E8', count: 3, urgent: true, tasks: [
    { id: 'REQ-0002', title: '部署计划-支付模块v2', agent: 'DevOps', priority: '高', status: '需人工审核', statusColor: '#D93025', reviewRoute: '/review/deploy' },
    { id: 'REQ-0001-A', title: '架构方案-用户系统重构', agent: '', priority: '', status: '待审核', statusColor: '#F59D0D', reviewRoute: '/review/architecture' },
    { id: 'REQ-0005-B', title: '权限管理代码合并', agent: '', priority: '', status: '待审核', statusColor: '#F59D0D', reviewRoute: '/review/code' },
  ]},
  { key: 'completed', label: '已完成', icon: 'lucide:check-circle', iconColor: '#0F8B5D', count: 5, tasks: [
    { id: 'REQ-0010', title: '登录页面优化', priority: '', status: '已完成' },
    { id: 'REQ-0009', title: '数据导出CSV功能', priority: '', status: '已完成' },
    { id: 'REQ-0003', title: '支付流程重构', priority: '', status: '已完成' },
  ]},
]

export const notifications = [
  { id: 1, title: '架构方案REQ-0008待审核', type: '审核通知', typeColor: '#F59D0D', typeBg: '#FEF6E8', icon: 'lucide:shield-check', borderColor: '#F59D0D', iconBg: '#FEF6E8', description: 'AI架构师Agent已完成多语言支持模块架构设计，需要您进行人工审核后方可进入开发阶段。', time: '30分钟前', route: '/review/architecture' },
  { id: 2, title: '部署计划REQ-0002需人工审核', type: '审核通知', typeColor: '#D93025', typeBg: '#FDECEA', icon: 'lucide:rocket', borderColor: '#D93025', iconBg: '#FDECEA', description: '支付模块v2部署计划已生成，涉及生产环境变更，需要人工审核确认后方可执行部署。', time: '2小时前', route: '/review/deploy' },
  { id: 3, title: '权限管理代码合并审核', type: '审核通知', typeColor: '#F59D0D', typeBg: '#FEF6E8', icon: 'lucide:git-merge', borderColor: '#F59D0D', iconBg: '#FEF6E8', description: 'REQ-0005权限管理代码已完成合并请求，需要人工审核代码变更后方可合入主干分支。', time: '4小时前', route: '/review/code' },
  { id: 4, title: 'REQ-0007任务已完成开发', type: '进度通知', typeColor: '#0F8B5D', typeBg: '#E8F7F0', icon: 'lucide:check-circle', borderColor: '#0F8B5D', iconBg: '#E8F7F0', description: '短信验证码API及前端表单开发任务已全部完成，已自动流转至测试阶段。', time: '', route: '/workflow/task/1' },
  { id: 5, title: '系统更新: AI架构师Agent v2.3已上线', type: '系统通知', typeColor: '#3367D6', typeBg: '#E8F0FE', icon: 'lucide:info', borderColor: '#3367D6', iconBg: '#E8F0FE', description: 'AI架构师Agent已升级至v2.3版本，新增微服务拆分建议能力和架构安全审查功能。', time: '', route: '/settings/ai-agent' },
  { id: 6, title: '审核超时提醒: REQ-0001架构审核已超过24小时', type: '超时提醒', typeColor: '#D93025', typeBg: '#FDECEA', icon: 'lucide:alert-triangle', borderColor: '#D93025', iconBg: '#FDECEA', description: 'REQ-0001用户系统重构架构方案审核已超过24小时未处理，相关开发任务处于阻塞状态，请尽快完成审核。', time: '', route: '/review/architecture' },
]

export const codeReviewFiles = [
  { name: 'verify.controller.ts', status: 'added', icon: 'lucide:file-plus', iconColor: '#0F8B5D', lines: '+85' },
  { name: 'verify.service.ts', status: 'modified', icon: 'lucide:file-pen', iconColor: '#F59D0D', lines: '+32 -8' },
  { name: 'redis.store.ts', status: 'added-issue', icon: 'lucide:file-plus', iconColor: '#0F8B5D', lines: '+45', hasIssue: true },
  { name: 'app.module.ts', status: 'modified', icon: 'lucide:file-pen', iconColor: '#F59D0D', lines: '+5 -2' },
  { name: 'old.verify.ts', status: 'deleted', icon: 'lucide:file-minus', iconColor: '#D93025', lines: '-60' },
]

export const codeReviewIssues = [
  { severity: 'severe', label: '严重', color: '#D93025', title: '验证码明文存储在Redis中', file: 'redis.store.ts:5', description: '验证码直接以明文存储在Redis中,存在安全风险。建议使用AES加密后再存储。' },
  { severity: 'warning', label: '警告', color: '#F59D0D', title: '验证码比对未使用加密比较', file: 'redis.store.ts:9', description: '直接比较验证码可能受到时序攻击。建议使用恒定时间比较函数。' },
  { severity: 'suggestion', label: '建议', color: '#3367D6', title: '添加验证码发送频率限制', file: 'verify.service.ts', description: '' },
]

export const deploySteps = [
  { label: '停止服务', icon: 'lucide:check', estimate: '2min', status: 'completed' },
  { label: '数据库迁移', icon: 'lucide:check', estimate: '3min', status: 'completed' },
  { label: '代码部署', icon: 'lucide:package', estimate: '5min', status: 'in-progress', progress: 40 },
  { label: '服务启动', icon: 'lucide:server', estimate: '3min', status: 'pending' },
  { label: '健康检查', icon: 'lucide:heart-pulse', estimate: '2min', status: 'pending' },
  { label: '流量切换', icon: 'lucide:repeat', estimate: '2min', status: 'pending' },
]

export const deployLogLines = [
  { time: '08:00:01', level: 'INFO', message: '部署流水线启动 - REQ-0002' },
  { time: '08:00:03', level: 'INFO', message: 'Step 1: 停止服务开始...' },
  { time: '08:00:15', level: 'INFO', message: '停止 auth-service 实例 (2/3)' },
  { time: '08:00:28', level: 'INFO', message: '停止 verify-service 实例 (1/1)' },
  { time: '08:02:01', level: 'INFO', message: 'Step 1 完成 - 停止服务 ✓ (耗时 2min)', highlight: true },
  { time: '08:02:03', level: 'INFO', message: 'Step 2: 数据库迁移开始...' },
  { time: '08:02:05', level: 'INFO', message: '执行迁移脚本: V20210606__create_verify_codes.sql' },
  { time: '08:02:12', level: 'INFO', message: 'CREATE TABLE verify_codes ... 成功' },
  { time: '08:02:30', level: 'INFO', message: '数据备份完成: auth数据库全量备份' },
  { time: '08:05:01', level: 'INFO', message: 'Step 2 完成 - 数据库迁移 ✓ (耗时 3min)', highlight: true },
  { time: '08:05:03', level: 'INFO', message: 'Step 3: 代码部署开始...' },
  { time: '08:05:10', level: 'INFO', message: '拉取镜像 verify-service:v2.1.0 ...' },
  { time: '08:05:25', level: 'INFO', message: '拉取镜像 auth-service:v2.1.0 ...' },
  { time: '08:06:40', level: 'DEPLOY', message: '部署 verify-service 容器 (1/3) ...' },
  { time: '08:07:55', level: 'DEPLOY', message: '部署 auth-service 容器 (2/3) ...' },
  { time: '08:08:32', level: 'DEPLOY', message: '部署 api-gateway 配置更新 ...' },
  { time: '08:08:32', level: 'WARN', message: '等待容器健康检查就绪...' },
  { time: '08:08:33', level: 'DEPLOY', message: '进行中 - 代码部署 (预计 5min)', isCurrent: true },
]

export const teamMembers = [
  { name: 'Alex Chen', role: '技术负责人', avatar: 'alex', status: 'online', permission: '管理员', activity: 95 },
  { name: 'Sarah Wang', role: '产品经理', avatar: 'sarah', status: 'online', permission: '成员', activity: 80, boundTeam: '需求分析师' },
  { name: 'Mike Liu', role: '前端开发', avatar: 'mike', status: 'offline', permission: '成员', activity: 60, boundTeam: '前端开发' },
  { name: 'Emma Zhang', role: '后端开发', avatar: 'emma', status: 'online', permission: '成员', activity: 85, boundTeam: '后端开发' },
  { name: 'David Li', role: 'DevOps工程师', avatar: 'david', status: 'offline', permission: '成员', activity: 45, boundTeam: 'DevOps' },
]

export const alexTeams = ['架构师', '测试']
export const sarahTeams = ['需求分析师']
export const mikeTeams = ['前端开发']
export const emmaTeams = ['后端开发']
export const davidTeams = ['DevOps']

export const aiAgentTeams = [
  { key: 'requirement-analyst', name: '需求分析师 Team', icon: 'lucide:file-text', status: '运行中', boundMember: 'Sarah Wang', boundAvatar: 'sarah', consensus: '2/3', consensusPercent: 67, expanded: true },
  { key: 'architect', name: '架构师 Team', icon: 'lucide:layout', status: '运行中', boundMember: 'Alex Chen', boundAvatar: 'alex', consensus: '2/2', consensusPercent: 100, expanded: false },
  { key: 'frontend-dev', name: '前端开发 Team', icon: 'lucide:monitor', status: '运行中', boundMember: 'Mike Liu', boundAvatar: 'mike', consensus: '1/2', consensusPercent: 50, expanded: false },
  { key: 'backend-dev', name: '后端开发 Team', icon: 'lucide:database', status: '运行中', boundMember: 'Emma Zhang', boundAvatar: 'emma', consensus: '2/2', consensusPercent: 100, expanded: false },
  { key: 'testing', name: '测试 Team', icon: 'lucide:test-tubes', status: '待命', boundMember: 'Alex Chen', boundAvatar: 'alex', consensus: '0/2', consensusPercent: 0, expanded: false },
  { key: 'devops', name: 'DevOps Team', icon: 'lucide:server', status: '运行中', boundMember: 'David Li', boundAvatar: 'david', consensus: '2/2', consensusPercent: 100, expanded: false },
]

export const connectedServices = [
  { name: 'GitHub', icon: 'lucide:github', description: '代码仓库与CI/CD集成', connected: true },
  { name: 'Slack', icon: 'lucide:message-circle', description: '团队协作与通知推送', connected: true },
  { name: 'AWS', icon: 'lucide:cloud', description: '云服务部署与资源管理', connected: false },
]

export const apiKeys = [
  { name: 'GitHub API Token', masked: 'ghp_9xK4**********************', created: '2025-06-15', permissions: 'repo, workflow' },
  { name: 'Slack Bot Token', masked: 'xoxb-2847**********************', created: '2025-08-20', permissions: 'channels:read, chat:write' },
]

export const reviewRules = [
  { name: '架构变更审核', description: '当架构方案涉及核心模块变更时触发审核', trigger: '核心模块变更', handler: '技术负责人审核', enabled: true },
  { name: '代码合并审核', description: '当合并请求涉及安全敏感文件时触发审核', trigger: '安全敏感文件', handler: '双人审核', enabled: true },
  { name: '生产部署审核', description: '所有生产环境部署必须经过审核', trigger: '生产环境部署', handler: '运维负责人审核', enabled: true },
  { name: '依赖变更审核', description: '当第三方依赖版本变更时触发安全审核', trigger: '依赖版本变更', handler: '自动化安全扫描', enabled: false },
]

export const auditLogEntries = [
  { time: '06-06 14:30', operator: 'Alex Chen', action: '修改配置', target: '自动化级别', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-06 12:15', operator: 'Sarah Wang', action: '创建需求', target: 'REQ-2026-042', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-06 10:00', operator: 'AI Agent(架构师)', action: '生成方案', target: 'REQ-2026-040', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-05 23:45', operator: 'David Li', action: '部署操作', target: '生产环境 v2.3.1', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-05 18:30', operator: 'Alex Chen', action: '审核通过', target: '架构方案 #38', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-05 16:20', operator: 'Emma Zhang', action: '代码合并', target: 'PR #127', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-05 14:10', operator: 'AI Agent(后端)', action: '代码生成', target: 'API模块', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-05 11:00', operator: 'Alex Chen', action: '添加密钥', target: 'AWS凭证', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-04 20:15', operator: 'Sarah Wang', action: '更新需求', target: 'REQ-2026-039', result: '成功', resultIcon: 'lucide:check-circle', resultColor: '#0D7C7C' },
  { time: '06-04 09:30', operator: 'AI Agent(测试)', action: '执行失败', target: '集成测试套件', result: '失败', resultIcon: 'lucide:x-circle', resultColor: '#E5484D' },
]

export const rbacPermissions = [
  { name: '查看需求', admin: true, member: true, observer: true },
  { name: '创建需求', admin: true, member: true, observer: false },
  { name: '审核架构方案', admin: true, member: false, observer: false },
  { name: '代码合并', admin: true, member: true, observer: false },
  { name: '生产部署', admin: true, member: false, observer: false },
  { name: '系统配置', admin: true, member: false, observer: false },
  { name: '成员管理', admin: true, member: false, observer: false },
]

export const deployChangeSummary = {
  newFeatures: ['短信验证码发送与验证', '手机号格式校验', '验证码频率限制'],
  bugFixes: ['登录超时重定向问题', '验证码过期未清理'],
  improvements: ['Redis连接池优化', 'API响应缓存策略'],
  affectedServices: ['auth-service (核心)', 'verify-service (新增)', 'api-gateway'],
}

export const deployEnvVars = [
  { key: 'SMS_API_KEY', value: '***脱敏显示***', status: '新增', statusBg: '#E8F7F0', statusText: '#0F8B5D' },
  { key: 'SMS_PROVIDER_URL', value: 'https://sms.service.com/api', status: '新增', statusBg: '#E8F7F0', statusText: '#0F8B5D' },
  { key: 'REDIS_TTL', value: '300 → 600', status: '修改', statusBg: '#FEF6E8', statusText: '#F59D0D' },
]

export const deployServiceConfig = [
  { label: '容器镜像', value: 'verify-service:v2.1.0' },
  { label: '资源配置', value: 'CPU: 0.5核 / 内存: 512MB' },
  { label: '副本数量', value: '2 → 3' },
  { label: '端口映射', value: '8080:3000' },
]

export const deployDependencies = [
  { name: 'Redis', value: '6.2 → 7.0', status: '升级', statusBg: '#FEF6E8', statusText: '#F59D0D', icon: 'lucide:package' },
  { name: '短信服务API', value: 'v1.3 (新增依赖)', status: '新增', statusBg: '#E8F7F0', statusText: '#0F8B5D', icon: 'lucide:link' },
]

export const deployHistory = [
  { version: 'v2.0.9', status: '成功', time: '2026-06-01 02:15', duration: '15min' },
  { version: 'v2.0.8', status: '成功', time: '2026-05-28 02:10', duration: '14min' },
  { version: 'v2.0.7', status: '成功', time: '2026-05-25 02:08', duration: '16min' },
]

export const deployMonitoringSteps = [
  { label: '停止服务', icon: 'lucide:check', estimate: '2min', status: 'completed' },
  { label: '数据库迁移', icon: 'lucide:check', estimate: '3min', status: 'completed' },
  { label: '代码部署', icon: 'lucide:package', estimate: '5min', status: 'in-progress', progress: 40 },
  { label: '服务启动', icon: 'lucide:server', estimate: '3min', status: 'pending' },
  { label: '健康检查', icon: 'lucide:heart-pulse', estimate: '2min', status: 'pending' },
  { label: '流量切换', icon: 'lucide:repeat', estimate: '2min', status: 'pending' },
]

export const healthMetrics = [
  { label: 'API响应时间', value: '128ms', percent: 25, color: '#0F8B5D', threshold: '<500ms' },
  { label: '错误率', value: '0.2%', percent: 4, color: '#0F8B5D', threshold: '<5%' },
  { label: 'CPU使用率', value: '62%', percent: 62, color: '#F59D0D', threshold: '部署期间CPU略高 (正常)' },
  { label: '内存使用率', value: '54%', percent: 54, color: '#0F8B5D', threshold: '<80%' },
]