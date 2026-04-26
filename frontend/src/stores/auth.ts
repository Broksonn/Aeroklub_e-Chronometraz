import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<{ name: string; role: 'pilot' | 'admin' } | null>(null)
  
  const isLoggedIn = computed(() => user.value !== null)
  const userRole = computed(() => user.value?.role)

  function login(name: string, role: 'pilot' | 'admin') {
    user.value = { name, role }
  }

  function logout() {
    user.value = null
  }

  return { user, isLoggedIn, userRole, login, logout }
})