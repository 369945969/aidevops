<template>
  <span class="flex items-center gap-1 text-[#0F8B5D] text-xs">
    <Icon icon="lucide:star" class="text-xs" />
    质量 {{ score }}
  </span>
</template>

<script setup>
import { Icon } from '@iconify/vue'

defineProps({
  score: { type: Number, required: true },
})
</script>