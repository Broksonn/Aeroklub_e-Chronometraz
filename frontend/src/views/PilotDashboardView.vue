<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useAirplanesStore } from '@/stores/airplanes'
// 1. IMPORTUJEMY WIDŻET POGODOWY
import WeatherWidget from '../components/WeatherWidget.vue'

const airplanesStore = useAirplanesStore()
const selectedDate = ref(new Date())

onMounted(() => {
  airplanesStore.fetchAirplanes()
})

const bookFlight = (plane: any) => {
  if (!selectedDate.value) {
    alert('Wybierz najpierw datę lotu z kalendarza!')
    return
  }
  const dateString = selectedDate.value.toLocaleDateString('pl-PL')
  alert(`Zarezerwowano samolot ${plane.model} (${plane.registration}) na dzień ${dateString}.`)
}
</script>

<template>
  <main class="pilot-panel">
    <h1 class="title">Panel Pilota - Rezerwacje statków powietrznych</h1>

    <section class="calendar-section">
      <div class="calendar-container">
        <label><strong>Wybierz datę lotu:</strong></label>
        
        <VDatePicker v-model="selectedDate" mode="date" color="blue" :min-date="new Date()">
          <template #default="{ inputValue, inputEvents }">
            <div class="custom-input-wrapper">
              <input
                :value="inputValue"
                v-on="inputEvents"
                class="modern-input"
                readonly
              />
              <span class="calendar-icon">📅</span>
            </div>
          </template>
        </VDatePicker>
      </div>
    </section>

    <WeatherWidget />

    <section class="planes-grid">
      <div v-for="plane in airplanesStore.airplanes" :key="plane.id" class="plane-card">
        <div class="plane-header">
          <h3>{{ plane.model }}</h3>
          <p class="registration">{{ plane.registration }}</p>
        </div>
        
        <div class="plane-body">
          <p class="price">Cena: <strong>{{ plane.pricePerMinute || 0 }} zł / min</strong></p>
          <p class="status" :class="{ 'unavailable': plane.status !== 'available' }">
            {{ plane.status === 'available' ? 'Dostępny' : 'W serwisie' }}
          </p>
        </div>

        <button 
          :disabled="plane.status !== 'available'" 
          @click="bookFlight(plane)"
          class="book-btn"
        >
          Rezerwuj ten samolot
        </button>
      </div>
    </section>
  </main>
</template>

<style scoped>
/* Twoje style zostają DOKŁADNIE TAKIE SAME, nie usunąłem ani jednej linijki! */
.pilot-panel {
  max-width: 1000px;
  margin: 0 auto;
  padding: 40px 20px;
  font-family: 'Inter', sans-serif;
}

.title {
  text-align: center;
  margin-bottom: 40px;
  color: #1f2937;
  font-weight: 800;
}

.calendar-section {
  background-color: #ffffff;
  padding: 25px;
  border-radius: 12px;
  margin-bottom: 40px;
  border: 1px solid #e5e7eb;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.calendar-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 15px;
}

.custom-input-wrapper {
  position: relative;
  width: 300px; 
  display: flex;
  align-items: center;
}

.modern-input {
  width: 100%;
  padding: 12px 40px; 
  border: 2px solid #3b82f6;
  border-radius: 8px;
  font-size: 1.1rem;
  text-align: center;
  cursor: pointer;
  background: white;
  transition: all 0.2s;
  font-weight: 500;
  outline: none;
}

.modern-input:hover {
  border-color: #2563eb;
  box-shadow: 0 0 10px rgba(59, 130, 246, 0.2);
}

.calendar-icon {
  position: absolute;
  right: 15px; 
  top: 50%;
  transform: translateY(-50%);
  pointer-events: none;
  font-size: 1.2rem;
  line-height: 1;
}

.planes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 25px;
}

.plane-card {
  background: white;
  border: 1px solid #e5e7eb;
  padding: 20px;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  transition: transform 0.2s, box-shadow 0.2s;
}

.plane-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
}

.plane-header h3 {
  font-size: 1.25rem;
  font-weight: 700;
  color: #111827;
  margin: 0;
}

.registration {
  color: #9ca3af;
  font-family: monospace;
  font-size: 0.9rem;
  margin-bottom: 15px;
}

.plane-body {
  margin-bottom: 20px;
}

.price {
  color: #4b5563;
}

.status {
  display: inline-block;
  margin-top: 10px;
  padding: 4px 12px;
  border-radius: 20px;
  background-color: #dcfce7;
  color: #166534;
  font-size: 0.85rem;
  font-weight: 600;
}

.status.unavailable {
  background-color: #fee2e2;
  color: #991b1b;
}

.book-btn {
  width: 100%;
  padding: 12px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;
  font-size: 1rem;
  transition: background 0.2s;
}

.book-btn:hover:not(:disabled) {
  background-color: #1d4ed8;
}

.book-btn:disabled {
  background-color: #d1d5db;
  cursor: not-allowed;
}
</style>