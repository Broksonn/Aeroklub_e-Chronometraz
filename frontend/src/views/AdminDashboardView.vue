<script setup lang="ts">
import { ref } from 'vue'
import mockAirplanes from '../data/mockAirplanes.json'

const airplanes = ref([...mockAirplanes])

const updatePrice = (id: number, newPrice: string) => {
  const plane = airplanes.value.find(p => p.id === id)
  if (plane) {
    plane.pricePerMinute = parseFloat(newPrice)
    alert(`Zaktualizowano cenę dla ${plane.model}`)
  }
}
</script>

<template>
  <main class="admin-panel">
    <h1>Panel Administratora - Zarządzanie Flotą</h1>
    <table>
      <thead>
        <tr>
          <th>Model</th>
          <th>Rejestracja</th>
          <th>Cena (zł/min)</th>
          <th>Akcje</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="plane in airplanes" :key="plane.id">
          <td>{{ plane.model }}</td>
          <td>{{ plane.registration }}</td>
          <td>
            <input type="number" v-model="plane.pricePerMinute" step="0.5" />
          </td>
          <td>
            <button @click="updatePrice(plane.id, plane.pricePerMinute.toString())">Zapisz</button>
          </td>
        </tr>
      </tbody>
    </table>
  </main>
</template>

<style scoped>
.admin-panel { max-width: 900px; margin: 0 auto; padding: 20px; }
table { width: 100%; border-collapse: collapse; background: white; border-radius: 8px; overflow: hidden; }
th, td { padding: 15px; border-bottom: 1px solid #eee; text-align: left; }
th { background-color: #1e293b; color: white; }
input { padding: 5px; width: 80px; }
button { background: #10b981; color: white; border: none; padding: 8px 12px; border-radius: 4px; cursor: pointer; }
</style>