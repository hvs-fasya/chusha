<template>
  <q-layout view="lHh Lpr lFf">
    <q-tabs align="justify"  color="primary" inverted>
      <q-route-tab
              to="/groups"
              exact
              slot="title">
        ГРУППЫ
      </q-route-tab>
      <q-route-tab
              to="/features"
              exact
              slot="title">
        ФИЧИ
      </q-route-tab>
      <q-route-tab
              to="/users"
              exact
              slot="title">
        ПОЛЬЗОВАТЕЛИ
      </q-route-tab>
    </q-tabs>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
  import { QTabs, QRouteTab } from 'quasar'

  export default {
    components: {
      QRouteTab,
      QTabs
    },
    name: 'LayoutDefault',
    data () {
      return {

      }
    },
    methods: {
    }
  }
</script>

<style>
</style>
