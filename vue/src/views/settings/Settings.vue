<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="automation" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <!-- Page Title + Actions -->
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">自动化配置</h1>
        <div class="flex items-center gap-3">
          <button @click="saveConfig" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Automation Level Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">自动化级别设置</h3>
          <p class="text-sm text-[#6B7680] mt-1">控制AI Agent的自动化程度,从保守到激进</p>
        </div>
        <div class="flex items-center gap-6">
          <label v-for="level in automationLevels" :key="level.key" class="flex items-center gap-2">
            <input type="radio" name="auto-level" class="sr-only peer" :value="level.key" v-model="config.level">
            <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
              <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
            </div>
            <span :class="['text-sm', config.level === level.key ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">{{ level.label }}</span>
          </label>
        </div>
        <div class="w-full">
          <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full">
            <div class="h-2 bg-[#0D7C7C] rounded-full transition" :style="{ width: levelProgress + '%' }"></div>
          </div>
          <div class="flex items-center justify-between text-xs text-[#9BA3AB] mt-1">
            <span>保守</span>
            <span>平衡</span>
            <span>激进</span>
            <span>全自动</span>
          </div>
        </div>
        <div class="bg-[hsla(200,20%,98%,1)] border border-[#0D7C7C] rounded-2xl p-4">
          <div class="flex items-center gap-2 mb-2">
            <Icon icon="lucide:info" class="text-base text-[#0D7C7C]" />
            <span class="text-sm font-semibold text-[#0D7C7C]">当前: {{ currentLevelLabel }}</span>
          </div>
          <p class="text-sm text-[#4A5259]">{{ currentLevelDescription }}</p>
        </div>
      </div>

      <!-- Human Review Nodes Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">人工审核节点配置</h3>
          <p class="text-sm text-[#6B7680] mt-1">选择哪些环节需要人工审核</p>
        </div>
        <div v-for="node in nodes" :key="node.name" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex flex-col gap-1">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ node.title }}</span>
              <p class="text-sm text-[#6B7680]">{{ node.description }}</p>
              <p :class="['text-xs flex items-center gap-1', node.enabled ? 'text-[#0D7C7C]' : 'text-[#F59D0D]']">
                <Icon icon="lucide:alert-circle" class="text-xs" :class="node.enabled ? 'text-[#0D7C7C]' : 'text-[#F59D0D]'" />
                影响: {{ node.impact }}
              </p>
              <template v-if="node.sub_option">
                <div class="flex items-center gap-2 mt-2">
                  <div v-for="opt in node.sub_option" :key="opt.label"
                    @click="selectNodeOption(node, opt)"
                    :class="['flex items-center gap-2 px-3 py-1 rounded-full text-xs cursor-pointer', opt.selected ? 'bg-[#0D7C7C] text-white/95' : 'bg-white border border-[#E1E6EA] text-[#4A5259]']">
                    <Icon :icon="opt.selected ? 'lucide:check-circle' : 'lucide:circle'" class="text-xs" :class="opt.selected ? '' : 'text-[#9BA3AB]'" />
                    {{ opt.label }}
                  </div>
                </div>
              </template>
              <template v-if="node.sub_input">
                <div class="mt-2">
                  <label class="text-xs text-[#4A5259]">{{ node.sub_input.label }}</label>
                  <input type="text" v-model="node.sub_input.value" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base mt-1">
                </div>
              </template>
            </div>
            <label class="flex items-center gap-3">
              <div class="relative">
                <input type="checkbox" class="sr-only peer" v-model="node.enabled">
                <div class="w-12 h-6 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition"></div>
                <div class="absolute top-0.5 left-0.5 w-5 h-5 bg-white rounded-full shadow-sm peer-checked:translate-x-6 transition"></div>
              </div>
            </label>
          </div>
        </div>
      </div>

      <!-- AI Behavior Preferences Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">AI行为偏好</h3>
          <p class="text-sm text-[#6B7680] mt-1">调整AI决策时的偏好</p>
        </div>
        <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <span class="text-sm font-semibold text-[#1A1F24]">技术选型偏好</span>
          <div class="flex items-center gap-4 mt-3">
            <label class="flex items-center gap-2">
              <input type="radio" name="tech-pref" class="sr-only peer" value="conservative" v-model="config.tech_preference">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.tech_preference === 'conservative' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">倾向保守</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="tech-pref" class="sr-only peer" value="neutral" v-model="config.tech_preference">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.tech_preference === 'neutral' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">中立</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="tech-pref" class="sr-only peer" value="innovative" v-model="config.tech_preference">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.tech_preference === 'innovative' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">倾向创新</span>
            </label>
          </div>
        </div>
        <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <span class="text-sm font-semibold text-[#1A1F24]">代码风格偏好</span>
          <div class="flex items-center gap-4 mt-3">
            <label class="flex items-center gap-2">
              <input type="radio" name="code-style" class="sr-only peer" value="concise" v-model="config.code_style">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.code_style === 'concise' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">简洁优先</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="code-style" class="sr-only peer" value="detailed" v-model="config.code_style">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.code_style === 'detailed' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">注释详尽</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="code-style" class="sr-only peer" value="performance" v-model="config.code_style">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.code_style === 'performance' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">性能优先</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="code-style" class="sr-only peer" value="readable" v-model="config.code_style">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.code_style === 'readable' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">可读性优先</span>
            </label>
          </div>
        </div>
        <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <span class="text-sm font-semibold text-[#1A1F24]">测试覆盖率要求</span>
          <div class="flex items-center gap-4 mt-3">
            <div class="flex-1">
              <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full">
                <div class="h-2 bg-[#0D7C7C] rounded-full transition" :style="{ width: config.test_coverage + '%' }"></div>
              </div>
            </div>
            <span class="text-sm font-semibold text-[#0D7C7C]">{{ config.test_coverage }}% (推荐)</span>
          </div>
          <div class="flex items-center justify-between text-xs text-[#9BA3AB] mt-1">
            <span>60%</span>
            <span>100%</span>
          </div>
        </div>
      </div>

      <!-- Timeout & Retry Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">超时与重试配置</h3>
        </div>
        <div class="grid grid-cols-2 gap-3">
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">需求分析超时(分钟)</label>
            <input type="number" v-model.number="config.analysis_timeout" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">架构设计超时(分钟)</label>
            <input type="number" v-model.number="config.architecture_timeout" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">代码开发超时(分钟)</label>
            <input type="number" v-model.number="config.development_timeout" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">测试执行超时(分钟)</label>
            <input type="number" v-model.number="config.test_execution_timeout" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
        </div>
        <div class="grid grid-cols-2 gap-3">
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">自动重试次数</label>
            <input type="number" v-model.number="config.auto_retry_count" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">重试间隔(分钟)</label>
            <input type="number" v-model.number="config.retry_interval" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
        </div>
        <div>
          <span class="text-sm text-[#4A5259]">连续失败后的处理</span>
          <div class="flex items-center gap-4 mt-2">
            <label class="flex items-center gap-2">
              <input type="radio" name="failure-handling" class="sr-only peer" value="notify_human" v-model="config.failure_handling">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.failure_handling === 'notify_human' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">自动通知人工介入</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="failure-handling" class="sr-only peer" value="pause_task" v-model="config.failure_handling">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.failure_handling === 'pause_task' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">暂停任务等待处理</span>
            </label>
            <label class="flex items-center gap-2">
              <input type="radio" name="failure-handling" class="sr-only peer" value="mark_failed" v-model="config.failure_handling">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95">
                <svg fill="currentColor" class="w-2.5 h-2.5"><circle cx="10" cy="10" r="5"/></svg>
              </div>
              <span :class="['text-sm', config.failure_handling === 'mark_failed' ? 'text-[#0D7C7C] font-semibold' : 'text-[#4A5259]']">标记失败并跳过</span>
            </label>
          </div>
        </div>
      </div>

    </main>
  </div>

  <!-- Bottom Save Bar -->
  <div v-if="hasChanges" class="fixed bottom-0 left-0 right-0 bg-[#FEF6E8] border-t border-[#F59D0D] px-8 py-3 flex items-center justify-between">
    <span class="text-sm text-[#F59D0D] font-semibold flex items-center gap-2">
      <Icon icon="lucide:alert-circle" class="text-base text-[#F59D0D]" />
      您有未保存的更改
    </span>
    <div class="flex items-center gap-3">
      <button @click="resetConfig" class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
        <Icon icon="lucide:x" class="text-base" />
        <span class="whitespace-nowrap">取消</span>
      </button>
      <button @click="saveConfig" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
        <Icon icon="lucide:save" class="text-base" />
        <span class="whitespace-nowrap">保存更改</span>
      </button>
    </div>
  </div>
  <Toast ref="toastRef" />
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import SettingsSidebar from '../../layout/SettingsSidebar.vue'
import Toast from '../../components/Toast.vue'
import { getAutomationConfig, updateAutomationConfig, getReviewNodes, updateReviewNode } from '../../api/settings'

