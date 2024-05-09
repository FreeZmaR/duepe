import { createRouter, createWebHistory } from 'vue-router'

import MainView from '@/views/MainView.vue'
import EnvironmentsView from '@/views/EnvironmentsView.vue'
import SettingsView from '@/views/SettingsView.vue'
import ProfileView from '@/views/ProfileView.vue'
import { usePageStore } from '@/stores/page'


const routes = [
  {
    path: '/',
    name: 'Main',
    component: MainView,
    meta: {
      title: 'Main'
    }
  },
  {
    path: '/services',
    name: 'services',
    component: MainView,
    meta: {
      title: 'Services'
    }
  },
  {
    path: '/environments',
    name: 'environments',
    component: EnvironmentsView,
    meta: {
      title: 'Environments'
    }
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfileView,
    meta: {
      title: 'Profile'
    }
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsView,
    meta: {
      title: 'Settings'
    }
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { left: 0, top: 0 }
  }
})

router.beforeEach((to, from, next) => {
  document.title = `Duepe Panel | ${to.meta.title}`

  if (typeof to.meta.title === 'string') {
    usePageStore().setTitle(to.meta.title)
  }

  next()
})

export default router
