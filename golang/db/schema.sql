-- AI DevOps Orchestrator - Settings Schema
-- Database: SQLite
-- Generated from Paraflow settings page designs

PRAGMA foreign_keys = ON;

-- ============================================================
-- 1. Automation Configuration (settings.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS automation_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    level TEXT NOT NULL DEFAULT 'balanced' CHECK(level IN ('conservative', 'balanced', 'aggressive', 'full-auto')),
    tech_preference TEXT NOT NULL DEFAULT 'neutral' CHECK(tech_preference IN ('conservative', 'neutral', 'innovative')),
    code_style TEXT NOT NULL DEFAULT 'concise' CHECK(code_style IN ('concise', 'detailed-comments', 'performance-first', 'readability-first')),
    test_coverage INTEGER NOT NULL DEFAULT 80 CHECK(test_coverage BETWEEN 60 AND 100),
    analysis_timeout INTEGER NOT NULL DEFAULT 30,
    architecture_timeout INTEGER NOT NULL DEFAULT 60,
    development_timeout INTEGER NOT NULL DEFAULT 120,
    test_execution_timeout INTEGER NOT NULL DEFAULT 60,
    auto_retry_count INTEGER NOT NULL DEFAULT 2,
    retry_interval INTEGER NOT NULL DEFAULT 5,
    failure_handling TEXT NOT NULL DEFAULT 'notify_human' CHECK(failure_handling IN ('notify_human', 'pause_task', 'mark_failed_skip')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS review_nodes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    enabled INTEGER NOT NULL DEFAULT 1,
    sub_option TEXT NOT NULL DEFAULT '',
    sub_input TEXT NOT NULL DEFAULT '',
    sort_order INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 2. AI Agent Configuration (settings-ai-agent.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS ai_teams (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    icon TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'running' CHECK(status IN ('running', 'standby', 'paused')),
    enabled INTEGER NOT NULL DEFAULT 1,
    bound_member_id INTEGER REFERENCES team_members(id) ON DELETE SET NULL,
    consensus_pass_rate REAL NOT NULL DEFAULT 0,
    total_tasks INTEGER NOT NULL DEFAULT 0,
    avg_duration INTEGER NOT NULL DEFAULT 0,
    success_rate REAL NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ai_agents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    team_id INTEGER NOT NULL REFERENCES ai_teams(id) ON DELETE CASCADE,
    role TEXT NOT NULL,
    primary_model_id INTEGER REFERENCES ai_models(id) ON DELETE SET NULL,
    fallback_model_id INTEGER REFERENCES ai_models(id) ON DELETE SET NULL,
    tech_preference TEXT NOT NULL DEFAULT 'structured_analysis' CHECK(tech_preference IN ('structured_analysis', 'free_exploration')),
    behavior_param REAL NOT NULL DEFAULT 0.7,
    trust_level REAL NOT NULL DEFAULT 0.85,
    rigor_level REAL NOT NULL DEFAULT 0.9,
    total_tasks INTEGER NOT NULL DEFAULT 0,
    avg_duration INTEGER NOT NULL DEFAULT 0,
    success_rate REAL NOT NULL DEFAULT 0,
    consensus_status TEXT NOT NULL DEFAULT 'pending' CHECK(consensus_status IN ('confirmed', 'pending', 'rejected')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 3. Development Environment (settings-dev-environment.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS code_repos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    provider TEXT NOT NULL CHECK(provider IN ('github', 'gitlab')),
    connected INTEGER NOT NULL DEFAULT 0,
    repo_url TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cicd_services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    provider TEXT NOT NULL CHECK(provider IN ('jenkins', 'gitlab_ci')),
    configured INTEGER NOT NULL DEFAULT 0,
    service_url TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cloud_credentials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    provider TEXT NOT NULL CHECK(provider IN ('aws', 'alibaba_cloud')),
    access_key_id TEXT NOT NULL DEFAULT '',
    secret_access_key TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS ssh_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    key_type TEXT NOT NULL DEFAULT 'RSA 2048',
    description TEXT NOT NULL DEFAULT '',
    key_content TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used_at DATETIME,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS env_vars (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key TEXT NOT NULL UNIQUE,
    value TEXT NOT NULL DEFAULT '',
    is_secret INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 4. Integration Configuration (settings-integrations.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS connected_services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    icon TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('connected', 'pending', 'disconnected')),
    config_data TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS api_keys (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    service TEXT NOT NULL,
    key_value TEXT NOT NULL DEFAULT '',
    permissions TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS webhook_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL DEFAULT '',
    event_types TEXT NOT NULL DEFAULT '',
    signature_key TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS sync_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    last_sync_at DATETIME,
    next_sync_at DATETIME,
    frequency TEXT NOT NULL DEFAULT 'hourly' CHECK(frequency IN ('hourly', 'daily', 'weekly', 'manual')),
    status TEXT NOT NULL DEFAULT 'running' CHECK(status IN ('running', 'paused', 'error')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 5. Model Management (settings-model-management.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS ai_models (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    provider TEXT NOT NULL DEFAULT '',
    is_builtin INTEGER NOT NULL DEFAULT 0,
    api_endpoint TEXT NOT NULL DEFAULT '',
    api_key TEXT NOT NULL DEFAULT '',
    max_context_length INTEGER NOT NULL DEFAULT 8192,
    connection_status TEXT NOT NULL DEFAULT 'not_tested' CHECK(connection_status IN ('connected', 'not_tested', 'error')),
    capability_tags TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 6. Notification Settings (settings-notifications.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS notification_channels (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    channel_type TEXT NOT NULL UNIQUE CHECK(channel_type IN ('email', 'slack', 'in_app', 'sms')),
    enabled INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notification_events (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_type TEXT NOT NULL UNIQUE,
    icon TEXT NOT NULL DEFAULT '',
    enabled INTEGER NOT NULL DEFAULT 1,
    severity TEXT NOT NULL DEFAULT 'info' CHECK(severity IN ('info', 'warning', 'critical')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notification_preferences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    frequency TEXT NOT NULL DEFAULT 'instant' CHECK(frequency IN ('instant', 'hourly_digest', 'daily_digest')),
    quiet_hours_enabled INTEGER NOT NULL DEFAULT 0,
    quiet_hours_start TEXT NOT NULL DEFAULT '22:00',
    quiet_hours_end TEXT NOT NULL DEFAULT '08:00',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 7. Review Rules (settings-review-rules.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS review_rules (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    trigger_condition TEXT NOT NULL DEFAULT '',
    handling_method TEXT NOT NULL DEFAULT '',
    enabled INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS review_templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS review_timeout_configs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    enabled INTEGER NOT NULL DEFAULT 1,
    threshold_hours INTEGER NOT NULL DEFAULT 24,
    strategy TEXT NOT NULL DEFAULT 'auto_approve_mark' CHECK(strategy IN ('auto_approve_mark', 'notify_admin', 'reject')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 8. Security & Permissions (settings-security.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    display_name TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS permissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    module TEXT NOT NULL,
    action TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS role_permissions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    access_level TEXT NOT NULL DEFAULT 'granted' CHECK(access_level IN ('granted', 'partial', 'denied')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS security_policies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    policy_type TEXT NOT NULL UNIQUE,
    enabled INTEGER NOT NULL DEFAULT 0,
    description TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS password_policies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    min_length INTEGER NOT NULL DEFAULT 12,
    expiry_days INTEGER NOT NULL DEFAULT 90,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS audit_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    operator TEXT NOT NULL,
    action TEXT NOT NULL,
    target TEXT NOT NULL DEFAULT '',
    result TEXT NOT NULL DEFAULT 'success' CHECK(result IN ('success', 'failure')),
    detail TEXT NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================
-- 9. Team Members (settings-team-members.html)
-- ============================================================

CREATE TABLE IF NOT EXISTS team_members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    title TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    avatar TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'offline' CHECK(status IN ('online', 'offline')),
    role_id INTEGER NOT NULL REFERENCES roles(id) ON DELETE SET NULL,
    activity_rate REAL NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS pending_invites (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    role_id INTEGER REFERENCES roles(id) ON DELETE SET NULL,
    invite_token TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending', 'accepted', 'expired', 'cancelled')),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS member_ai_team_bindings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    member_id INTEGER NOT NULL REFERENCES team_members(id) ON DELETE CASCADE,
    team_id INTEGER NOT NULL REFERENCES ai_teams(id) ON DELETE CASCADE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(member_id, team_id)
);

-- ============================================================
-- Initial seed data
-- ============================================================

-- Roles
INSERT INTO roles (name, display_name) VALUES ('admin', '超级管理员');
INSERT INTO roles (name, display_name) VALUES ('tech_lead', '技术负责人');
INSERT INTO roles (name, display_name) VALUES ('developer', '开发者');
INSERT INTO roles (name, display_name) VALUES ('observer', '观察者');

-- Permissions
INSERT INTO permissions (module, action) VALUES ('system_config', 'view');
INSERT INTO permissions (module, action) VALUES ('system_config', 'edit');
INSERT INTO permissions (module, action) VALUES ('requirement', 'view');
INSERT INTO permissions (module, action) VALUES ('requirement', 'create');
INSERT INTO permissions (module, action) VALUES ('code_review', 'approve');
INSERT INTO permissions (module, action) VALUES ('deploy', 'execute');
INSERT INTO permissions (module, action) VALUES ('knowledge', 'manage');

-- Role-Permissions (admin=all granted)
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 1, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 2, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 3, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 4, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 5, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 6, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (1, 7, 'granted');

-- tech_lead: partial system_config, full requirement & code_review
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 1, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 2, 'partial');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 3, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 4, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 5, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 6, 'partial');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (2, 7, 'granted');

-- developer: view + create requirement, code_review partial
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 1, 'partial');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 2, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 3, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 4, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 5, 'partial');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 6, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (3, 7, 'partial');

-- observer: only view
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 1, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 2, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 3, 'granted');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 4, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 5, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 6, 'denied');
INSERT INTO role_permissions (role_id, permission_id, access_level) VALUES (4, 7, 'denied');

-- Security policies
INSERT INTO security_policies (policy_type, enabled, description) VALUES ('db_encryption', 1, '数据库加密 (AES-256)');
INSERT INTO security_policies (policy_type, enabled, description) VALUES ('api_encryption', 1, 'API通信加密 (TLS 1.3)');
INSERT INTO security_policies (policy_type, enabled, description) VALUES ('two_factor_auth', 1, '双因素认证 (TOTP)');
INSERT INTO security_policies (policy_type, enabled, description) VALUES ('ip_whitelist', 0, 'IP白名单');

-- Notification channels
INSERT INTO notification_channels (channel_type, enabled) VALUES ('email', 1);
INSERT INTO notification_channels (channel_type, enabled) VALUES ('slack', 1);
INSERT INTO notification_channels (channel_type, enabled) VALUES ('in_app', 1);
INSERT INTO notification_channels (channel_type, enabled) VALUES ('sms', 0);

-- Notification events
INSERT INTO notification_events (event_type, icon, enabled, severity) VALUES ('review_passed', 'check-circle', 1, 'info');
INSERT INTO notification_events (event_type, icon, enabled, severity) VALUES ('merge_request', 'git-merge', 1, 'info');
INSERT INTO notification_events (event_type, icon, enabled, severity) VALUES ('deploy_completed', 'rocket', 1, 'info');
INSERT INTO notification_events (event_type, icon, enabled, severity) VALUES ('review_timeout', 'clock', 0, 'warning');
INSERT INTO notification_events (event_type, icon, enabled, severity) VALUES ('agent_exception', 'alert-triangle', 1, 'critical');

-- Default notification preferences
INSERT INTO notification_preferences (frequency, quiet_hours_enabled, quiet_hours_start, quiet_hours_end) VALUES ('instant', 1, '22:00', '08:00');

-- Default password policy
INSERT INTO password_policies (min_length, expiry_days) VALUES (12, 90);

-- Default review timeout config
INSERT INTO review_timeout_configs (enabled, threshold_hours, strategy) VALUES (1, 24, 'auto_approve_mark');

-- Default automation config
INSERT INTO automation_configs (level, tech_preference, code_style, test_coverage, analysis_timeout, architecture_timeout, development_timeout, test_execution_timeout, auto_retry_count, retry_interval, failure_handling) VALUES ('balanced', 'neutral', 'concise', 80, 30, 60, 120, 60, 2, 5, 'notify_human');

-- Review nodes
INSERT INTO review_nodes (name, description, enabled, sub_option, sub_input, sort_order) VALUES ('requirement_confirm', '需求理解确认', 1, '', '', 1);
INSERT INTO review_nodes (name, description, enabled, sub_option, sub_input, sort_order) VALUES ('architecture_review', '架构方案审核', 1, '', '', 2);
INSERT INTO review_nodes (name, description, enabled, sub_option, sub_input, sort_order) VALUES ('code_merge_review', '代码合并审核', 1, 'core_modules_only', 'auth, payment, user', 3);
INSERT INTO review_nodes (name, description, enabled, sub_option, sub_input, sort_order) VALUES ('test_confirm', '测试结果确认', 0, '', '', 4);
INSERT INTO review_nodes (name, description, enabled, sub_option, sub_input, sort_order) VALUES ('deploy_review', '集群部署审核', 1, 'production_only', '', 5);

-- Built-in AI models
INSERT INTO ai_models (name, provider, is_builtin, api_endpoint, max_context_length, connection_status, capability_tags) VALUES ('Claude Opus', 'Anthropic', 1, '', 200000, 'connected', 'complex_reasoning,long_text,multi_step_reasoning');
INSERT INTO ai_models (name, provider, is_builtin, api_endpoint, max_context_length, connection_status, capability_tags) VALUES ('Claude Sonnet', 'Anthropic', 1, '', 128000, 'connected', 'balanced_performance,fast_output,code_generation');
INSERT INTO ai_models (name, provider, is_builtin, api_endpoint, max_context_length, connection_status, capability_tags) VALUES ('GPT-4o', 'OpenAI', 1, '', 128000, 'connected', 'multi_modal,fast_reasoning,general');
INSERT INTO ai_models (name, provider, is_builtin, api_endpoint, max_context_length, connection_status, capability_tags) VALUES ('Gemini Pro', 'Google', 1, '', 1000000, 'not_tested', 'multi_modal,long_context');
INSERT INTO ai_models (name, provider, is_builtin, api_endpoint, max_context_length, connection_status, capability_tags) VALUES ('DeepSeek', 'DeepSeek', 1, '', 64000, 'connected', 'code_generation,math_reasoning,low_cost');

-- Review rules
INSERT INTO review_rules (name, description, trigger_condition, handling_method, enabled) VALUES ('架构变更审核', '核心架构模块变更时触发审核', '核心模块变更', '技术负责人审核', 1);
INSERT INTO review_rules (name, description, trigger_condition, handling_method, enabled) VALUES ('代码合并审核', '安全敏感文件的合并请求', '安全敏感文件', '双人审核', 1);
INSERT INTO review_rules (name, description, trigger_condition, handling_method, enabled) VALUES ('生产部署审核', '生产环境部署操作', '生产环境部署', '运维负责人审核', 1);
INSERT INTO review_rules (name, description, trigger_condition, handling_method, enabled) VALUES ('依赖变更审核', '依赖版本变更', '依赖版本变更', '自动安全扫描', 0);

-- Review template
INSERT INTO review_templates (name, description) VALUES ('安全优先模板', '包含6个预定义审核规则，涵盖架构变更、代码合并、生产部署等，适用于高安全要求项目');