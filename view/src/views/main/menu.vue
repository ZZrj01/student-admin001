<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-25 07:59:42
 * @LastEditTime: 2021-12-11 14:31:27
 * @Description: 
-->
<template>
  <el-menu
    default-active="2"
    class="el-menu-vertical-demo"
    background-color="#0d3f67"
    text-color="#f2f4f6"
  >
    <div class="divtop">学生管理系统</div>
    <el-menu-item
      v-for="item in routers"
      :key="item.path"
      @click="menuHandler(item.path)"
    >
      <i class="el-icon-guide"></i>
      <span slot="title">{{ item.meta.title }}</span>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: "MainMenu",
  data() {
    return {};
  },
  computed: {
    routers() {
      var routes = [];
      for (const item of this.$router.options.routes) {
        if (!item.hide) {
          console.log(Object.prototype.toString.call(item.children));
          if (item.children.length > 0) {
            for (const obj of item.children) {
              if(obj.meta.title != "考勤详情"){
                routes.push(obj);
              }
            }
          }
        }
      }
      return routes;
    },
  },
  mounted() {
    console.log("isturn的",this.$route);
  },
  methods: {
    // 实现路由跳转
    menuHandler(path) {
      if (this.$router.currentRoute.path != path) {
        this.$router.replace(path);
      }
    },
  },
};
</script>

<style scoped>
.el-menu{
  width: 100%;
}
.divtop {
  color: white;
  font-size: 25px;
  text-align: center;
  font-weight: bold;
  margin-top: 10px;
}
.el-menu-item {
  text-align: center;
  color: white;
}
.el-menu-item :hover {
  color: white;
  /* padding: 10px 50px; */
}
</style>
