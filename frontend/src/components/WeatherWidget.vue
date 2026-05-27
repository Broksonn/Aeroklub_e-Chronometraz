<template>
  <div class="weather-widget">
    <h3>🌤️ Prognoza pogody - Lotnisko Szymanów (EWSZ)</h3>
    
    <div v-if="loading" class="loading">
      Pobieranie danych pogodowych...
    </div>
    
    <div v-else-if="error" class="error">
      {{ error }}
    </div>
    
    <div v-else class="forecast-container">
      <div v-for="(day, index) in forecast" :key="index" class="forecast-day">
        <span class="date">{{ day.date }}</span>
        <span class="icon">{{ day.icon }}</span>
        <span class="temp">
          <span class="max">{{ day.maxTemp }}°C</span>
          <span class="min">{{ day.minTemp }}°C</span>
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Typowanie TypeScript dla naszych danych
interface DailyForecast {
  date: string;
  maxTemp: number;
  minTemp: number;
  icon: string;
}

const forecast = ref<DailyForecast[]>([])
const loading = ref(true)
const error = ref('')

// Funkcja do mapowania kodów pogodowych WMO z API na ładne emotki
const getWeatherEmoji = (code: number) => {
  if (code === 0) return '☀️'; // Bezchmurnie
  if (code === 1 || code === 2 || code === 3) return '⛅'; // Częściowe zachmurzenie
  if (code === 45 || code === 48) return '🌫️'; // Mgła
  if (code >= 51 && code <= 67) return '🌧️'; // Deszcz / Mżawka
  if (code >= 71 && code <= 77) return '❄️'; // Śnieg
  if (code >= 80 && code <= 82) return '🌦️'; // Przelotne opady
  if (code >= 95) return '⛈️'; // Burza
  return '☁️';
}

// Formatowanie daty do ładnego polskiego formatu (np. "pon., 12.05")
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('pl-PL', { weekday: 'short', day: '2-digit', month: '2-digit' });
}

// Główna funkcja pobierająca pogodę
const fetchWeather = async () => {
  try {
    loading.value = true;
    // Współrzędne: Lotnisko Szymanów / Wrocław
    const url = 'https://api.open-meteo.com/v1/forecast?latitude=51.20&longitude=16.98&daily=weather_code,temperature_2m_max,temperature_2m_min&timezone=Europe%2FWarsaw';
    
    const response = await fetch(url);
    if (!response.ok) throw new Error('Błąd pobierania pogody');
    
    const data = await response.json();
    
    // Mapowanie odpowiedzi z API na naszą tablicę, bierzemy 7 dni
    forecast.value = data.daily.time.map((time: string, index: number) => ({
      date: formatDate(time),
      maxTemp: Math.round(data.daily.temperature_2m_max[index]),
      minTemp: Math.round(data.daily.temperature_2m_min[index]),
      icon: getWeatherEmoji(data.daily.weather_code[index])
    }));
    
  } catch (e: any) {
    error.value = 'Nie udało się załadować pogody z serwera.';
    console.error(e);
  } finally {
    loading.value = false;
  }
}

// Wywołaj pobieranie od razu po załadowaniu komponentu
onMounted(() => {
  fetchWeather();
})
</script>

<style scoped>
.weather-widget {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  padding: 20px;
  margin: 20px 0;
  box-shadow: 0 4px 6px rgba(0,0,0,0.05);
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.weather-widget h3 {
  margin-top: 0;
  margin-bottom: 20px;
  font-size: 1.2rem;
  color: #1e293b;
  text-align: center;
}

.forecast-container {
  display: flex;
  justify-content: space-between;
  overflow-x: auto;
  gap: 10px;
  padding-bottom: 10px;
}

.forecast-day {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #f8fafc;
  padding: 15px 10px;
  border-radius: 8px;
  min-width: 80px;
  border: 1px solid #f1f5f9;
}

.date {
  font-size: 0.85rem;
  color: #64748b;
  margin-bottom: 8px;
  text-transform: capitalize;
}

.icon {
  font-size: 2rem;
  margin-bottom: 8px;
}

.temp {
  display: flex;
  gap: 8px;
  font-size: 0.95rem;
}

.max {
  font-weight: bold;
  color: #0f172a;
}

.min {
  color: #94a3b8;
}

.loading, .error {
  text-align: center;
  padding: 20px;
  color: #64748b;
}
</style>