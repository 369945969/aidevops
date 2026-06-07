<template>
  <TopNavBar activeNav="settings" />
  <div class="w-full flex min-h-[900px] bg-[hsla(200,8%,99%,1)]">
    <SettingsSidebar activeSection="dev-environment" />
    <main class="flex-1 overflow-x-hidden flex flex-col px-8 py-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-2xl font-semibold text-[#1A1F24]">开发环境</h1>
        <div class="flex items-center gap-3">
          <button @click="handleSaveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:save" class="text-base" />
            <span class="whitespace-nowrap">保存更改</span>
          </button>
          <router-link to="/settings" class="text-sm text-[#6B7680] hover:text-[#0D7C7C]">恢复默认设置</router-link>
        </div>
      </div>

      <!-- Code Repository Connection Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">代码仓库连接</h3>
          <p class="text-sm text-[#6B7680] mt-1">配置代码仓库连接以支持自动化工作流</p>
        </div>
        <div v-for="repo in codeRepos" :key="repo.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center gap-3 mb-3">
            <Icon :icon="repo.icon" class="text-base text-[#1A1F24]" />
            <span class="text-sm font-semibold text-[#1A1F24]">{{ repo.name }}</span>
            <div :class="['flex items-center gap-2 px-3 py-1 rounded-full text-xs', repo.connected ? 'bg-[#0D7C7C] text-white/95' : 'bg-white border border-[#E1E6EA] text-[#9BA3AB]']">
              <Icon :icon="repo.connected ? 'lucide:check-circle' : 'lucide:circle'" class="text-xs" />
              {{ repo.connected ? '已连接' : '未连接' }}
            </div>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">仓库URL</label>
            <input type="text" v-model="repo.url" :placeholder="repo.placeholder || ''" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
        </div>
      </div>

      <!-- CI/CD Configuration Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">CI/CD配置</h3>
          <p class="text-sm text-[#6B7680] mt-1">配置持续集成和持续部署服务</p>
        </div>
        <div v-for="cicd in cicdServices" :key="cicd.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center gap-3 mb-3">
            <Icon :icon="cicd.icon" class="text-base text-[#1A1F24]" />
            <span class="text-sm font-semibold text-[#1A1F24]">{{ cicd.name }}</span>
            <div :class="['flex items-center gap-2 px-3 py-1 rounded-full text-xs', cicd.configured ? 'bg-[#0D7C7C] text-white/95' : 'bg-white border border-[#E1E6EA] text-[#9BA3AB]']">
              <Icon :icon="cicd.configured ? 'lucide:check-circle' : 'lucide:circle'" class="text-xs" />
              {{ cicd.configured ? '已配置' : '未配置' }}
            </div>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-xs text-[#4A5259]">{{ cicd.urlLabel }}</label>
            <input type="text" v-model="cicd.url" :placeholder="cicd.placeholder || ''" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
          </div>
        </div>
      </div>

      <!-- Cloud Service Credentials Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div>
          <h3 class="text-base font-semibold text-[#1A1F24]">云服务凭证</h3>
          <p class="text-sm text-[#6B7680] mt-1">管理云服务访问密钥(脱敏显示)</p>
        </div>
        <div v-for="cloud in cloudServices" :key="cloud.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center gap-3 mb-3">
            <Icon :icon="cloud.icon" class="text-base text-[#1A1F24]" />
            <span class="text-sm font-semibold text-[#1A1F24]">{{ cloud.name }}</span>
          </div>
          <div class="grid grid-cols-2 gap-3">
            <div class="flex flex-col gap-1">
              <label class="text-xs text-[#4A5259]">{{ cloud.keyIdLabel }}</label>
              <input type="text" v-model="cloud.keyId" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
            </div>
            <div class="flex flex-col gap-1">
              <label class="text-xs text-[#4A5259]">{{ cloud.secretLabel }}</label>
              <input type="password" v-model="cloud.secret" class="w-full px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
            </div>
          </div>
        </div>
      </div>

      <!-- SSH Key Management Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-semibold text-[#1A1F24]">SSH密钥管理</h3>
            <p class="text-sm text-[#6B7680] mt-1">管理用于服务器访问的SSH密钥</p>
          </div>
          <button @click="handleAddSSHKey" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:plus" class="text-base" />
            <span class="whitespace-nowrap">添加密钥</span>
          </button>
        </div>
        <div v-for="ssh in sshKeys" :key="ssh.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="flex items-center justify-between">
            <div class="flex flex-col gap-1">
              <span class="text-sm font-semibold text-[#1A1F24]">{{ ssh.displayName }}</span>
              <p class="text-sm text-[#6B7680]">{{ ssh.description }}</p>
              <span class="text-xs text-[#9BA3AB]">{{ ssh.created }} | {{ ssh.lastUsed }}</span>
            </div>
            <div class="flex items-center gap-3">
              <button class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259] transition hover:opacity-80">
                <Icon icon="lucide:eye" class="text-xs" />
                查看
              </button>
              <button @click="handleDeleteSSHKey(ssh.id)" class="flex items-center gap-2 px-3 py-1 bg-white border border-[#E1E6EA] rounded-full text-xs text-[#4A5259] transition hover:opacity-80">
                <Icon icon="lucide:trash-2" class="text-xs" />
                删除
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Environment Variables Card -->
      <div class="bg-white shadow-[0_2px_8px_rgba(0,0,0,0.10)] rounded-2xl flex flex-col p-5 gap-4 mb-6">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-semibold text-[#1A1F24]">环境变量列表</h3>
            <p class="text-sm text-[#6B7680] mt-1">配置运行时环境变量(键值对)</p>
          </div>
          <button @click="handleAddEnvVar" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
            <Icon icon="lucide:plus" class="text-base" />
            <span class="whitespace-nowrap">添加变量</span>
          </button>
        </div>
        <div v-for="env in envVars" :key="env.id" class="bg-[hsla(200,15%,95%,1)] rounded-xl px-5 py-4">
          <div class="grid grid-cols-[1fr_2fr_auto] gap-3 items-center">
            <input type="text" v-model="env.key" class="px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
            <input :type="env.isPassword ? 'password' : 'text'" v-model="env.value" class="px-3 py-2 bg-white border border-[#E1E6EA] rounded-lg focus:outline-none focus:border-[#CBD3DA] transition text-base">
            <button @click="handleDeleteEnvVar(env.id)" class="flex items-center justify-center cursor-pointer">
              <Icon icon="lucide:trash-2" class="text-base text-[#9BA3AB] hover:text-[#E5484D]" />
            </button>
          </div>
        </div>
      </div>
    </main>
  </div>

  <!-- Bottom Save Bar -->
  <div v-if="hasUnsavedChanges" class="fixed bottom-0 left-0 right-0 bg-[#FEF6E8] border-t border-[#F59D0D] px-8 py-3 flex items-center justify-between">
    <span class="text-sm text-[#F59D0D] font-semibold flex items-center gap-2">
      <Icon icon="lucide:alert-circle" class="text-base text-[#F59D0D]" />
      您有未保存的更改
    </span>
    <div class="flex items-center gap-3">
      <button @click="handleCancel" class="flex items-center gap-2 px-4 py-2 bg-[hsla(200,15%,95%,1)] text-[#4A5259] border border-[#E1E6EA] rounded-full transition hover:opacity-80">
        <Icon icon="lucide:x" class="text-base" />
        <span class="whitespace-nowrap">取消</span>
      </button>
      <button @click="handleSaveAll" class="flex items-center gap-2 px-4 py-2 bg-[#0D7C7C] text-white/95 rounded-full transition hover:opacity-80">
        <Icon icon="lucide:save" class="text-base" />
        <span class="whitespace-nowrap">保存更改</span>
      </button>
    </div>
  </div>
  <Toast ref="toastRef" />
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Icon } from '@iconify/vue'
import TopNavBar from '../../layout/TopNavBar.vue'
import SettingsSidebar from '../../layout/SettingsSidebar.vue'
import Toast from '../../components/Toast.vue'
import {
  getCodeRepos,
  updateCodeRepo,
  getCICDServices,
  updateCICDService,
  getCloudCredentials,
  updateCloudCredential,
  getSSHKeys,
  createSSHKey,
  deleteSSHKey,
  getEnvVars,
  createEnvVar,
  updateEnvVar,
  deleteEnvVar,
} from '../../api/settings'

const toastRef = ref(null)

// --- Provider display mappings ---
const repoProviderMap = {
  github: { name: 'GitHub', icon: 'lucide:github' },
  gitlab: { name: 'GitLab', icon: 'lucide:git-branch' },
  bitbucket: { name: 'Bitbucket', icon: 'lucide:bitbucket' },
  gitee: { name: 'Gitee', icon: 'lucide:code-2' },
}

const cicdProviderMap = {
  jenkins: { name: 'Jenkins', icon: 'lucide:settings-2', urlLabel: 'Jenkins URL' },
  gitlab_ci: { name: 'GitLab CI', icon: 'lucide:workflow', urlLabel: 'GitLab CI URL' },
  github_actions: { name: 'GitHub Actions', icon: 'lucide:rocket', urlLabel: 'GitHub Actions URL' },
}

const cloudProviderMap = {
  aws: { name: 'AWS', icon: 'lucide:cloud', keyIdLabel: 'Access Key ID', secretLabel: 'Secret Access Key' },
  aliyun: { name: '阿里云', icon: 'lucide:server', keyIdLabel: 'Access Key ID', secretLabel: 'Access Key Secret' },
  azure: { name: 'Azure', icon: 'lucide:cloud-sun', keyIdLabel: 'Tenant ID', secretLabel: 'Client Secret' },
  gcp: { name: 'Google Cloud', icon: 'lucide:cloud-rain', keyIdLabel: 'Project ID', secretLabel: 'Service Account Key' },
}

// --- Reactive state ---
const codeRepos = ref([])
const cicdServices = ref([])
const cloudServices = ref([])
const sshKeys = ref([])
const envVars = ref([])
const loading = ref(false)
const error = ref(null)

// --- Original data snapshots for detecting changes ---
const originalCodeRepos = ref([])
const originalCicdServices = ref([])
const originalCloudServices = ref([])
const originalEnvVars = ref([])

const hasUnsavedChanges = computed(() => {
  const stringify = (arr) => JSON.stringify(arr)
  return (
    stringify(codeRepos.value) !== stringify(originalCodeRepos.value) ||
    stringify(cicdServices.value) !== stringify(originalCicdServices.value) ||
    stringify(cloudServices.value) !== stringify(originalCloudServices.value) ||
    stringify(envVars.value) !== stringify(originalEnvVars.value)
  )
})

// --- Transform API data to display format ---
function transformRepo(raw) {
  const display = repoProviderMap[raw.provider] || { name: raw.provider, icon: 'lucide:code-2' }
  return {
    id: raw.id,
    provider: raw.provider,
    name: display.name,
    icon: display.icon,
    connected: raw.connected,
    url: raw.repo_url || '',
    placeholder: raw.repo_url ? '' : `输入${display.name}仓库URL`,
  }
}

function transformCicd(raw) {
  const display = cicdProviderMap[raw.provider] || { name: raw.provider, icon: 'lucide:settings-2', urlLabel: `${raw.provider} URL` }
  return {
    id: raw.id,
    provider: raw.provider,
    name: display.name,
    icon: display.icon,
    configured: raw.configured,
    url: raw.service_url || '',
    urlLabel: display.urlLabel,
    placeholder: raw.service_url ? '' : `输入${display.name}URL`,
  }
}

function transformCloud(raw) {
  const display = cloudProviderMap[raw.provider] || { name: raw.provider, icon: 'lucide:cloud', keyIdLabel: 'Access Key ID', secretLabel: 'Secret Access Key' }
  return {
    id: raw.id,
    provider: raw.provider,
    name: display.name,
    icon: display.icon,
    keyId: raw.access_key_id || '',
    secret: raw.secret_access_key || '',
    keyIdLabel: display.keyIdLabel,
    secretLabel: display.secretLabel,
  }
}

function transformSSHKey(raw) {
  return {
    id: raw.id,
    displayName: `${raw.name} (${raw.key_type})`,
    description: raw.description || '',
    created: `创建于 ${raw.created_at ? formatDate(raw.created_at) : '未知'}`,
    lastUsed: `最后使用 ${raw.last_used_at ? formatDate(raw.last_used_at) : '从未'}`,
    keyContent: raw.key_content || '',
  }
}

function transformEnvVar(raw) {
  return {
    id: raw.id,
    key: raw.key,
    value: raw.value || '',
    isPassword: raw.is_secret,
  }
}

function formatDate(dateStr) {
  const d = new Date(dateStr)
  const y = d.getFullYear()
  const m = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  return `${y}-${m}-${day}`
}

// --- Fetch all data on mount ---
async function fetchAllData() {
  loading.value = true
  error.value = null
  try {
    const [repos, cicds, clouds, sshs, envs] = await Promise.all([
      getCodeRepos(),
      getCICDServices(),
      getCloudCredentials(),
      getSSHKeys(),
      getEnvVars(),
    ])
    codeRepos.value = (repos || []).map(transformRepo)
    cicdServices.value = (cicds || []).map(transformCicd)
    cloudServices.value = (clouds || []).map(transformCloud)
    sshKeys.value = (sshs || []).map(transformSSHKey)
    envVars.value = (envs || []).map(transformEnvVar)

    // Snapshot originals for change detection
    originalCodeRepos.value = JSON.parse(JSON.stringify(codeRepos.value))
    originalCicdServices.value = JSON.parse(JSON.stringify(cicdServices.value))
    originalCloudServices.value = JSON.parse(JSON.stringify(cloudServices.value))
    originalEnvVars.value = JSON.parse(JSON.stringify(envVars.value))
  } catch (e) {
    error.value = e
    console.error('Failed to fetch dev environment settings:', e)
  } finally {
    loading.value = false
  }
}

onMounted(fetchAllData)

// --- Save handlers ---
async function handleSaveAll() {
  loading.value = true
  error.value = null
  try {
    const repoPromises = codeRepos.value.map((repo) =>
      updateCodeRepo(repo.id, {
        provider: repo.provider,
        connected: repo.connected,
        repo_url: repo.url,
      })
    )
    const cicdPromises = cicdServices.value.map((cicd) =>
      updateCICDService(cicd.id, {
        provider: cicd.provider,
        configured: cicd.configured,
        service_url: cicd.url,
      })
    )
    const cloudPromises = cloudServices.value.map((cloud) =>
      updateCloudCredential(cloud.id, {
        provider: cloud.provider,
        access_key_id: cloud.keyId,
        secret_access_key: cloud.secret,
      })
    )
    const envPromises = envVars.value.map((env) =>
      updateEnvVar(env.id, {
        key: env.key,
        value: env.value,
        is_secret: env.isPassword,
      })
    )

    await Promise.all([...repoPromises, ...cicdPromises, ...cloudPromises, ...envPromises])

    // Re-fetch to get server-transformed data and update snapshots
    await fetchAllData()
    toastRef.value?.success('开发环境设置已保存')
  } catch (e) {
    error.value = e
    console.error('Failed to save dev environment settings:', e)
    toastRef.value?.error('保存失败，请重试')
  } finally {
    loading.value = false
  }
}

function handleCancel() {
  // Reset to original snapshots
  codeRepos.value = JSON.parse(JSON.stringify(originalCodeRepos.value))
  cicdServices.value = JSON.parse(JSON.stringify(originalCicdServices.value))
  cloudServices.value = JSON.parse(JSON.stringify(originalCloudServices.value))
  envVars.value = JSON.parse(JSON.stringify(originalEnvVars.value))
}

// --- SSH key handlers ---
async function handleAddSSHKey() {
  const newKeyData = {
    name: '新密钥',
    key_type: 'Ed25519',
    description: '',
  }
  try {
    const created = await createSSHKey(newKeyData)
    sshKeys.value.push(transformSSHKey(created))
    toastRef.value?.success('SSH密钥已添加')
  } catch (e) {
    console.error('Failed to create SSH key:', e)
    toastRef.value?.error('添加密钥失败')
  }
}

async function handleDeleteSSHKey(id) {
  try {
    await deleteSSHKey(id)
    sshKeys.value = sshKeys.value.filter((ssh) => ssh.id !== id)
    toastRef.value?.success('SSH密钥已删除')
  } catch (e) {
    console.error('Failed to delete SSH key:', e)
    toastRef.value?.error('删除密钥失败')
  }
}

// --- Env var handlers ---
async function handleAddEnvVar() {
  const newVarData = {
    key: '',
    value: '',
    is_secret: false,
  }
  try {
    const created = await createEnvVar(newVarData)
    envVars.value.push(transformEnvVar(created))
    toastRef.value?.success('环境变量已添加')
  } catch (e) {
    console.error('Failed to create env var:', e)
    toastRef.value?.error('添加变量失败')
  }
}

async function handleDeleteEnvVar(id) {
  try {
    await deleteEnvVar(id)
    envVars.value = envVars.value.filter((env) => env.id !== id)
    toastRef.value?.success('环境变量已删除')
  } catch (e) {
    console.error('Failed to delete env var:', e)
    toastRef.value?.error('删除变量失败')
  }
}
</script>