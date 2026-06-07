<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 text-sm text-[#6B7680] mb-4">
        <router-link to="/workflow/dashboard" class="text-[#0D7C7C]">工作流看板</router-link>
        <span>/</span>
        <router-link to="/workflow/task/1" class="text-[#0D7C7C]">任务详情</router-link>
        <span>/</span>
        <span class="text-[#1A1F24] font-semibold">TASK-007-A</span>
      </div>

      <!-- Task Overview -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-sm font-semibold text-[#0D7C7C]">TASK-007-A</span>
            <span class="text-lg font-semibold text-[#1A1F24]">短信验证码API接口</span>
            <StatusBadge label="开发中" />
            <PriorityBadge label="高" />
            <span class="bg-[hsla(200,15%,95%,1)] text-[#4A5259] px-3 py-1 rounded-full text-xs font-semibold">后端</span>
          </div>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 text-[#0D7C7C] rounded-full transition hover:opacity-80 text-sm border border-[#E1E6EA]">
              <Icon icon="lucide:refresh-cw" class="text-sm" />
              <span class="whitespace-nowrap">刷新状态</span>
            </button>
            <button class="flex items-center gap-2 px-4 py-2 bg-[#FEF6E8] text-[#F59D0D] rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:pause" class="text-sm" />
              <span class="whitespace-nowrap">暂停任务</span>
            </button>
          </div>
        </div>
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <div class="flex items-center justify-center w-10 h-10 rounded-full bg-[#E8F7F0]">
              <Icon icon="lucide:bot" class="text-lg text-[#0D7C7C]" />
            </div>
            <div class="flex flex-col gap-0.5">
              <span class="text-sm font-semibold text-[#1A1F24]">后端开发Agent</span>
              <span class="text-xs text-[#9BA3AB]">AI Agent 正在工作中</span>
            </div>
            <div class="flex items-center gap-2 ml-4">
              <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full w-32">
                <div class="h-2 bg-[#0D7C7C] rounded-full" style="width:75%"></div>
              </div>
              <span class="text-sm font-semibold text-[#0D7C7C]">75%</span>
            </div>
          </div>
          <div class="flex items-center gap-2 text-sm text-[#4A5259]">
            <Icon icon="lucide:clock" class="text-sm text-[#F59D0D]" />
            <span class="text-[#F59D0D] font-semibold">预计完成时间: 1小时后</span>
          </div>
        </div>
        <div class="flex items-center gap-2">
          <div class="w-2 h-2 rounded-full bg-[#0F8B5D] animate-pulse"></div>
          <span class="text-xs text-[#4A5259]">实时同步中 · 最后更新: 30秒前</span>
        </div>
      </div>

      <!-- Task Requirements -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
          <Icon icon="lucide:target" class="text-sm text-[#0D7C7C]" />
          任务需求与范围
        </h3>
        <p class="text-sm text-[#4A5259] leading-[1.8]">
          实现短信验证码发送与校验的API接口，包括: 验证码生成、短信发送调用、Redis缓存存储(5分钟有效期)、频率限制(每手机号每天5次)、以及6位随机数字验证码生成逻辑。
        </p>
        <div class="flex flex-col gap-2">
          <div v-for="scope in scopes" :key="scope" class="flex items-center gap-2 bg-[hsla(200,15%,95%,1)] rounded-lg px-3 py-2 text-sm text-[#4A5259]">
            <Icon icon="lucide:plus-circle" class="text-sm text-[#0F8B5D]" />
            <span>{{ scope }}</span>
          </div>
        </div>
      </div>

      <!-- AI Agent Work Log -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2">
            <Icon icon="lucide:scroll-text" class="text-sm text-[#0D7C7C]" />
            AI Agent 实时工作日志
          </h3>
          <div class="flex items-center gap-2">
            <div class="w-2 h-2 rounded-full bg-[#0F8B5D] animate-pulse"></div>
            <span class="text-xs text-[#4A5259]">实时更新</span>
          </div>
        </div>
        <div class="flex flex-col gap-3 max-h-[240px] overflow-y-auto pr-2">
          <div v-for="(log, idx) in logEntries" :key="idx"
            :class="[
              'flex items-start gap-3',
              idx === logEntries.length - 1 ? 'bg-[#E8F7F0] rounded-lg px-3 py-2 -mx-3' : ''
            ]">
            <div class="flex flex-col items-center">
              <div :class="[
                'w-6 h-6 rounded-full flex items-center justify-center',
                idx === logEntries.length - 1 ? 'bg-[#E8F7F0]' : 'bg-[hsla(200,15%,95%,1)]'
              ]">
                <Icon :icon="log.icon" :class="['text-xs', idx === logEntries.length - 1 ? 'text-[#0D7C7C]' : log.iconColor ? `text-[${log.iconColor}]` : 'text-[#6B7680]']" />
              </div>
            </div>
            <div class="flex flex-col gap-1">
              <div class="flex items-center gap-2">
                <span :class="[
                  'text-xs font-semibold px-2 py-0.5 rounded-full',
                  idx === logEntries.length - 1 ? 'text-[#0D7C7C] bg-[#E8F0FE]' : 'text-[#9BA3AB] bg-[hsla(200,15%,95%,1)]'
                ]">{{ log.time }}</span>
                <span :class="[
                  'text-sm',
                  idx === logEntries.length - 1 ? 'text-[#0D7C7C] font-semibold' : 'text-[#1A1F24]'
                ]">{{ log.message }}</span>
                <div v-if="idx === logEntries.length - 1" class="w-2 h-2 rounded-full bg-[#0D7C7C] animate-pulse"></div>
              </div>
              <span class="text-xs text-[#9BA3AB]">{{ log.detail }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Code Output Preview -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col overflow-hidden mb-6">
        <div class="flex items-center justify-between px-5 py-3 bg-[hsla(200,15%,96%,1)] border-b border-[#E1E6EA]">
          <div class="flex items-center gap-3">
            <Icon icon="lucide:file-code" class="text-base text-[#0D7C7C]" />
            <span class="text-sm font-semibold text-[#1A1F24]">src/verify/verify.controller.ts</span>
            <span class="px-3 py-1 text-xs font-semibold bg-[#E8F7F0] text-[#0F8B5D] rounded-full">实时输出</span>
          </div>
          <div class="flex items-center gap-2">
            <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:copy" class="text-sm" />
              <span class="whitespace-nowrap">复制代码</span>
            </button>
            <router-link to="/review/code" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80 text-sm">
              <Icon icon="lucide:git-pull-request" class="text-sm" />
              <span class="whitespace-nowrap">查看完整代码</span>
            </router-link>
          </div>
        </div>
        <div class="p-5 font-mono text-sm leading-[1.8] overflow-y-auto max-h-[260px]">
          <div v-for="(line, idx) in codeLines" :key="idx"
            :class="[
              'flex items-center gap-3 px-2 py-0.5',
              line.highlight ? line.bgClass : 'hover:bg-[hsla(200,15%,96%,1)]'
            ]">
            <span :class="['text-xs w-8 text-right shrink-0', line.numColor]">{{ line.num }}</span>
            <span :class="line.textColor">{{ line.text }}</span>
          </div>
          <div class="flex items-center justify-center py-2 text-xs text-[#9BA3AB]">
            <Icon icon="lucide:more-horizontal" class="text-base" />
            <span class="ml-1">Agent正在继续编写...</span>
          </div>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center gap-3 mb-6">
        <router-link to="/review/code" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:git-pull-request" class="text-sm" />
          <span class="whitespace-nowrap">查看代码审核</span>
        </router-link>
        <router-link to="/requirements/REQ-0007" class="flex items-center gap-2 px-4 py-2 text-[#0D7C7C] rounded-full transition hover:opacity-80 border border-[#E1E6EA]">
          <Icon icon="lucide:link" class="text-sm" />
          <span class="whitespace-nowrap">查看关联需求</span>
        </router-link>
        <router-link to="/workflow/dashboard" class="flex items-center gap-2 px-4 py-2 text-[#4A5259] rounded-full transition hover:opacity-80 border border-[#E1E6EA]">
          <Icon icon="lucide:arrow-left" class="text-sm" />
          <span class="whitespace-nowrap">返回看板</span>
        </router-link>
      </div>
    </main>

    <!-- Right Sidebar -->
    <aside class="flex-shrink-0 min-w-fit bg-[hsla(200,15%,97.5%,1)] px-6 py-6 flex flex-col gap-6 overflow-y-auto" style="width:320px">
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:git-branch" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">任务依赖关系</span>
        </div>
        <div class="flex flex-col gap-2">
          <span class="text-xs text-[#9BA3AB] font-semibold">依赖的任务</span>
          <div class="flex items-center justify-between bg-[hsla(200,15%,95%,1)] rounded-lg px-3 py-2">
            <div class="flex items-center gap-2">
              <span class="text-sm font-semibold text-[#0D7C7C]">REQ-0008</span>
              <span class="text-sm text-[#4A5259]">多语言支持模块</span>
            </div>
            <StatusBadge label="已完成" />
          </div>
          <div class="flex items-center justify-between bg-[hsla(200,15%,95%,1)] rounded-lg px-3 py-2">
            <div class="flex items-center gap-2">
              <span class="text-sm font-semibold text-[#0D7C7C]">REQ-0005-A</span>
              <span class="text-sm text-[#4A5259]">用户权限管理API</span>
            </div>
            <StatusBadge label="代码审查" />
          </div>
        </div>
        <div class="flex flex-col gap-2">
          <span class="text-xs text-[#9BA3AB] font-semibold">阻塞的任务</span>
          <div class="flex items-center justify-between bg-[hsla(200,15%,95%,1)] rounded-lg px-3 py-2">
            <div class="flex items-center gap-2">
              <span class="text-sm font-semibold text-[#0D7C7C]">TASK-007-B</span>
              <span class="text-sm text-[#4A5259]">手机号验证前端表单</span>
            </div>
            <StatusBadge label="待审核" />
          </div>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:bot" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">AI Agent能力说明</span>
        </div>
        <p class="text-sm text-[#4A5259] leading-[1.8]">
          后端开发Agent专长于Node.js后端服务开发,擅长API接口设计与实现、数据库操作、缓存策略和安全防护逻辑。
        </p>
        <div class="flex flex-wrap gap-2">
          <span class="px-3 py-1 text-xs font-semibold bg-[#E8F0FE] text-[#3367D6] rounded-full">API开发</span>
          <span class="px-3 py-1 text-xs font-semibold bg-[#E8F0FE] text-[#3367D6] rounded-full">数据库设计</span>
          <span class="px-3 py-1 text-xs font-semibold bg-[#E8F0FE] text-[#3367D6] rounded-full">缓存策略</span>
          <span class="px-3 py-1 text-xs font-semibold bg-[#E8F0FE] text-[#3367D6] rounded-full">安全防护</span>
          <span class="px-3 py-1 text-xs font-semibold bg-[#E8F7F0] text-[#0F8B5D] rounded-full">单元测试</span>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:layers" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">技术栈信息</span>
        </div>
        <div class="flex flex-col gap-3">
          <div v-for="tech in techStack" :key="tech.label" class="flex items-center justify-between text-sm">
            <span class="text-[#4A5259] flex items-center gap-2">
              <Icon :icon="tech.icon" :class="`text-base ${tech.iconColor}`" />
              {{ tech.label }}
            </span>
            <span class="font-semibold text-[#1A1F24]">{{ tech.value }}</span>
          </div>
        </div>
      </div>

      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center gap-2">
          <Icon icon="lucide:check-square" class="text-base text-[#0D7C7C]" />
          <span class="text-sm font-semibold text-[#1A1F24]">完成标准</span>
        </div>
        <div class="flex flex-col gap-3 text-sm text-[#4A5259]">
          <div v-for="(c, idx) in criteria" :key="idx" class="flex items-start gap-2">
            <Icon :icon="c.icon" :class="`text-sm ${c.iconColor}`" />
            <span :class="c.textClass">{{ c.text }}</span>
          </div>
        </div>
        <div class="flex items-center justify-between pt-2 border-t border-[#E1E6EA]">
          <span class="text-sm text-[#4A5259]">完成标准达成</span>
          <div class="flex items-center gap-2">
            <div class="h-2 bg-[hsla(200,15%,95%,1)] rounded-full w-24">
              <div class="h-2 bg-[#0D7C7C] rounded-full" style="width:50%"></div>
            </div>
            <span class="text-sm font-semibold text-[#0D7C7C]">3/6</span>
          </div>
        </div>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import StatusBadge from '../../components/StatusBadge.vue'
