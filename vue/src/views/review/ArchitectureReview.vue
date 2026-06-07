<template>
  <TopNavBar activeNav="workflow" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center gap-2 text-sm mb-3">
        <router-link to="/requirements" class="text-[#0D7C7C] hover:opacity-80 transition">需求管理</router-link>
        <span>/</span>
        <router-link to="/requirements/REQ-0008" class="text-[#0D7C7C] hover:opacity-80 transition">需求详情</router-link>
        <span>/</span>
        <span class="text-[#1A1F24]">架构审核</span>
      </div>
      <div class="flex items-center justify-between bg-[#FEF6E8] border border-[#F59D0D] rounded-xl px-5 py-3 mb-6">
        <div class="flex items-center gap-3">
          <Icon icon="lucide:alert-circle" class="text-base text-[#F59D0D]" />
          <span class="px-3 py-1 bg-[#F59D0D]/15 text-[#F59D0D] text-sm font-semibold rounded-full">待审核</span>
          <span class="text-sm text-[#6B7680]">AI生成于 2026-06-05 10:30</span>
        </div>
        <span class="text-sm text-[#F59D0D] font-medium">审核超时提醒: 已超过24小时</span>
      </div>
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="px-3 py-1 bg-[#0D7C7C]/10 text-[#0D7C7C] text-sm font-semibold rounded-full">REQ-0008</span>
            <span class="text-lg font-semibold text-[#1A1F24]">多语言支持模块</span>
          </div>
          <router-link to="/requirements/REQ-0008" class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
            <Icon icon="lucide:external-link" class="text-base" /> <span class="whitespace-nowrap text-sm">查看完整需求</span>
          </router-link>
        </div>
        <p class="text-sm text-[#6B7680]">系统需要支持中英文双语界面切换，需要设计可扩展的国际化架构，支持动态语言包加载和翻译管理。</p>
      </div>
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-lg font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:blocks" class="text-base text-[#0D7C7C]" /> AI架构方案详情</h3>
        <h4 class="text-base font-semibold text-[#1A1F24]">技术选型</h4>
        <div class="flex flex-col gap-3">
          <div v-for="tech in techSelections" :key="tech.title" class="flex items-start gap-4 bg-[hsla(200,15%,95%,1)] rounded-lg p-4">
            <div class="w-8 h-8 rounded-lg flex items-center justify-center" :style="{ backgroundColor: tech.bg }">
              <Icon :icon="tech.icon" :class="`text-sm text-[${tech.color}]`" />
            </div>
            <div>
              <span class="text-sm font-semibold text-[#1A1F24]">{{ tech.title }}</span>
              <p class="text-sm text-[#6B7680] mt-1">{{ tech.description }}</p>
            </div>
          </div>
        </div>
        <!-- Architecture Diagram -->
        <div class="mt-4 bg-white rounded-lg p-4 border border-[#E1E6EA]">
          <h5 class="text-sm font-semibold text-[#1A1F24] mb-3 flex items-center gap-2"><Icon icon="lucide:layout-dashboard" class="text-base text-[#0D7C7C]" /> 系统架构图</h5>
          <div class="flex flex-col items-center gap-3">
            <div class="w-full bg-[#E8F0FE] rounded-lg px-4 py-3 text-center border border-[#3367D6]/30">
              <div class="flex items-center justify-center gap-2"><Icon icon="lucide:monitor" class="text-base text-[#3367D6]" /><span class="text-sm font-semibold text-[#3367D6]">客户端层 (React + i18next)</span></div>
              <div class="flex justify-center gap-4 mt-2">
                <span class="text-xs bg-white rounded px-2 py-1 text-[#4A5259]">语言切换UI</span>
                <span class="text-xs bg-white rounded px-2 py-1 text-[#4A5259]">翻译组件</span>
                <span class="text-xs bg-white rounded px-2 py-1 text-[#4A5259]">语言包缓存</span>
              </div>
            </div>
            <div class="flex items-center gap-1 text-[#6B7680]"><Icon icon="lucide:arrow-down" class="text-sm" /><span class="text-xs">REST API</span><Icon icon="lucide:arrow-down" class="text-sm" /></div>
            <div class="w-full bg-[#E8F7F0] rounded-lg px-4 py-3 text-center border border-[#0F8B5D]/30">
              <div class="flex items-center justify-center gap-2"><Icon icon="lucide:server" class="text-base text-[#0F8B5D]" /><span class="text-sm font-semibold text-[#0F8B5D]">服务层 (Node.js 翻译管理API)</span></div>
            </div>
            <div class="flex items-center gap-1 text-[#6B7680]"><Icon icon="lucide:arrow-down" class="text-sm" /><span class="text-xs">SQL / JSONB</span><Icon icon="lucide:arrow-down" class="text-sm" /></div>
            <div class="w-full bg-[#FEF6E8] rounded-lg px-4 py-3 text-center border border-[#F59D0D]/30">
              <div class="flex items-center justify-center gap-2"><Icon icon="lucide:database" class="text-base text-[#F59D0D]" /><span class="text-sm font-semibold text-[#F59D0D]">数据层 (PostgreSQL + JSONB)</span></div>
            </div>
          </div>
        </div>
      </div>
      <!-- Task Breakdown -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:list-checks" class="text-base text-[#0D7C7C]" /> 任务拆解预览</h3>
        <div class="flex flex-col gap-3">
          <div v-for="task in taskBreakdown" :key="task.title" class="flex items-center justify-between bg-[hsla(200,15%,95%,1)] rounded-lg px-4 py-3">
            <div class="flex items-center gap-3"><span class="w-2 h-2 rounded-full" :style="{ backgroundColor: task.dotColor }"></span><span class="text-sm text-[#1A1F24]">{{ task.title }}</span></div>
            <div class="flex items-center gap-3 text-sm text-[#6B7680]"><span>预估 {{ task.estimate }}h</span><span :class="task.tagClass" class="px-2 py-0.5 rounded-full text-xs">{{ task.tag }}</span></div>
          </div>
        </div>
      </div>
      <!-- Annotation Tool -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <h3 class="text-base font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:message-square" class="text-base text-[#0D7C7C]" /> 批注工具</h3>
        <div class="flex items-center gap-2">
          <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80"><Icon icon="lucide:message-plus" class="text-base" /><span class="whitespace-nowrap text-sm">添加评论</span></button>
          <button class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80"><Icon icon="lucide:flag" class="text-base" /><span class="whitespace-nowrap text-sm">标记修改</span></button>
          <button class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80"><Icon icon="lucide:lightbulb" class="text-base" /><span class="whitespace-nowrap text-sm">插入替代方案</span></button>
        </div>
      </div>
    </main>

    <!-- Right Sidebar -->
    <aside class="flex-shrink-0 min-w-fit w-[340px] bg-[hsla(200,15%,97.5%,1)] border-l border-[#E1E6EA] p-6 flex flex-col gap-6 overflow-y-auto">
      <div class="flex flex-col gap-3">
        <h3 class="text-sm font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:brain" class="text-base text-[#0D7C7C]" /> AI决策依据</h3>
        <div class="flex flex-col gap-3">
          <div v-for="basis in aiBasis" :key="basis.title" class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl p-4">
            <span class="text-sm text-[#1A1F24] font-medium">{{ basis.title }}</span>
            <p class="text-xs text-[#6B7680] mt-1">{{ basis.description }}</p>
          </div>
        </div>
      </div>
      <div class="flex flex-col gap-3">
        <h3 class="text-sm font-semibold text-[#1A1F24] flex items-center gap-2"><Icon icon="lucide:alert-triangle" class="text-base text-[#F59D0D]" /> 风险提示</h3>
        <div class="flex flex-col gap-2">
          <div class="flex items-start gap-2 bg-[#FDECEA] rounded-lg px-3 py-2"><span class="w-2 h-2 rounded-full bg-[#D93025] mt-1.5"></span><div><span class="px-2 py-0.5 bg-[#D93025]/10 text-[#D93025] text-xs font-semibold rounded-full">高风险</span><p class="text-xs text-[#6B7680] mt-1">语言包JSON文件可能过大影响加载性能</p></div></div>
          <div class="flex items-start gap-2 bg-[#FEF6E8] rounded-lg px-3 py-2"><span class="w-2 h-2 rounded-full bg-[#F59D0D] mt-1.5"></span><div><span class="px-2 py-0.5 bg-[#F59D0D]/10 text-[#F59D0D] text-xs font-semibold rounded-full">中风险</span><p class="text-xs text-[#6B7680] mt-1">翻译内容同步更新可能存在延迟</p></div></div>
        </div>
      </div>
      <div class="flex flex-col gap-3">
        <button class="flex items-center gap-2 w-full justify-center px-4 py-2 bg-[#0F8B5D] text-white/95 rounded-full transition hover:opacity-80"><Icon icon="lucide:check-circle" class="text-base" /><span class="whitespace-nowrap">批准方案</span></button>
        <button class="flex items-center gap-2 w-full justify-center px-4 py-2 bg-[#F59D0D] text-white/95 rounded-full transition hover:opacity-80"><Icon icon="lucide:pencil" class="text-base" /><span class="whitespace-nowrap">要求修改</span></button>
        <button class="flex items-center gap-2 w-full justify-center px-4 py-2 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80"><Icon icon="lucide:x-circle" class="text-base" /><span class="whitespace-nowrap">驳回方案</span></button>
      </div>
    </aside>
  </div>
</template>

<script setup>
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'

const techSelections = [
  { title: '前端: React + i18next', description: '选择i18next作为国际化框架,支持JSON语言包加载、插值翻译和惰性加载', icon: 'simple-icons:react', bg: '#E8F0FE', color: '#3367D6' },
  { title: '后端: Node.js + 翻译管理API', description: '提供翻译管理后台API,支持语言包CRUD和版本管理', icon: 'simple-icons:nodedotjs', bg: '#E8F7F0', color: '#0F8B5D' },
  { title: '数据库: PostgreSQL + JSONB', description: '使用JSONB字段存储翻译内容,支持按语言高效查询和更新', icon: 'simple-icons:postgresql', bg: '#FEF6E8', color: '#F59D0D' },
]

const taskBreakdown = [
  { title: '前端语言切换组件开发', estimate: '3h', dotColor: '#3367D6', tag: '前端', tagClass: 'bg-[#3367D6]/10 text-[#3367D6]' },
  { title: '翻译管理后台API开发', estimate: '4h', dotColor: '#0F8B5D', tag: '后端', tagClass: 'bg-[#0F8B5D]/10 text-[#0F8B5D]' },
  { title: '数据库迁移脚本', estimate: '1h', dotColor: '#F59D0D', tag: '数据库', tagClass: 'bg-[#F59D0D]/10 text-[#F59D0D]' },
]

const aiBasis = [
  { title: '参考项目: REQ-0003', description: '支付模块同样使用了React+Node.js+PostgreSQL技术栈' },
  { title: '技术文档引用', description: 'i18next官方文档, React国际化最佳实践' },
  { title: '团队技术栈偏好', description: '前端React, 后端Node.js, 数据库PostgreSQL' },
]
</script>
