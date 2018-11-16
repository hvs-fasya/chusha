<template>
  <q-layout view="lHh Lpr lFf">
  <LoginForm />
  <q-layout-header :reveal="headerReveal">
    <q-toolbar color="primary">
      <q-toolbar-title>
        CHUSHA
      </q-toolbar-title>
      <span v-if="$store.state.loggedIn">{{$store.state.user.email}}</span>
      <q-btn v-if="$store.state.loggedIn" class="float-right" flat @click="Logout()">Выйти</q-btn>
      <q-toggle
              color="black"
              checked-icon="edit"
              unchecked-icon="visibility_off"
              v-if="$store.state.loggedIn && $store.state.user.role.role === 'admin'"
              class="float-right"
              v-model="editMode"
              left-label
              label="Режим редактирования" />
      <q-btn v-if="editMode" flat round dense icon="menu" @click="drawerOpen = !drawerOpen" class="q-ml-sm">
          <q-tooltip>
              Меню редактирования
          </q-tooltip>
      </q-btn>
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
  </q-layout-header>
    <q-layout-drawer v-if="editMode" side="right" v-model="drawerOpen" no-hide-on-route-change behavior="desktop" :content-class="$q.theme === 'mat' ? 'bg-pink-1' : null">
        <EditDrawer />
    </q-layout-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
  //todo: adjustable colors
  //todo: adjustable brand name
  //todo: adjustable tabs visibility
  //todo: pull to refresh
  //todo: tabs adaptive
  //todo: header adaptive
  import {EventBus} from "../event-bus";
  import LoginForm from "../components/LoginForm"
  import EditDrawer from "../components/EditDrawer"
  import * as utils from './../utils';
  export default {
    data () {
      return {
        editMode: true,
        drawerOpen: true
      }
    },
    components: {
      LoginForm,
      EditDrawer
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
