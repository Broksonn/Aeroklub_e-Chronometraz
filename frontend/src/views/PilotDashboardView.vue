<script setup lang="ts">
import { ref } from 'vue'
// Importujemy nasze sztuczne dane
import mockAirplanes from '../data/mockAirplanes.json'

// Stan naszej aplikacji
const airplanes = ref(mockAirplanes)
const selectedDate = ref('') // Tu przechowamy wybraną datę z kalendarza

// Funkcja wywoływana po kliknięciu "Rezerwuj"
const bookFlight = (plane: any) => {
  if (!selectedDate.value) {
    alert('Wybierz najpierw datę lotu z kalendarza!')
    return
  }
  alert(`Zarezerwowano samolot ${plane.model} (${plane.registration}) na dzień ${selectedDate.value}.`)
}
</script>

<template>
  <main class="pilot-panel">
    <h1>Panel Pilota - Rezerwacje statków powietrznych</h1>

    <section class="calendar-section">
      <label for="flight-date"><strong>Wybierz datę lotu: </strong></label>
      <input type="date" id="flight-date" v-model="selectedDate" />
    </section>

    <section class="planes-grid">
      <div v-for="plane in airplanes" :key="plane.id" class="plane-card">
        <h3>{{ plane.model }}</h3>
        <p class="registration">{{ plane.registration }}</p>
        <p>Cena: <strong>{{ plane.pricePerMinute }} zł / min</strong></p>
        
        <p class="status" :class="{ 'unavailable': plane.status !== 'Dostępny' }">
          Status: {{ plane.status }}
        </p>

        <button 
          :disabled="plane.status !== 'Dostępny'" 
          @click="bookFlight(plane)"
        >
          Rezerwuj ten samolot
        </button>
      </div>
    </section>
  </main>
</template>

<style scoped>
.pilot-panel {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  font-family: sans-serif;
}

.calendar-section {
  background-color: #f0f4f8;
  padding: 15px;
  border-radius: 8px;
  margin-bottom: 30px;
  border: 1px solid #d1d5db;
}

input[type="date"] {
  padding: 8px;
  margin-left: 10px;
  border-radius: 4px;
  border: 1px solid #ccc;
}

.planes-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.plane-card {
  border: 1px solid #e5e7eb;
  padding: 15px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.registration {
  color: #6b7280;
  font-size: 0.9em;
  margin-top: -10px;
}

.status.unavailable {
  color: #dc2626;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

button:disabled {
  background-color: #9ca3af;
  cursor: not-allowed;
}

button:hover:not(:disabled) {
  background-color: #1d4ed8;
}
</style>