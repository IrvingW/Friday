import Vue from 'vue'
import Router from 'vue-router'

import IndexView from '@/views/index'
import LogInView from '@/views/login'
import RegisterView from '@/views/register'

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/',
      name: 'index',
      component: IndexView
    },
    {
      path: '/login',
      name: 'login',
      component: LogInView
    },
    {
      path: '/register',
      name: 'register',
      component: RegisterView
    }
  ]
})
