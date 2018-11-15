<template>
  <q-layout view="lHh Lpr lFf">
    <LoginForm />
    <q-toolbar color="primary">
      <q-toolbar-title>
        CHUSHA
      </q-toolbar-title>
      <q-btn flat @click="Logout()">Выйти</q-btn>
      <q-btn flat @click="ShowSignUp()">Зарегистрироваться</q-btn>
      <q-btn flat @click="ShowLogin()">Войти</q-btn>
    </q-toolbar>

    <q-tabs align="justify"  color="primary" inverted>
      <q-route-tab v-for="t in this.$store.state.tabs" :key="t.id"
              :to="t.tab_type.type"
              exact
              slot="title">
        {{ t.title }}
      </q-route-tab>
    </q-tabs>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
  import {EventBus} from "../event-bus";
  import LoginForm from "../components/LoginForm"
  export default {
    data () {
      return {
      }
    },
    components: {
      LoginForm
    },
    name: 'LayoutDefault',
    mounted(){
    },
    methods: {
      ShowLogin: function () {
        EventBus.$emit('user-form-open', 'login')
      },
      ShowSignUp: function () {
        EventBus.$emit('user-form-open', 'signup')
      },
      Logout: function () {
        console.log("LOGOUT: " + this.user);
        this.axios.delete('session',{},{jar: true, withCredentials: true})
          .then(response => {
            console.log(response.data)
          })
          .catch(e => {
            console.log("ERROR: " + e)
          })
      }
    }
  }
</script>

<style>
</style>
