<template>
  <TopNavBar activeNav="requirements" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-12 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">需求管理</h1>
        <router-link to="/requirements/new" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:plus" class="text-sm" />
          <span class="whitespace-nowrap">新建需求</span>
        </router-link>
      </div>
      <div class="flex items-center justify-between mb-6">
        <div class="flex gap-2 overflow-x-auto">
          <label v-for="status in statusFilters" :key="status" class="cursor-pointer">
            <input type="radio" name="statusFilter" class="sr-only peer" :checked="status === '全部'">
            <div class="bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full peer-checked:bg-[#0D7C7C] peer-checked:text-white/95 hover:opacity-80 transition whitespace-nowrap text-sm">{{ status }}</div>
          </label>
        </div>
        <div class="flex items-center gap-3">
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full">
            <Icon icon="lucide:arrow-up-down" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#4A5259] whitespace-nowrap">最新创建</span>
          </div>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full">
            <Icon icon="lucide:search" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#9BA3AB] whitespace-nowrap">搜索标题或ID</span>
          </div>
        </div>
      </div>
      <div class="flex flex-col gap-4">
        <router-link v-for="req in requirements" :key="req.id" :to="`/requirements/${req.id}`"
          :class="[
            'bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 cursor-pointer transition hover:shadow-[0_4px_12px_rgba(0,0,0,0.12)] no-underline',
            req.status === '已完成' ? 'opacity-70' : ''
          ]">
          <div class="flex items-start justify-between">
            <div class="flex flex-col gap-3 flex-1">
              <div class="flex items-center gap-3">
                <span class="text-sm font-semibold text-[#0D7C7C]">{{ req.id }}</span>
                <span class="text-lg font-semibold text-[#1A1F24]">{{ req.title }}</span>
                <StatusBadge :label="req.status" />
                <PriorityBadge :label="req.priority" />
              </div>
              <p class="text-sm text-[#6B7680] line-clamp-2">{{ req.description }}</p>
              <div class="flex items-center gap-4 text-sm text-[#9BA3AB]">
                <span class="flex items-center gap-2">
                  <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${req.authorAvatar}`" class="w-6 h-6 rounded-full" />
                  {{ req.author }}
                </span>
                <span>创建于 {{ req.createdAt }}</span>
                <template v-if="req.updatedAt">
                  <span>更新于 {{ req.updatedAt }}</span>
                </template>
                <span class="flex items-center gap-1">
                  <Icon icon="lucide:tasks" class="text-xs text-[#9BA3AB]" />
                  {{ req.taskCount }}个任务
                </span>
              </div>
            </div>
            <div class="flex items-center gap-2 ml-4">
              <div class="flex items-center justify-center w-10 h-10 rounded-full transition hover:opacity-80 bg-white border border-[#E1E6EA]">
                <Icon icon="lucide:eye" class="text-sm text-[#4A5259]" />
              </div>
            </div>
          </div>
        </router-link>
      </div>
      <Pagination :currentPage="1" :totalPages="3" :total="15" />
    </main>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import StatusBadge from '../../components/StatusBadge.vue'
import PriorityBadge from '../../components/PriorityBadge.vue'
import Pagination from '../../components/Pagination.vue'
import { requirements } from '../../data/mockData'

const statusFilters = ['全部', '进行中', '待审核', '已完成', '已归档']
</script>
