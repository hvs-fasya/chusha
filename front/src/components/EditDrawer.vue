<template>
    <q-list highlight>
        <q-list-header>ВКЛАДКИ</q-list-header>
        <q-item v-for="tab in allTabs" :key="tab.id">
            <q-item-main :label="tab.title" />
            <q-item-side right>
                <q-toggle
                        checked-icon="flare"
                        unchecked-icon="visibility_off"
                        v-model="tab.enabled"
                        @input="ToggleTab(tab.id, tab.enabled)"/>
            </q-item-side>
        </q-item>
        <q-item-separator />
        <q-list-header>ЦВЕТА</q-list-header>
        <q-item>
            <q-item-side avatar="statics/guy-avatar.png" />
            <q-item-main label="Jack Doe" />
        </q-item>
    </q-list>
</template>

<script>
  import { EventBus } from './../event-bus.js'
  import { Notify } from 'quasar'
  import * as utils from './../utils'
  export default {
    name: "EditDrawer",
    data () {
      return {
        allTabs: []
      }
    },
    mounted(){
      this.GetAllTabs();
    },
    beforeDestroy(){},
    methods: {
      GetAllTabs: function () {
        this.axios.get('tabs?enabled=false')
          .then(response => {
            this.allTabs = response.data;
          })
          .catch(e => {
            //todo: handle errors
              console.log("ERROR: " + e)
            }
          )
      },
      ToggleTab: function(id, enabled){
        this.$store.commit('toggleTab', {id, enabled})
      }
    }
  }
</script>

<style scoped>

</style>