<template>
  <TopNavBar activeNav="knowledge" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <KnowledgeSidebar activeCategory="code-patterns" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h1 class="text-2xl font-bold text-[#1A1F24]">代码模式</h1>
          <div class="flex items-center gap-4 mt-1 text-sm text-[#6B7680]">
            <span>总条目数 <span class="text-[#1A1F24] font-semibold">{{ filteredEntries.length }}</span></span>
          </div>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full" style="width:280px">
            <Icon icon="lucide:search" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#6B7680]">搜索...</span>
          </div>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon icon="lucide:arrow-up-down" class="text-base text-[#6B7680]" />
            <span class="text-sm text-[#4A5259] whitespace-nowrap">最新创建</span>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-3 gap-5">
        <router-link v-for="entry in filteredEntries" :key="entry.id" :to="`/knowledge/detail/${entry.id}`"
          class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 no-underline hover:shadow-[0_4px_12px_rgba(0,0,0,0.12)] transition">
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
        </router-link>
      </div>

      <Pagination :currentPage="1" :totalPages="5" :total="filteredEntries.length" />
    </main>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import KnowledgeSidebar from '../../layout/KnowledgeSidebar.vue'
import Pagination from '../../components/Pagination.vue'
import QualityScore from '../../components/QualityScore.vue'
import { knowledgeEntries, knowledgeTypeConfig } from '../../data/mockData'

const typeMap = {
  'requirement-docs': '需求文档',
  'architecture': '架构方案',
  'code-patterns': '代码模式',
  'solutions': '问题解决方案',
  'deploy-config': '部署配置',
  'favorites': null,
}

const filteredEntries = computed(() => {
  const type = typeMap['code-patterns']
  if (!type) return knowledgeEntries
  return knowledgeEntries.filter(e => e.type === type)
})
</script>