import PriorityBadge from '../../components/PriorityBadge.vue'

const scopes = [
  'POST /api/verify/send — 验证码发送接口',
  'POST /api/verify/check — 验证码校验接口',
  'Redis存储模块 — 验证码缓存与过期管理',
  '频率限制模块 — IP与手机号双重限制',
]

const logEntries = [
  { time: '10:00', message: '开始分析代码结构...', detail: '分析现有项目结构，确定接口位置与模块依赖关系', icon: 'lucide:eye', iconColor: '#6B7680' },
  { time: '10:15', message: '设计Redis存储方案...', detail: '确定验证码缓存策略: key格式、过期时间、频率计数器设计', icon: 'lucide:layout-list', iconColor: '#6B7680' },
  { time: '10:30', message: '编写验证码发送接口...', detail: '实现 POST /api/verify/send 接口，集成第三方短信服务SDK', icon: 'lucide:code-2', iconColor: '#0D7C7C' },
  { time: '10:45', message: '实现频率限制逻辑...', detail: '基于Redis计数器实现IP与手机号双重频率限制，防止滥用', icon: 'lucide:shield', iconColor: '#F59D0D' },
  { time: '11:00', message: '正在编写单元测试...', detail: '为验证码发送与校验接口编写单元测试用例', icon: 'lucide:flask-conical', isCurrent: true },
]

