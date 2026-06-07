<template>
  <span :class="['px-3 py-1 rounded-full text-xs font-semibold whitespace-nowrap']" :style="{ backgroundColor: config?.bg, color: config?.text }">{{ label }}</span>
</template>

<script setup>
import { computed } from 'vue'
import { priorityConfig } from '../data/mockData'

const props = defineProps({
  label: { type: String, required: true },
})

const config = computed(() => priorityConfig[props.label])
</script>