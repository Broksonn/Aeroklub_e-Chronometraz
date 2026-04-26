<script setup lang="ts">
import { RouterView, RouterLink } from 'vue-router'
</script>

<template>
  <header>
    <nav class="navbar">
      <h2>✈️ E-Chronometraż</h2>
      
      <div class="nav-links">
        <RouterLink to="/">Logowanie</RouterLink>
        <RouterLink to="/pilot">Panel Pilota</RouterLink>
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