<!--
 * @Author: WUMIANHUA
 * @Date: 2021-11-07 11:18:47
 * @LastEditTime: 2021-11-30 19:21:11
 * @Description: 
-->
<template>
    <div class="contain">
        <el-card>
            <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-form-inline">
                <el-form-item
                    label="手机号码"
                    prop="phone"
                    :rules="[
                        { required:true,message:'手机号码不能为空'},
                        { type:'number',message:'年龄必须为数字值'},
                        { min:11, max:11,message:'长度在 11 位',trigger: 'blur'}
                    ]"
                >
                <el-input v-model="ruleForm.name"></el-input>
            </el-form-item>
              <el-form-item label="考勤类型">
                <el-input v-model="ruleForm.name" label-width="50px"></el-input>
              </el-form-item>
              <el-button type="primary">查询</el-button>
            </el-form>
            <el-table :data="tableData" style="width: 100%">
              <el-table-column prop="createAt" label="时间"> </el-table-column>
              <el-table-column type="index" label="状态"> </el-table-column>
              <el-table-column prop="name" label="节次"> </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>

<script>
export default {
    name:'studentDetails', 
    data() {
      return {
        ruleForm: {
          phone: '',
        },
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
    }
  }
</script>

<style scoped>
.contain{
    width: 100%;
    height: 100%;
    background-image: url("../../assets/u0_div.png");
    background-size: 100%;
    position: relative;
}
.el-card {
  width: 80%;
  margin: auto;
  filter: opacity(90%);
  position: absolute;
  left: 10%;
  top: 10%;
  filter: opacity(90%);
}
</style>