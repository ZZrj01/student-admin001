<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-24 22:41:20
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2021-12-13 08:25:43
-->
<template>
  <div class="main-contain">
    <div class="main-left"><MainMenu /></div>
    <div class="main-top"><MainNav /></div>
    <div class="main-right">
      <router-view></router-view>
    </div>
  </div>
</template>
<script>
import MainMenu from "./menu";
import MainNav from "./nav";
export default {
  name: "Main",
  components: {
    MainMenu,
    MainNav,
  },
};
</script>
<style lang="scss" scoped>
.main-contain {
  position: relative;
  height: 100%;
}
.main-left {
  position: absolute;
  left: 0;
  top: 0;
  width: 200px;
  height: 100%;
  color: white;
  background-color: #0d3f67;
  z-index: 2;
}
.main-top {
  position: absolute;
  left: 200px;
  top: 0;
  // css 计算公式
  width: calc(100% - 200px);
  height: 50px;
  background-color: rgb(255, 255, 255);
  box-shadow: 0 1px 5px black;
  z-index: 1;
}
.main-right {
  position: absolute;
  top: 50px;
  left: 200px;
  width: calc(100% - 200px);
  height: calc(100% - 50px);
  background-image: url('../../assets/navimg.jpg');
  background-size: 100%;
}
</style>
