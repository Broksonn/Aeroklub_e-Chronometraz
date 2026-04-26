import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import PilotDashboardView from '../views/PilotDashboardView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: LoginView, // Strona główna to teraz logowanie
    },
    {
      path: '/pilot',
      name: 'pilot-dashboard',
      component: PilotDashboardView, // Panel pilota przenosimy pod /pilot
    }
  ],
})

export default router