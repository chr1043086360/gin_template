// 配置中心
import Vue from "vue";
import App from "./App.vue";
import router from "./router/index";
import store from "./store";
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
import axios from "axios";

Vue.use(ElementUI);

Vue.config.productionTip = false;

Vue.prototype.$axios = axios;

new Vue({
  // 如果前后一样可以简写
  router,
  store,
  // 为了解耦要这么写,都是简写的方式
  render: h => h(App)
}).$mount("#app");
