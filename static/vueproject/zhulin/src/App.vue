<template>
  <div id="app">
    <!-- <h1>单页面应用App</h1> -->
    <MyHeader></MyHeader>
    <router-view></router-view>
    <MyFooter class="footer"></MyFooter>
  </div>
</template>

<script>
import MyHeader from "./components/MyHeader";
import MyFooter from "./components/MyFooter";
export default {
  name: "App",
  components: {
    MyHeader,
    MyFooter
  }
};
</script>

<style>
.footer {
  position: fixed;
  bottom: 0;
}
</style>

