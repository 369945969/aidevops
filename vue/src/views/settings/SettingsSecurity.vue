<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="security" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">安全与权限</h1>
        <div class="flex items-center gap-3">
          <button @click="saveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- RBAC Role Definition -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">RBAC角色定义</h3>
          <p class="text-sm text-[#6B7680] mt-1">基于角色的访问控制权限矩阵</p>
        </div>

        <div class="overflow-hidden rounded-xl border border-[#E1E6EA]">
          <table class="w-full text-sm">
            <thead class="bg-[hsla(200,15%,95%,1)]">
              <tr>
                <th class="px-4 py-3 text-left font-semibold text-[#1A1F24]">功能模块</th>
                <th v-for="role in roleNames" :key="role" class="px-4 py-3 text-center font-semibold text-[#1A1F24]">{{ role }}</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="row in permissionMatrix" :key="row.module" class="border-t border-[#E1E6EA]">
                <td class="px-4 py-3 text-[#1A1F24] font-medium">{{ row.module }}</td>
                <td v-for="role in roleNames" :key="role" class="px-4 py-3 text-center">
                  <div @click="togglePermission(row, role)" :class="['w-5 h-5 rounded-full mx-auto cursor-pointer', row[role] ? 'bg-[#0F8B5D]' : 'bg-[#E1E6EA]']"></div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Security Policies -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">安全策略配置</h3>
          <p class="text-sm text-[#6B7680] mt-1">设置全局安全策略和访问控制规则</p>
        </div>

        <div v-for="policy in policies" :key="policy.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <Icon :icon="policy.icon" class="text-base text-[#0D7C7C]" />
              <div>
                <span class="text-sm font-semibold text-[#1A1F24]">{{ policy.name }}</span>
                <p class="text-xs text-[#6B7680]">{{ policy.desc }}</p>
              </div>
            </div>
            <div @click="togglePolicy(policy)" :class="['w-12 h-6 rounded-full transition-colors cursor-pointer', policy.enabled ? 'bg-[#0D7C7C]' : 'bg-[#E1E6EA]']">
              <div :class="['w-5 h-5 rounded-full bg-white shadow transition-transform', policy.enabled ? 'translate-x-6' : 'translate-x-0.5']"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Audit Log -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4">
        <div class="flex items-center justify-between">
          <h3 class="text-base font-semibold text-[#1A1F24]">操作审计日志</h3>
          <div class="flex items-center gap-2 px-4 py-2 bg-white border border-[#E1E6EA] rounded-full cursor-pointer">
            <Icon icon="lucide:arrow-up-down" class="text-sm text-[#6B7680]" />
            <span class="text-sm text-[#4A5259]">最近7天</span>
          </div>
        </div>
        <div v-for="log in auditLogs" :key="log.id" class="flex items-start gap-3 border-t border-[#E1E6EA] py-3">
          <div class="w-2 h-2 rounded-full bg-[#0D7C7C] mt-1.5 shrink-0"></div>
          <div class="flex-1">
            <span class="text-sm text-[#1A1F24]">{{ log.action }}</span>
            <span class="text-xs text-[#9BA3AB] ml-2">{{ log.time }}</span>
          </div>
          <span class="text-xs text-[#6B7680]">{{ log.user }}</span>
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
  getRoles,
  getPermissions,
  getRolePermissions,
  updateRolePermission,
  getSecurityPolicies,
  updateSecurityPolicy,
  getAuditLogs,
} from '../../api/settings'

const toastRef = ref(null)

const policyTypeMap = {
  db_encryption: { name: '数据加密', desc: 'API密钥和凭证加密存储', icon: 'lucide:lock' },
  api_encryption: { name: 'API加密', desc: 'API通信加密传输', icon: 'lucide:lock' },
  two_factor_auth: { name: '双因素认证', desc: '登录时需要验证码二次确认', icon: 'lucide:shield' },
  ip_whitelist: { name: 'IP白名单', desc: '限制访问来源IP范围', icon: 'lucide:globe' },
}

const roleNames = ref([])
const permissionMatrix = ref([])
const policies = ref([])
const auditLogs = ref([])

async function fetchRBAC() {
  const rolesData = await getRoles()
  const permissionsData = await getPermissions()
  const rolePermsData = await getRolePermissions()

  roleNames.value = rolesData.map(r => r.name)

  // Build permission matrix from role-permissions data
  const matrix = permissionsData.map(perm => {
    const row = { module: perm.name, _permId: perm.id }
    rolesData.forEach(role => {
      const rp = rolePermsData.find(
        item => item.role_id === role.id && item.permission_id === perm.id
      )
      row[role.name] = rp ? rp.allowed : false
      if (rp) row[`_rpId_${role.name}`] = rp.id
    })
    return row
  })
  permissionMatrix.value = matrix
}

async function fetchPolicies() {
  const data = await getSecurityPolicies()
  policies.value = data.map(p => ({
    id: p.id,
    policy_type: p.policy_type,
    name: policyTypeMap[p.policy_type]?.name || p.policy_type,
    desc: policyTypeMap[p.policy_type]?.desc || p.description || '',
    icon: policyTypeMap[p.policy_type]?.icon || 'lucide:shield',
    enabled: p.enabled,
  }))
}

async function fetchAuditLogs() {
  const data = await getAuditLogs({ period: '7d' })
  auditLogs.value = data.map(log => ({
    id: log.id,
    action: log.action,
    target: log.target,
    result: log.result,
    detail: log.detail,
    time: log.created_at,
    user: log.operator,
  }))
}

async function togglePermission(row, roleName) {
  row[roleName] = !row[roleName]
  const rpId = row[`_rpId_${roleName}`]
  if (rpId) {
    await updateRolePermission(rpId, { allowed: row[roleName] })
  }
}

async function togglePolicy(policy) {
  policy.enabled = !policy.enabled
  await updateSecurityPolicy(policy.id, { enabled: policy.enabled })
}

async function saveAll() {
  try {
    await Promise.all(policies.value.map(p =>
      updateSecurityPolicy(p.id, { enabled: p.enabled })
    ))
    toastRef.value?.success('安全设置已保存')
  } catch (err) {
    console.error('Failed to save security settings:', err)
    toastRef.value?.error('保存失败，请重试')
  }
}

onMounted(() => {
  fetchRBAC()
  fetchPolicies()
  fetchAuditLogs()
})
</script>