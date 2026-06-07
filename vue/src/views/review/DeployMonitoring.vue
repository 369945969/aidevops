<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 mb-4 text-sm">
        <router-link to="/workflow/dashboard" class="text-[#0D7C7C] hover:underline">工作流看板</router-link>
        <span class="text-[#9BA3AB]">/</span>
        <router-link to="/review/deploy" class="text-[#0D7C7C] hover:underline">部署审核</router-link>
        <span class="text-[#9BA3AB]">/</span>
        <span class="text-[#1A1F24] font-semibold">REQ-0002</span>
      </div>

      <!-- Deploy Status Banner -->
      <div class="bg-[#E8F7F0] border border-[#0F8B5D]/30 shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex items-center gap-4 px-6 py-4 mb-6">
        <div class="w-3 h-3 rounded-full bg-[#0F8B5D] animate-pulse"></div>
        <div class="flex items-center gap-3 flex-1">
          <span class="text-lg font-semibold text-[#0F8B5D]">部署进行中</span>
          <span class="text-sm text-[#4A5259]">REQ-0002 - 支付模块 v2.1.0 部署至生产环境</span>
        </div>
        <div class="flex items-center gap-2 text-sm text-[#6B7680]">
          <Icon icon="lucide:clock" class="text-sm text-[#0F8B5D]" />
          <span>已运行 8:32</span>
        </div>
      </div>

      <!-- Real-time Metrics -->
      <div class="grid grid-cols-4 gap-4 mb-6">
        <div v-for="metric in metrics" :key="metric.label" class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl p-5">
          <div class="flex items-center gap-2 mb-2">
            <Icon :icon="metric.icon" :class="['text-base', metric.iconColor]" />
            <span class="text-sm text-[#6B7680]">{{ metric.label }}</span>
          </div>
          <span :class="['text-2xl font-semibold', metric.valueColor]">{{ metric.value }}</span>
          <p class="text-xs text-[#9BA3AB] mt-1">{{ metric.sub }}</p>
        </div>
      </div>

      <!-- Deploy Logs -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <h3 class="text-base font-semibold text-[#1A1F24]">实时部署日志</h3>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:download" class="text-sm" />
              <span>下载日志</span>
            </button>
            <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:refresh-cw" class="text-sm" />
              <span>刷新</span>
            </button>
          </div>
        </div>
        <div class="bg-[hsla(200,15%,95%,1)] rounded-xl p-4 font-mono text-xs space-y-1.5 overflow-y-auto" style="max-height:300px">
          <div v-for="log in logs" :key="log.time" class="flex gap-2">
            <span class="text-[#9BA3AB] shrink-0">{{ log.time }}</span>
            <span :class="log.color">{{ log.msg }}</span>
          </div>
        </div>
      </div>

      <!-- Service Health -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <h3 class="text-base font-semibold text-[#1A1F24]">服务健康状态</h3>
        <div class="grid grid-cols-2 gap-4">
          <div v-for="svc in services" :key="svc.name" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
            <div class="flex items-center justify-between mb-2">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ svc.name }}</span>
              <div :class="['w-2 h-2 rounded-full', svc.healthy ? 'bg-[#0F8B5D]' : 'bg-[#D93025]']"></div>
            </div>
            <div class="flex items-center justify-between text-xs text-[#6B7680]">
              <span>响应时间: {{ svc.responseTime }}</span>
              <span>错误率: {{ svc.errorRate }}</span>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Right Sidebar -->
    <aside class="flex-shrink-0 min-w-fit w-[300px] bg-[hsla(200,15%,97.5%,1)] border-l border-[#E1E6EA] p-6 flex flex-col gap-6">
      <!-- Deploy Progress -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:activity" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">部署进度</span>
        </div>
        <div class="flex flex-col gap-3">
          <div v-for="(step, idx) in progressSteps" :key="idx">
            <div class="flex items-center justify-between text-xs mb-1">
              <span :class="step.done ? 'text-[#0F8B5D]' : 'text-[#4A5259]'">{{ step.name }}</span>
              <span class="text-[#9BA3AB]">{{ step.pct }}%</span>
            </div>
            <div class="h-1.5 bg-[hsla(200,15%,95%,1)] rounded-full overflow-hidden">
              <div :class="['h-full rounded-full transition-all', step.done ? 'bg-[#0F8B5D]' : 'bg-[#F59D0D]']" :style="{ width: step.pct + '%' }"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Alerts -->
      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3">告警信息</h3>
        <div class="flex flex-col gap-2">
          <div class="bg-[#FEF6E8] border border-[#F59D0D]/30 rounded-lg p-3">
            <div class="flex items-center gap-2 mb-1">
              <Icon icon="lucide:alert-triangle" class="text-sm text-[#F59D0D]" />
              <span class="text-xs font-semibold text-[#F59D0D]">警告</span>
            </div>
            <p class="text-xs text-[#4A5259]">CPU使用率短暂上升至78%,已自动恢复</p>
            <span class="text-xs text-[#9BA3AB]">8:28 前</span>
          </div>
          <div class="bg-[#E8F7F0] border border-[#0F8B5D]/30 rounded-lg p-3">
            <div class="flex items-center gap-2 mb-1">
              <Icon icon="lucide:check-circle" class="text-sm text-[#0F8B5D]" />
              <span class="text-xs font-semibold text-[#0F8B5D]">正常</span>
            </div>
            <p class="text-xs text-[#4A5259]">所有健康检查通过,服务稳定运行</p>
            <span class="text-xs text-[#9BA3AB]">8:25 前</span>
          </div>
        </div>
      </div>

      <!-- Quick Actions -->
      <div class="flex flex-col gap-2">
        <button class="flex items-center gap-2 px-4 py-2.5 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80 text-sm justify-center">
          <Icon icon="lucide:x-circle" class="text-sm" />
          <span>紧急停止</span>
        </button>
        <button class="flex items-center gap-2 px-4 py-2.5 bg-[#F59D0D] text-white/95 rounded-full transition hover:opacity-80 text-sm justify-center">
          <Icon icon="lucide:rotate-ccw" class="text-sm" />
          <span>回滚部署</span>
        </button>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'