const toastRef = ref(null)

const automationLevels = [
  { key: 'conservative', label: '保守' },
  { key: 'balanced', label: '平衡(推荐)' },
  { key: 'aggressive', label: '激进' },
  { key: 'full-auto', label: '全自动' },
]

const levelDescriptions = {
  conservative: '所有关键节点都需要人工审核,每一步都需要人工确认后才执行。最安全但效率最低。',
  balanced: '关键节点需要人工审核(架构设计、代码合并、部署发布),其余环节AI自动执行。在保证质量的前提下最大化自动化效率。',
  aggressive: '仅在部署发布时需要人工审核,其余环节AI自动执行。追求效率,风险相对可控。',
  'full-auto': 'AI全程自动执行,无需任何人工干预。效率最高,但需要完善的监控和回滚机制。',
}

const levelProgressMap = { conservative: 25, balanced: 50, aggressive: 75, 'full-auto': 100 }

const loading = ref(false)
const saving = ref(false)
const hasChanges = ref(false)

const config = reactive({
  level: 'balanced',
  tech_preference: 'neutral',
  code_style: 'concise',
  test_coverage: 80,
  analysis_timeout: 30,
  architecture_timeout: 60,
  development_timeout: 120,
  test_execution_timeout: 60,
  auto_retry_count: 2,
  retry_interval: 5,
  failure_handling: 'notify_human',
})

