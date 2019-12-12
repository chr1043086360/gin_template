<template>
  <div>
    <!-- <h1>头部组件</h1> -->
    <!-- <router-link to="/">首页</router-link>
    <router-link to="/course">课程</router-link>
    <router-link to="/degree">学位课程</router-link> -->
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :router="true"
    >
      <el-menu-item index="/">首页</el-menu-item>
      <el-menu-item index="/course">免费课程</el-menu-item>
      <el-menu-item index="/light">轻课</el-menu-item>
      <el-menu-item index="4">学位课程</el-menu-item>
      <el-menu-item index="5">智能题库</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
// 默认抛出并命名
export default {
  name: "MyHeader",
  data() {
    return {
      activeIndex: "/"
    };
  }
};
</script>
<style>
.el-menu {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