const codeLines = [
  { num: '1', numColor: 'text-[#9BA3AB]', text: 'import { Controller, Post, Body } from \'@nestjs/common\';', textColor: 'text-[#4A5259]' },
  { num: '2', numColor: 'text-[#9BA3AB]', text: 'import { VerifyService } from \'./verify.service\';', textColor: 'text-[#4A5259]' },
  { num: '3', numColor: 'text-[#9BA3AB]', text: 'import { RedisVerifyStore } from \'./redis.store\';', textColor: 'text-[#4A5259]' },
  { num: '5', numColor: 'text-[#0F8B5D]', text: '@Controller(\'verify\')', textColor: 'text-[#0F8B5D]', highlight: true, bgClass: 'bg-[#E8F7F0]' },
  { num: '6', numColor: 'text-[#0F8B5D]', text: 'export class VerifyController {', textColor: 'text-[#0F8B5D]', highlight: true, bgClass: 'bg-[#E8F7F0]' },
  { num: '9', numColor: 'text-[#0F8B5D]', text: '@Post(\'send\')', textColor: 'text-[#0F8B5D]', highlight: true, bgClass: 'bg-[#E8F7F0]' },
  { num: '14', numColor: 'text-[#3367D6]', text: '@Post(\'check\')', textColor: 'text-[#3367D6]', highlight: true, bgClass: 'bg-[#E8F0FE]' },
]

