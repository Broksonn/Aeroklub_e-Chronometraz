<script setup lang="ts">
import { RouterView, RouterLink, useRouter, useRoute } from 'vue-router'
import { useAuthStore } from './stores/auth'

const auth = useAuthStore()
const router = useRouter()
const route = useRoute()

const handleLogout = () => {
  auth.logout()
  router.push('/')
}
</script>

<template>
  <header v-if="route.path !== '/'">
    <nav class="navbar">
      <h2>✈️ E-Chronometraż</h2>
      <div class="nav-links">
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
  padding: 0;
  font-family: sans-serif;
}

header {
  background-color: #1e293b;
  color: white;
  padding: 1rem 2rem;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.nav-links {
  display: flex;
  gap: 20px;
}

.nav-links a {
  color: white;
  text-decoration: none;
  font-weight: bold;
}

.nav-links a.router-link-exact-active {
  color: #60a5fa; 
}
</style>