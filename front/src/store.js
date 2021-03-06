import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    tabs:[],
    user: null,
    loggedIn: false
  },
  mutations: {
    setTabs (state, tabs) {
      state.tabs = tabs
    },
    setUser (state, user) {
      state.user = user
    },
    setLoggedIn (state) {
      state.loggedIn = true
    },
    setLoggedOut (state) {
      state.loggedIn = false
      state.user = null
    }
  },
  actions: {
    setTabs({ commit }, tabs) {
      commit('setTabs', tabs)
    },
    setUser({ commit }, user) {
      commit('setUser', user)
    },
    setLoggedIn ({ commit }) {
      commit('setLoggedIn')
    },
    setLoggedOut ({ commit }) {
      commit('setLoggedIn')
    }
  }
})
