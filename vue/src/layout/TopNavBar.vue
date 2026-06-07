<template>
  <header class="w-full bg-[hsla(200,15%,96%,1)]">
    <div class="flex items-center justify-between w-full px-8 py-3">
      <div class="flex items-center gap-8">
        <router-link to="/" class="flex items-center gap-2">
          <Icon icon="lucide:workflow" class="text-xl text-[#0D7C7C]" />
          <span class="text-base font-semibold text-[#0D7C7C]">AI DevOps</span>
        </router-link>
        <div class="flex items-center gap-6">
          <router-link v-for="item in navItems" :key="item.key" :to="item.route"
            :class="[
              'flex items-center gap-2 text-sm cursor-pointer transition',
              activeNav === item.key ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259] hover:text-[#0D7C7C]'
            ]">
            <Icon :icon="item.icon" class="text-sm" />
            <span>{{ item.label }}</span>
          </router-link>
        </div>
      </div>
      <div class="flex items-center gap-4">
        <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full">
          <Icon icon="lucide:search" class="text-sm text-[#6B7680]" />
          <span class="text-sm text-[#9BA3AB]">搜索任务或需求...</span>
        </div>
        <router-link to="/requirements/new" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:plus" class="text-sm" />
          <span class="whitespace-nowrap text-sm">新建需求</span>
        </router-link>
        <router-link to="/notifications" class="relative flex items-center justify-center w-10 h-10">
          <Icon icon="lucide:bell" class="w-5 h-5 text-[#4A5259]" />
          <div class="absolute -top-1 -right-1 w-5 h-5 bg-[#D93025] rounded-full flex items-center justify-center">
            <span class="text-xs text-white/95">3</span>
          </div>
        </router-link>
        <div class="flex items-center gap-2">
          <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=alex" class="w-9 h-9 rounded-full" />
          <Icon icon="lucide:chevron-down" class="w-4 h-4 text-[#6B7680]" />
        </div>
      </div>
    </div>
  </header>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { navItems } from '../data/mockData'

defineProps({
  activeNav: { type: String, default: 'workflow' }
})
</script>