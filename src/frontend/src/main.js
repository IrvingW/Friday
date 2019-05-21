// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import VueResource from 'vue-resource'
import VueCookies from 'vue-cookies'
import VCharts from 'v-charts'
import MavonEditor from 'mavon-editor'
import 'element-ui/lib/theme-chalk/index.css'
import 'mavon-editor/dist/css/index.css'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueResource)
Vue.use(VueCookies)
Vue.use(VCharts)
Vue.use(MavonEditor)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
