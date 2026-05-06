// import { createRouter, createWebHistory } from 'vue-router'
// import LoginView from '../views/LoginView.vue'
// import PilotDashboardView from '../views/PilotDashboardView.vue'
// import AdminDashboardView from '../views/AdminDashboardView.vue'
// import { useAuthStore } from '../stores/auth'

// const router = createRouter({
//   history: createWebHistory(import.meta.env.BASE_URL),
//   routes: [
//     {
//       path: '/',
//       name: 'login',
//       component: LoginView,
//     },
//     {
//       path: '/pilot',
//       name: 'pilot-dashboard',
//       component: PilotDashboardView,
//       // Добавляем флаг, что страница требует входа
//       meta: { requiresAuth: true }
//     },
//     {
//       path: '/admin',
//       name: 'admin-dashboard',
//       component: AdminDashboardView,
//       // Требует входа и роль админа
//       meta: { requiresAuth: true, role: 'admin' }
//     }
//   ],
// })

// // Наш "Strażnik" теперь работает по-настоящему:
// router.beforeEach((to, from, next) => {
//   const auth = useAuthStore()

//   // 1. Если страница требует авторизации, а пользователь не залогинен — на выход (на логин)
//   if (to.meta.requiresAuth && !auth.isLoggedIn) {
//     next('/')
//   } 
//   // 2. Если страница только для админа, а роль пользователя не 'admin' — тоже на выход
//   else if (to.meta.role === 'admin' && auth.userRole !== 'admin') {
//     next('/')
//   } 
//   // 3. Во всех остальных случаях — добро пожаловать
//   else {
//     next()
//   }
// })

// export default router

import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import PilotDashboardView from '../views/PilotDashboardView.vue'
import AdminDashboardView from '../views/AdminDashboardView.vue'

// --- NOWE ZMIANY: Importujemy nasz magazyn autoryzacji ---
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/pilot',
      name: 'pilot',
      component: PilotDashboardView,
      // --- NOWE ZMIANY: Dodajemy meta tagi dla ochrony ścieżki ---
      meta: { 
        requiresAuth: true // Wymaga bycia zalogowanym
      }
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminDashboardView,
      // --- NOWE ZMIANY: Ten widok wymaga logowania ORAZ roli admina ---
      meta: { 
        requiresAuth: true, 
        requiresAdmin: true 
      }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    }
  ]
})

// --- NOWE ZMIANY: Globalny strażnik nawigacji (Navigation Guard) ---
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  // Sprawdzamy czy użytkownik ma token (czy jest zalogowany)
  const isAuthenticated = !!authStore.token
  // Pobieramy rolę użytkownika (jeśli user istnieje)
  const userRole = authStore.user?.role

  // 1. Jeśli ścieżka wymaga logowania (requiresAuth), a użytkownik NIE jest zalogowany
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'login' }) // Przekieruj na stronę logowania
  } 
  // 2. Jeśli ścieżka wymaga roli administratora (requiresAdmin), a użytkownik nią NIE jest
  else if (to.meta.requiresAdmin && userRole !== 'admin') {
    alert('Brak uprawnień administratora!')
    next({ name: 'pilot' }) // Odesłanie zwykłego piloti do jego panelu
  } 
  // 3. Jeśli użytkownik jest już zalogowany i próbuje wejść na stronę logowania
  else if (to.name === 'login' && isAuthenticated) {
    // Przekieruj go od razu do odpowiedniego panelu w zależności od roli
    if (userRole === 'admin') {
      next({ name: 'admin' })
    } else {
      next({ name: 'pilot' })
    }
  }
  // 4. We wszystkich innych przypadkach - przepuść dalej normalnie
  else {
    next()
  }
})

export default router