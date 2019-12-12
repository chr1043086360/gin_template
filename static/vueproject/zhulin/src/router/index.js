import Vue from "vue";
// 凡是在项目里安装的东西都不用路径
import VueRouter from "vue-router";
// import Home from "../views/Home.vue";
import HomeThis from "../views/HomeThis";
import Course from "../views/Course";
import Degree from "../views/Degree";
import Light from "../views/Light";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "HomeThis",
    component: HomeThis
  },
  {
    path: "/course",
    name: "Course",
    component: Course
  },
  {
    path: "/light",
    name: "Light",
    component: Light
  },
  {
    path: "/degree",
    name: "Degree",
    component: Degree
  }
];

const router = new VueRouter({
  routes
});

export default router;
