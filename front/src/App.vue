<template>
  <div id="app">
    <router-view/>
  </div>
</template>
<script>
  export default {
    created(){
      this.SetTabs()
    },
    methods: {
      SetTabs: function () {
      this.axios.get('tabs')
        .then(response => {
          console.log(response.data);
          let tabs = response.data;
this.$store.dispatch('setTabs').then(
  this.$router.push("/")
)
        })
        .catch(e => {
          console.log("ERROR: " + e)
          }
        )
      }
    }
  }
</script>
<style>
</style>
