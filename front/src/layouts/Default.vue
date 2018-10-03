<template>
  <q-layout view="lHh Lpr lFf">
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
  import { QTabs, QRouteTab } from 'quasar'

  export default {
    components: {
      QRouteTab,
      QTabs
    },
    name: 'LayoutDefault',
    data: () => ({
    }),
    mounted(){
    },
    methods: {
    }
  }
</script>

<style>
</style>
