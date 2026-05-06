import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useAuthStore } from './auth'

export const usePricingStore = defineStore('pricing', () => {
  const prices = ref<any[]>([])
  const authStore = useAuthStore()

  // Pobieranie aktualnego cennika
  async function fetchPrices() {
    const response = await fetch('http://localhost:8080/api/pricing', {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (response.ok) {
      prices.value = await response.json()
    }
  }

  // Aktualizacja cen (dostępna tylko dla roli Administratora)
  async function updatePrices(newPrices: any[]) {
    const response = await fetch('http://localhost:8080/api/pricing', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(newPrices)
    })
    
    if (response.ok) {
      alert('Cennik został zaktualizowany!') 
      // Odświeżamy listę po pomyślnym zapisie w bazie
      await fetchPrices() 
    } else {
      alert('Błąd! Brak uprawnień do edycji cennika.') 
    }
  }

  return { prices, fetchPrices, updatePrices }
})