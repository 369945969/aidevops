<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-12 py-6">
      <div class="flex items-center justify-between mb-5">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">工作流看板</h1>
        <div class="flex items-center gap-4">
          <div class="flex gap-2 overflow-x-auto">
            <label v-for="filter in filters" :key="filter" class="cursor-pointer">
              <input type="radio" name="filter1" class="sr-only peer" :checked="filter === '全部任务'">
              <div class="bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full peer-checked:bg-[#0D7C7C] peer-checked:text-white/95 hover:opacity-80 transition whitespace-nowrap text-sm">{{ filter }}</div>
            </label>
          </div>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon icon="lucide:layout-grid" class="text-sm text-[#0D7C7C]" />
            <span class="text-sm text-[#0D7C7C] font-semibold">看板</span>
            <span class="text-[#E1E6EA]">|</span>
            <Icon icon="lucide:list" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#6B7680]">列表</span>
          </div>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon icon="lucide:calendar" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#4A5259]">最近7天</span>
          </div>
        </div>
      </div>

      <div class="flex gap-5 overflow-x-auto pb-4" style="min-height:620px">
        <div v-for="col in kanbanColumns" :key="col.key"
          :class="[
            'flex flex-col gap-3 min-w-[260px] rounded-2xl p-5',
            col.bg ? col.bg : 'bg-[hsla(200,15%,95%,1)]'
          ]">
          <div class="flex items-center justify-between px-1 mb-1">
            <span class="flex items-center gap-2 text-sm font-semibold" :class="col.iconColor ? `text-[${col.iconColor}]` : 'text-[#4A5259]'">
              <Icon :icon="col.icon" class="text-sm" :class="`text-[${col.iconColor || '#4A5259'}]`" />
              {{ col.label }}
            </span>
            <div :class="[
              'px-3 py-1 rounded-full text-xs font-semibold',
              col.urgent ? 'bg-[#FDECEA] text-[#D93025]' : 'bg-[hsla(200,20%,92%,1)] text-[#4A5259]'
            ]">
              {{ col.count }}
              <template v-if="col.warning || col.urgent"> !</template>
            </div>
          </div>

          <router-link v-for="task in col.tasks" :key="task.id"
            :to="task.reviewRoute || (task.status === '已完成' ? '/requirements/REQ-0007' : '/workflow/task/1')"
            :class="[
              'bg-white shadow-[0_1px_3px_rgba(0,0,0,0.08)] rounded-2xl flex flex-col p-5 gap-3 no-underline hover:shadow-[0_4px_12px_rgba(0,0,0,0.12)] transition',
              task.status === '已完成' ? 'opacity-70' : '',
              task.status === '需人工审核' ? 'border-l-4 border-l-[#D93025]' : '',
              task.status === '待审核' ? 'border-l-4 border-l-[#F59D0D]' : ''
            ]">
            <span class="text-sm text-[#9BA3AB]">{{ task.id }}</span>
            <span class="text-lg font-semibold text-[#1A1F24]">{{ task.title }}</span>
            <template v-if="task.status && task.status !== '已完成'">
              <div :style="{ backgroundColor: statusConfig[task.status]?.bg, color: statusConfig[task.status]?.text }"
                class="px-3 py-1 rounded-full text-xs font-semibold">{{ task.status }}</div>
            </template>
            <template v-if="task.agent">
              <div class="flex items-center justify-between">
                <span class="flex items-center gap-2 text-sm text-[#6B7680]">
                  <Icon icon="lucide:bot" class="text-sm text-[#0D7C7C]" />
                  {{ task.agent }}
                </span>
                <template v-if="task.priority">
                  <PriorityBadge :label="task.priority" />
                </template>
              </div>
            </template>
            <template v-if="task.progress && task.status !== '已完成'">
              <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full">
                <div class="h-2 rounded-full" :class="task.progress >= 100 && task.status === '待审核' ? 'bg-[#F59D0D]' : 'bg-[#0D7C7C]'" :style="{ width: task.progress + '%' }"></div>
              </div>
            </template>
            <template v-if="task.estimate">
              <span class="text-sm text-[#9BA3AB]">{{ task.estimate }}</span>
            </template>
            <template v-if="task.extra">
              <span class="text-sm text-[#9BA3AB]">{{ task.extra }}</span>
            </template>
            <template v-if="task.reviewRoute && task.status === '需人工审核'">
              <router-link :to="task.reviewRoute" class="flex items-center gap-2 px-4 py-2 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80 text-sm">
                <span class="whitespace-nowrap">审核部署</span>
              </router-link>
            </template>
            <template v-if="task.reviewRoute && task.status === '待审核' && col.key === 'architecture-design'">
              <router-link :to="task.reviewRoute" class="flex items-center gap-2 px-4 py-2 bg-[#F59D0D] text-white/95 rounded-full transition hover:opacity-80 text-sm">
                <span class="whitespace-nowrap">审核</span>
              </router-link>
            </template>
          </router-link>
        </div>
      </div>

      <div class="flex items-center justify-end gap-2 mt-5">
        <div class="flex items-center gap-2 bg-white border border-[#E1E6EA] rounded-full px-4 py-2">
          <div class="w-2 h-2 rounded-full bg-[#0F8B5D] animate-pulse"></div>
          <span class="text-sm text-[#4A5259]">实时同步中</span>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import PriorityBadge from '../../components/PriorityBadge.vue'
import { kanbanColumns, statusConfig } from '../../data/mockData'

const filters = ['全部任务', '我的审核', '今日更新']
</script>