<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="team-members" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">团队成员</h1>
        <button @click="inviteMember" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
          <Icon icon="lucide:user-plus" class="text-base" />
          <span class="whitespace-nowrap">邀请新成员</span>
        </button>
      </div>

      <!-- Member List -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">成员列表</h3>
          <p class="text-sm text-[#6B7680] mt-1">管理团队成员及角色权限分配</p>
        </div>

        <div v-for="member in members" :key="member.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <img :src="member.avatar" :alt="member.name" class="w-10 h-10 rounded-full">
              <div>
                <div class="flex items-center gap-2">
                  <span class="text-sm font-semibold text-[#1A1F24]">{{ member.name }}</span>
                  <span :class="['px-2 py-0.5 rounded-full text-xs font-semibold', member.roleBadge.bg, member.roleBadge.color]">
                    {{ member.roleName }}
                  </span>
                </div>
                <p class="text-xs text-[#6B7680]">{{ member.email }}</p>
              </div>
            </div>
            <div class="flex items-center gap-2">
              <button @click="editMember(member)" class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full transition hover:opacity-80 text-sm text-[#4A5259]">
                <Icon icon="lucide:edit" class="text-sm" />
                <span>编辑</span>
              </button>
              <button @click="removeMember(member)" class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full transition hover:opacity-80 text-sm text-[#D93025]">
                <Icon icon="lucide:trash-2" class="text-sm" />
                <span>移除</span>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Pending Invitations -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">待处理邀请</h3>
          <p class="text-sm text-[#6B7680] mt-1">查看已发送但尚未接受的邀请</p>
        </div>

        <div v-for="invite in pendingInvites" :key="invite.id" class="bg-[#FEF6E8] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Icon icon="lucide:mail" class="text-base text-[#F59D0D]" />
              <div>
                <span class="text-sm font-semibold text-[#1A1F24]">{{ invite.email }}</span>
                <p class="text-xs text-[#6B7680]">邀请时间: {{ invite.time }}</p>
              </div>
            </div>
            <div class="flex items-center gap-2">
              <button class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full transition hover:opacity-80 text-sm text-[#4A5259]">
                <Icon icon="lucide:copy" class="text-sm" />
                <span>复制链接</span>
              </button>
              <button @click="cancelInvite(invite)" class="flex items-center gap-2 px-4 py-2 bg-[#D93025] text-white/95 rounded-full transition hover:opacity-80 text-sm">
                <Icon icon="lucide:x" class="text-sm" />
                <span>取消</span>
              </button>
            </div>
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
  getTeamMembers,
  createTeamMember,
  updateTeamMember,
  deleteTeamMember,
  getPendingInvites,
  createPendingInvite,
  deletePendingInvite,
} from '../../api/settings'

const toastRef = ref(null)

const roleBadgeMap = {
  admin: { bg: 'bg-[#E8F0FE]', color: 'text-[#3367D6]', name: '超级管理员' },
  tech_lead: { bg: 'bg-[#E8F7F0]', color: 'text-[#0F8B5D]', name: '技术负责人' },
  developer: { bg: 'bg-[#FEF6E8]', color: 'text-[#F59D0D]', name: '开发者' },
  observer: { bg: 'bg-[#F5F5F5]', color: 'text-[#6B7680]', name: '观察者' },
}

const roleIdToKey = {}

const members = ref([])
const pendingInvites = ref([])

async function fetchMembers() {
  const data = await getTeamMembers()
  members.value = data.map(m => {
    // If API returns role preload (e.g. m.Role.name), use it; otherwise map role_id
    const roleKey = m.role?.name || m.role_id
    const badge = roleBadgeMap[roleKey] || { bg: 'bg-[#F5F5F5]', color: 'text-[#6B7680]', name: roleKey }
    return {
      id: m.id,
      name: m.name,
      email: m.email,
      avatar: m.avatar || `https://api.dicebear.com/7.x/avataaars/svg?seed=${m.name}`,
      title: m.title,
      status: m.status,
      role_id: m.role_id,
      roleName: badge.name,
      roleBadge: { bg: badge.bg, color: badge.color },
      activity_rate: m.activity_rate,
    }
  })
}

async function fetchInvites() {
  const data = await getPendingInvites()
  pendingInvites.value = data.map(inv => ({
    id: inv.id,
    email: inv.email,
    role_id: inv.role_id,
    status: inv.status,
    time: inv.created_at,
  }))
}

async function editMember(member) {
  try {
    await updateTeamMember(member.id, { name: member.name, email: member.email, role_id: member.role_id })
    toastRef.value?.success('成员信息已更新')
  } catch (err) {
    console.error('Failed to update member:', err)
    toastRef.value?.error('更新失败，请重试')
  }
}

async function removeMember(member) {
  try {
    await deleteTeamMember(member.id)
    members.value = members.value.filter(m => m.id !== member.id)
    toastRef.value?.success('成员已移除')
  } catch (err) {
    console.error('Failed to remove member:', err)
    toastRef.value?.error('移除失败，请重试')
  }
}

async function inviteMember() {
  try {
    const newInvite = await createPendingInvite({
      email: '',
      role_id: '',
    })
    pendingInvites.value.push({
      id: newInvite.id,
      email: newInvite.email,
      role_id: newInvite.role_id,
      status: newInvite.status,
      time: newInvite.created_at,
    })
    toastRef.value?.success('邀请已发送')
  } catch (err) {
    console.error('Failed to invite member:', err)
    toastRef.value?.error('邀请失败，请重试')
  }
}

async function cancelInvite(invite) {
  try {
    await deletePendingInvite(invite.id)
    pendingInvites.value = pendingInvites.value.filter(i => i.id !== invite.id)
    toastRef.value?.success('邀请已取消')
  } catch (err) {
    console.error('Failed to cancel invite:', err)
    toastRef.value?.error('取消失败，请重试')
  }
}

onMounted(() => {
  fetchMembers()
  fetchInvites()
})
</script>