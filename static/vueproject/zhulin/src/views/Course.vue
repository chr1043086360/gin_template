<template>
  <div>
    <h1>课程页面</h1>
    <div class="vxname">
      <h1>Vuex:{{ name1 }}</h1>
      <h1>{{ name2 }}</h1>
      <h1>{{ name3 }}</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: "course",
  // data中放的是静态数据,在别处更改这里获取不到
  // data() {
  //   return {
  //     hoddy: this.$store.state.name
  //   };
  // }
  computed: {
    name1: function() {
      return this.$store.state.name;
    },
    name2: function() {
      return this.$store.getters.getName;
    },
    name3: function() {
      return this.$store.getters.getNamePlus;
    }
  }
};
</script>

<style>
.vxname {
  color: indianred;
}
</style>
