<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-26 17:52:54
 * @LastEditTime: 2021-12-14 19:36:32
 * @Description: 
-->
<template>
<div class="big_box">
  <div class="login_container">
    <el-card class="box-card">
      <el-form>
        <el-form-item>
          <el-input v-model="userInfo.name" placeholder="用户名"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input v-model="userInfo.passwd" placeholder="用户密码"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="warning" :loading="btnLoading" @click="loginHandler"
            >登陆</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</div>
  
</template>
<script>
var vm;
export default {
  name: "Login",
  data() {
    return {
      btnLoading: false,
      userInfo: {
        name: "admin",
        passwd: "123",
      },
    };
  },
  created() {
    vm = this;
  },
  methods: {
    // 点击登录
    loginHandler() {
      // loading 设置为true
      this.btnLoading = true;
      this.$POST("/login", this.userInfo, function(res) {
        console.log(res);
        if (res.code == 200) {
          vm.$message({
            message: res.message,
            type: "success",
          });
          //设置token
          sessionStorage.setItem("token", res.data.id);
          // 存储登录时的 name，用于签到
          sessionStorage.setItem("tokename", res.data.name);
          sessionStorage.setItem("tokenphone", res.data.phone);
          sessionStorage.setItem("tokenip", res.data.ip);
          //延时跳转，为了效果而延时，正常环境不用延时
          setTimeout(function() {
            vm.btnLoading = false;
            // 跳转到首页去
            if(res.data.name == "admin"){
              vm.$router.replace("/");
            }else{
              vm.$router.replace("/studentMain");
            }
          }, 1000);
        }
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.big_box{
  width: 100%;
  height: 100%;
  background-image: url('../../assets/loginimg.png');
  background-size: 100%;
  position: relative;
}
.login_container {
  width: 30%;
  min-width: 200px;
  position: absolute;
  top: 30%;
  left: 10%;
  ::v-deep .el-card__body {
    // 深度作用element-ui 带__的样式
    background-image: linear-gradient(to bottom left,#8ac6d1, #fff591);
    filter: opacity(100%);
  }
}
</style>