const originalConfig = reactive({ ...config })

const nodes = ref([])

const levelProgress = computed(() => levelProgressMap[config.level] || 50)

const currentLevelLabel = computed(() => {
  const level = automationLevels.find(l => l.key === config.level)
  return level ? level.label : '未知'
})

const currentLevelDescription = computed(() => levelDescriptions[config.level] || '')

function selectNodeOption(node, opt) {
  if (node.sub_option) {
    node.sub_option.forEach(o => o.selected = o.label === opt.label)
  }
  hasChanges.value = true
}

function markChanged() {
  hasChanges.value = true
}

// Watch for config changes
const configFields = Object.keys(config)
for (const field of configFields) {
  // We rely on v-model to update config fields; the save bar visibility
  // is checked by comparing config to originalConfig in hasChanges logic.
  // Since reactive objects are mutated in place, we detect changes manually.
}

function checkHasChanges() {
  for (const field of configFields) {
    if (config[field] !== originalConfig[field]) {
      hasChanges.value = true
      return
    }
  }
  // Check nodes for changes
  for (const node of nodes.value) {
    if (node._originalEnabled !== node.enabled) {
      hasChanges.value = true
      return
    }
    if (node.sub_option) {
      for (const opt of node.sub_option) {
        if (opt._originalSelected !== opt.selected) {
          hasChanges.value = true
          return
        }
      }
    }
    if (node.sub_input) {
      if (node.sub_input._originalValue !== node.sub_input.value) {
        hasChanges.value = true
        return
      }
    }
  }
  hasChanges.value = false
}

// Node display mapping: API fields -> display properties
const nodeTitleMap = {
  requirement_confirm: '需求理解确认',
  architecture_review: '架构方案审核',
  code_merge_review: '代码合并审核',
  test_confirm: '测试结果确认',
  deploy_review: '集群部署审核',
}

