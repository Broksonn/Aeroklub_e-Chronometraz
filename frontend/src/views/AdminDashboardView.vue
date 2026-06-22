<script setup lang="ts">
import { ref } from 'vue'
import mockAirplanes from '../data/mockAirplanes.json'

const activeTab = ref('fleet') 
const showPlaneModal = ref(false)
const showUserModal = ref(false)

const airplanes = ref([...mockAirplanes])
const users = ref([
  { id: 1, name: 'Jan Kowalski', role: 'admin' },
  { id: 2, name: 'Marek Nowak', role: 'pilot' }
])

const newPlane = ref({ model: '', registration: '', pricePerMinute: 0 })
const newUser = ref({ name: '', role: 'user' })

const addPlane = () => {
  if (!newPlane.value.model || !newPlane.value.registration) return
  airplanes.value.push({ id: Date.now(), ...newPlane.value })
  newPlane.value = { model: '', registration: '', pricePerMinute: 0 }
  showPlaneModal.value = false
}

const addUser = () => {
  if (!newUser.value.name) return
  users.value.push({ id: Date.now(), ...newUser.value })
  newUser.value = { name: '', role: 'user' }
  showUserModal.value = false
}

const removePlane = (id: number) => {
  airplanes.value = airplanes.value.filter(p => p.id !== id)
}

const removeUser = (id: number) => {
  users.value = users.value.filter(u => u.id !== id)
}
</script>

<template>
  <main class="admin-panel">
    <div class="tabs">
      <button @click="activeTab = 'fleet'" :class="{ active: activeTab === 'fleet' }">Flota</button>
      <button @click="activeTab = 'users'" :class="{ active: activeTab === 'users' }">Użytkownicy</button>
    </div>

    <div v-if="activeTab === 'fleet'" class="table-container">
      <div class="actions-bar">
        <button @click="showPlaneModal = true" class="add-btn">+ Dodaj statek</button>
      </div>
      <table>
        <thead><tr><th>Model statku</th><th>Rejestracja</th><th>Stawka</th><th class="text-right">Akcje</th></tr></thead>
        <tbody>
          <tr v-for="plane in airplanes" :key="plane.id">
            <td>{{ plane.model }}</td>
            <td><span class="reg-tag">{{ plane.registration }}</span></td>
            <td>{{ plane.pricePerMinute }} PLN/min</td>
            <td class="text-right"><button @click="removePlane(plane.id)" class="del-btn">Usuń</button></td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="activeTab === 'users'" class="table-container">
      <div class="actions-bar">
        <button @click="showUserModal = true" class="add-btn">+ Dodaj użytkownika</button>
      </div>
      <table>
        <thead><tr><th>Imię</th><th>Rola</th><th class="text-right">Akcje</th></tr></thead>
        <tbody>
          <tr v-for="user in users" :key="user.id">
            <td>{{ user.name }}</td>
            <td><span class="reg-tag">{{ user.role }}</span></td>
            <td class="text-right"><button @click="removeUser(user.id)" class="del-btn">Usuń</button></td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showPlaneModal || showUserModal" class="modal-overlay">
      <div class="modal-content">
        <h3>{{ showPlaneModal ? 'Dodaj nowy statek' : 'Dodaj użytkownika' }}</h3>
        
        <template v-if="showPlaneModal">
          <input v-model="newPlane.model" placeholder="Model statku" />
          <input v-model="newPlane.registration" placeholder="Znak rej." />
          <input v-model.number="newPlane.pricePerMinute" type="number" placeholder="Stawka PLN/min" />
        </template>

        <template v-if="showUserModal">
          <input v-model="newUser.name" placeholder="Imię użytkownika" />
          <select v-model="newUser.role" class="custom-select">
            <option value="user">User</option>
            <option value="admin">Admin</option>
            <option value="pilot">Pilot</option>
          </select>
        </template>

        <div class="modal-actions">
          <button @click="showPlaneModal = false; showUserModal = false" class="cancel-btn">Anuluj</button>
          <button @click="showPlaneModal ? addPlane() : addUser()" class="save-btn">Zapisz</button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped>
.admin-panel { background-color: rgb(15, 23, 42); min-height: 100vh; padding: 40px 20px; font-family: sans-serif; color: rgb(241, 245, 249); }
.tabs { display: flex; gap: 10px; margin-bottom: 20px; max-width: 1000px; margin: 0 auto 20px auto; }
.tabs button { background: rgb(30, 41, 59); color: #94a3b8; border: none; padding: 12px 24px; border-radius: 8px; cursor: pointer; font-weight: 600; }
.tabs button.active { background: #3b82f6; color: white; }
.table-container { background: rgb(30, 41, 59); border-radius: 12px; padding: 30px; max-width: 1000px; margin: 0 auto; }
.actions-bar { display: flex; justify-content: flex-end; margin-bottom: 20px; }
table { width: 100%; border-collapse: collapse; }
td, th { padding: 16px 12px; text-align: left; }
.reg-tag { background: rgb(51, 65, 85); padding: 4px 8px; border-radius: 4px; font-family: monospace; }
.add-btn { background: #3b82f6; color: white; border: none; padding: 10px 20px; border-radius: 6px; cursor: pointer; }
.del-btn { background: #ef4444; color: white; border: none; padding: 6px 12px; border-radius: 4px; cursor: pointer; }
.text-right { text-align: right; }
.modal-overlay { position: fixed; top:0; left:0; width:100%; height:100%; background: rgba(0,0,0,0.7); display: flex; justify-content: center; align-items: center; z-index: 1000; }
.modal-content { background: rgb(30, 41, 59); padding: 40px; border-radius: 12px; width: 100%; max-width: 400px; display: flex; flex-direction: column; gap: 15px; }
.modal-content input, .custom-select { background: rgb(15, 23, 42); border: 1px solid rgb(51, 65, 85); color: white; padding: 12px; border-radius: 6px; }
.modal-actions { display: flex; gap: 10px; margin-top: 15px; }
.cancel-btn {  background: #475569; color: white; border: none; padding: 12px 20px; border-radius: 6px; cursor: pointer; }
.save-btn {margin-left: auto; background: #10b981; color: white; border: none; padding: 12px 20px; border-radius: 6px; cursor: pointer; }
</style>