import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PilotDashboardView from '../views/PilotDashboardView.vue'
import AdminDashboardView from '../views/AdminDashboardView.vue'

import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', name: 'login', component: LoginView },
    {
      path: '/pilot',
      name: 'pilot',
      component: PilotDashboardView,
      meta: { requiresAuth: true }
    },
    {
      path: '/admin',
      name: 'admin',
      component: AdminDashboardView,
      meta: { requiresAuth: true, requiresAdmin: true }
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  
  const isAuthenticated = !!authStore.token
  const userRole = authStore.user?.role

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'login' })
  } 
  else if (to.meta.requiresAdmin && !['admin', 'superadmin'].includes(userRole)) {
    alert('Brak uprawnień administratora!')
    next({ name: 'pilot' })
  } 
  else if (to.name === 'login' && isAuthenticated) {
    if (userRole === 'admin' || userRole === 'superadmin') {
      next({ name: 'admin' })
    } else {
      next({ name: 'pilot' })
    }
  }
  else {
    next()
  }
})

export default router