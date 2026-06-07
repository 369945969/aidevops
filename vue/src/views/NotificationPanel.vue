<template>
  <TopNavBar />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-12 py-6">
      <div class="flex items-center justify-between mb-5">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">通知中心</h1>
        <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full transition hover:opacity-80 text-sm">
          <Icon icon="lucide:check-check" class="text-sm text-[#4A5259]" />
          <span class="whitespace-nowrap">标记全部已读</span>
        </button>
      </div>

      <!-- Filter Tabs -->
      <div class="flex gap-2 mb-6">
        <label v-for="filter in filters" :key="filter" class="cursor-pointer">
          <input type="radio" name="notif-filter" class="sr-only peer" :checked="filter === '待审核通知'">
          <div class="bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full peer-checked:bg-[#0D7C7C] peer-checked:text-white/95 hover:opacity-80 transition whitespace-nowrap text-sm">{{ filter }}</div>
        </label>
      </div>

      <!-- Notification List -->
      <div class="flex flex-col gap-4">
        <router-link v-for="notif in notifications" :key="notif.id" :to="notif.link"
          :class="['bg-white shadow-[0_1px_3px_rgba(0,0,0,0.08)] rounded-2xl flex items-center gap-5 p-5 border-l-4 transition hover:shadow-[0_2px_8px_rgba(0,0,0,0.12)] cursor-pointer group no-underline', notif.borderColor]">
          <div :class="['flex items-center justify-center w-10 h-10 rounded-xl shrink-0', notif.iconBg]">
            <Icon :icon="notif.icon" :class="['text-lg', notif.iconColor]" />
          </div>
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-3 mb-1">
              <span class="text-base font-semibold text-[#1A1F24]">{{ notif.title }}</span>
              <div :class="['px-3 py-1 rounded-full text-xs font-semibold shrink-0', notif.badgeBg, notif.badgeColor]">{{ notif.badge }}</div>
            </div>
            <p class="text-sm text-[#6B7680]">{{ notif.desc }}</p>
          </div>
          <div class="flex items-center gap-3 shrink-0">
            <span class="text-sm text-[#9BA3AB]">{{ notif.time }}</span>
            <Icon icon="lucide:chevron-right" class="w-5 h-5 text-[#9BA3AB] group-hover:text-[#0D7C7C] transition" />
          </div>
        </router-link>
      </div>
    </main>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../layout/TopNavBar.vue'

const filters = ['待审核通知', '系统通知', '全部']

const notifications = [
  {
    id: 1,
    title: '架构方案REQ-0008待审核',
    desc: 'AI架构师Agent已完成多语言支持模块架构设计，需要您进行人工审核后方可进入开发阶段。',
    link: '/review/architecture',
    icon: 'lucide:shield-check',
    iconBg: 'bg-[#FEF6E8]',
    iconColor: 'text-[#F59D0D]',
    badge: '审核通知',
    badgeBg: 'bg-[#FEF6E8]',
    badgeColor: 'text-[#F59D0D]',
    borderColor: 'border-l-[#F59D0D]',
    time: '30分钟前',
  },
  {
    id: 2,
    title: '部署计划REQ-0002需人工审核',
    desc: '支付模块v2部署计划已生成，涉及生产环境变更，需要人工审核确认后方可执行部署。',
    link: '/review/deploy',
    icon: 'lucide:rocket',
    iconBg: 'bg-[#FDECEA]',
    iconColor: 'text-[#D93025]',
    badge: '审核通知',
    badgeBg: 'bg-[#FDECEA]',
    badgeColor: 'text-[#D93025]',
    borderColor: 'border-l-[#D93025]',
    time: '2小时前',
  },
  {
    id: 3,
    title: '权限管理代码合并审核',
    desc: 'REQ-0005权限管理代码已完成合并请求，需要人工审核代码变更后方可合入主干分支。',
    link: '/review/code',
    icon: 'lucide:git-merge',
    iconBg: 'bg-[#FEF6E8]',
    iconColor: 'text-[#F59D0D]',
    badge: '审核通知',
    badgeBg: 'bg-[#FEF6E8]',
    badgeColor: 'text-[#F59D0D]',
    borderColor: 'border-l-[#F59D0D]',
    time: '3小时前',
  },
  {
    id: 4,
    title: '知识库新增条目: 短信服务架构模式',
    desc: '从REQ-0007需求中自动提炼沉淀知识条目，已通过AI质量评分，等待人工确认入库。',
    link: '/knowledge',
    icon: 'lucide:book-open',
    iconBg: 'bg-[#E8F0FE]',
    iconColor: 'text-[#3367D6]',
    badge: '系统通知',
    badgeBg: 'bg-[#E8F0FE]',
    badgeColor: 'text-[#3367D6]',
    borderColor: 'border-l-[#3367D6]',
    time: '5小时前',
  },
  {
    id: 5,
    title: '支付流程重构已完成',
    desc: 'REQ-0003支付流程重构任务已全部完成，包含5个子任务，所有代码已合并并通过测试。',
    link: '/requirements/REQ-0003',
    icon: 'lucide:check-circle',
    iconBg: 'bg-[#E8F7F0]',
    iconColor: 'text-[#0F8B5D]',
    badge: '完成通知',
    badgeBg: 'bg-[#E8F7F0]',
    badgeColor: 'text-[#0F8B5D]',
    borderColor: 'border-l-[#0F8B5D]',
    time: '1天前',
  },
]
</script>