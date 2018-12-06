<template>
    <q-list highlight>
        <q-item role="button" style="">
            <q-item-main
                    align="center"
                    >
                <q-btn
                        size="lg"
                        dense
                        outline
                        rounded
                        text-color="primary"
                        label="Сохранить все изменения"
                        class="q-mt-sm q-mb-sm"
                        @click="SaveAllChanges()"
                />
            </q-item-main>
        </q-item>
        <q-item-separator />
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
        <q-item>
            <q-item-main/>
            <q-item-side right>
                <q-btn
                    size="sm"
                    dense
                    outline
                    text-color="primary"
                    label="Сохранить"
                    class="q-mt-sm q-mb-sm"
                    @click="SaveTabsState()"
                />
            </q-item-side>
        </q-item>
        <q-item-separator />
        <q-list-header>ЦВЕТА</q-list-header>
        <q-item>
            <q-item-side avatar="statics/guy-avatar.png" />
            <q-item-main label="Jack Doe" />
        </q-item>
        <q-item>
            <q-item-main/>
            <q-item-side right>
                <q-btn
                        size="sm"
                        dense
                        outline
                        text-color="primary"
                        label="Сохранить"
                        class="q-mt-sm q-mb-sm"
                        @click="SaveTabsState()"
                />
            </q-item-side>
        </q-item>
    </q-list>
</template>

<script>
  import { EventBus } from './../event-bus.js'
  import { Notify } from 'quasar'
  import * as utils from './../utils'
  import QBtn from "quasar-framework/src/components/btn/QBtn";
  export default {
    name: "EditDrawer",
    components: {QBtn},
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
      },
      SaveTabsState: function(){
        console.log(JSON.stringify(this.$store.state.tabs))
        this.axios.put('tabs',
          this.$store.state.tabs,
          {
            jar:true,
            withCredentials: true
        })
          .then(response =>{
            //todo:
            console.log(response.data)
          })
          .catch(e => {
            //todo:
            console.log(e)
          })
      },
      SaveAllChanges: function(){
        this.SaveTabsState();
      }
    }
  }
</script>

<style scoped>

</style>