<template>
  <q-layout view="lHh Lpr lFf">
    <LoginForm />
    <q-toolbar color="primary">
      <q-toolbar-title>
        CHUSHA
      </q-toolbar-title>
      <q-btn v-if="$store.state.loggedIn" flat @click="Logout()">Выйти</q-btn>
      <q-btn v-if="!$store.state.loggedIn" flat @click="ShowSignUp()">Зарегистрироваться</q-btn>
      <q-btn v-if="!$store.state.loggedIn" flat @click="ShowLogin()">Войти</q-btn>
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
  import * as utils from './../utils';
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
        this.axios.delete('session',{},{jar: true, withCredentials: true})
          .then(response => {
            this.$store.commit('setLoggedOut');
          })
          .catch(e => {
            if (e.response.status === 401) {
              utils.notify401();
            } else {
              utils.notify500();
            }
          })
      }
    }
  }
</script>

<style>
</style>
