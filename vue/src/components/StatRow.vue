<template>
  <div class="flex items-center gap-4 mt-1 text-sm text-[#6B7680]">
    <span v-for="stat in stats" :key="stat.label">
      {{ stat.label }}
      <span :class="stat.highlight ? 'text-[#0D7C7C] font-semibold' : 'text-[#1A1F24] font-semibold'">{{ stat.value }}</span>
    </span>
  </div>
</template>

<script setup>
defineProps({
  stats: { type: Array, default: () => [] },
})
</script>