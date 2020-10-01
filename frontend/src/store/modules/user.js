const state = () => ({
  info: {
      id: "",
      name: "Guset",
      isLogin: false
  },
})

// getters
const getters = {
}

// actions
const actions = {
  login ({ commit }, payload) {
    let info = {
        id: payload.id,
        name: payload.name,
        isLogin: true
    }
    commit('setUserInfo', { items: info })
  },

  logout ({ commit },) {
    let info = {
        id: "",
        name: "Guset",
        isLogin: false
    }
    commit('setUserInfo', { items: info })
  }
}

// mutations
const mutations = {
  setUserInfo (state, { items }) {
    state.info = items
  },
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}