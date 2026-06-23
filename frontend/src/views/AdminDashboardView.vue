<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useAirplanesStore } from '@/stores/airplanes'
import { useAuthStore } from '@/stores/auth'

const airplanesStore = useAirplanesStore()
const authStore = useAuthStore()

const activeTab = ref<'fleet' | 'users'>('fleet')
const isSuperAdmin = computed(() => authStore.user?.role === 'superadmin')

const newAirplane = ref({ model: '', registration: '', status: 'available' })
const newUser = ref({ username: '', password: '', role: 'pilot' })

const showDeleteModal = ref(false)
const airplaneToDelete = ref<number | null>(null)
const confirmText = ref('')

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
      headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${authStore.token}` },
      body: JSON.stringify(payload)
    })
    if (response.ok) {
      alert('Ceny zaktualizowane pomyślnie!')
      airplanesStore.fetchAirplanes()
    } else {
      alert('Błąd aktualizacji cen.')
    }
  } catch (err) { alert('Błąd sieci.') }
}

const addNewAirplane = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/airplanes', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${authStore.token}` },
      body: JSON.stringify(newAirplane.value)
    })
    if (response.ok) {
      alert('Samolot dodany do floty!')
      airplanesStore.fetchAirplanes()
      newAirplane.value = { model: '', registration: '', status: 'available' }
    } else { alert('Błąd dodawania samolotu.') }
  } catch (err) { alert('Błąd serwera.') }
}

const changeAirplaneStatus = async (id: number, newStatus: string) => {
  try {
    const response = await fetch(`http://localhost:8080/api/airplanes/${id}/status`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${authStore.token}` },
      body: JSON.stringify({ status: newStatus })
    })
    if (!response.ok) {
      alert('Nie udało się zaktualizować statusu.')
      airplanesStore.fetchAirplanes()
    }
  } catch (err) {
    console.error(err)
    airplanesStore.fetchAirplanes()
  }
}

const promptDelete = (id: number) => {
  airplaneToDelete.value = id
  confirmText.value = ''
  showDeleteModal.value = true
}

const confirmDelete = async () => {
  if (confirmText.value.toLowerCase() !== 'potwierdzam') return
  try {
    const response = await fetch(`http://localhost:8080/api/airplanes/${airplaneToDelete.value}`, {
      method: 'DELETE',
      headers: { 'Authorization': `Bearer ${authStore.token}` }
    })
    if (response.ok) {
      alert('Samolot usunięty.')
      airplanesStore.fetchAirplanes()
    }
  } catch (err) { console.error(err) } 
  finally { showDeleteModal.value = false; airplaneToDelete.value = null }
}

