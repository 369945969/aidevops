<template>
  <Transition name="toast">
    <div v-if="visible" :class="['fixed top-6 right-6 z-50 flex items-center gap-3 px-5 py-3 rounded-xl shadow-lg transition-all', typeClass]">
      <Icon :icon="iconName" class="text-base" />
      <span class="text-sm font-semibold">{{ message }}</span>
    </div>
  </Transition>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import { Icon } from '@iconify/vue'

const visible = ref(false)
const message = ref('')
const type = ref('success')
const timer = ref(null)

const typeClass = computed(() => {
  if (type.value === 'success') return 'bg-[#E8F7F0] text-[#0F8B5D] border border-[#0F8B5D]/20'
  if (type.value === 'error') return 'bg-[#FDECEA] text-[#D93025] border border-[#D93025]/20'
  return 'bg-[#FEF6E8] text-[#F59D0D] border border-[#F59D0D]/20'
})

const iconName = computed(() => {
  if (type.value === 'success') return 'lucide:check-circle'
  if (type.value === 'error') return 'lucide:x-circle'
  return 'lucide:alert-circle'
})

function show(msg, t = 'success', duration = 3000) {
  message.value = msg
  type.value = t
  visible.value = true
  if (timer.value) clearTimeout(timer.value)
  timer.value = setTimeout(() => { visible.value = false }, duration)
}

function success(msg) { show(msg, 'success') }
function error(msg) { show(msg, 'error', 5000) }
function warning(msg) { show(msg, 'warning') }

defineExpose({ show, success, error, warning })
</script>

<style scoped>
.toast-enter-active { transition: all 0.3s ease-out; }
.toast-leave-active { transition: all 0.3s ease-in; }
.toast-enter-from { opacity: 0; transform: translateY(-20px); }
.toast-leave-to { opacity: 0; transform: translateY(-20px); }
</style>