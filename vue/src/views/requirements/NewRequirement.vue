<template>
  <TopNavBar activeNav="requirements" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-12 py-6">
      <div class="flex items-center gap-2 text-sm mb-4">
        <router-link to="/requirements" class="text-[#0D7C7C] hover:opacity-80 transition">需求管理</router-link>
        <span class="text-[#6B7680]">/</span>
        <span class="text-[#1A1F24]">新建需求</span>
      </div>
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">新建需求</h1>
        <router-link to="/requirements" class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
          <Icon icon="lucide:arrow-left" class="text-sm" />
          <span class="whitespace-nowrap text-sm">返回列表</span>
        </router-link>
      </div>

      <!-- Section 1 -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-5 mb-6">
        <div class="flex items-center gap-3">
          <Icon icon="lucide:pencil-line" class="text-base text-[#0D7C7C]" />
          <h2 class="text-lg font-semibold text-[#1A1F24]">描述您的需求</h2>
          <span class="bg-[#E8F0FE] text-[#3367D6] px-3 py-1 rounded-full text-xs font-semibold">自然语言输入</span>
        </div>
        <p class="text-sm text-[#6B7680]">用自然语言描述产品需求,AI将自动分析并生成结构化需求文档和开发计划</p>
        <textarea class="w-full h-[160px] bg-[hsla(200,15%,95%,1)] border border-[#E1E6EA] rounded-xl px-4 py-3 text-base text-[#1A1F24] placeholder-[#9BA3AB] resize-none focus:border-[#0D7C7C] focus:outline-none transition"
          placeholder="请用自然语言描述您的产品需求..."
          v-model="requirementText"></textarea>
        <div class="flex items-center justify-between">
          <span class="text-sm text-[#9BA3AB]">建议包含: 功能目标、用户场景、关键约束条件</span>
          <span class="text-sm text-[#9BA3AB]">{{ requirementText.length }} / 2000 字</span>
        </div>

        <div class="flex flex-col gap-2">
          <span class="text-sm font-semibold text-[#1A1F24]">优先级</span>
          <div class="flex items-center gap-4">
            <label v-for="p in priorities" :key="p.label" class="flex items-center gap-2 cursor-pointer">
              <input type="radio" name="priority" class="sr-only peer" :checked="p.label === '高'">
              <div class="w-5 h-5 bg-[hsla(200,15%,95%,1)] rounded-full flex items-center justify-center peer-checked:bg-[#0D7C7C] text-transparent peer-checked:text-white/95 transition">
                <svg class="w-2 h-2" viewBox="0 0 8 8" fill="currentColor"><circle cx="4" cy="4" r="4"/></svg>
              </div>
              <span class="text-sm text-[#1A1F24]">{{ p.label }}</span>
              <span v-if="p.badge" :class="p.badgeClass" class="px-2 py-0.5 rounded-full text-xs font-semibold">{{ p.badge }}</span>
            </label>
          </div>
        </div>

        <div class="flex items-center gap-3 mt-2">
          <button class="flex items-center gap-2 px-6 py-3 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:sparkles" class="text-base" />
            <span class="whitespace-nowrap font-semibold">提交需求</span>
          </button>
          <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-sm" />
            <span class="whitespace-nowrap text-sm">保存草稿</span>
          </button>
        </div>
      </div>

      <!-- Section 2: AI Dialog -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-5 mb-6">
        <div class="flex items-center gap-3">
          <Icon icon="lucide:bot" class="text-base text-[#0D7C7C]" />
          <h2 class="text-lg font-semibold text-[#1A1F24]">AI需求澄清对话</h2>
          <span class="bg-[#E8F7F0] text-[#0F8B5D] px-3 py-1 rounded-full text-xs font-semibold">分析中</span>
        </div>
        <div class="flex flex-col gap-4">
          <div v-for="(q, idx) in aiQuestions" :key="idx" class="flex flex-col gap-3 bg-[hsla(200,15%,95%,1)] rounded-xl p-4">
            <div class="flex items-center gap-2">
              <div class="flex items-center justify-center bg-[#0D7C7C] w-8 h-8 rounded-full">
                <Icon icon="lucide:bot" class="text-base text-white/95" />
              </div>
              <span class="text-sm font-semibold text-[#0D7C7C]">需求分析师</span>
              <span class="text-xs text-[#9BA3AB]">提问 {{ idx + 1 }}/3</span>
            </div>
            <div class="flex items-center gap-2 bg-white rounded-lg px-4 py-3">
              <Icon icon="lucide:help-circle" class="text-base text-[#3367D6]" />
              <span class="text-sm text-[#1A1F24]">{{ q.question }}</span>
              <span class="text-xs text-[#6B7680]">{{ q.detail }}</span>
            </div>
            <div v-if="q.answer" class="flex items-start gap-2 ml-8">
              <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=sarah" class="w-6 h-6 rounded-full mt-0.5" />
              <div class="flex flex-col gap-1 bg-[hsla(200,15%,95%,1)] rounded-lg px-4 py-2">
                <span class="text-xs font-semibold text-[#0D7C7C]">Sarah</span>
                <span class="text-sm text-[#1A1F24]">{{ q.answer }}</span>
              </div>
            </div>
          </div>
        </div>
        <div class="flex items-center gap-3 mt-2">
          <button class="flex items-center gap-2 px-6 py-3 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:check-circle" class="text-base" />
            <span class="whitespace-nowrap font-semibold">完成回答并提交分析</span>
          </button>
          <span class="text-sm text-[#9BA3AB] flex items-center gap-2">
            <Icon icon="lucide:info" class="text-xs" /> 已回答 2/3 个问题
          </span>
        </div>
      </div>

      <!-- Flow Steps -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center gap-4 bg-[hsla(200,15%,95%,1)] rounded-xl p-4">
          <div v-for="(step, idx) in flowSteps" :key="idx" class="flex items-center gap-2">
            <div :class="['w-6 h-6 rounded-full flex items-center justify-center', step.bg]">
              <Icon :icon="step.icon" class="text-xs" :class="step.textColor" />
            </div>
            <span :class="['text-sm', step.textClass]">{{ step.label }}</span>
            <div v-if="idx < flowSteps.length - 1" class="flex-1 h-0.5" :class="step.lineClass"></div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'

