<template>
  <aside class="flex-shrink-0 min-w-fit w-[220px] bg-[hsla(200,15%,97.5%,1)] border-r border-[#E1E6EA] py-4">
    <nav class="flex flex-col">
      <router-link v-for="section in settingsSections" :key="section.key" :to="section.route"
        :class="[
          'flex items-center gap-3 px-5 py-3 text-sm cursor-pointer transition no-underline',
          activeSection === section.key
            ? 'font-semibold text-[#0D7C7C] border-l-2 border-[#0D7C7C] bg-[hsla(200,20%,98%,1)]'
            : 'text-[#4A5259] border-l-2 border-transparent hover:text-[#0D7C7C]'
        ]">
        <Icon :icon="section.icon" class="text-base" />
        <span>{{ section.label }}</span>
      </router-link>
    </nav>
  </aside>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import { settingsSections } from '../data/mockData'

defineProps({
  activeSection: { type: String, default: 'automation' }
})
</script>