<template>
  <TopNavBar activeNav="requirements" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 text-sm text-[#6B7680] mb-4">
        <router-link to="/requirements" class="text-[#0D7C7C]">需求管理</router-link>
        <span>/</span>
        <router-link to="/requirements/REQ-0007" class="text-[#0D7C7C]">需求详情</router-link>
        <span>/</span>
        <span class="text-[#1A1F24] font-semibold">REQ-0007</span>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-sm font-semibold text-[#0D7C7C]">REQ-0007</span>
            <span class="text-lg font-semibold text-[#1A1F24]">短信验证码功能</span>
            <StatusBadge label="开发中" />
            <PriorityBadge label="高" />
          </div>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 text-[#0D7C7C] rounded-full transition hover:opacity-80 text-sm border border-[#E1E6EA]">
              <Icon icon="lucide:pencil" class="text-sm" />
              <span class="whitespace-nowrap">编辑</span>
            </button>
          </div>
        </div>
        <div class="flex items-center gap-6 text-sm text-[#6B7680]">
          <span class="flex items-center gap-2"><Icon icon="lucide:user" class="text-sm" /> 创建者: Sarah Okonkwo</span>
          <span class="flex items-center gap-2"><Icon icon="lucide:calendar" class="text-sm" /> 创建时间: 2026-05-30 14:30</span>
          <span class="flex items-center gap-2"><Icon icon="lucide:clock" class="text-sm" /> 最后更新: 2小时前</span>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
          <Icon icon="lucide:file-text" class="text-sm text-[#0D7C7C]" /> 原始需求描述
        </h3>
        <p class="text-sm text-[#4A5259] leading-[1.8]">
          用户注册流程中需要增加手机号短信验证码验证功能，确保用户身份真实性，减少虚假账号注册。需要支持中国大陆手机号格式（11位数字），验证码有效期5分钟，每个手机号每天最多发送5次验证码。需要在注册、登录和修改密码场景中使用。
        </p>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
          <Icon icon="lucide:bot" class="text-sm text-[#0D7C7C]" /> AI生成的需求文档
        </h3>
        <div><h4 class="text-sm font-semibold text-[#1A1F24] mb-2">业务目标</h4><p class="text-sm text-[#4A5259]">通过手机号短信验证码增强用户身份验证，降低虚假账号注册率，提升平台用户数据真实性和安全性。</p></div>
        <div><h4 class="text-sm font-semibold text-[#1A1F24] mb-2">功能范围</h4>
          <div class="flex flex-col gap-2">
            <div v-for="scope in scopes" :key="scope" class="flex items-center gap-2 bg-[hsla(200,15%,95%,1)] rounded-lg px-3 py-2 text-sm text-[#4A5259]">
              <Icon icon="lucide:plus-circle" class="text-sm text-[#0F8B5D]" /> <span>{{ scope }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
          <Icon icon="lucide:blocks" class="text-sm text-[#0D7C7C]" /> 技术架构方案
        </h3>
        <div class="flex flex-col gap-3">
          <div class="text-sm text-[#4A5259] flex items-start gap-2">
            <Icon icon="lucide:wrench" class="text-xs text-[#0D7C7C]" />
            <span><span class="font-semibold text-[#1A1F24]">技术选型:</span> 第三方短信服务API + Redis缓存 + Node.js后端</span>
          </div>
          <div class="text-sm text-[#4A5259] flex items-start gap-2">
            <Icon icon="lucide:boxes" class="text-xs text-[#0D7C7C]" />
            <span><span class="font-semibold text-[#1A1F24]">模块划分:</span> 短信发送模块、验证码存储模块、校验逻辑模块、频率限制模块</span>
          </div>
        </div>
        <router-link to="/review/architecture" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:shield-check" class="text-sm" /> <span class="whitespace-nowrap">提交架构审核</span>
        </router-link>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
          <Icon icon="lucide:list-checks" class="text-sm text-[#0D7C7C]" /> 任务拆解列表
        </h3>
        <div class="flex flex-col gap-3">
          <div v-for="task in tasks" :key="task.id" class="flex items-center justify-between bg-[hsla(200,15%,95%,1)] rounded-lg px-4 py-3">
            <div class="flex items-center gap-3">
              <span class="text-sm font-semibold text-[#0D7C7C]">{{ task.id }}</span>
              <span class="text-sm text-[#1A1F24]">{{ task.title }}</span>
              <StatusBadge :label="task.status" />
            </div>
            <div class="flex items-center gap-3 text-sm text-[#6B7680]">
              <span>{{ task.type }}</span>
              <span>预估{{ task.estimate }}h</span>
              <Icon icon="lucide:bot" class="text-xs text-[#0D7C7C]" />
              <router-link to="/workflow/task/1" class="text-sm text-[#0D7C7C] hover:opacity-80">进入任务</router-link>
            </div>
          </div>
        </div>
      </div>
    </main>

    <aside class="flex-shrink-0 min-w-fit bg-[hsla(200,15%,97.5%,1)] px-6 py-6 flex flex-col gap-6 overflow-y-auto" style="width:320px">
      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3 flex items-center gap-2">
          <Icon icon="lucide:clock" class="text-sm text-[#0D7C7C]" /> 时间线
        </h3>
        <div class="flex flex-col gap-0">
          <div v-for="(item, idx) in timeline" :key="idx" class="flex items-start gap-3 pb-4">
            <div class="flex flex-col items-center">
              <div :class="['w-6 h-6 rounded-full flex items-center justify-center', item.bg]">
                <Icon :icon="item.icon" class="text-xs" :class="item.iconColor" />
              </div>
              <div v-if="idx < timeline.length - 1" class="w-0.5 h-full bg-[#E1E6EA] mt-1"></div>
            </div>
            <div class="flex flex-col gap-1">
              <span :class="['text-sm', item.fontWeight]">{{ item.label }}</span>
              <span class="text-xs text-[#9BA3AB]">{{ item.date }}</span>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3 flex items-center gap-2">
          <Icon icon="lucide:bar-chart-2" class="text-sm text-[#0D7C7C]" /> 统计信息
        </h3>
        <div class="flex flex-col gap-3">
          <div class="flex items-center justify-between text-sm">
            <span class="text-[#4A5259] flex items-center gap-2"><Icon icon="lucide:tasks" class="text-xs text-[#6B7680]" /> 关联任务</span>
            <span class="font-semibold text-[#1A1F24]">4个</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-[#4A5259] flex items-center gap-2"><Icon icon="lucide:check-circle" class="text-xs text-[#6B7680]" /> 已完成任务</span>
            <span class="font-semibold text-[#0F8B5D]">0个</span>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-[#4A5259] flex items-center gap-2"><Icon icon="lucide:clock" class="text-xs text-[#6B7680]" /> 预计总工时</span>
            <span class="font-semibold text-[#1A1F24]">11h</span>
          </div>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import StatusBadge from '../../components/StatusBadge.vue'
import PriorityBadge from '../../components/PriorityBadge.vue'

const scopes = ['短信验证码发送接口', '验证码校验接口', '手机号格式校验（中国大陆11位）', '发送频率限制（5次/天/号）']

const tasks = [
  { id: 'TASK-007-A', title: '短信验证码API接口', status: '开发中', type: '后端', estimate: 4 },
  { id: 'TASK-007-B', title: '手机号验证前端表单', status: '开发中', type: '前端', estimate: 3 },
  { id: 'TASK-007-C', title: 'Redis验证码存储逻辑', status: '开发中', type: '后端', estimate: 2 },
  { id: 'TASK-007-D', title: '验证码功能单元测试', status: '待开始', type: '测试', estimate: 2 },
]

const timeline = [
  { label: '需求创建', date: '2026-05-30 14:30', icon: 'lucide:file-plus', bg: 'bg-[#E8F7F0]', iconColor: 'text-[#0F8B5D]', fontWeight: 'font-semibold text-[#1A1F24]' },
  { label: 'AI分析完成', date: '2026-05-30 15:00', icon: 'lucide:bot', bg: 'bg-[#E8F7F0]', iconColor: 'text-[#0F8B5D]', fontWeight: 'font-semibold text-[#1A1F24]' },
  { label: '用户补充说明', date: '2026-05-30 16:00', icon: 'lucide:message-circle', bg: 'bg-[hsla(200,15%,95%,1)]', iconColor: 'text-[#6B7680]', fontWeight: 'text-[#4A5259]' },
  { label: '架构方案生成', date: '2026-05-31 10:00', icon: 'lucide:blocks', bg: 'bg-[#E8F0FE]', iconColor: 'text-[#3367D6]', fontWeight: 'font-semibold text-[#1A1F24]' },
  { label: '架构审核通过', date: '2026-05-31 14:00', icon: 'lucide:shield-check', bg: 'bg-[#E8F7F0]', iconColor: 'text-[#0F8B5D]', fontWeight: 'text-[#4A5259]' },
  { label: '任务分配,开发开始', date: '2026-06-01 09:00', icon: 'lucide:git-branch', bg: 'bg-[hsla(200,15%,95%,1)]', iconColor: 'text-[#6B7680]', fontWeight: 'text-[#4A5259]' },
]
</script>
