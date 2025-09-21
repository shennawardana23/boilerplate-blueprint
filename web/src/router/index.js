import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: 'Boilerplate Blueprint - AI-Powered Project Generator'
    }
  },
  {
    path: '/generator',
    name: 'Generator',
    component: () => import('@/views/Generator.vue'),
    meta: {
      title: 'Project Generator - Boilerplate Blueprint'
    }
  },
  {
    path: '/project/:id',
    name: 'Project',
    component: () => import('@/views/Project.vue'),
    props: true,
    meta: {
      title: 'Project Details - Boilerplate Blueprint'
    }
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue'),
    meta: {
      title: 'About - Boilerplate Blueprint'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
    meta: {
      title: 'Page Not Found - Boilerplate Blueprint'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// Update page title
router.beforeEach((to, from, next) => {
  document.title = to.meta.title || 'Boilerplate Blueprint'
  next()
})

export default router