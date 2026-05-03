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
    const role = password.value === 'admin' ? 'admin' : 'pilot'
    auth.login(username.value, role)
    
    if (role === 'admin') {
      router.push('/admin')
    } else {
      router.push('/pilot')
    }
  }
}
</script>

<template>
  <main class="login-page">
    <div class="overlay"></div>

    <div class="login-box">
      <div class="login-header">
        <div class="logo-circle">
          <span class="airplane-icon">✈️</span>
        </div>
        <h2>E-Chronometraż</h2>
        <p>System zarządzania operacjami</p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div class="input-group">
          <label for="username">Login / Nr Licencji</label>
          <div class="input-wrapper">
            <span class="field-icon">👤</span>
            <input 
              type="text" 
              id="username" 
              v-model="username" 
              required 
              placeholder="Wpisz swój login" 
            />
          </div>
        </div>

        <div class="input-group">
          <label for="password">Hasło</label>
          <div class="input-wrapper">
            <span class="field-icon">🔒</span>
            <input 
              type="password" 
              id="password" 
              v-model="password" 
              required 
              placeholder="••••••••" 
            />
          </div>
        </div>

        <button type="submit" class="login-btn">
          Zaloguj się
        </button>
      </form>
      
      <div class="login-footer">
        <p>© 2026 Aeroklub - Panel Operacyjny</p>
      </div>
    </div>
  </main>
</template>

<style scoped>
.login-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  width: 100%;
  background-image: url('https://images.unsplash.com/photo-1544016768-982d1554f0b9?q=80&w=2070&auto=format&fit=crop');
  background-size: cover;
  background-position: center;
  position: relative;
  margin: 0;
  padding: 0;
}


.overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, rgba(15, 23, 42, 0.8) 0%, rgba(30, 41, 59, 0.6) 100%);
  z-index: 1;
}

.login-box {
  background: rgba(255, 255, 255, 0.98);
  padding: 40px;
  border-radius: 20px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4);
  width: 100%;
  max-width: 400px;
  position: relative;
  z-index: 2; 
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.logo-circle {
  width: 70px;
  height: 70px;
  background: #f1f5f9;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 15px;
}

.airplane-icon {
  font-size: 2rem;
}

h2 {
  font-size: 1.6rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}

.login-header p {
  color: #64748b;
  font-size: 0.9rem;
  margin-top: 5px;
}

.input-group {
  margin-bottom: 18px;
}

label {
  display: block;
  margin-bottom: 6px;
  font-weight: 600;
  font-size: 0.8rem;
  color: #475569;
  text-transform: uppercase;
}

.input-wrapper {
  position: relative;
}

.field-icon {
  position: absolute;
  left: 12px;
  top: 50%;
  transform: translateY(-50%);
  opacity: 0.6;
}

input {
  width: 100%;
  padding: 12px 12px 12px 40px;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 1rem;
  outline: none;
  transition: border-color 0.2s;
}

input:focus {
  border-color: #3b82f6;
}

.login-btn {
  width: 100%;
  padding: 14px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 10px;
  cursor: pointer;
  font-weight: 700;
  font-size: 1rem;
  margin-top: 10px;
  transition: background 0.3s;
}

.login-btn:hover {
  background-color: #1d4ed8;
}

.login-footer {
  margin-top: 25px;
  text-align: center;
  font-size: 0.75rem;
  color: #94a3b8;
}
</style>