<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 mb-4 text-sm">
        <router-link to="/workflow/dashboard" class="text-[#0D7C7C] hover:underline">工作流看板</router-link>
        <span class="text-[#9BA3AB]">/</span>
        <span class="text-[#1A1F24] font-semibold">部署审核</span>
        <span class="text-[#9BA3AB]">/</span>
        <span class="text-[#1A1F24] font-semibold">REQ-0002</span>
      </div>

      <!-- Deploy Overview Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-sm font-semibold text-[#0D7C7C]">REQ-0002</span>
            <span class="text-lg font-semibold text-[#1A1F24]">部署计划 - 支付模块 v2.1.0</span>
          </div>
          <router-link to="/requirements/REQ-0007" class="text-sm text-[#0D7C7C] flex items-center gap-2 hover:underline">
            关联需求: REQ-0007
            <Icon icon="lucide:external-link" class="text-xs" />
          </router-link>
        </div>
        <p class="text-sm text-[#6B7680]">将支付模块从v1.8.0升级到v2.1.0，包含多渠道支付支持、支付网关切换、退款流程优化等核心功能变更。</p>
        <div class="flex items-center gap-6 text-sm text-[#6B7680]">
          <span class="flex items-center gap-2">
            <Icon icon="lucide:calendar" class="text-sm text-[#6B7680]" />
            计划时间: 2026-06-06 02:00 (凌晨低峰)
          </span>
          <span class="flex items-center gap-2">
            <Icon icon="lucide:clock" class="text-sm text-[#6B7680]" />
            预计时长: 45分钟
          </span>
          <span class="flex items-center gap-2">
            <Icon icon="lucide:server" class="text-sm text-[#6B7680]" />
            目标环境: 生产环境
          </span>
        </div>
      </div>

      <!-- Risk Assessment -->
      <div class="bg-[#FEF6E8] border border-[#F59D0D]/30 shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center gap-3">
          <Icon icon="lucide:alert-triangle" class="text-lg text-[#F59D0D]" />
          <span class="text-base font-semibold text-[#F59D0D]">需人工审核 — 涉及生产环境变更</span>
        </div>
        <div class="grid grid-cols-3 gap-4">
          <div class="bg-white rounded-xl p-4">
            <span class="text-xs font-semibold text-[#D93025] block mb-1">高风险项</span>
            <div class="flex flex-col gap-1 text-xs text-[#4A5259]">
              <span>• 数据库迁移脚本需验证</span>
              <span>• 支付网关切换影响在线交易</span>
            </div>
          </div>
          <div class="bg-white rounded-xl p-4">
            <span class="text-xs font-semibold text-[#F59D0D] block mb-1">回滚方案</span>
            <div class="flex flex-col gap-1 text-xs text-[#4A5259]">
              <span>• 数据库备份: deploy-backup-20260606.sql</span>
              <span>• 回滚脚本: rollback-v2.1.0.sh</span>
            </div>
          </div>
          <div class="bg-white rounded-xl p-4">
            <span class="text-xs font-semibold text-[#0F8B5D] block mb-1">前置条件</span>
            <div class="flex flex-col gap-1 text-xs text-[#4A5259]">
              <span class="flex items-center gap-1"><Icon icon="lucide:check" class="text-xs text-[#0F8B5D]" />代码审核已通过</span>
              <span class="flex items-center gap-1"><Icon icon="lucide:check" class="text-xs text-[#0F8B5D]" />集成测试100%通过</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Deployment Steps -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24]">部署步骤详情</h3>
        <div v-for="(step, idx) in steps" :key="idx" class="flex items-start gap-3">
          <div :class="['w-7 h-7 rounded-full flex items-center justify-center text-xs font-semibold shrink-0', step.done ? 'bg-[#E8F7F0] text-[#0F8B5D]' : 'bg-[#E8F0FE] text-[#3367D6]']">
            {{ idx + 1 }}
          </div>
          <div class="flex-1">
            <div class="flex items-center gap-2">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ step.title }}</span>
              <span :class="['px-2 py-0.5 rounded-full text-xs font-semibold', step.done ? 'bg-[#E8F7F0] text-[#0F8B5D]' : 'bg-[#FEF6E8] text-[#F59D0D]']">{{ step.status }}</span>
            </div>
            <p class="text-xs text-[#6B7680] mt-0.5">{{ step.desc }}</p>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center gap-3">
        <button class="flex items-center gap-2 px-6 py-2.5 bg-[#0F8B5D] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:check-circle" class="text-base" />
          <span class="whitespace-nowrap">批准部署</span>
        </button>
        <button class="flex items-center gap-2 px-6 py-2.5 bg-[#F59D0D] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:clock" class="text-base" />
          <span class="whitespace-nowrap">延期部署</span>
        </button>
        <button class="flex items-center gap-2 px-6 py-2.5 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:x-circle" class="text-base" />
          <span class="whitespace-nowrap">拒绝部署</span>
        </button>
      </div>
    </main>

    <!-- Right Sidebar -->
    <aside class="flex-shrink-0 min-w-fit w-[300px] bg-[hsla(200,15%,97.5%,1)] border-l border-[#E1E6EA] p-6 flex flex-col gap-6">
      <!-- AI Deploy Report -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:bot" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">AI部署风险评估</span>
        </div>
        <div class="flex items-center justify-between">
          <span class="text-sm text-[#4A5259]">风险评分</span>
          <span class="text-sm font-semibold text-[#F59D0D]">68/100</span>
        </div>
        <div class="flex flex-col gap-2 text-sm">
          <div class="flex items-center justify-between">
            <span class="text-[#4A5259]">影响范围</span>
            <span class="px-2 py-0.5 text-xs font-semibold bg-[#FDECEA] text-[#D93025] rounded-full">大</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-[#4A5259]">回滚难度</span>
            <span class="px-2 py-0.5 text-xs font-semibold bg-[#FEF6E8] text-[#F59D0D] rounded-full">中</span>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-[#4A5259]">历史成功率</span>
            <span class="px-2 py-0.5 text-xs font-semibold bg-[#E8F7F0] text-[#0F8B5D] rounded-full">92%</span>
          </div>
        </div>
      </div>

      <!-- Recent Deploy History -->
      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3">历史部署记录</h3>
        <div class="flex flex-col gap-2">
          <div class="bg-white rounded-lg p-3">
            <div class="flex items-center gap-2 mb-1">
              <span class="text-sm font-semibold text-[#0F8B5D]">v2.0.0</span>
              <span class="px-2 py-0.5 text-xs font-semibold bg-[#E8F7F0] text-[#0F8B5D] rounded-full">成功</span>
            </div>
            <span class="text-xs text-[#9BA3AB]">2026-05-20 · 耗时 38分钟</span>
          </div>
          <div class="bg-white rounded-lg p-3">
            <div class="flex items-center gap-2 mb-1">
              <span class="text-sm font-semibold text-[#D93025]">v1.9.2-hotfix</span>
              <span class="px-2 py-0.5 text-xs font-semibold bg-[#E8F7F0] text-[#0F8B5D] rounded-full">成功</span>
            </div>
            <span class="text-xs text-[#9BA3AB]">2026-05-10 · 耗时 12分钟</span>
          </div>
          <div class="bg-white rounded-lg p-3">
            <div class="flex items-center gap-2 mb-1">
              <span class="text-sm font-semibold text-[#4A5259]">v1.9.1</span>
              <span class="px-2 py-0.5 text-xs font-semibold bg-[#FDECEA] text-[#D93025] rounded-full">失败</span>
            </div>
            <span class="text-xs text-[#9BA3AB]">2026-05-08 · 已回滚</span>
          </div>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'

const steps = [
  { title: '数据库备份', desc: '创建生产数据库快照备份', status: '已完成', done: true },
  { title: '数据库迁移', desc: '执行迁移脚本更新表结构,新增 payment_channels 字段', status: '待执行', done: false },
  { title: '停止旧服务', desc: '优雅停止 v1.8.0 实例,等待现有请求处理完毕', status: '待执行', done: false },
  { title: '部署新服务', desc: '启动 v2.1.0 容器实例,运行健康检查', status: '待执行', done: false },
  { title: '切换流量', desc: '将负载均衡流量指向新服务实例,观察10分钟稳定性', status: '待执行', done: false },
  { title: '验证监控', desc: '确认生产监控指标正常,无异常错误日志', status: '待执行', done: false },
]
</script>