<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-25 07:59:42
 * @LastEditTime: 2021-12-07 19:42:11
 * @Description: 
-->
<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="dialogVisible = true"
        >添加学生</el-button
      >
      <el-table :data="tableData" style="width: 100%">
        <el-table-column type="index" width="50" label="序号"> </el-table-column>
        <!-- <el-table-column prop="id" label="学生ID"> </el-table-column> -->
        <el-table-column prop="name" label="学生姓名"> </el-table-column>
        <el-table-column prop="createAt" label="创建时间"> </el-table-column>
        <el-table-column prop="sex" label="性别"> </el-table-column>
        <el-table-column prop="professional" label="班级"> </el-table-column>
        <el-table-column prop="phone" label="手机号"> </el-table-column>
        <el-table-column prop="group" label="分组"> </el-table-column>
        <el-table-column prop="seat" label="座位"> </el-table-column>
        <el-table-column label="操作" width="150">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <!-- 弹窗 -->
    <el-dialog
      v-loading="loading"
      title="添加学生"
      :visible.sync="dialogVisible"
      width="30%"
    >
    <el-form label-width="80px">
      <el-form-item label="学生名称">
        <el-input v-model="studentName"></el-input>
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
        <el-input v-model="studentPhone"></el-input>
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
        <el-input v-model="studentSeat"></el-input>
      </el-form-item>
    </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="commitStudentHandler">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 编辑弹窗 -->
    <el-dialog
      v-loading="loading"
      title="编辑学生"
      :visible.sync="modifyDialog"
      width="30%"
    >
      <el-form label-width="150px">
      <el-form-item label="学生名称">
        <el-input v-model="modifyForm.name"></el-input>
      </el-form-item>
      <el-form-item label="学生性别">
        <el-radio-group v-model="modifyForm.sex">
          <el-radio label="男"></el-radio>
          <el-radio label="女"></el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="学生专业">
      <el-select v-model="modifyForm.professional" placeholder="请选择活动区域">
        <el-option label="软件技术" value="软件技术"></el-option>
        <el-option label="人力资源" value="人力资源"></el-option>
        <el-option label="会计" value="会计"></el-option>
        <el-option label="艺术" value="艺术"></el-option>
        <el-option label="机器人" value="机器人"></el-option>
      </el-select>
      </el-form-item>
      <el-form-item label="手机号码">
        <el-input v-model="modifyForm.phone"></el-input>
      </el-form-item>
      <el-form-item label="所属分组">
      <el-select v-model="modifyForm.group" placeholder="请选择活动区域">
        <el-option label="一" value="G1"></el-option>
        <el-option label="二" value="G2"></el-option>
        <el-option label="三" value="G3"></el-option>
        <el-option label="四" value="G4"></el-option>
      </el-select>
      </el-form-item>
      <el-form-item label="座位号">
        <el-input v-model="modifyForm.seat"></el-input>
      </el-form-item>
    </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="modifyDialog = false">取 消</el-button>
        <el-button type="primary" @click="modifyStudentHandler">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: "",
  data() {
    return {
      loading: false, // dialog loading
      studentName: "",
      studentSex: "",
      studentProfessional:"",
      studentPhone:"",
      studentGroup:"",
      studentSeat:"",
      dialogVisible: false, // 显示弹窗
      tableData: [],
      //modifyForm编辑时候的对象
      modifyForm: {},
      modifyDialog: false, //控制编辑
      modifyIndex: 0,
    };
  },
  // 准备要显示的数据
  created() {
    this.loadData();
    //el初始化，创建data里面的数据
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
        // 收到服务器反馈时候就停止loading
        that.loading = false;
        if (res.code != 200) {
          that.$message.error(res.message);
        } else {
          that.$message.success(res.message);
          // 关闭弹窗
          that.dialogVisible = false;
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
    },
    // 加载列表数据
    loadData() {
      var that = this;
      this.$POST("/student", {}, function(res) {
        console.log("student的vc：",res);
        if (res.code == 200) {
          that.tableData = res.data;
        }
      });
    },
    // 删除班级
    handleDelete(index, row) {
      console.log(index, row);
      var that = this;
      this.$confirm("你确定要删除该数据?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$POST("/delstudent", { id: row.id }, function(res) {
            if (res.code == 200) {
              that.tableData.splice(index, 1);
            } else {
              that.$message.error(res.message);
            }
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    //编辑班级
    handleEdit(index, row) {
      //console.log(index, row);
      this.modifyIndex = index;
      //assign合并两个对象生成一个新的对象
      this.modifyForm = Object.assign({}, row);
      this.modifyDialog = true;
    },
    //提交编辑班级
    modifyStudentHandler() {
      var that = this;
      this.$POST("/modifystudent", this.modifyForm, function(res) {
        if (res.code == 200) {
          //1.使用js深浅拷贝实现创建新变量
          var temp = JSON.parse(JSON.stringify(that.tableData));
          temp[that.modifyIndex] = that.modifyForm;
          that.tableData = temp;
          console.log(that.tableData);
          that.modifyDialog = false;
        } else {
          that.$message.error(res.message);
        }
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.el-card {
  margin: 10px;
  filter: opacity(90%);
}
</style>
