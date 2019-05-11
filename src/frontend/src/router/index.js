import Vue from 'vue'
import Router from 'vue-router'

import HomeView from '@/views/home'
import LogInView from '@/views/login'
import RegisterView from '@/views/register'
import IndexView from '@/components/Index'

Vue.use(Router)

export default new Router({
  mode: "history",
  routes: [
    {
      path: '/friday/login',
      name: 'login',
      component: LogInView
    },
    {
      path: '/friday/register',
      name: 'register',
      component: RegisterView
    },
    {
      path: '/',
      redirect: '/friday/home/index'
    },
    {
      path: '/friday/home',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: 'index',
          component: IndexView
        }
      ]
    }
  ]
})
