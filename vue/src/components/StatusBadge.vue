<template>
  <span :class="['px-3 py-1 rounded-full text-xs font-semibold whitespace-nowrap', customClass]" :style="customStyle">{{ label }}</span>
</template>

<script setup>
import { computed } from 'vue'
import { statusConfig, priorityConfig } from '../data/mockData'

const props = defineProps({
  label: { type: String, required: true },
  type: { type: String, default: 'status' }, // 'status' or 'priority'
})

const config = computed(() => {
  if (props.type === 'priority') return priorityConfig[props.label]
  return statusConfig[props.label]
})

const customClass = computed(() => '')
const customStyle = computed(() => {
  if (!config.value) return {}
  return { backgroundColor: config.value.bg, color: config.value.text }
})
</script>