const nodeImpactMap = {
  requirement_confirm: '避免需求理解偏差导致后续工作无效',
  architecture_review: '确保技术选型符合团队技术栈和长期规划',
  code_merge_review: '保证代码质量,防止安全漏洞',
  test_confirm: '增加一道质量保障,但会延长交付时间',
  deploy_review: '避免误操作导致生产事故',
}

function mapNode(apiNode) {
  const mapped = {
    id: apiNode.id,
    name: apiNode.name,
    title: nodeTitleMap[apiNode.name] || apiNode.name,
    description: apiNode.description,
    impact: nodeImpactMap[apiNode.name] || '',
    enabled: apiNode.enabled,
    _originalEnabled: apiNode.enabled,
    sub_option: null,
    sub_input: null,
  }

  if (apiNode.sub_option) {
    mapped.sub_option = apiNode.sub_option.map(opt => ({
      label: opt.label,
      selected: opt.selected,
      _originalSelected: opt.selected,
    }))
  }

  if (apiNode.sub_input) {
    mapped.sub_input = {
      label: apiNode.sub_input.label,
      value: apiNode.sub_input.value,
      _originalValue: apiNode.sub_input.value,
    }
  }

  return mapped
}

async function fetchData() {
  loading.value = true
  try {
    const [configData, nodesData] = await Promise.all([
      getAutomationConfig(),
      getReviewNodes(),
    ])

    // Populate config from API
    Object.assign(config, configData)
    Object.assign(originalConfig, configData)

    // Populate review nodes from API
    nodes.value = (nodesData || []).map(mapNode)

    hasChanges.value = false
  } catch (err) {
    console.error('Failed to fetch settings:', err)
  } finally {
    loading.value = false
  }
}

async function saveConfig() {
  saving.value = true
  try {
    // Save automation config
    await updateAutomationConfig({
      level: config.level,
      tech_preference: config.tech_preference,
      code_style: config.code_style,
      test_coverage: config.test_coverage,
      analysis_timeout: config.analysis_timeout,
      architecture_timeout: config.architecture_timeout,
      development_timeout: config.development_timeout,
      test_execution_timeout: config.test_execution_timeout,
      auto_retry_count: config.auto_retry_count,
      retry_interval: config.retry_interval,
      failure_handling: config.failure_handling,
    })

    // Save changed review nodes
    for (const node of nodes.value) {
      const payload = {
        name: node.name,
        description: node.description,
        enabled: node.enabled,
      }
      if (node.sub_option) {
        payload.sub_option = node.sub_option.map(opt => ({
          label: opt.label,
          selected: opt.selected,
        }))
      }
      if (node.sub_input) {
        payload.sub_input = {
          label: node.sub_input.label,
          value: node.sub_input.value,
        }
      }
      await updateReviewNode(node.id, payload)
    }

    // Update originals so hasChanges resets
    Object.assign(originalConfig, config)
    for (const node of nodes.value) {
      node._originalEnabled = node.enabled
      if (node.sub_option) {
        node.sub_option.forEach(opt => opt._originalSelected = opt.selected)
      }
      if (node.sub_input) {
        node.sub_input._originalValue = node.sub_input.value
      }
    }

    hasChanges.value = false
    toastRef.value?.success('设置已保存')
  } catch (err) {
    console.error('Failed to save settings:', err)
    toastRef.value?.error('保存失败，请重试')
  } finally {
    saving.value = false
  }
}

function resetConfig() {
  Object.assign(config, originalConfig)
  for (const node of nodes.value) {
    node.enabled = node._originalEnabled
    if (node.sub_option) {
      node.sub_option.forEach(opt => opt.selected = opt._originalSelected)
    }
    if (node.sub_input) {
      node.sub_input.value = node.sub_input._originalValue
    }
  }
  hasChanges.value = false
}

// Override v-model reactive tracking to detect changes
const originalSetHasChanges = () => { hasChanges.value = true }

// We use a custom approach: since v-model on reactive doesn't trigger watchers easily,
// we add @change handlers where possible. For number inputs we rely on the save button.
// The bottom bar will show based on hasChanges ref which is set by interactions.

onMounted(fetchData)
</script>