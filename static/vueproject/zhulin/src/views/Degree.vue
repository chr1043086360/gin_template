<template>
  <div>
    <h1>学位课程页面</h1>
    <button @click="change_click">点击改变仓库中Name值</button>
    <!-- <h1>{{django_zhulin}}</h1> -->
    <!-- post请求过来的数据 -->
    <h1>{{gin_zhulin}}</h1>
  </div>
</template>

<script>
export default {
  name: "degree",
  data() {
    return {
      django_zhulin: "",
      gin_zhulin: ""
    };
  },
  methods: {
    change_click: function() {
      this.$store.commit("change", "Chr");
    }
  },
  // 默认发请求
  mounted() {
    let that = this;
    // this.$axios
    // .request({
    //   url: "http://127.0.0.1:8000/test",
    //   method: "get"
    // })
    // .then(function(data) {
    //   console.log(data);
    //   that.django_zhulin = data.data.code;
    // })
    // .catch(function(data) {
    //   console.log(data);
    // });

    // put测试复杂请求
    // this.$axios
    //   .request({
    //     method: "put",
    //     url: "http://127.0.0.1:3000/api/v1/ping"
    //   })
    //   .then(function(data) {
    //     console.log(data);
    //     that.django_zhulin = data.data;
    //   })
    //   .catch(function(data) {
    //     console.log(data);
    //   });
    this.$axios
      .request({
        method: "post",
        url: "http://127.0.0.1:3000/api/v1/login",
        data: {
          user: "cui",
          password: "123" // 注意这里只能填入字符串,number类型后端识别不了
        }
      })
      .then(function(data) {
        console.log(data);
        that.django_zhulin = data.data;
      })
      .catch(function(data) {
        console.log(data);
      });
  }
};
</script>

<style>
</style>