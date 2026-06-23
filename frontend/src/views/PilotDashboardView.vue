<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAirplanesStore } from '@/stores/airplanes'
import { useAuthStore } from '@/stores/auth'
import WeatherWidget from '../components/WeatherWidget.vue'

const airplanesStore = useAirplanesStore()
const authStore = useAuthStore()
const router = useRouter()

const selectedDate = ref(new Date())
const reservations = ref<any[]>([])

const generateTimeSlots = () => {
  const slots = []
  for (let h = 8; h <= 17; h++) {
    slots.push(`${h.toString().padStart(2, '0')}:00`)
    slots.push(`${h.toString().padStart(2, '0')}:30`)
  }
  slots.push('18:00')
  return slots
}

const timeSlots = generateTimeSlots()
const selectedStartTime = ref('08:00')
const selectedEndTime = ref('09:00')

const isAdminOrSuper = computed(() => {
  return authStore.user?.role === 'admin' || authStore.user?.role === 'superadmin'
})

const getFormattedDate = (dateObj: Date) => {
  const year = dateObj.getFullYear()
  const month = String(dateObj.getMonth() + 1).padStart(2, '0')
  const day = String(dateObj.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const fetchReservations = async () => {
  try {
    const res = await fetch('http://localhost:8080/api/reservations', {
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (res.ok) {
      reservations.value = await res.json()
    }
  } catch (err) {
    console.error('Błąd pobierania rezerwacji', err)
  }
}

onMounted(() => {
  airplanesStore.fetchAirplanes()
  fetchReservations()
})

const isReserved = (planeId: number) => {
  if (!selectedDate.value) return false
  const targetDate = getFormattedDate(selectedDate.value)
  const start = selectedStartTime.value
  const end = selectedEndTime.value

  return reservations.value.some(r => {
    return r.airplane_id === planeId &&
           r.date === targetDate &&
           start < r.end_time && 
           end > r.start_time    
  })
}

const bookFlight = async (plane: any) => {
  if (!selectedDate.value) {
    alert('Wybierz najpierw datę lotu!')
    return
  }

  if (selectedStartTime.value >= selectedEndTime.value) {
    alert('Czas zakończenia rezerwacji musi być późniejszy niż czas rozpoczęcia!')
    return
  }

  // Sprawdzanie wyprzedzenia 30-minutowego
  const [startHour, startMinute] = selectedStartTime.value.split(':').map(Number)
  const reservationDateTime = new Date(selectedDate.value)
  reservationDateTime.setHours(startHour, startMinute, 0, 0)

  const now = new Date()
  const diffInMinutes = (reservationDateTime.getTime() - now.getTime()) / (1000 * 60)

  if (diffInMinutes < 30) {
    alert('BŁĄD: Rezerwację należy złożyć z co najmniej 30-minutowym wyprzedzeniem w stosunku do aktualnego czasu! Rezerwacje w przeszłości również są zabronione.')
    return
  }
  
  const targetDate = getFormattedDate(selectedDate.value)

  try {
    const response = await fetch('http://localhost:8080/api/reservations', {
      method: 'POST',
      headers: { 
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}` 
      },
      body: JSON.stringify({
        airplane_id: plane.id,
        date: targetDate,
        start_time: selectedStartTime.value,
        end_time: selectedEndTime.value
      })
    })

    const data = await response.json()

    if (response.ok) {
      alert(`Sukces: Zarezerwowano ${plane.model} od ${selectedStartTime.value} do ${selectedEndTime.value}`)
      fetchReservations() 
    } else {
      alert(`Błąd z serwera: ${data.error}`)
    }
  } catch (err) {
    alert('Błąd sieci podczas rezerwacji.')
  }
}
</script>

<template>
  <main class="pilot-panel">
    <h1 class="title">Panel Pilota - Rezerwacje statków powietrznych</h1>

    <section class="calendar-section">
      <div class="calendar-container">
        <label><strong>Wybierz datę i godziny lotu:</strong></label>
        
        <VDatePicker v-model="selectedDate" mode="date" color="blue" :min-date="new Date()">
          <template #default="{ inputValue, inputEvents }">
            <div class="custom-input-wrapper">
              <input :value="inputValue" v-on="inputEvents" class="modern-input" readonly />
              <span class="calendar-icon">📅</span>
            </div>
          </template>
        </VDatePicker>

        <div class="time-selectors">
          <div class="time-group">
            <label>Od godziny:</label>
            <select v-model="selectedStartTime" class="time-select">
              <option v-for="time in timeSlots" :key="time" :value="time" :disabled="time === '18:00'">
                {{ time }}
              </option>
            </select>
          </div>
          
          <span class="time-separator">-</span>
          
          <div class="time-group">
            <label>Do godziny:</label>
            <select v-model="selectedEndTime" class="time-select">
              <option v-for="time in timeSlots" :key="time" :value="time" :disabled="time <= selectedStartTime">
                {{ time }}
              </option>
            </select>
          </div>
        </div>
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
          
          <p v-if="plane.status !== 'available'" class="status unavailable">W serwisie</p>
          <p v-else-if="isReserved(plane.id)" class="status reserved">Zajęty w tym czasie</p>
          <p v-else class="status available">Dostępny do rezerwacji</p>
        </div>

        <button 
          :disabled="plane.status !== 'available' || isReserved(plane.id)" 
          @click="bookFlight(plane)"
          class="book-btn"
        >
          {{ isReserved(plane.id) ? 'Brak wolnych miejsc' : 'Zarezerwuj' }}
        </button>
      </div>
    </section>
  </main>
</template>

<style scoped>
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

/* Kalendarz i inputy */
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

/* Nowe selektory czasu */
.time-selectors {
  display: flex;
  align-items: center;
  gap: 15px;
  margin-top: 10px;
  background: #f8fafc;
  padding: 15px 25px;
  border-radius: 8px;
  border: 1px solid #e2e8f0;
}

.time-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
  align-items: center;
}

.time-group label {
  font-size: 0.85rem;
  color: #64748b;
  font-weight: 600;
  text-transform: uppercase;
}

.time-select {
  padding: 10px 15px;
  border: 2px solid #cbd5e1;
  border-radius: 8px;
  font-size: 1.1rem;
  font-weight: 600;
  color: #1e293b;
  cursor: pointer;
  outline: none;
  transition: border-color 0.2s;
}

.time-select:focus {
  border-color: #3b82f6;
}

.time-separator {
  font-weight: bold;
  color: #94a3b8;
  font-size: 1.5rem;
  margin-top: 15px;
}

/* Siatka samolotów */
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

/* Statusy */
.status {
  display: inline-block;
  margin-top: 10px;
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
}

.status.available {
  background-color: #dcfce7;
  color: #166534;
}

.status.unavailable {
  background-color: #fee2e2;
  color: #991b1b;
}

.status.reserved {
  background-color: #fef3c7;
  color: #b45309;
}

/* Przycisk rezerwacji */
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
  background-color: #f1f5f9;
  color: #9ca3af;
  border: 1px solid #e2e8f0;
  cursor: not-allowed;
}
</style>