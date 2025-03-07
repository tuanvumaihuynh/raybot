import type { RouteRecordRaw } from 'vue-router'
import { useNProgress } from '@/lib/nprogress'
import { createRouter, createWebHistory } from 'vue-router'
import 'nprogress/nprogress.css'

const MainLayout = () => import('@/layouts/main-layout/MainLayout.vue')

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/home',
  },
  {
    path: '/home',
    component: MainLayout,
    children: [
      {
        path: '',
        component: () => import('@/views/Home.vue'),
      },
    ],
  },
  {
    path: '/system',
    component: MainLayout,
    children: [
      {
        path: '',
        component: () => import('@/views/System.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

const nprogress = useNProgress()

router.beforeEach((to, _, next) => {
  let title = import.meta.env.VITE_APP_NAME
  if (to.meta.title) {
    title = `${to.meta.title} | ${title}`
  }
  document.title = title

  nprogress.start()
  next()
})

router.afterEach(() => {
  nprogress.done()
})

export default router
