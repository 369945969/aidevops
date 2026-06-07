<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="ai-agent" />
    <main class="flex-1 overflow-x-hidden overflow-y-auto flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">AI Team 配置</h1>
        <div class="flex items-center gap-3">
          <button @click="handleSave" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Team Overview Stats Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">Team概览</h3>
          <p class="text-sm text-[#6B7680] mt-1">6个AI Team，12个Agent运行中</p>
        </div>
        <div class="grid grid-cols-4 gap-3">
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-4 py-3 flex flex-col items-center gap-1">
            <Icon icon="lucide:users-round" class="text-lg text-[#0D7C7C]" />
            <span class="text-xs text-[#6B7680]">活跃Team</span>
            <span class="text-sm font-semibold text-[#1A1F24]">5 / 6</span>
          </div>
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-4 py-3 flex flex-col items-center gap-1">
            <Icon icon="lucide:bot" class="text-lg text-[#0D7C7C]" />
            <span class="text-xs text-[#6B7680]">Agent总数</span>
            <span class="text-sm font-semibold text-[#1A1F24]">12</span>
          </div>
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-4 py-3 flex flex-col items-center gap-1">
            <Icon icon="lucide:vote" class="text-lg text-[#0D7C7C]" />
            <span class="text-xs text-[#6B7680]">全票通过率</span>
            <span class="text-sm font-semibold text-[#1A1F24]">78%</span>
          </div>
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-4 py-3 flex flex-col items-center gap-1">
            <Icon icon="lucide:link" class="text-lg text-[#0D7C7C]" />
            <span class="text-xs text-[#6B7680]">人员绑定</span>
            <span class="text-sm font-semibold text-[#1A1F24]">5 / 6</span>
          </div>
        </div>
      </div>

      <!-- Expanded Team: 需求分析师Team -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col mb-6">
        <div class="px-5 py-4 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="flex items-center justify-center w-8 h-8 bg-[#0D7C7C] rounded-lg">
              <Icon icon="lucide:file-text" class="text-base text-white/95" />
            </div>
            <span class="text-sm font-semibold text-[#1A1F24]">需求分析师 Team</span>
            <div class="flex items-center gap-2 px-3 py-1 bg-[#0D7C7C] text-white/95 rounded-full text-xs">
              <Icon icon="lucide:check-circle" class="text-xs" />
              运行中
            </div>
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=sarah" class="w-6 h-6 rounded-full border-2 border-white ml-2" />
            <div class="flex items-center gap-2 px-3 py-1 bg-[hsla(200,15%,95%,1)] border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
              <Icon icon="lucide:user" class="text-xs text-[#0D7C7C]" />
              Sarah Wang
            </div>
          </div>
          <div class="flex items-center gap-3">
            <div class="flex items-center gap-2 text-xs text-[#4A5259]">
              <Icon icon="lucide:vote" class="text-xs text-[#0D7C7C]" />
              全票通过
            </div>
            <div class="w-24">
              <div class="h-1.5 bg-[hsla(200,15%,95%,1)] rounded-full"><div class="h-1.5 bg-[#0D7C7C] rounded-full" style="width: 67%;"></div></div>
            </div>
            <span class="text-xs font-semibold text-[#0D7C7C]">2/3</span>
          </div>
          <div class="flex items-center gap-4">
            <label class="flex items-center gap-2">
              <span class="text-xs text-[#6B7680]">启用</span>
              <div class="relative">
                <input type="checkbox" class="sr-only peer" checked>
                <div class="w-9 h-5 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition"></div>
                <div class="absolute top-0.5 left-0.5 w-4 h-4 bg-white rounded-full shadow-sm peer-checked:translate-x-4 transition"></div>
              </div>
            </label>
            <Icon icon="lucide:chevron-up" class="text-base text-[#6B7680]" />
          </div>
        </div>

        <div class="px-5 pb-5 space-y-3">
          <!-- 撰写 Agent -->
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4 border-l-4 border-[#0D7C7C]">
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-3">
                <div class="flex items-center gap-2 px-3 py-1 bg-[#0D7C7C] text-white/95 rounded-full text-xs font-semibold">
                  <Icon icon="lucide:pen-line" class="text-xs" />
                  撰写
                </div>
                <span class="text-sm font-semibold text-[#1A1F24]">需求分析撰写</span>
                <div class="flex items-center gap-2 px-3 py-1 bg-[#0D7C7C] text-white/95 rounded-full text-xs">
                  <Icon icon="lucide:check-circle" class="text-xs" />
                  运行中
                </div>
              </div>
            </div>
            <!-- Primary Model -->
            <div class="mt-3">
              <span class="text-sm font-semibold text-[#1A1F24]">主模型</span>
              <div class="flex gap-2 mt-2 overflow-x-auto">
                <label v-for="model in primaryModels" :key="model.name" class="flex items-center gap-2">
                  <input type="radio" name="req-writer-primary-model" class="sr-only peer" :checked="model.selected">
                  <div :class="['bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full peer-checked:bg-[#0D7C7C] peer-checked:text-white/95 hover:opacity-80 transition whitespace-nowrap flex items-center gap-2', model.selected ? 'bg-[#0D7C7C] text-white/95' : '']">
                    <Icon :icon="model.icon" class="text-sm" />
                    {{ model.name }}
                  </div>
                </label>
              </div>
            </div>
            <!-- Fallback Model -->
            <div class="mt-3">
              <div class="flex items-center gap-2">
                <span class="text-sm font-semibold text-[#1A1F24]">备模型</span>
                <div class="flex items-center gap-2 px-3 py-1 bg-[#FEF6E8] border border-[#F59D0D] rounded-full text-xs text-[#F59D0D]">
                  <Icon icon="lucide:alert-triangle" class="text-xs" />
                  主模型不可用时自动切换
                </div>
              </div>
              <div class="flex gap-2 mt-2 overflow-x-auto">
                <label v-for="model in fallbackModels" :key="model.name" class="flex items-center gap-2">
                  <input type="radio" name="req-writer-fallback-model" class="sr-only peer" :checked="model.selected">
                  <div :class="['bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-4 py-2 rounded-full hover:opacity-80 transition whitespace-nowrap flex items-center gap-2', model.selected ? 'bg-[#5B7C92] text-white/95' : '']">
                    <Icon :icon="model.icon" class="text-sm" />
                    {{ model.name }}
                  </div>
                </label>
              </div>
            </div>
            <!-- Stats -->
            <div class="flex items-center gap-4 mt-3">
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:check-circle" class="text-xs text-[#0D7C7C]" />
                任务完成: 42
              </div>
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:clock" class="text-xs text-[#9BA3AB]" />
                平均时长: 15分钟
              </div>
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:trending-up" class="text-xs text-[#0D7C7C]" />
                成功率: 97%
              </div>
            </div>
          </div>

          <!-- 审查 Agent -->
          <div class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4 border-l-4 border-[#5B7C92]">
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-3">
                <div class="flex items-center gap-2 px-3 py-1 bg-[#5B7C92] text-white/95 rounded-full text-xs font-semibold">
                  <Icon icon="lucide:scan-search" class="text-xs" />
                  审查
                </div>
                <span class="text-sm font-semibold text-[#1A1F24]">需求分析审查</span>
                <div class="flex items-center gap-2 px-3 py-1 bg-[#5B7C92] text-white/95 rounded-full text-xs">
                  <Icon icon="lucide:check-circle" class="text-xs" />
                  运行中
                </div>
              </div>
            </div>
            <!-- Stats -->
            <div class="flex items-center gap-4 mt-3">
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:check-circle" class="text-xs text-[#5B7C92]" />
                任务完成: 38
              </div>
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:clock" class="text-xs text-[#9BA3AB]" />
                平均时长: 10分钟
              </div>
              <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
                <Icon icon="lucide:trending-up" class="text-xs text-[#5B7C92]" />
                成功率: 95%
              </div>
            </div>
          </div>

          <!-- Add Agent Button -->
          <div class="flex items-center justify-center mt-2">
            <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] border border-[#E1E6EA] rounded-full text-sm text-[#4A5259] hover:opacity-80 transition">
              <Icon icon="lucide:plus" class="text-base" />
              <span class="whitespace-nowrap">添加Agent</span>
            </button>
          </div>

          <!-- 全票通过机制 -->
          <div class="bg-[hsla(200,20%,98%,1)] border border-[#0D7C7C] rounded-xl p-4 mt-2">
            <div class="flex items-center justify-between mb-3">
              <div class="flex items-center gap-2">
                <Icon icon="lucide:vote" class="text-base text-[#0D7C7C]" />
                <span class="text-sm font-semibold text-[#0D7C7C]">全票通过机制</span>
              </div>
              <span class="text-xs text-[#6B7680]">Team输出需所有Agent确认后方可采纳</span>
            </div>
            <div class="flex items-center gap-4">
              <div class="flex items-center gap-2">
                <div class="w-2.5 h-2.5 bg-[#0D7C7C] rounded-full"></div>
                <span class="text-xs text-[#1A1F24] font-semibold">撰写Agent</span>
                <Icon icon="lucide:check-circle" class="text-xs text-[#0F8B5D]" />
                <span class="text-xs text-[#0F8B5D]">已确认</span>
              </div>
              <div class="flex items-center gap-2">
                <div class="w-2.5 h-2.5 bg-[#5B7C92] rounded-full"></div>
                <span class="text-xs text-[#1A1F24] font-semibold">审查Agent</span>
                <Icon icon="lucide:check-circle" class="text-xs text-[#0F8B5D]" />
                <span class="text-xs text-[#0F8B5D]">已确认</span>
              </div>
              <div class="flex items-center gap-2">
                <div class="w-2.5 h-2.5 bg-[#9BA3AB] rounded-full"></div>
                <span class="text-xs text-[#9BA3AB]">补充Agent</span>
                <Icon icon="lucide:circle" class="text-xs text-[#9BA3AB]" />
                <span class="text-xs text-[#9BA3AB]">待确认</span>
              </div>
            </div>
          </div>

          <!-- 负责人绑定 -->
          <div class="border-t border-[#E1E6EA] pt-4 mt-3">
            <div class="flex items-center justify-between">
              <div class="flex items-center gap-2">
                <Icon icon="lucide:link" class="text-base text-[#0D7C7C]" />
                <span class="text-sm font-semibold text-[#1A1F24]">负责人绑定</span>
              </div>
              <router-link to="/settings/team-members" class="text-xs text-[#0D7C7C] hover:text-[#096363]">管理团队成员 &rarr;</router-link>
            </div>
            <div class="flex items-center gap-4 mt-2">
              <div class="flex items-center gap-3 px-4 py-2 bg-white border border-[#E1E6EA] rounded-xl">
                <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=sarah" class="w-8 h-8 rounded-full">
                <div class="flex flex-col gap-0.5">
                  <span class="text-sm font-semibold text-[#1A1F24]">Sarah Wang</span>
                  <span class="text-xs text-[#6B7680]">产品经理</span>
                </div>
                <div class="flex items-center gap-2 px-2 py-1 bg-[#0D7C7C] text-white/95 rounded-full text-xs">
                  <Icon icon="lucide:check-circle" class="text-xs" />
                  在线
                </div>
              </div>
            </div>
          </div>

          <!-- Team Stats Footer -->
          <div class="flex items-center gap-4 mt-3 border-t border-[#E1E6EA] pt-3">
            <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
              <Icon icon="lucide:check-circle" class="text-xs text-[#0D7C7C]" />
              Team任务: 42
            </div>
            <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
              <Icon icon="lucide:clock" class="text-xs text-[#9BA3AB]" />
              平均时长: 20分钟
            </div>
            <div class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
              <Icon icon="lucide:trending-up" class="text-xs text-[#0D7C7C]" />
              成功率: 93%
            </div>
          </div>
        </div>
      </div>

      <!-- Collapsed Teams -->
      <div v-for="team in collapsedTeams" :key="team.key" class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col mb-6">
        <div class="px-5 py-4 flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="flex items-center justify-center w-8 h-8 bg-[#0D7C7C] rounded-lg">
              <Icon :icon="team.icon" class="text-base text-white/95" />
            </div>
            <span class="text-sm font-semibold text-[#1A1F24]">{{ team.name }}</span>
            <div :class="['flex items-center gap-2 px-3 py-1 rounded-full text-xs', team.status === '运行中' ? 'bg-[#0D7C7C] text-white/95' : 'bg-white border border-[#E1E6EA] text-[#9BA3AB]']">
              <Icon :icon="team.status === '运行中' ? 'lucide:check-circle' : 'lucide:pause'" class="text-xs" />
              {{ team.status }}
            </div>
            <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${team.avatarSeed}`" class="w-6 h-6 rounded-full border-2 border-white ml-2" />
            <div class="flex items-center gap-2 px-3 py-1 bg-[hsla(200,15%,95%,1)] border border-[#E1E6EA] rounded-full text-xs text-[#4A5259]">
              <Icon icon="lucide:user" class="text-xs text-[#0D7C7C]" />
              {{ team.boundPerson }}
            </div>
          </div>
          <div class="flex items-center gap-3">
            <div class="flex items-center gap-2 text-xs text-[#4A5259]">
              <Icon :icon="team.status === '运行中' ? 'lucide:vote' : 'lucide:vote'" :class="team.status === '运行中' ? 'text-xs text-[#0D7C7C]' : 'text-xs text-[#9BA3AB]'" />
              全票通过
            </div>
            <div class="w-24">
              <div class="h-1.5 bg-[hsla(200,15%,95%,1)] rounded-full"><div :class="['h-1.5 rounded-full', team.status === '运行中' ? 'bg-[#0D7C7C]' : 'bg-[#9BA3AB]']" :style="{ width: team.consensusPercent + '%' }"></div></div>
            </div>
            <span :class="['text-xs font-semibold', team.status === '运行中' ? 'text-[#0D7C7C]' : 'text-[#9BA3AB]']">{{ team.consensus }}</span>
          </div>
          <div class="flex items-center gap-4">
            <label class="flex items-center gap-2">
              <span class="text-xs text-[#6B7680]">启用</span>
              <div class="relative">
                <input type="checkbox" class="sr-only peer" checked>
                <div class="w-9 h-5 bg-[hsla(200,15%,95%,1)] rounded-full peer-checked:bg-[#0D7C7C] transition"></div>
                <div class="absolute top-0.5 left-0.5 w-4 h-4 bg-white rounded-full shadow-sm peer-checked:translate-x-4 transition"></div>
              </div>
            </label>
            <Icon icon="lucide:chevron-down" class="text-base text-[#6B7680]" />
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
      <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
        <Icon icon="lucide:save" class="text-base" />
        <span class="whitespace-nowrap">保存更改</span>
      </button>
    </div>
  </div>
  <Toast ref="toastRef" />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import SettingsSidebar from '../../layout/SettingsSidebar.vue'
import Toast from '../../components/Toast.vue'
import { getAITeams, updateAITeam } from '../../api/settings'

const toastRef = ref(null)
const teams = ref([])
const primaryModels = [
  { name: 'Claude Opus', icon: 'lucide:brain', selected: true },
  { name: 'GPT-4o', icon: 'lucide:sparkles', selected: false },
  { name: 'Gemini Pro', icon: 'lucide:diamond', selected: false },
  { name: 'DeepSeek', icon: 'lucide:server', selected: false },
]

const fallbackModels = [
  { name: 'GPT-4o', icon: 'lucide:sparkles', selected: true },
  { name: 'Gemini Pro', icon: 'lucide:diamond', selected: false },
  { name: 'DeepSeek', icon: 'lucide:server', selected: false },
]

const collapsedTeams = [
  { key: 'architect', name: '架构师 Team', icon: 'lucide:layout', status: '运行中', avatarSeed: 'alex', boundPerson: 'Alex Chen', consensus: '2/2', consensusPercent: 100 },
  { key: 'frontend', name: '前端开发 Team', icon: 'lucide:monitor', status: '运行中', avatarSeed: 'mike', boundPerson: 'Mike Liu', consensus: '1/2', consensusPercent: 50 },
  { key: 'backend', name: '后端开发 Team', icon: 'lucide:database', status: '运行中', avatarSeed: 'emma', boundPerson: 'Emma Zhang', consensus: '2/2', consensusPercent: 100 },
  { key: 'testing', name: '测试 Team', icon: 'lucide:test-tubes', status: '待命', avatarSeed: 'alex', boundPerson: 'Alex Chen', consensus: '0/2', consensusPercent: 0 },
  { key: 'devops', name: 'DevOps Team', icon: 'lucide:server', status: '运行中', avatarSeed: 'david', boundPerson: 'David Li', consensus: '2/2', consensusPercent: 100 },
]

async function handleSave() {
  try {
    for (const team of teams.value) {
      await updateAITeam(team.id, { enabled: team.enabled })
    }
    toastRef.value?.success('AI Team 配置已保存')
  } catch (err) {
    console.error('Failed to save AI agent settings:', err)
    toastRef.value?.error('保存失败，请重试')
  }
}

onMounted(async () => {
  try {
    const data = await getAITeams()
    if (data) teams.value = data
  } catch (err) {
    console.error('Failed to fetch AI teams:', err)
  }
})
</script>