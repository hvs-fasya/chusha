import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tabs:[],
    user: {
      nickname: "",
      role: ""
    }
  },
  mutations: {
    setTabs (state, tabs) {
      state.tabs = tabs
    },
    setUser (state, user) {
      state.user = user
    }
  },
  actions: {
    setTabs({ commit }, tabs) {
      commit('setTabs', tabs)
    },
    setUser({ commit }, user) {
      commit('setUser', user)
    }
  }
})
