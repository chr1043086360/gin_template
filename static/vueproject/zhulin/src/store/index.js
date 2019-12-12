import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

// 在这里直接暴露就可以
export default new Vuex.Store({
  state: {
    name: "CHR"
  },
  // mutations: {},
  actions: {},
  modules: {},
  // 对state的数据做二次处理
  getters: {
    getName: function(state) {
      return state.name + ":getters";
    },
    getNamePlus: function(state, getters) {
      return getters.getName + "Plus";
    }
  },
  // 接受提交的事件
  mutations: {
    change(state, data) {
      state.name = data;
    }
  }
});
