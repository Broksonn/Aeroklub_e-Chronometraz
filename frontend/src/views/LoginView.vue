<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const auth = useAuthStore()

const username = ref('')
const password = ref('')

const handleLogin = () => {
  if (username.value && password.value) {
    // Prosta logika ról dla testów
    const role = password.value === 'admin' ? 'admin' : 'pilot'
    auth.login(username.value, role)
    
    alert(`Zalogowano jako ${auth.user?.name} (Rola: ${role})`)
    
    if (role === 'admin') {
      router.push('/admin')
    } else {
      router.push('/pilot')
    }
  }
}
</script>

<template>
  <main class="login-container">
    <div class="login-box">
      <h2>Logowanie</h2>
      <p>Wprowadź swoje dane, aby uzyskać dostęp do E-Chronometrażu.</p>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="input-group">
          <label for="username">Login (np. numer licencji):</label>
          <input type="text" id="username" v-model="username" required placeholder="Wpisz login" />
        </div>

        <div class="input-group">
          <label for="password">Hasło:</label>
          <input type="password" id="password" v-model="password" required placeholder="Wpisz hasło" />
        </div>

        <button type="submit">Zaloguj się</button>
      </form>
    </div>
  </main>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 70vh;
}

.login-box {
  background-color: white;
  padding: 40px;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
}

h2 {
  margin-top: 0;
  color: #1e293b;
}

.input-group {
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 5px;
  font-weight: bold;
  font-size: 0.9em;
  color: #4b5563;
}

input {
  padding: 10px;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  font-size: 1em;
}

button {
  width: 100%;
  padding: 12px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
  font-size: 1em;
  margin-top: 10px;
}

button:hover {
  background-color: #1d4ed8;
}
</style>