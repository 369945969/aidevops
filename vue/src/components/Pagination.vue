<template>
  <div class="flex items-center justify-center gap-2">
    <button class="flex items-center justify-center w-9 h-9 rounded-full border border-[#E1E6EA] text-[#6B7680] transition hover:bg-[hsla(200,15%,92%,1)] cursor-pointer">
      <Icon icon="lucide:chevron-left" class="text-sm" />
    </button>
    <button v-for="page in displayedPages" :key="page"
      :class="[
        'flex items-center justify-center w-9 h-9 rounded-full cursor-pointer text-sm',
        page === currentPage ? 'bg-[#0D7C7C] text-white/95 font-semibold' : 'border border-[#E1E6EA] text-[#4A5259] transition hover:bg-[hsla(200,15%,92%,1)]'
      ]">
      {{ page }}
    </button>
    <span v-if="showEllipsis" class="text-sm text-[#9BA3AB] px-2">...</span>
    <button class="flex items-center justify-center w-9 h-9 rounded-full border border-[#E1E6EA] text-[#6B7680] transition hover:bg-[hsla(200,15%,92%,1)] cursor-pointer">
      <Icon icon="lucide:chevron-right" class="text-sm" />
    </button>
    <span class="text-sm text-[#6B7680] ml-4 whitespace-nowrap">共 {{ total }} 条记录</span>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Icon } from '@iconify/vue'

const props = defineProps({
  currentPage: { type: Number, default: 1 },
  totalPages: { type: Number, default: 11 },
  total: { type: Number, default: 156 },
})

const displayedPages = computed(() => {
  const pages = []
  for (let i = 1; i <= Math.min(3, props.totalPages); i++) pages.push(i)
  if (props.totalPages > 3) pages.push(props.totalPages)
  return pages
})

const showEllipsis = computed(() => props.totalPages > 4)
</script>