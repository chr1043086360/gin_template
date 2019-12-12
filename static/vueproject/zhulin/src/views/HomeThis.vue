<template>
  <div>
    <!-- <h1>首页</h1> -->
    <el-carousel :interval="3000" type="card">
      <!-- key:要绑定一个唯一的变量 -->
      <el-carousel-item v-for="(item, index) in img_list" :key="index">
        <img :src="item" alt="" />
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  name: "homethis",
  data() {
    return {
      img_list: [
        "https://cdn.sspai.com/2019/11/29/289a9a82c269221f2350bd08768e8202.png?imageMogr2/quality/95/thumbnail/!600x300r/gravity/Center/crop/600x300/interlace/1",
        "https://cdn.sspai.com/2019/12/06/d83ed1ac875d70b010b0ade55a2c3ce0.png?imageMogr2/quality/95/thumbnail/!600x300r/gravity/Center/crop/600x300/interlace/1",
        "https://cdn.sspai.com/article/4cd8ad89-37fe-44fa-1aaf-a8b32376081e.jpg?imageMogr2/quality/95/thumbnail/!600x300r/gravity/Center/crop/600x300/interlace/1"
      ]
    };
  }
};
</script>

<style>
.el-carousel-item img {
  height: 100%;
  width: 100%;
}
/* .el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
} */
</style>
