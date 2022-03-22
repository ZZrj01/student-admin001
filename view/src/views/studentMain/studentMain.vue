<!--
 * @Author: WUMIANHUA
 * @Date: 2021-11-06 08:55:55
 * @LastEditTime: 2021-12-16 23:01:05
 * @Description: 
-->
<template>
  <div class="contain">
      <el-card>
        <h2>学生名称</h2>
        <p>{{tokename}}</p>
        <h2>手机号码</h2>
        <p>{{tokephone}}</p>
        <h2>IP地址</h2>
        <p>{{tokenip}}</p>
        <el-button type="primary"  @click="dialogVisible = true">提交</el-button>
      </el-card>
      <el-dialog :visible.sync="dialogVisible" title="提示">
        <h2>签到成功 😀</h2>
        <span slot="footer" class="dialog-footer">
          <router-link :to="{name:'studentDetails'}">
            <el-button>查 看</el-button>
          </router-link>  
          <router-link :to="{name:'login'}">
            <el-button type="danger" @click="modifyStudentHandler">退 出</el-button>
          </router-link>
        </span>
      </el-dialog>
  </div>
</template>

<script>
export default {
    name:"studentMain",
    data() {
        return {
          //获得登录界面的学生名称数据
          tokename:sessionStorage.getItem('tokename'),
          tokephone:sessionStorage.getItem('tokenphone'),
          tokenip:sessionStorage.getItem('tokenip'),
          dialogVisible: false, // 显示弹窗
        }
    },
    methods: {
        readData(){
            console.log(sessionStorage.getItem('tokename'))
        },
    }
}
</script>

<style scoped>
.contain{
    width: 100%;
    height: 100%;
    background-image: url("../../assets/studentRegister.jpg");
    background-size: 100%;
    position: relative;
}
.el-card {
  width: 30%;
  margin: auto;
  filter: opacity(90%);
  position: absolute;
  left: 35%;
  top: 10%;
  filter: opacity(90%);
}
</style>