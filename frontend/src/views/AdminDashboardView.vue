<script setup lang="ts">
import { onMounted } from 'vue'
import { useAirplanesStore } from '@/stores/airplanes'
import { useAuthStore } from '@/stores/auth'

const airplanesStore = useAirplanesStore()
const authStore = useAuthStore()

onMounted(() => {
  airplanesStore.fetchAirplanes()
})

const updatePrice = async () => {
  const uniqueTypes = new Map()
  for (const p of airplanesStore.airplanes) {
    uniqueTypes.set(p.model, p.pricePerMinute)
  }

  const payload = Array.from(uniqueTypes.entries()).map(([type, price]) => ({
    type: type,
    pricePerMinute: price
  }))

  try {
    const response = await fetch('http://localhost:8080/api/pricing', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(payload)
    })

    if (response.ok) {
      alert('Ceny zostały zaktualizowane pomyślnie!')
      airplanesStore.fetchAirplanes()
    } else {
      alert('Błąd podczas aktualizacji cen.')
    }
  } catch (err) {
    console.error(err)
    alert('Wystąpił błąd sieci.')
  }
}
</script>

<template>
  <main class="admin-panel">
    <header class="admin-header">
      <div class="header-content">
        <h1>Zarządzanie Flotą</h1>
        <p>Panel administracyjny stawek operacyjnych</p>
      </div>
      <div class="header-stats">
        <span class="stat-label">Statki w systemie:</span>
        <span class="stat-value">{{ airplanesStore.airplanes.length }}</span>
      </div>
    </header>

    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Model statku</th>
            <th>Znak rejestracyjny</th>
            <th>Stawka (PLN/min)</th>
            <th class="text-right">Zarządzanie</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="plane in airplanesStore.airplanes" :key="plane.id">
            <td class="model-cell">{{ plane.model }}</td>
            <td><span class="reg-tag">{{ plane.registration }}</span></td>
            <td>
              <div class="price-input-group">
                <input type="number" v-model.number="plane.pricePerMinute" step="0.1" min="0" />
                <span class="currency">PLN</span>
              </div>
            </td>
            <td class="text-right">
              <button @click="updatePrice()" class="save-btn">Zapisz</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </main>
</template>

<style scoped>
.admin-panel {
  max-width: 1000px;
  margin: 50px auto;
  padding: 0 20px;
  font-family: sans-serif;
}

.admin-header {
  background-color: #1e293b !important;
  padding: 20px 30px;
  border-radius: 8px 8px 0 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white !important; /* Принудительно белый для всей шапки */
}

.admin-header h1 {
  color: #ffffff !important;
  margin: 0;
  font-size: 1.5rem;
}

.admin-header p {
  color: #94a3b8 !important;
  margin: 2px 0 0 0;
  font-size: 0.85rem;
}

.header-stats {
  background: rgba(255, 255, 255, 0.1);
  padding: 8px 16px;
  border-radius: 6px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  display: flex;
  gap: 10px;
  align-items: center;
}

.stat-label {
  color: #cbd5e1 !important; /* Светло-серый текст */
  font-size: 0.8rem;
  text-transform: uppercase;
  font-weight: 600;
}

.stat-value {
  color: #ffffff !important; /* Чисто белый для цифры */
  font-weight: 800;
  font-size: 1.1rem;
}

/* Таблица */
.table-container {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 0 0 8px 8px;
  overflow: hidden;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  background-color: #f8fafc;
  padding: 14px 24px;
  text-align: left;
  font-size: 0.75rem;
  text-transform: uppercase;
  color: #64748b;
  border-bottom: 1px solid #e2e8f0;
}

td {
  padding: 12px 24px;
  border-bottom: 1px solid #f1f5f9;
  color: #1e293b;
}

.reg-tag {
  background: #f1f5f9;
  padding: 3px 8px;
  border-radius: 4px;
  font-family: monospace;
  font-weight: bold;
  border: 1px solid #cbd5e1;
}

input[type="number"] {
  padding: 6px;
  width: 70px;
  border: 1px solid #cbd5e1;
  border-radius: 4px;
}

.save-btn {
  background-color: #10b981;
  color: white;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.save-btn:hover {
  background-color: #059669;
}

.text-right { text-align: right; }
</style>