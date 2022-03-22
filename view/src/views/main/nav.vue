<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-19 21:25:00
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2021-11-08 16:56:19
-->
<template>
  <div class="nav-top">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页 </el-breadcrumb-item>
      <el-breadcrumb-item>{{ $route.meta.title }}</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="nav-top-right">
      <el-avatar
        :size="size"
        alt="用户头像"
        :src="require('../../assets/adminImg.jpg')"
      ></el-avatar>
      <el-dropdown>
        <span class="el-dropdown-link">
          管理员<i class="el-icon-caret-bottom"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item>个人中心</el-dropdown-item>
          <router-link :to="{name:'login'}" class="quit">
            <el-dropdown-item>退出</el-dropdown-item>
          </router-link>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>
<script>
export default {
  name: "MainNav",
  data() {
    return {
      size: 25,
      routervalue: "",
    };
  },
};
</script>
<style lang="scss" scoped>
.nav-top {
  display: flex;
  justify-content: space-between;
  height: 100%;
  align-items: center;
}
.el-breadcrumb {
  margin-left: 20px;
}
.el-dropdown {
  margin-left: 10px;
  margin-right: 20px;
}
.nav-top-right {
  height: 100%;
  display: flex;
  align-items: center;
}
</style>
