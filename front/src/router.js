import Vue from 'vue'
import Router from 'vue-router'
import DefaultLayout from './layouts/Default.vue'
import store from './store'

Vue.use(Router);

let Tabs = () => {
  return store.state.tabs
};

export default new Router({
  routes: [
    {
      path: '/',
      component: DefaultLayout,
      children: []
    }
  ]
})