const createClubUser = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${authStore.token}` },
      body: JSON.stringify(newUser.value)
    })
    if (response.ok) {
      alert(`Utworzono konto: ${newUser.value.username}`)
      newUser.value = { username: '', password: '', role: 'pilot' }
    } else {
      const data = await response.json()
      alert(data.error || 'Błąd rejestracji.')
    }
  } catch (err) { alert('Błąd sieci.') }
}
</script>

<template>
  <main class="admin-panel">
    <header class="admin-header">
      <div class="header-content">
        <h1>Panel Zarządzania</h1>
        <p>Zalogowano jako: <strong>{{ authStore.user?.role?.toUpperCase() }}</strong></p>
      </div>
      
      <div class="admin-navigation">
        <button @click="activeTab = 'fleet'" :class="['nav-btn', { active: activeTab === 'fleet' }]">Flota i Stawki</button>
        <button @click="activeTab = 'users'" :class="['nav-btn', { active: activeTab === 'users' }]">Zarządzanie Użytkownikami</button>
      </div>
    </header>

    <section v-if="activeTab === 'fleet'" class="tab-content">
      <div class="admin-card">
        <h3>Dodaj statek powietrzny</h3>
        <form @submit.prevent="addNewAirplane" class="inline-form">
          <input v-model="newAirplane.model" type="text" placeholder="Model (np. C172)" required />
          <input v-model="newAirplane.registration" type="text" placeholder="Znaki (np. SP-ABC)" required />
          <select v-model="newAirplane.status" required>
            <option value="available">Dostępny</option>
            <option value="maintenance">Serwis</option>
          </select>
          <button type="submit" class="action-btn success">Dodaj</button>
        </form>
      </div>

      <div class="table-container">
        <table>
          <thead>
            <tr>
              <th>Model</th>
              <th>Znaki</th>
              <th>Status Operacyjny</th>
              <th>Stawka (PLN/min)</th>
              <th class="text-right">Operacje</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="plane in airplanesStore.airplanes" :key="plane.id">
              <td class="model-cell">{{ plane.model }}</td>
              <td><span class="reg-tag">{{ plane.registration }}</span></td>
              <td>
                <select 
                  v-model="plane.status" 
                  @change="changeAirplaneStatus(plane.id, plane.status)"
                  :class="['status-select', plane.status]"
                >
                  <option value="available">Dostępny</option>
                  <option value="maintenance">W serwisie</option>
                </select>
              </td>
              <td>
                <div class="price-input-group">
                  <input type="number" v-model.number="plane.pricePerMinute" step="0.1" min="0" />
                  <span class="currency">PLN</span>
                </div>
              </td>
              <td class="text-right actions-cell">
                <button @click="updatePrice()" class="action-btn save">Zapisz cenę</button>
                <button @click="promptDelete(plane.id)" class="action-btn danger">Usuń</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>

    <section v-if="activeTab === 'users'" class="tab-content">
      <div class="admin-card">
        <h3>Rejestracja członka aeroklubu</h3>
        <p v-if="!isSuperAdmin" class="helper-text">Masz uprawnienia tylko do tworzenia kont pilotów.</p>

        <form @submit.prevent="createClubUser" class="vertical-form">
          <div class="form-group">
            <label>Login:</label>
            <input v-model="newUser.username" type="text" placeholder="Wpisz login..." required />
          </div>
          <div class="form-group">
            <label>Hasło (min. 6 znaków):</label>
            <input v-model="newUser.password" type="password" placeholder="Hasło..." required minlength="6" />
          </div>
          <div class="form-group">
            <label>Rola systemowa:</label>
            <select v-model="newUser.role" required>
              <option value="pilot">Pilot</option>
              <option v-if="isSuperAdmin" value="admin">Administrator</option>
              <option v-if="isSuperAdmin" value="superadmin">Super Administrator</option>
            </select>
          </div>
          <button type="submit" class="action-btn success wide">Utwórz konto</button>
        </form>
      </div>
    </section>

    <div v-if="showDeleteModal" class="modal-overlay">
      <div class="modal-content">
        <div class="modal-icon">⚠️</div>
        <h3>Krytyczna operacja</h3>
        <p>Aby usunąć samolot, wpisz <strong>potwierdzam</strong> poniżej:</p>
        <input v-model="confirmText" type="text" placeholder="Wpisz potwierdzam" class="captcha-input" />
        <div class="modal-actions">
          <button @click="showDeleteModal = false" class="action-btn cancel">Anuluj</button>
          <button @click="confirmDelete" :disabled="confirmText.toLowerCase() !== 'potwierdzam'" class="action-btn danger">Usuń trwale</button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.admin-panel {
  max-width: 1100px;
  margin: 40px auto;
  padding: 0 20px;
  font-family: 'Segoe UI', Roboto, sans-serif;
}

.admin-header {
  background-color: #1e293b !important;
  padding: 24px 30px;
  border-radius: 8px 8px 0 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: white !important;
  flex-wrap: wrap;
  gap: 20px;
}

.admin-header h1 {
  color: #ffffff !important;
  margin: 0;
  font-size: 1.6rem;
}

.admin-header p {
  color: #94a3b8 !important;
  margin: 4px 0 0 0;
  font-size: 0.85rem;
}

/* Przełączniki zakładek */
.admin-navigation {
  display: flex;
  gap: 10px;
}

.nav-btn {
  background: rgba(255, 255, 255, 0.08);
  border: 1px solid rgba(255, 255, 255, 0.15);
  color: #cbd5e1;
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
  font-weight: 600;
  transition: all 0.2s ease;
}

.nav-btn:hover {
  background: rgba(255, 255, 255, 0.15);
  color: white;
}

.nav-btn.active {
  background: #3b82f6;
  border-color: #3b82f6;
  color: white;
}

/* Karty */
.admin-card {
  background: #ffffff;
  padding: 24px;
  border-left: 1px solid #e2e8f0;
  border-right: 1px solid #e2e8f0;
  border-bottom: 1px solid #e2e8f0;
}

.admin-card h3 {
  margin-top: 0;
  margin-bottom: 15px;
  color: #0f172a;
}

.card-subtitle {
  font-size: 0.85rem;
  color: #64748b;
  margin-top: -10px;
  margin-bottom: 20px;
}

/* Formularze */
.inline-form {
  display: flex;
  gap: 15px;
  flex-wrap: wrap;
}

.inline-form input, .inline-form select {
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  flex: 1;
  min-width: 180px;
}

.vertical-form {
  max-width: 500px;
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.form-group label {
  font-size: 0.9rem;
  font-weight: 600;
  color: #334155;
}

.vertical-form input, .vertical-form select {
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 6px;
  font-size: 1rem;
}

/* Przyciski operacyjne */
.action-btn {
  padding: 10px 18px;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-weight: bold;
  font-size: 0.9rem;
  transition: background 0.2s ease;
}

.action-btn.success { background-color: #10b981; color: white; }
.action-btn.success:hover { background-color: #059669; }

.action-btn.save { background-color: #3b82f6; color: white; padding: 8px 14px; }
.action-btn.save:hover { background-color: #2563eb; }

.action-btn.danger { background-color: #ef4444; color: white; padding: 8px 14px; }
.action-btn.danger:hover:not(:disabled) { background-color: #dc2626; }
.action-btn.danger:disabled { background-color: #fca5a5; cursor: not-allowed; }

.action-btn.cancel { background-color: #94a3b8; color: white; }
.action-btn.cancel:hover { background-color: #64748b; }

.action-btn.wide { width: 100%; margin-top: 10px; font-size: 1rem; }

/* Tabela */
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
  padding: 16px 24px;
  text-align: left;
  font-size: 0.75rem;
  text-transform: uppercase;
  color: #64748b;
  border-bottom: 1px solid #e2e8f0;
}

td {
  padding: 14px 24px;
  border-bottom: 1px solid #f1f5f9;
  color: #1e293b;
}

.reg-tag {
  background: #f1f5f9;
  padding: 4px 10px;
  border-radius: 4px;
  font-family: monospace;
  font-weight: bold;
  border: 1px solid #cbd5e1;
}

.price-input-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.price-input-group input {
  padding: 6px;
  width: 80px;
  border: 1px solid #cbd5e1;
  border-radius: 4px;
}

.actions-cell {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.text-right { text-align: right; }

/* Modale */
.modal-overlay {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: rgba(15, 23, 42, 0.6);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
}

.modal-content {
  background: white;
  padding: 30px;
  border-radius: 12px;
  max-width: 450px;
  text-align: center;
  box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1);
}

.modal-icon {
  font-size: 3rem;
  margin-bottom: 10px;
}

.captcha-instruction {
  margin: 20px 0 8px;
  font-size: 0.95rem;
  color: #475569;
}

.captcha-input {
  width: 100%;
  padding: 12px;
  margin-bottom: 20px;
  border: 2px solid #cbd5e1;
  border-radius: 6px;
  text-align: center;
  font-size: 1.1rem;
  font-weight: bold;
  box-sizing: border-box;
}

.captcha-input:focus {
  border-color: #ef4444;
  outline: none;
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  gap: 15px;
}

.modal-actions .action-btn {
  flex: 1;
}
</style>