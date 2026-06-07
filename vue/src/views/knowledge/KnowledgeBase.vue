<template>
  <TopNavBar activeNav="knowledge" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <KnowledgeSidebar activeCategory="all" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="text-2xl font-bold text-[#1A1F24]">知识库</h1>
          <div class="flex items-center gap-4 mt-1 text-sm text-[#6B7680]">
            <span>总条目数 <span class="text-[#1A1F24] font-semibold">156</span></span>
            <span class="text-[#E1E6EA]">|</span>
            <span>本月新增 <span class="text-[#0D7C7C] font-semibold">12</span></span>
            <span class="text-[#E1E6EA]">|</span>
            <span>引用次数 <span class="text-[#1A1F24] font-semibold">892</span></span>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full" style="width:280px">
            <Icon icon="lucide:search" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#6B7680]">全文搜索知识库...</span>
            <button class="text-xs text-[#0D7C7C] ml-auto whitespace-nowrap cursor-pointer">高级筛选</button>
          </div>
          <div class="flex items-center gap-2 bg-white border border-[#E1E6EA] rounded-full px-3 py-2">
            <Icon icon="lucide:layout-grid" class="text-base text-[#0D7C7C]" />
            <span class="text-sm text-[#0D7C7C] font-semibold whitespace-nowrap">卡片</span>
            <span class="text-sm text-[#E1E6EA]">|</span>
            <Icon icon="lucide:list" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#6B7680] whitespace-nowrap">列表</span>
          </div>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon icon="lucide:arrow-up-down" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#4A5259] whitespace-nowrap">最新创建</span>
          </div>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-3 mb-5">
        <div class="flex items-center justify-between">
          <span class="text-sm font-semibold text-[#1A1F24]">高级筛选</span>
          <button class="text-xs text-[#0D7C7C] cursor-pointer whitespace-nowrap">重置筛选</button>
        </div>
        <div class="flex items-center gap-3">
          <div v-for="filter in filterOptions" :key="filter.label" class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon :icon="filter.icon" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#4A5259] whitespace-nowrap">{{ filter.label }}</span>
            <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
          </div>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:check" class="text-base" />
            <span class="text-sm whitespace-nowrap">应用筛选</span>
          </button>
        </div>
      </div>

      <div class="grid grid-cols-3 gap-5">
        <router-link v-for="entry in knowledgeEntries" :key="entry.id" :to="`/knowledge/detail/${entry.id}`"
          class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 no-underline hover:shadow-[0_4px_12px_rgba(0,0,0,0.12)] transition relative">
          <div v-if="entry.aiRecommended" class="absolute top-3 right-3 bg-[#0D7C7C] text-white/95 text-xs px-3 py-1 rounded-full flex items-center gap-1">
            <Icon icon="lucide:bot" class="text-xs" />
            <span class="whitespace-nowrap">相似知识推荐</span>
          </div>
          <div v-if="entry.aiRecommended" class="flex items-center gap-1 text-xs text-[#0D7C7C] bg-[hsla(200,20%,92%,1)] rounded-full px-2.5 py-1">
            <Icon icon="lucide:bot" class="text-xs" />
            <span class="whitespace-nowrap font-semibold">AI推荐</span>
          </div>
          <div class="flex items-center gap-2">
            <div class="w-8 h-8 rounded-lg flex items-center justify-center" :style="{ backgroundColor: knowledgeTypeConfig[entry.type]?.bg }">
              <Icon :icon="knowledgeTypeConfig[entry.type]?.icon" class="text-sm" :style="{ color: knowledgeTypeConfig[entry.type]?.text }" />
            </div>
            <span class="text-xs px-3 py-1 rounded-full font-semibold whitespace-nowrap" :style="{ backgroundColor: knowledgeTypeConfig[entry.type]?.bg, color: knowledgeTypeConfig[entry.type]?.text }">{{ entry.type }}</span>
          </div>
          <h3 class="text-base font-semibold text-[#1A1F24]">{{ entry.title }}</h3>
          <p class="text-sm text-[#6B7680] line-clamp-3">{{ entry.description }}</p>
          <div class="flex items-center gap-2">
            <span v-for="tag in entry.tags" :key="tag" class="text-xs bg-[hsla(200,15%,95%,1)] rounded-full px-2.5 py-1 text-[#4A5259] whitespace-nowrap">{{ tag }}</span>
          </div>
          <div class="flex items-center justify-between text-xs text-[#9BA3AB]">
            <span>引用 {{ entry.citations }}次</span>
            <QualityScore :score="entry.quality" />
          </div>
          <div class="flex items-center justify-between">
            <span class="text-xs text-[#9BA3AB]">来源 {{ entry.source }} · {{ entry.date }}</span>
            <router-link :to="`/knowledge/detail/${entry.id}`" class="flex items-center gap-2 px-3 py-1.5 text-xs text-[#0D7C7C] rounded-full transition hover:bg-[hsla(200,20%,92%,1)]">
              <span class="whitespace-nowrap">查看详情</span>
            </router-link>
          </div>
          <div v-if="entry.aiRecommended" class="text-xs text-[#0D7C7C] flex items-center gap-1">
            <Icon icon="lucide:lightbulb" class="text-xs" />
            <span class="whitespace-nowrap">推荐理由: {{ entry.recommendationReason }}</span>
          </div>
        </router-link>
      </div>

      <Pagination :currentPage="1" :totalPages="11" :total="156" />
    </main>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import KnowledgeSidebar from '../../layout/KnowledgeSidebar.vue'
import Pagination from '../../components/Pagination.vue'
import QualityScore from '../../components/QualityScore.vue'
import { knowledgeEntries, knowledgeTypeConfig } from '../../data/mockData'

const filterOptions = [
  { label: '类别', icon: 'lucide:filter' },
  { label: '质量评分', icon: 'lucide:star' },
  { label: '时间范围', icon: 'lucide:calendar' },
  { label: '技术标签', icon: 'lucide:tag' },
  { label: '知识类型', icon: 'lucide:layers' },
  { label: '创建者', icon: 'lucide:user' },
]
</script>
