<script setup lang="ts">
import { RouterView, RouterLink, useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth' // Importujemy nasz stan

const auth = useAuthStore()
const router = useRouter()

// Funkcja wylogowania
const handleLogout = () => {
  auth.logout()
  router.push('/') // Po wylogowaniu wracamy na stronę logowania
}
</script>

<template>
  <header>
    <nav class="navbar">
      <h2>✈️ E-Chronometraż</h2>
      
      <div class="nav-links">
        <RouterLink v-if="!auth.isLoggedIn" to="/">Logowanie</RouterLink>
        
        <RouterLink v-if="auth.isLoggedIn" to="/pilot">Panel Pilota</RouterLink>

        <RouterLink v-if="auth.userRole === 'admin'" to="/admin">Panel Admina</RouterLink>

        <a href="#" v-if="auth.isLoggedIn" @click.prevent="handleLogout" class="logout-btn">
          Wyloguj ({{ auth.user?.name }})
        </a>
      </div>
    </nav>
  </header>

  <RouterView /> 
</template>

<style>
body {
  margin: 0;
  font-family: sans-serif;
  background-color: #f9fafb;
}

header {
  background-color: #1e293b;
  color: white;
  padding: 1rem 2rem;
  margin-bottom: 2rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

nav h2 {
  margin: 0;
  font-size: 1.5rem;
}

.nav-links {
  display: flex;
  gap: 20px; /* Tworzy równe odstępy między linkami */
}

.nav-links a {
  color: white; /* Zmienia kolor linków na biały */
  text-decoration: none;
  font-weight: bold;
}

.nav-links a:hover {
  text-decoration: underline;
}

.nav-links a.router-link-exact-active {
  color: #60a5fa; 
}
</style>