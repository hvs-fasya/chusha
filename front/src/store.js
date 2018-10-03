import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tabs:[]
  },
  mutations: {
    setTabs (state, tabs) {
      state.tabs = tabs
    }
  },
  actions: {
    setTabs({ commit }, tabs) {
      commit('setTabs', tabs)
    }
  }
})
