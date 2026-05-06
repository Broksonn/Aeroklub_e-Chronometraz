// import { defineStore } from 'pinia'
// import { ref, computed } from 'vue'

// export const useAuthStore = defineStore('auth', () => {
//   const user = ref<{ name: string; role: 'pilot' | 'admin' } | null>(null)
  
//   const isLoggedIn = computed(() => user.value !== null)
//   const userRole = computed(() => user.value?.role)

//   function login(name: string, role: 'pilot' | 'admin') {
//     user.value = { name, role }
//   }

//   function logout() {
//     user.value = null
//   }

//   return { user, isLoggedIn, userRole, login, logout }
// })


import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
  // Próbujemy pobrać dane z localStorage (jeśli użytkownik był już wcześniej zalogowany)
  const token = ref<string | null>(localStorage.getItem('token'))
  const user = ref<any>(JSON.parse(localStorage.getItem('user') || 'null'))

  // Funkcja logowania
  async function login(username: string, password: string) {
    try {
      const response = await fetch('http://localhost:8080/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username, password })
      })

      if (!response.ok) {
        throw new Error('Nieprawidłowy login lub hasło') 
      }

      const data = await response.json()
      
      // Zapisujemy dane w stanie (state) Pinia
      token.value = data.token
      user.value = data.user

      // Zapisujemy w pamięci przeglądarki (aby sesja przetrwała odświeżenie strony - F5)
      localStorage.setItem('token', data.token)
      localStorage.setItem('user', JSON.stringify(data.user))

      return true
    } catch (error) {
      console.error('Błąd logowania:', error)
      return false
    }
  }

  // Funkcja wylogowywania
  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
  }

  return { user, token, login, logout }
})