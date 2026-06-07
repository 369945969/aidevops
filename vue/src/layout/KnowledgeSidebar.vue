<template>
  <aside class="flex-shrink-0 w-[220px] bg-[hsla(200,15%,97.5%,1)] border-r border-[#E1E6EA] flex flex-col p-4 gap-1 overflow-y-auto">
    <div class="flex flex-col gap-1">
      <router-link v-for="cat in knowledgeCategories" :key="cat.key" :to="cat.route"
        :class="[
          'flex items-center gap-3 px-3 py-2.5 rounded-lg cursor-pointer transition no-underline',
          activeCategory === cat.key
            ? 'bg-[hsla(200,20%,92%,1)] text-[#0D7C7C] font-semibold relative'
            : 'text-[#4A5259] hover:bg-[hsla(200,15%,92%,1)]'
        ]">
        <div v-if="activeCategory === cat.key" class="absolute left-0 top-1/2 -translate-y-1/2 w-1.5 h-5 bg-[#0D7C7C] rounded-full"></div>
        <Icon :icon="cat.icon" class="text-base" />
        <span class="text-sm whitespace-nowrap">{{ cat.label }}</span>
      </router-link>
    </div>
  </aside>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { knowledgeCategories } from '../data/mockData'

defineProps({
  activeCategory: { type: String, default: 'all' }
})
</script>