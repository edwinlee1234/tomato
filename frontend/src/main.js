import Vue from 'vue'
import Vuex from 'vuex'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import VueAxios from 'vue-axios'
import VCharts from 'v-charts'
import { BootstrapVue, BootstrapVueIcons } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

axios.defaults.withCredentials = true

Vue.use(Vuex)
Vue.use(BootstrapVue)
Vue.use(BootstrapVueIcons)
Vue.use(VueAxios, axios)
Vue.use(VCharts)

Vue.config.productionTip = false
Vue.prototype.$axios = axios

// 全域變數
Vue.prototype.APIURL = "http://localhost:9090"
Vue.prototype.BASE = "http://localhost:8080"

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