const metrics = [
  { label: 'CPU使用率', icon: 'lucide:cpu', iconColor: 'text-[#0D7C7C]', value: '42%', valueColor: 'text-[#0F8B5D]', sub: '正常范围' },
  { label: '内存占用', icon: 'lucide:hard-drive', iconColor: 'text-[#3367D6]', value: '2.8GB', valueColor: 'text-[#0F8B5D]', sub: '峰值 3.1GB' },
  { label: '请求延迟', icon: 'lucide:zap', iconColor: 'text-[#F59D0D]', value: '128ms', valueColor: 'text-[#F59D0D]', sub: 'P95延迟' },
  { label: '错误率', icon: 'lucide:alert-circle', iconColor: 'text-[#0F8B5D]', value: '0.02%', valueColor: 'text-[#0F8B5D]', sub: '低于阈值' },
]

const logs = [
  { time: '10:32:45', color: 'text-[#0F8B5D]', msg: '✓ 数据库迁移完成,新增 payment_channels 字段' },
  { time: '10:31:20', color: 'text-[#4A5259]', msg: '执行迁移脚本 migrate-v2.1.0.sql...' },
  { time: '10:30:55', color: 'text-[#0F8B5D]', msg: '✓ 数据库连接池建立,连接数: 20' },
  { time: '10:30:40', color: 'text-[#4A5259]', msg: '启动数据库迁移工具 v3.2.1...' },
  { time: '10:30:15', color: 'text-[#0F8B5D]', msg: '✓ 数据库备份完成,快照ID: snap-20260606-103015' },
  { time: '10:30:00', color: 'text-[#4A5259]', msg: '开始执行部署计划 REQ-0002-deploy-v2.1.0' },
  { time: '10:29:50', color: 'text-[#0D7C7C]', msg: '[INFO] 部署任务已启动,操作者: deploy-bot' },
]

const services = [
  { name: '支付网关', healthy: true, responseTime: '89ms', errorRate: '0.01%' },
  { name: '用户认证服务', healthy: true, responseTime: '45ms', errorRate: '0.00%' },
  { name: '订单处理服务', healthy: true, responseTime: '156ms', errorRate: '0.03%' },
  { name: '通知服务', healthy: true, responseTime: '210ms', errorRate: '0.05%' },
]

const progressSteps = [
  { name: '数据库备份', pct: 100, done: true },
  { name: '数据库迁移', pct: 100, done: true },
  { name: '服务部署', pct: 75, done: false },
  { name: '健康检查', pct: 0, done: false },
  { name: '流量切换', pct: 0, done: false },
]
</script>