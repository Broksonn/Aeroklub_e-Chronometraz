import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PilotDashboardView from '../views/PilotDashboardView.vue'
// 1. Dodajemy import widoku administratora:
import AdminDashboardView from '../views/AdminDashboardView.vue'
// 2. Dodajemy import magazynu autoryzacji:
import { useAuthStore } from '../stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView,
    },
    {
      path: '/pilot',
      name: 'pilot-dashboard',
      component: PilotDashboardView,
    },
    // 3. Rejestrujemy nową ścieżkę dla admina:
    {
      path: '/admin',
      name: 'admin-dashboard',
      component: AdminDashboardView,
    }
  ],
})

export default router

// Nasz "Strażnik" sprawdzający uprawnienia
router.beforeEach((to, from, next) => {
  const auth = useAuthStore()
  
  // 1. Blokada dla niezalogowanych próbujących wejść do Panelu Pilota
  if (to.path === '/pilot' && !auth.isLoggedIn) {
    alert('Musisz się zalogować, aby uzyskać dostęp.')
    next('/') // Wyrzucamy na stronę logowania
  }
  // 2. Blokada dla osób bez roli admina próbujących wejść do Panelu Admina
  else if (to.path === '/admin' && auth.userRole !== 'admin') {
    alert('Brak uprawnień! Tylko dla administratorów.')
    next('/') // Wyrzucamy na stronę logowania
  } 
  // 3. Wpuszczamy całą resztę
  else {
    next() 
  }
})