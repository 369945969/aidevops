<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <!-- Left: File Tree -->
    <aside class="flex-shrink-0 min-w-fit w-[200px] bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 m-3 overflow-y-auto">
      <div class="flex items-center gap-2">
        <Icon icon="lucide:folder-git-2" class="text-base text-[#0D7C7C]" />
        <span class="text-sm font-semibold text-[#1A1F24]">变更文件</span>
      </div>
      <div class="flex items-center gap-2">
        <span class="px-3 py-1 text-xs bg-[#0D7C7C] text-white/95 rounded-full cursor-pointer">全部</span>
        <span class="px-3 py-1 text-xs bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full cursor-pointer">仅问题文件</span>
      </div>
      <div class="flex flex-col gap-2">
        <div v-for="file in codeReviewFiles" :key="file.name"
          :class="[
            'flex items-center gap-2 px-3 py-2 bg-[hsla(200,15%,95%,1)] rounded-lg cursor-pointer hover:shadow-[0_2px_8px_rgba(0,0,0,0.10)] transition',
            file.hasIssue ? 'border-l-2 border-[#D93025]' : ''
          ]">
          <Icon :icon="file.icon" :class="`text-base text-[${file.iconColor}]`" />
          <span class="text-sm text-[#1A1F24] truncate">{{ file.name }}</span>
          <span :class="`text-xs whitespace-nowrap text-[${file.iconColor}]`">{{ file.lines }}</span>
          <Icon v-if="file.hasIssue" icon="lucide:alert-circle" class="text-base text-[#D93025]" />
        </div>
      </div>
      <div class="flex items-center justify-between text-xs text-[#6B7680] pt-2">
        <span>5 文件变更</span>
        <span>+167 -70</span>
      </div>
    </aside>

    <!-- Center: Code Diff -->
    <main class="flex-1 overflow-x-hidden flex flex-col m-3 gap-3">
      <div class="flex items-center gap-2 text-sm">
        <router-link to="/workflow/dashboard" class="text-[#0D7C7C] hover:opacity-80">工作流看板</router-link>
        <span class="text-[#6B7680]">/</span>
        <span class="text-[#4A5259]">代码审核</span>
        <span class="text-[#6B7680]">/</span>
        <span class="text-[#1A1F24] font-semibold">TASK-007-A</span>
      </div>
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-sm font-semibold text-[#0D7C7C]">TASK-007-A</span>
            <span class="text-base font-semibold text-[#1A1F24]">短信验证码API接口</span>
            <span class="px-3 py-1 text-xs font-semibold bg-[#FEF6E8] text-[#F59D0D] rounded-full">需要修改</span>
          </div>
          <div class="flex items-center gap-3 text-sm text-[#6B7680]">
            <router-link to="/requirements/REQ-0007" class="text-[#0D7C7C] cursor-pointer hover:opacity-80">REQ-0007</router-link>
            <span>· 后端开发Agent</span>
            <span>· 提交于 2026-06-05 16:00</span>
          </div>
        </div>
      </div>

      <!-- Code Diff -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col overflow-hidden flex-1 min-h-0">
        <div class="flex items-center justify-between px-5 py-3 bg-[hsla(200,15%,96%,1)] border-b border-[hsla(200,15%,95%,1)]">
          <div class="flex items-center gap-3">
            <Icon icon="lucide:file-code" class="text-base text-[#0D7C7C]" />
            <span class="text-sm font-semibold text-[#1A1F24]">src/verify/redis.store.ts</span>
            <span class="px-3 py-1 text-xs font-semibold bg-[#FDECEA] text-[#D93025] rounded-full">有问题</span>
          </div>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:columns" class="text-base" /> <span class="whitespace-nowrap">并排对比</span>
            </button>
          </div>
        </div>
        <div class="flex-1 overflow-y-auto p-5 font-mono text-sm leading-[1.8]">
          <div v-for="(line, idx) in diffLines" :key="idx"
            :class="['flex items-center gap-3 px-2 py-0.5', line.bgClass]">
            <span :class="['text-xs w-8 text-right shrink-0', line.numColor]">{{ line.num }}</span>
            <span :class="line.textColor" v-html="line.text"></span>
            <Icon v-if="line.hasIssue" icon="lucide:alert-circle" class="text-base text-[#D93025] absolute right-4 cursor-pointer" />
          </div>
        </div>
      </div>

      <!-- Annotation -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-3">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:message-square-plus" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">人工批注</span>
        </div>
        <textarea class="w-full min-h-[80px] px-3 py-2 text-sm text-[#1A1F24] bg-[hsla(200,15%,96%,1)] border border-[hsla(200,15%,95%,1)] rounded-lg resize-none focus:outline-none focus:border-[#0D7C7C] placeholder:text-[#9BA3AB]"
          placeholder="输入评论内容..."></textarea>
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-1 px-3 py-1.5 text-xs bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full hover:bg-[hsla(200,15%,90%,1)] transition">
              <Icon icon="lucide:at-sign" class="text-sm" /> <span>@提及</span>
            </button>
          </div>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80 text-sm">
            <Icon icon="lucide:send" class="text-base" /> <span>提交评论</span>
          </button>
        </div>
      </div>
    </main>

    <!-- Right: Review Report -->
    <aside class="flex-shrink-0 min-w-fit w-[300px] m-3 flex flex-col gap-3 overflow-y-auto">
      <!-- AI Review -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:bot" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">AI代码审查摘要</span>
        </div>
        <div class="flex flex-col gap-3">
          <div class="flex items-center justify-between">
            <span class="text-sm text-[#4A5259]">代码质量评分</span>
            <div class="flex items-center gap-2">
              <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full w-32"><div class="bg-[#F59D0D] h-2 rounded-full" style="width:68%"></div></div>
              <span class="text-sm font-semibold text-[#F59D0D]">68/100</span>
            </div>
          </div>
          <div class="flex items-center justify-between text-sm">
            <span class="text-[#4A5259]">测试覆盖率</span>
            <div class="flex items-center gap-2">
              <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full w-24"><div class="bg-[#F59D0D] h-2 rounded-full" style="width:45%"></div></div>
              <span class="text-sm font-semibold text-[#F59D0D]">45%</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Issues -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:list-checks" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">问题列表</span>
          <span class="px-3 py-1 text-xs font-semibold bg-[#FDECEA] text-[#D93025] rounded-full">3</span>
        </div>
        <div class="flex flex-col gap-2">
          <div v-for="issue in codeReviewIssues" :key="issue.title"
            :class="[
              'border rounded-lg px-3 py-2 cursor-pointer hover:shadow-[0_2px_8px_rgba(0,0,0,0.10)] transition',
              issue.severity === 'severe' ? 'bg-[#FDECEA] border-[#D93025]' : issue.severity === 'warning' ? 'bg-[#FEF6E8] border-[#F59D0D]' : 'bg-[#E8F0FE] border-[#3367D6]'
            ]">
            <div class="flex items-center gap-2">
              <Icon :icon="issue.severity === 'severe' ? 'lucide:alert-circle' : issue.severity === 'warning' ? 'lucide:alert-triangle' : 'lucide:lightbulb'" class="text-base" :style="{ color: issue.color }" />
              <span :class="`px-2 py-0.5 text-xs font-semibold bg-[${issue.color}] text-white rounded-full`">{{ issue.label }}</span>
            </div>
            <p class="text-xs text-[#1A1F24] mt-1">{{ issue.title }}</p>
            <span class="text-xs text-[#9BA3AB]">{{ issue.file }}</span>
          </div>
        </div>
      </div>

      <!-- Review Actions -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:git-pull-request" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">审核操作</span>
        </div>
        <div class="flex flex-col gap-2">
          <button class="flex items-center gap-2 px-4 py-2 bg-[#0F8B5D] text-white/95 rounded-full transition hover:opacity-80 opacity-50 cursor-not-allowed w-full justify-center" disabled>
            <Icon icon="lucide:check-circle" class="text-base" /> <span class="whitespace-nowrap">批准合并</span>
          </button>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#F59D0D] text-white/95 rounded-full transition hover:opacity-80 w-full justify-center">
            <Icon icon="lucide:pencil" class="text-base" /> <span class="whitespace-nowrap">要求修改</span>
          </button>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80 w-full justify-center">
            <Icon icon="lucide:x-circle" class="text-base" /> <span class="whitespace-nowrap">拒绝合并</span>
          </button>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import { codeReviewFiles, codeReviewIssues } from '../../data/mockData'

const diffLines = [
  { num: '1', numColor: 'text-[#9BA3AB]', text: 'import { RedisClient } from \'../config/redis.config\';', textColor: 'text-[#4A5259]', bgClass: '' },
  { num: '2', numColor: 'text-[#9BA3AB]', text: 'import { generateCode } from \'../utils/code-generator\';', textColor: 'text-[#4A5259]', bgClass: '' },
  { num: '+3', numColor: 'text-[#0F8B5D]', text: 'export class RedisVerifyStore {', textColor: 'text-[#0F8B5D]', bgClass: 'bg-[#E8F7F0]' },
  { num: '+4', numColor: 'text-[#0F8B5D]', text: '&nbsp;&nbsp;private readonly EXPIRY_SECONDS = 300;', textColor: 'text-[#0F8B5D]', bgClass: 'bg-[#E8F7F0]' },
  { num: '+5', numColor: 'text-[#D93025]', text: '&nbsp;&nbsp;async storeCode(phone: string, code: string) {', textColor: 'text-[#D93025]', bgClass: 'bg-[#FDECEA]', hasIssue: true },
  { num: '+6', numColor: 'text-[#0F8B5D]', text: '&nbsp;&nbsp;&nbsp;&nbsp;const key = `verify:${phone}`;', textColor: 'text-[#0F8B5D]', bgClass: 'bg-[#E8F7F0]' },
  { num: '+7', numColor: 'text-[#0F8B5D]', text: '&nbsp;&nbsp;&nbsp;&nbsp;await this.redis.set(key, code, \'EX\', this.EXPIRY_SECONDS);', textColor: 'text-[#0F8B5D]', bgClass: 'bg-[#E8F7F0]' },
  { num: '+9', numColor: 'text-[#F59D0D]', text: '&nbsp;&nbsp;async verifyCode(phone: string, code: string) {', textColor: 'text-[#F59D0D]', bgClass: 'bg-[#FEF6E8]' },
  { num: '+10', numColor: 'text-[#0F8B5D]', text: '&nbsp;&nbsp;&nbsp;&nbsp;const key = `verify:${phone}`;', textColor: 'text-[#0F8B5D]', bgClass: 'bg-[#E8F7F0]' },
  { num: '-15', numColor: 'text-[#D93025]', text: 'export class OldVerifyService {', textColor: 'text-[#D93025]', bgClass: 'bg-[#FDECEA]' },
]
</script>
