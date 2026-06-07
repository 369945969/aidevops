<template>
  <TopNavBar activeNav="knowledge" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 text-sm text-[#6B7680] mb-4">
        <router-link to="/knowledge" class="text-[#0D7C7C] hover:opacity-80">知识库</router-link>
        <span>/</span>
        <span class="text-[#0D7C7C]">知识详情</span>
        <span>/</span>
        <span class="text-[#1A1F24] font-semibold">短信服务架构设计模式</span>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <h1 class="text-2xl font-semibold text-[#1A1F24]">短信服务架构设计模式</h1>
            <span class="text-xs px-3 py-1 rounded-full bg-[#FEF6E8] text-[#F59D0D] font-semibold whitespace-nowrap">架构方案</span>
          </div>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 text-[#0D7C7C] rounded-full transition hover:opacity-80 text-sm border border-[#E1E6EA]">
              <Icon icon="lucide:pencil" class="text-sm" /> <span class="whitespace-nowrap">编辑</span>
            </button>
            <button class="flex items-center justify-center w-10 h-10 rounded-full transition hover:opacity-80 border border-[#E1E6EA]">
              <Icon icon="lucide:share-2" class="text-sm text-[#6B7680]" />
            </button>
          </div>
        </div>
        <div class="flex items-center gap-6 text-sm text-[#6B7680]">
          <span class="flex items-center gap-2"><Icon icon="lucide:link" class="text-sm" /> 来自 <router-link to="/requirements/REQ-0007" class="text-[#0D7C7C] font-semibold">REQ-0007</router-link></span>
          <span class="flex items-center gap-2"><Icon icon="lucide:calendar" class="text-sm" /> 创建于 2026-06-01</span>
          <span class="flex items-center gap-2"><Icon icon="lucide:quote" class="text-sm" /> 引用 28次</span>
          <span class="flex items-center gap-2 text-[#0F8B5D]"><Icon icon="lucide:star" class="text-sm" /> 质量 4.5/5.0</span>
        </div>
        <div class="flex items-center gap-2">
          <span class="text-xs bg-[hsla(200,15%,95%,1)] rounded-full px-2.5 py-1 text-[#4A5259] whitespace-nowrap">Redis</span>
          <span class="text-xs bg-[hsla(200,15%,95%,1)] rounded-full px-2.5 py-1 text-[#4A5259] whitespace-nowrap">Node.js</span>
          <span class="text-xs bg-[hsla(200,15%,95%,1)] rounded-full px-2.5 py-1 text-[#4A5259] whitespace-nowrap">短信API</span>
          <span class="text-xs bg-[hsla(200,15%,95%,1)] rounded-full px-2.5 py-1 text-[#4A5259] whitespace-nowrap">验证码</span>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:file-text" class="text-sm text-[#0D7C7C]" /> 概述</h3>
        <p class="text-sm text-[#4A5259] leading-[1.8]">基于Redis缓存的验证码存储架构,包含发送频率限制、过期管理、加密存储等核心设计要点,适用于短信验证码类功能开发。该架构方案从REQ-0007需求中提炼沉淀,经过AI架构师和人工审核双重验证后入库,已在28个相关需求场景中被引用复用,有效减少了重复设计工作,平均节省开发时间约40%。</p>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:blocks" class="text-sm text-[#0D7C7C]" /> 核心设计要点</h3>
        <div class="flex flex-col gap-3">
          <div v-for="point in keyPoints" :key="point.title" class="flex items-start gap-3 bg-[hsla(200,15%,95%,1)] rounded-lg px-4 py-3">
            <div class="w-8 h-8 rounded-lg flex items-center justify-center flex-shrink-0" :style="{ backgroundColor: point.bg }">
              <Icon :icon="point.icon" class="text-sm" :style="{ color: point.color }" />
            </div>
            <div class="flex flex-col gap-1">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ point.title }}</span>
              <span class="text-sm text-[#4A5259]">{{ point.description }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="flex items-center gap-3 mb-6">
        <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:star" class="text-sm" /> <span class="whitespace-nowrap text-sm">收藏</span>
        </button>
        <router-link to="/requirements/new" class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] text-[#0D7C7C] rounded-full transition hover:opacity-80">
          <Icon icon="lucide:quote" class="text-sm" /> <span class="whitespace-nowrap text-sm">引用到新需求</span>
        </router-link>
      </div>
    </main>

    <aside class="flex-shrink-0 min-w-fit bg-[hsla(200,15%,97.5%,1)] px-6 py-6 flex flex-col gap-6 overflow-y-auto" style="width:300px">
      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3 flex items-center gap-2"><Icon icon="lucide:star" class="text-sm text-[#0D7C7C]" /> 质量评分详情</h3>
        <div class="bg-white shadow-[0_1px_3px_rgba(0,0,0,0.08)] rounded-lg p-4 mb-3">
          <div class="flex items-center justify-between mb-2">
            <span class="text-sm font-semibold text-[#1A1F24]">综合评分</span>
            <span class="text-lg font-semibold text-[#0F8B5D] flex items-center gap-1"><Icon icon="lucide:star" class="text-sm" /> 4.5/5.0</span>
          </div>
        </div>
      </div>
      <div>
        <h3 class="text-sm font-semibold text-[#1A1F24] mb-3 flex items-center gap-2"><Icon icon="lucide:git-branch" class="text-sm text-[#0D7C7C]" /> 来源追溯</h3>
        <div class="bg-white shadow-[0_1px_3px_rgba(0,0,0,0.08)] rounded-lg p-3">
          <div class="flex items-center gap-2 mb-1">
            <span class="text-sm font-semibold text-[#0D7C7C]">REQ-0007</span>
            <span class="text-xs px-2 py-0.5 rounded-full bg-[#E8F0FE] text-[#3367D6] font-semibold whitespace-nowrap">开发中</span>
          </div>
          <router-link to="/requirements/REQ-0007" class="text-sm text-[#4A5259] cursor-pointer hover:text-[#0D7C7C] transition">短信验证码功能</router-link>
          <span class="text-xs text-[#9BA3AB] mt-1 block">该知识条目从REQ-0007的架构方案中沉淀生成</span>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'

const keyPoints = [
  { title: '验证码过期管理', description: '使用Redis的TTL机制,设置验证码5分钟自动过期', icon: 'lucide:timer', bg: '#E8F7F0', color: '#0F8B5D' },
  { title: '发送频率限制', description: '基于Redis计数器实现每手机号每日5次发送上限', icon: 'lucide:shield', bg: '#E8F0FE', color: '#3367D6' },
  { title: '加密存储设计', description: '验证码在Redis中以SHA256哈希形式存储', icon: 'lucide:lock', bg: '#FEF6E8', color: '#F59D0D' },
  { title: '降级容灾策略', description: 'Redis不可用时自动降级到内存缓存', icon: 'lucide:git-merge', bg: 'hsla(200,15%,95%,1)', color: '#0D7C7C' },
]
</script>
