import Vue from 'vue'
import Router from 'vue-router'

import MainLayout from '@/views/layout/MainLayout'

Vue.use(Router)

const mainRouter = [
  {
    path: '/',
    component: MainLayout,
    children: [{
      path: '',
      component: () => import('@/views/home/index')
    }]
  }
]

const router = new Router({
  routes: mainRouter
})

export default router
