<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="review-rules" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">审核规则</h1>
        <div class="flex items-center gap-3">
          <button @click="saveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Review Rules List Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-semibold text-[#1A1F24]">审核规则配置</h3>
            <p class="text-sm text-[#6B7680] mt-1">定义触发审核的条件和处理方式</p>
          </div>
          <button @click="addRule" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:plus" class="text-base" />
            <span class="whitespace-nowrap">新增规则</span>
          </button>
        </div>
        <div v-for="rule in reviewRules" :key="rule.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex flex-col gap-1">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ rule.title }}</span>
              <p class="text-sm text-[#6B7680]">{{ rule.description }}</p>
              <div class="flex items-center gap-4 mt-2">
                <div class="flex flex-col gap-1">
                  <label class="text-xs text-[#4A5259]">触发条件</label>
                  <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
                    <span class="text-sm text-[#1A1F24]">{{ rule.trigger }}</span>
                    <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
                  </div>
                </div>
                <div class="flex flex-col gap-1">
                  <label class="text-xs text-[#4A5259]">处理方式</label>
                  <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
                    <span class="text-sm text-[#1A1F24]">{{ rule.handler }}</span>
                    <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
                  </div>
                </div>
              </div>
            </div>
            <label class="flex items-center gap-3">
              <div class="relative">
                <input type="checkbox" class="sr-only peer" :checked="rule.enabled" @change="toggleRule(rule)">
                <div class="w-12 h-6 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition"></div>
                <div class="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow-sm peer-checked:translate-x-6 transition"></div>
              </div>
            </label>
          </div>
        </div>
      </div>

      <!-- Rule Template Selection Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">规则模板选择</h3>
          <p class="text-sm text-[#6B7680] mt-1">从预定义模板快速创建审核规则</p>
        </div>
        <div class="flex flex-col gap-1">
          <label class="text-xs text-[#4A5259]">选择模板</label>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <span class="text-sm text-[#1A1F24]">{{ selectedTemplate?.name || '安全优先模板' }}</span>
            <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
          </div>
        </div>
        <div v-if="selectedTemplate" class="bg-[hsla(200,20%,98%,1)] border border-[#0D7C7C] rounded-2xl p-4">
          <div class="flex items-center gap-2 mb-2">
            <Icon icon="lucide:info" class="text-base text-[#0D7C7C]" />
            <span class="text-sm font-semibold text-[#0D7C7C]">{{ selectedTemplate.name }}说明</span>
          </div>
          <p class="text-sm text-[#4A5259]">{{ selectedTemplate.description }}</p>
        </div>
      </div>

      <!-- Review Timeout Auto-processing Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">审核超时自动处理</h3>
          <p class="text-sm text-[#6B7680] mt-1">当审核请求超过指定时间未响应时自动处理</p>
        </div>
        <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex flex-col gap-1">
              <span class="text-sm font-semibold text-[#1A1F24]">启用超时自动处理</span>
              <p class="text-sm text-[#6B7680]">审核请求超过时限后自动按预设策略处理</p>
            </div>
            <label class="flex items-center gap-3">
              <div class="relative">
                <input type="checkbox" class="sr-only peer" :checked="timeoutConfig.enabled" @change="toggleTimeout">
                <div class="w-12 h-6 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition"></div>
                <div class="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow-sm peer-checked:translate-x-6 transition"></div>
              </div>
            </label>
          </div>
        </div>
        <div class="grid grid-cols-2 gap-3">
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">超时时限(小时)</label>
            <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
              <span class="text-sm text-[#1A1F24]">{{ timeoutConfig.threshold_hours }}小时</span>
              <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
            </div>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">超时处理策略</label>
            <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
              <span class="text-sm text-[#1A1F24]">{{ timeoutConfig.strategy }}</span>
              <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
            </div>
          </div>
        </div>
      </div>

    </main>
  </div>

  <!-- Bottom Save Bar -->
  <div class="fixed bottom-0 left-0 right-0 bg-[#FEF6E8] border-t border-[#F59D0D] px-8 py-3 flex items-center justify-between">
    <span class="text-sm text-[#F59D0D] font-semibold flex items-center gap-2">
      <Icon icon="lucide:alert-circle" class="text-base text-[#F59D0D]" />
      您有未保存的更改
    </span>
    <div class="flex items-center gap-3">
      <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
        <Icon icon="lucide:x" class="text-base" />
        <span class="whitespace-nowrap">取消</span>
      </button>
      <button @click="saveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
        <Icon icon="lucide:save" class="text-base" />
        <span class="whitespace-nowrap">保存更改</span>
      </button>
    </div>
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
  getReviewRules,
  createReviewRule,
  updateReviewRule,
  deleteReviewRule,
  getReviewTemplates,
  getReviewTimeoutConfig,
  updateReviewTimeoutConfig,
} from '../../api/settings'

const toastRef = ref(null)

const reviewRules = ref([])
const templates = ref([])
const selectedTemplate = ref(null)
const timeoutConfig = reactive({
  enabled: true,
  threshold_hours: 24,
  strategy: '自动通过并标记',
})

async function fetchRules() {
  const data = await getReviewRules()
  reviewRules.value = data.map(rule => ({
    id: rule.id,
    title: rule.name,
    description: rule.description,
    trigger: rule.trigger_condition,
    handler: rule.handling_method,
    enabled: rule.enabled,
  }))
}

async function fetchTemplates() {
  const data = await getReviewTemplates()
  templates.value = data
  if (data.length > 0) {
    selectedTemplate.value = data[0]
  }
}

async function fetchTimeoutConfig() {
  const data = await getReviewTimeoutConfig()
  if (data) {
    timeoutConfig.enabled = data.enabled
    timeoutConfig.threshold_hours = data.threshold_hours
    timeoutConfig.strategy = data.strategy
  }
}

async function toggleRule(rule) {
  rule.enabled = !rule.enabled
  await updateReviewRule(rule.id, { enabled: rule.enabled })
}

async function addRule() {
  const newRule = await createReviewRule({
    name: '新审核规则',
    description: '请编辑规则描述',
    trigger_condition: '请选择触发条件',
    handling_method: '请选择处理方式',
    enabled: false,
  })
  reviewRules.value.push({
    id: newRule.id,
    title: newRule.name,
    description: newRule.description,
    trigger: newRule.trigger_condition,
    handler: newRule.handling_method,
    enabled: newRule.enabled,
  })
}

async function toggleTimeout() {
  timeoutConfig.enabled = !timeoutConfig.enabled
  await updateReviewTimeoutConfig({
    enabled: timeoutConfig.enabled,
    threshold_hours: timeoutConfig.threshold_hours,
    strategy: timeoutConfig.strategy,
  })
}

async function saveAll() {
  try {
    await Promise.all(reviewRules.value.map(rule =>
      updateReviewRule(rule.id, {
        name: rule.title,
        description: rule.description,
        trigger_condition: rule.trigger,
        handling_method: rule.handler,
        enabled: rule.enabled,
      })
    ))
    await updateReviewTimeoutConfig(timeoutConfig)
    toastRef.value?.success('审核规则已保存')
  } catch (err) {
    console.error('Failed to save review rules:', err)
    toastRef.value?.error('保存失败，请重试')
  }
}

onMounted(() => {
  fetchRules()
  fetchTemplates()
  fetchTimeoutConfig()
})
</script>