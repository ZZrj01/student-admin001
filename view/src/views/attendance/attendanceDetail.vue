<!--
 * @Author: WUMIANHUA
 * @Date: 2021-11-09 19:19:54
 * @LastEditTime: 2021-12-11 08:52:44
 * @Description: 
-->
<template>
  <div>
    <el-card>
      <el-card label-width='30px' padding='10px'>
        <p style="color:green;">已签</p>
        <div class="icon-warp">
          <i class="el-icon-user-solid"></i>
        </div>
        <h2>001-张三</h2>
      </el-card>
      <el-card label-width='30px' padding='10px'>
        <p style="color:green;">已签</p>
        <div class="icon-warp">
          <i class="el-icon-user-solid"></i>
        </div>
        <h2>002-马龙</h2>
      </el-card>
      <el-card label-width='30px' padding='10px'>
        <p style="color:green;">已签</p>
        <div class="icon-warp">
          <i class="el-icon-user-solid"></i>
        </div>
        <h2>003-陈梦</h2>
      </el-card>
      <el-card label-width='30px' padding='10px'>
        <p style="color:red;">旷课</p>
        <div class="icon-warp">
          <i class="el-icon-user-solid"></i>
        </div>
        <h2>004-许昕</h2>
      </el-card>
      <el-card label-width='30px' padding='10px'>
        <p style="color:blue;">请假</p>
        <div class="icon-warp">
          <i class="el-icon-user-solid"></i>
        </div>
        <h2>005-莎莎</h2>
      </el-card>
    </el-card>
  </div>
</template>

<script>
export default {
    name:'attendanceDetail',
    
}
</script>

<style scoped>
.el-card {
  margin: 10px;
  filter: opacity(90%);
}
.el-card .el-card{
  margin: 10px;
  width: 13%;
  height: 150px;
  position: relative;
  float:left;
  margin-left: 50px;
}
.el-card .el-card p{
  position: absolute;
  top: -5px;
  right: 5px;
}
.icon-warp{
  font-size: 60px;
  margin-left: 20px;
}
</style>