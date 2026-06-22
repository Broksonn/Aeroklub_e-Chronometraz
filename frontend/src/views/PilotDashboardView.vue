<template>
  <main class="pilot-panel">
    <h1 class="title">Panel Pilota - Rezerwacje statków powietrznych</h1>

    <section class="calendar-section">
      <div class="calendar-container">
        <label><strong>Wybierz datę lotu:</strong></label>
        <VDatePicker v-model="selectedDate" mode="date" color="blue" :min-date="new Date()">
          <template #default="{ inputValue, inputEvents }">
            <div class="custom-input-wrapper">
              <input :value="inputValue" v-on="inputEvents" class="modern-input" readonly />
              <span class="calendar-icon">📅</span>
            </div>
          </template>
        </VDatePicker>
      </div>
    </section>

    <section class="planes-grid">
      <div v-for="plane in airplanes" :key="plane.id" class="plane-card">
        <div class="plane-header">
          <h3>{{ plane.model }}</h3>
          <p class="registration">{{ plane.registration }}</p>
        </div>
        <div class="plane-body">
          <p class="price">Cena: <strong>{{ plane.pricePerMinute }} zł / min</strong></p>
          <p class="status" :class="{ 'unavailable': plane.status !== 'Dostępny' }">
            {{ plane.status }}
          </p>
        </div>
        <button :disabled="plane.status !== 'Dostępny'" @click="bookFlight(plane)" class="book-btn">
          Rezerwuj ten samolot
        </button>
      </div>
    </section>

    <section class="history-section">
      <h2>Historia moich lotów</h2>
      <table class="history-table">
        <thead>
          <tr>
            <th>Data</th>
            <th>Samolot</th>
            <th>Czas</th>
            <th>Koszt</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="flight in flightHistory" :key="flight.id">
            <td>{{ flight.date }}</td>
            <td>{{ flight.plane }}</td>
            <td>{{ flight.duration }}</td>
            <td>{{ flight.cost }}</td>
          </tr>
        </tbody>
      </table>
    </section>
  </main>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import mockAirplanes from '../data/mockAirplanes.json'

const airplanes = ref(mockAirplanes)
const selectedDate = ref(new Date())

// Фейковые данные истории (в будущем придут с сервера)
const flightHistory = ref([
  { id: 1, date: '2026-06-15', plane: 'Cessna 172', duration: '60 min', cost: '600 zł' },
  { id: 2, date: '2026-06-18', plane: 'Piper Archer', duration: '90 min', cost: '900 zł' }
])

const bookFlight = (plane: any) => {
  if (!selectedDate.value) {
    alert('Wybierz najpierw datę lotu z kalendarza!')
    return
  }
  const dateString = selectedDate.value.toLocaleDateString('pl-PL')
  alert(`Zarezerwowano samolot ${plane.model} (${plane.registration}) na dzień ${dateString}.`)
}
</script>

<style scoped>
.pilot-panel { max-width: 1000px; margin: 0 auto; padding: 40px 20px; font-family: 'Inter', sans-serif; }
.title { text-align: center; margin-bottom: 40px; color: #1f2937; font-weight: 800; }

.calendar-section { background-color: #ffffff; padding: 25px; border-radius: 12px; margin-bottom: 40px; border: 1px solid #e5e7eb; box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); }
.calendar-container { display: flex; flex-direction: column; align-items: center; gap: 15px; }
.custom-input-wrapper { position: relative; width: 300px; display: flex; align-items: center; }
.modern-input { width: 100%; padding: 12px 40px; border: 2px solid #3b82f6; border-radius: 8px; font-size: 1.1rem; text-align: center; cursor: pointer; background: white; outline: none; }
.calendar-icon { position: absolute; right: 15px; top: 50%; transform: translateY(-50%); pointer-events: none; }

.planes-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(280px, 1fr)); gap: 25px; }
.plane-card { background: white; border: 1px solid #e5e7eb; padding: 20px; border-radius: 12px; transition: transform 0.2s; }
.plane-card:hover { transform: translateY(-5px); box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1); }
.registration { color: #9ca3af; font-size: 0.9rem; margin-bottom: 15px; }
.status { display: inline-block; padding: 4px 12px; border-radius: 20px; background-color: #dcfce7; color: #166534; font-size: 0.85rem; font-weight: 600; }
.status.unavailable { background-color: #fee2e2; color: #991b1b; }
.book-btn { width: 100%; padding: 12px; background-color: #2563eb; color: white; border: none; border-radius: 8px; cursor: pointer; font-weight: bold; }
.book-btn:disabled { background-color: #d1d5db; }

.history-section { margin-top: 60px; background: white; padding: 30px; border-radius: 12px; border: 1px solid #e5e7eb; }
.history-table { width: 100%; border-collapse: collapse; margin-top: 20px; }
.history-table th, .history-table td { padding: 15px; text-align: left; border-bottom: 1px solid #e5e7eb; }
</style>