const techStack = [
  { label: '运行环境', icon: 'simple-icons:nodedotjs', iconColor: 'text-[#0F8B5D]', value: 'Node.js 18' },
  { label: '缓存服务', icon: 'simple-icons:redis', iconColor: 'text-[#D93025]', value: 'Redis 7' },
  { label: '框架', icon: 'lucide:server', iconColor: 'text-[#0D7C7C]', value: 'NestJS' },
  { label: '测试框架', icon: 'lucide:flask-conical', iconColor: 'text-[#6B7680]', value: 'Jest' },
  { label: '语言', icon: 'simple-icons:typescript', iconColor: 'text-[#3367D6]', value: 'TypeScript' },
]

const criteria = [
  { icon: 'lucide:check-circle', iconColor: 'text-[#0F8B5D]', text: '验证码发送接口返回正确响应格式', textClass: '' },
  { icon: 'lucide:check-circle', iconColor: 'text-[#0F8B5D]', text: 'Redis存储方案: 5分钟过期 + 频率计数器', textClass: '' },
  { icon: 'lucide:check-circle', iconColor: 'text-[#0F8B5D]', text: '频率限制: IP + 手机号双重验证', textClass: '' },
  { icon: 'lucide:loader', iconColor: 'text-[#0D7C7C]', text: '单元测试覆盖率 >= 80%', textClass: 'text-[#0D7C7C] font-semibold' },
  { icon: 'lucide:circle', iconColor: 'text-[#9BA3AB]', text: '代码通过安全扫描无高危问题', textClass: 'text-[#6B7680]' },
  { icon: 'lucide:circle', iconColor: 'text-[#9BA3AB]', text: '接口文档自动生成(Swagger)', textClass: 'text-[#6B7680]' },
]
</script>