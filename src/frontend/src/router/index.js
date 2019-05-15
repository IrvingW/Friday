import Vue from 'vue'
import Router from 'vue-router'

import HomeView from '@/views/home'
import LogInView from '@/views/login'
import RegisterView from '@/views/register'
import IndexView from '@/components/Index'
import UserView from '@/components/UserView'
import TaskList from '@/components/TaskList'
import ClusterView from '@/components/ClusterView'
import ImageView from '@/components/ImageView'
import MachineView from '@/components/MachineView'
import ClusterList from '@/components/ClusterList'
import NodeView from '@/components/NodeView'
import AddCluster from '@/components/AddCluster'
import AddMachine from '@/components/AddMachine'
import ConfigCluster from '@/components/ConfigCluster'
import ConfigKube from '@/components/ConfigKube'
import Done from '@/components/Done'
import CreateTask from '@/components/CreateTask'
import UploadImage from '@/components/UploadImage'
import ImageList from '@/components/ImageList'


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
      path: '/friday',
      redirect: '/friday/home/index'
    },
    {
      path: '/friday/home',
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
          path: 'cluster_list',
          component: ClusterList
        },
        {
          path: 'add_cluster',
          component: AddCluster,
          children: [
            {
              path: '',
              redirect: '/friday/home/add_cluster/master'
            },
            {
              path: 'master',
              component: AddMachine
            },
            {
              path: 'kube_config',
              component: ConfigKube
            },
            {
              path: 'cluster_config',
              component: ConfigCluster
            },
            {
              path: 'done',
              component: Done
            }
          ]
        },
        {
          path: 'task_list',
          component: TaskList
        },
        {
          path: 'create_task',
          component: CreateTask
        },
        {
          path: 'node_view',
          component: NodeView
        },
        {
          path: 'image_list',
          component: ImageList
        },
        {
          path: 'image_view',
          component: ImageView
        },
        {
          path: 'upload_image',
          component: UploadImage
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
