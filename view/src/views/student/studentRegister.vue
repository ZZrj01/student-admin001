<!--
 * @Author: WUMIANHUA
 * @Date: 2021-11-06 10:31:24
 * @LastEditTime: 2021-12-11 14:24:42
 * @Description: 
-->
<template>
<div class="contain">
    <el-card>
    <el-form label-width="80px">
      <el-form-item label="学生名称" prop="studentName">
        <el-input v-model="studentName" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="学生性别">
        <el-radio-group v-model="studentSex">
          <el-radio label="男"></el-radio>
          <el-radio label="女"></el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="学生专业">
      <el-select v-model="studentProfessional" placeholder="请选择活动区域">
        <el-option label="软件技术" value="软件技术"></el-option>
        <el-option label="人力资源" value="人力资源"></el-option>
        <el-option label="会计" value="会计"></el-option>
        <el-option label="艺术" value="艺术"></el-option>
        <el-option label="机器人" value="机器人"></el-option>
        <el-option label="体育" value="体育"></el-option>
      </el-select>
      </el-form-item>
      <el-form-item label="手机号码">
        <el-input v-model="studentPhone" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item label="所属分组">
      <el-select v-model="studentGroup" placeholder="请选择活动区域">
        <el-option label="一" value="G1"></el-option>
        <el-option label="二" value="G2"></el-option>
        <el-option label="三" value="G3"></el-option>
        <el-option label="四" value="G4"></el-option>
      </el-select>
      </el-form-item>
      <el-form-item label="座位号">
        <el-input v-model="studentSeat" placeholder="请输入内容"></el-input>
      </el-form-item>
      <el-form-item>
        <router-link :to="{name:'system'}">
          <el-button type="danger">退出</el-button>
        </router-link>
        <router-link :to="{name:'student'}">
          <el-button type="primary" @click="commitStudentHandler">提交</el-button>
        </router-link>
    </el-form-item>
    </el-form> 
  </el-card>
</div>
  
</template>

<script>
export default {
    name:"studentRegister",
    data() {
      return {
        studentName: "",
        studentSex: "",
        studentProfessional:"",
        studentPhone:"",
        studentGroup:"",
        studentSeat:"",
        tableData: [],
      }
    },
    methods: {
    // 提交班級
    commitStudentHandler() {
      this.loading = true;
      var that = this;
      // 提交到後台
      this.$POST("/addstudent", {
        name: this.studentName ,
        sex:this.studentSex,
        professional:this.studentProfessional,
        phone:this.studentPhone,
        group:this.studentGroup,
        seat:this.studentSeat
      }, function(res) {
        if (res.code != 200) {
          that.$message.error(res.message);
        } else {
          that.$message.success(res.message);
          that.tableData.push({
            id: res.data.id,
            name: res.data.name,
            createAt: res.data.createAt,
            sex: res.data.sex,
            professional: res.data.professional,
            phone: res.data.phone,
            group: res.data.group,
            seat: res.data.seat,
          });
        }
      });
    }
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
  width: 60%;
  margin: auto;
  filter: opacity(90%);
  position: absolute;
  left: 20%;
  top: 10%;
}
.el-button{
  margin-left: 10px;
}
</style>