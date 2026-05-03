import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PilotDashboardView from '../views/PilotDashboardView.vue'
import AdminDashboardView from '../views/AdminDashboardView.vue'
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
      // Добавляем флаг, что страница требует входа
      meta: { requiresAuth: true }
    },
    {
      path: '/admin',
      name: 'admin-dashboard',
      component: AdminDashboardView,
      // Требует входа и роль админа
      meta: { requiresAuth: true, role: 'admin' }
    }
  ],
})

// Наш "Strażnik" теперь работает по-настоящему:
router.beforeEach((to, from, next) => {
  const auth = useAuthStore()

  // 1. Если страница требует авторизации, а пользователь не залогинен — на выход (на логин)
  if (to.meta.requiresAuth && !auth.isLoggedIn) {
    next('/')
  } 
  // 2. Если страница только для админа, а роль пользователя не 'admin' — тоже на выход
  else if (to.meta.role === 'admin' && auth.userRole !== 'admin') {
    next('/')
  } 
  // 3. Во всех остальных случаях — добро пожаловать
  else {
    next()
  }
})

export default router