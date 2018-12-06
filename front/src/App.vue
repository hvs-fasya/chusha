<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>
<script>
  import * as utils from './utils'
  import Blog from './views/Blog'
  import Webinar from './views/Webinar'
  import Home from './views/Home'
  export default {
    created(){
      this.SetTabs();
      this.RefreshAuth();
      //open wss connection to server
      let socket = new WebSocket("wss://" + process.env.VUE_APP_WS_URL);
      this.$store.dispatch('setSocket', socket)
    },
    components:{Blog, Webinar, Home},
    methods: {
      SetTabs: function () {
      this.axios.get('tabs?enabled=false')
        .then(response => {
          let tabs = response.data;
          //add tabs to router
          let { routes } = this.$router.options;
          let routeData = routes.find(r => r.path === "/");
            tabs.forEach(tab => {
              routeData.children.push({
                path: '/' + tab.tab_type.type,
                name: tab.tab_type.type,
                component: this.$options.components[utils.capitalize(tab.tab_type.type)]
              })
          });
          this.$router.addRoutes([routeData]);
          //add tabs list to store
          this.$store.dispatch('setTabs', tabs)
        })
        .catch(e => {
          console.log("ERROR: " + e)
          }
        )
      },
      RefreshAuth: function () {
        this.axios.get('session')
          .then(response => {
            this.$store.commit('setUser', response.data);
            this.$store.commit('setLoggedIn');
          })
          .catch(e => {
            console.log(e)
          })
      }
    }
  }
</script>
<style>
</style>
