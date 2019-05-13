import Vue from 'vue'
import Router from 'vue-router'

import HomeView from '@/views/home'
import LogInView from '@/views/login'
import RegisterView from '@/views/register'
import IndexView from '@/components/Index'
import UserView from '@/components/UserView'
import TaskView from '@/components/TaskView'
import ClusterView from '@/components/ClusterView'
import ImageView from '@/components/ImageView'
import MachineView from '@/components/MachineView'

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
        },
        {
          path: 'cluster_view',
          component: ClusterView
        },
        {
          path: 'task_view',
          component: TaskView
        },
        {
          path: 'image_view',
          component: ImageView
        },
        {
          path: 'user_view',
          component: UserView
        },
        {
          path: 'machine_view',
          component: MachineView
        }
      ]
    }
  ]
})