const requirementText = ref('用户注册流程中需要增加手机号短信验证码验证功能，确保用户身份真实性，减少虚假账号注册。支持中国大陆手机号格式，验证码为6位数字，有效期5分钟。用户输入手机号后点击"获取验证码"按钮，系统发送短信，用户输入验证码完成验证。验证成功后自动跳转到下一步注册流程。')

const priorities = [
  { label: '高', badge: '紧急', badgeClass: 'bg-[#FDECEA] text-[#D93025]' },
  { label: '中', badge: null },
  { label: '低', badge: null },
]

const aiQuestions = [
  { question: '验证码有效期?', detail: '您的描述中提到"有效期5分钟",是否确认验证码有效期为5分钟?是否需要提供过期后重新发送的机制?', answer: '确认验证码有效期为5分钟。过期后用户需要重新获取验证码,原验证码自动失效。建议在UI上显示倒计时。' },
  { question: '是否支持国际号码?', detail: '您提到"支持中国大陆手机号格式",是否需要考虑海外用户的国际号码验证?', answer: '当前版本只支持中国大陆手机号(+86),国际号码支持作为v2.0的后续需求。架构上需要预留国际号码扩展能力。' },
  { question: '发送失败重试机制?', detail: '短信发送可能因网络或服务商问题失败,是否需要自动重试机制?', answer: null },
]

const flowSteps = [
  { label: '需求描述', icon: 'lucide:check', bg: 'bg-[#0F8B5D]', textColor: 'text-white/95', textClass: 'text-[#0F8B5D] font-semibold', lineClass: 'bg-[#0F8B5D]' },
  { label: 'AI澄清', icon: 'lucide:check', bg: 'bg-[#0F8B5D]', textColor: 'text-white/95', textClass: 'text-[#0F8B5D] font-semibold', lineClass: 'bg-[#0F8B5D]' },
  { label: '文档生成', icon: 'lucide:file-check', bg: 'bg-[#0D7C7C]', textColor: 'text-white/95', textClass: 'text-[#0D7C7C] font-semibold', lineClass: 'bg-[#E1E6EA]' },
  { label: '架构审核', icon: 'lucide:blocks', bg: 'bg-[hsla(200,15%,95%,1)]', textColor: 'text-[#9BA3AB]', textClass: 'text-[#9BA3AB]', lineClass: 'bg-[#E1E6EA]' },
  { label: '开发', icon: 'lucide:code-2', bg: 'bg-[hsla(200,15%,95%,1)]', textColor: 'text-[#9BA3AB]', textClass: 'text-[#9BA3AB]', lineClass: 'bg-[#E1E6EA]' },
]
</script>
