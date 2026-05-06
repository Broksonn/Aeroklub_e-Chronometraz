import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useAuthStore } from './auth'

export const useAirplanesStore = defineStore('airplanes', () => {
  const airplanes = ref<any[]>([])
  const authStore = useAuthStore()

  async function fetchAirplanes() {
    try {
      // Wysyłamy zapytanie do backendu po listę samolotów
      const response = await fetch('http://localhost:8080/api/airplanes', {
        method: 'GET',
        headers: {
          // WAŻNE: Dołączamy nasz token JWT, aby przejść przez bramkę autoryzacji (middleware) na backendzie!
          'Authorization': `Bearer ${authStore.token}`
        }
      })

      if (response.ok) {
        airplanes.value = await response.json()
      } else {
        console.error('Brak dostępu do samolotów')
      }
    } catch (error) {
      console.error('Błąd pobierania danych:', error)
    }
  }

  return { airplanes, fetchAirplanes }
})