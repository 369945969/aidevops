<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="integrations" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">集成配置</h1>
        <div class="flex items-center gap-3">
          <button @click="handleSave" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Connected Services -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-semibold text-[#1A1F24]">已连接服务</h3>
            <p class="text-sm text-[#6B7680] mt-1">管理第三方服务集成状态</p>
          </div>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:plus" class="text-base" />
            <span class="whitespace-nowrap text-sm">添加集成</span>
          </button>
        </div>

        <div v-for="svc in services" :key="svc.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <div :class="['flex items-center justify-center w-10 h-10 rounded-xl', svc.bg]">
                <Icon :icon="svc.icon" :class="['text-lg', svc.iconColor]" />
              </div>
              <div>
                <span class="text-sm font-semibold text-[#1A1F24]">{{ svc.name }}</span>
                <p class="text-xs text-[#6B7680]">{{ svc.desc }}</p>
              </div>
            </div>
            <div class="flex items-center gap-3">
              <div :class="['px-3 py-1 rounded-full text-xs font-semibold', svc.connected ? 'bg-[#E8F7F0] text-[#0F8B5D]' : 'bg-[#FDECEA] text-[#D93025]']">
                {{ svc.connected ? '已连接' : '未连接' }}
              </div>
              <button v-if="svc.connected" class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full transition hover:opacity-80 text-sm text-[#4A5259]">
                <Icon icon="lucide:settings" class="text-sm" />
                <span>配置</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- API Configuration -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">API密钥管理</h3>
          <p class="text-sm text-[#6B7680] mt-1">管理各集成服务的API密钥和认证信息</p>
        </div>
        <div v-for="key in apiKeys" :key="key.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <Icon :icon="key.icon" class="text-base text-[#6B7680]" />
            <span class="text-sm font-semibold text-[#1A1F24]">{{ key.service }}</span>
          </div>
          <div class="flex items-center gap-2">
            <span class="text-xs text-[#9BA3AB]">{{ key.masked }}</span>
            <button class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full transition hover:opacity-80 text-sm text-[#4A5259]">
              <Icon icon="lucide:copy" class="text-sm" />
              <span>复制</span>
            </button>
          </div>
        </div>
      </div>
    </main>
  </div>
  <Toast ref="toastRef" />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import SettingsSidebar from '../../layout/SettingsSidebar.vue'
import Toast from '../../components/Toast.vue'
import {
  getConnectedServices,
  updateConnectedService,
  getAPIKeys,
} from '../../api/settings'

const toastRef = ref(null)
const services = ref([])
const apiKeys = ref([])

function maskKeyValue(value) {
  if (!value) return ''
  const visibleStart = Math.min(4, value.length)
  const visibleEnd = Math.min(3, value.length)
  const start = value.slice(0, visibleStart)
  const end = value.slice(-visibleEnd)
  return `${start}****${end}`
}

function mapService(raw) {
  const config = raw.config_data || {}
  return {
    id: raw.id,
    name: raw.name,
    desc: config.desc || '',
    icon: raw.icon || 'lucide:plug',
    bg: config.bg || 'bg-[#E8F0FE]',
    iconColor: config.iconColor || 'text-[#3367D6]',
    connected: raw.status === 'connected',
    raw,
  }
}

function mapAPIKey(raw) {
  return {
    id: raw.id,
    service: raw.name || raw.service,
    icon: raw.icon || 'lucide:key',
    masked: maskKeyValue(raw.key_value),
  }
}

async function fetchServices() {
  try {
    const res = await getConnectedServices()
    services.value = (res || []).map(mapService)
  } catch {
    services.value = []
  }
}

async function fetchAPIKeys() {
  try {
    const res = await getAPIKeys()
    apiKeys.value = (res || []).map(mapAPIKey)
  } catch {
    apiKeys.value = []
  }
}

async function handleSave() {
  try {
    for (const svc of services.value) {
      await updateConnectedService(svc.id, svc.raw)
    }
    toastRef.value?.success('集成配置已保存')
  } catch (err) {
    console.error('Failed to save integrations:', err)
    toastRef.value?.error('保存失败，请重试')
  }
}

onMounted(() => {
  fetchServices()
  fetchAPIKeys()
})
</script>