<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="notifications" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">通知设置</h1>
        <div class="flex items-center gap-3">
          <button @click="saveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Notification Channels -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">通知渠道</h3>
          <p class="text-sm text-[#6B7680] mt-1">选择接收通知的方式</p>
        </div>

        <div v-for="channel in channels" :key="channel.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Icon :icon="channel.icon" :class="['text-base', channel.enabled ? 'text-[#0D7C7C]' : 'text-[#6B7680]']" />
              <div>
                <span class="text-sm font-semibold text-[#1A1F24]">{{ channel.name }}</span>
                <p class="text-xs text-[#6B7680]">{{ channel.desc }}</p>
              </div>
            </div>
            <div @click="toggleChannel(channel)" :class="['w-12 h-6 rounded-full transition-colors cursor-pointer', channel.enabled ? 'bg-[#0D7C7C]' : 'bg-[#E1E6EA]']">
              <div :class="['w-5 h-5 rounded-full bg-white shadow transition-transform', channel.enabled ? 'translate-x-6' : 'translate-x-0.5']"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Notification Types -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">通知类型设置</h3>
          <p class="text-sm text-[#6B7680] mt-1">选择需要接收通知的事件类型</p>
        </div>

        <div v-for="event in notifEvents" :key="event.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Icon :icon="event.icon" class="text-base text-[#0D7C7C]" />
              <span class="text-sm font-semibold text-[#1A1F24]">{{ event.name }}</span>
            </div>
            <div class="flex items-center gap-4">
              <label v-for="ch in channelLabels" :key="ch.value" class="flex items-center gap-1.5 text-xs text-[#4A5259]">
                <input type="checkbox" :checked="event.channels.includes(ch.value)" @change="toggleEventChannel(event, ch.value)" class="accent-[#0D7C7C] w-3.5 h-3.5">
                <span>{{ ch.label }}</span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <!-- Quiet Hours -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">免打扰时段</h3>
          <p class="text-sm text-[#6B7680] mt-1">设置免打扰时间段,暂停所有通知推送</p>
        </div>
        <div class="flex items-center gap-4">
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full">
            <Icon icon="lucide:clock" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#4A5259]">{{ preferences.quiet_hours_start || '22:00' }}</span>
          </div>
          <span class="text-sm text-[#6B7680]">至</span>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full">
            <Icon icon="lucide:clock" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#4A5259]">{{ preferences.quiet_hours_end || '08:00' }}</span>
          </div>
          <div @click="toggleQuietHours" :class="['w-12 h-6 rounded-full cursor-pointer', preferences.quiet_hours_enabled ? 'bg-[#0D7C7C]' : 'bg-[#E1E6EA]']">
            <div :class="['w-5 h-5 rounded-full bg-white shadow transition-transform', preferences.quiet_hours_enabled ? 'translate-x-6' : 'translate-x-0.5']"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
  <Toast ref="toastRef" />
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import SettingsSidebar from '../../layout/SettingsSidebar.vue'
import Toast from '../../components/Toast.vue'
import {
  getNotificationChannels,
  updateNotificationChannel,
  getNotificationEvents,
  updateNotificationEvent,
  getNotificationPreferences,
  updateNotificationPreferences,
} from '../../api/settings'

const toastRef = ref(null)

const channelTypeMap = {
  email: { name: '邮件通知', desc: '通过邮件接收重要通知', icon: 'lucide:mail' },
  slack: { name: 'Slack通知', desc: '推送至Slack频道', icon: 'lucide:message-circle' },
  in_app: { name: '站内通知', desc: '系统内弹窗与消息提醒', icon: 'lucide:bell' },
  sms: { name: '短信通知', desc: '紧急事件短信提醒', icon: 'lucide:smartphone' },
}

const eventTypeMap = {
  review: { name: '审核通知', icon: 'lucide:shield-check' },
  deploy: { name: '部署通知', icon: 'lucide:rocket' },
  error: { name: '错误报警', icon: 'lucide:alert-circle' },
  task_complete: { name: '任务完成', icon: 'lucide:check-circle' },
  requirement_change: { name: '需求变更', icon: 'lucide:file-text' },
}

const channelLabels = [
  { value: '邮件', label: '邮件' },
  { value: '站内', label: '站内' },
  { value: 'Slack', label: 'Slack' },
]

const channels = ref([])
const notifEvents = ref([])
const preferences = reactive({
  quiet_hours_enabled: false,
  quiet_hours_start: '22:00',
  quiet_hours_end: '08:00',
  frequency: '',
})

async function fetchChannels() {
  const data = await getNotificationChannels()
  channels.value = data.map(ch => ({
    id: ch.id,
    channel_type: ch.channel_type,
    name: channelTypeMap[ch.channel_type]?.name || ch.channel_type,
    desc: channelTypeMap[ch.channel_type]?.desc || '',
    icon: channelTypeMap[ch.channel_type]?.icon || 'lucide:bell',
    enabled: ch.enabled,
  }))
}

async function fetchEvents() {
  const data = await getNotificationEvents()
  notifEvents.value = data.map(ev => ({
    id: ev.id,
    event_type: ev.event_type,
    name: eventTypeMap[ev.event_type]?.name || ev.event_type,
    icon: eventTypeMap[ev.event_type]?.icon || 'lucide:bell',
    channels: ev.channels || [],
    enabled: ev.enabled,
    severity: ev.severity,
  }))
}

async function fetchPreferences() {
  const data = await getNotificationPreferences()
  if (data) {
    preferences.quiet_hours_enabled = data.quiet_hours_enabled
    preferences.quiet_hours_start = data.quiet_hours_start
    preferences.quiet_hours_end = data.quiet_hours_end
    preferences.frequency = data.frequency
  }
}

async function toggleChannel(channel) {
  channel.enabled = !channel.enabled
  await updateNotificationChannel(channel.id, { enabled: channel.enabled })
}

async function toggleEventChannel(event, channelName) {
  const idx = event.channels.indexOf(channelName)
  if (idx >= 0) {
    event.channels.splice(idx, 1)
  } else {
    event.channels.push(channelName)
  }
  await updateNotificationEvent(event.id, { channels: event.channels })
}

async function toggleQuietHours() {
  preferences.quiet_hours_enabled = !preferences.quiet_hours_enabled
  await updateNotificationPreferences({
    quiet_hours_enabled: preferences.quiet_hours_enabled,
    quiet_hours_start: preferences.quiet_hours_start,
    quiet_hours_end: preferences.quiet_hours_end,
  })
}

async function saveAll() {
  try {
    await updateNotificationPreferences(preferences)
    toastRef.value?.success('通知设置已保存')
  } catch (err) {
    console.error('Failed to save notification settings:', err)
    toastRef.value?.error('保存失败，请重试')
  }
}

onMounted(() => {
  fetchChannels()
  fetchEvents()
  fetchPreferences()
})
</script>