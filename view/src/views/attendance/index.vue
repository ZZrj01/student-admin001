<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-25 07:59:42
 * @LastEditTime: 2021-12-17 00:06:25
 * @Description: 
-->
<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="dialogVisible = true">立即点名</el-button>
      <el-table :data="tableData" style="width: 100%">
        <el-table-column type="index" width="50" label="序号"> </el-table-column>
        <!-- <el-table-column prop="id" label="班级ID"> </el-table-column> -->
        <el-table-column prop="create_at" label="开始时间"> </el-table-column>
        <el-table-column prop="finish_at" label="结束时间"> </el-table-column>
        <el-table-column prop="session" label="节次"> </el-table-column>
        <!-- <el-table-column prop="classnum" label="班级人数"> </el-table-column> -->
        <el-table-column label="操作">
          <template slot-scope="scope">
            <!-- <router-link to="/attendanceDetail" class="routerlink"> -->
            <el-button size="mini" type="primary" @click="attendanaceShow">详情</el-button>
            <!-- </router-link> -->
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
              >删除
              </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <!-- 弹窗 -->
    <el-dialog title="立即点名" :visible.sync="dialogVisible" width="30%"> 
      <el-form label-width="30%">
        <el-form-item label="点名时间限制">
          <el-form-item label="10分钟"></el-form-item>
        <!-- <el-select v-model="attendanceTime" placeholder="10分钟">
          <el-option label="10分钟" value="10"></el-option>
          <el-option label="15分钟" value="15"></el-option>
          <el-option label="20分钟" value="20"></el-option>
          <el-option label="30分钟" value="30"></el-option>
        </el-select> -->
      </el-form-item>
      <el-form-item label="节次">
        <el-select v-model="attendanceSession" placeholder="1-2">
          <el-option label="1-2" value="1-2"></el-option>
          <el-option label="3-4" value="3-4"></el-option>
          <el-option label="5-6" value="5-6"></el-option>
          <el-option label="7-8" value="7-8"></el-option>
          <el-option label="9-10" value="9-10"></el-option>
          <el-option label="11-12" value="11-12"></el-option>
        </el-select>
      </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button
          type="primary" 
          @click="commitAttendence"
        >确 定</el-button>
      </span>
    </el-dialog>
    <!-- 签到弹窗 -->
    <el-dialog title="提示" :visible.sync="dialogQiandao" width="30%">
      <h2>请访问 http://192.168.1.100/sign 进行签到</h2>
      <el-button type="primary" @click="commitAttendanceHandler">确定</el-button>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: "AttendanceAdmin",
  data() {
    return {
      dialogVisible: false, // 显示弹窗
      dialogAttendance: false, // 显示签到时间确定弹窗
      dialogQiandao:false, //显示点击签到弹窗
      tableData: [],
      attendanceTime:"",
      attendanceCreate_at:"",
      attendanceFinish_at:"",
      attendanceSession:this.session,
      attendanceClassNum:"",
    };
  },
  // 准备要显示的数据
  created() {
    this.loadData();
    //el初始化，创建data里面的数据
  },
  mounted() {
    console.log("finish_at的",this.attendanceFinish_at);
  },
  methods: {
    commitAttendence(){
      this.dialogVisible = false
      this.dialogAttendance = true
      this.dialogQiandao = true
    },
    commitAttendanceHandler(){
      // console.log("seesion",this.session);
      this.loading = true;
      this.dialogQiandao = false
      var that = this;
      // 提交到後台
      this.$POST("/addattendance", {
        create_at: this.attendanceCreate_at ,
        finish_at:this.attendanceFinish_at,
        attendanceTime:this.attendanceTime,
        session:this.attendanceSession,
        classnum:this.attendanceClassNum,
      }, function(res) {
        // 收到服务器反馈时候就停止loading
        that.loading = false;
        if (res.code != 200) {
          that.$message.error(res.message);
        } else {
          that.$message.success(res.message);
          // 关闭弹窗
          that.tableData.push({
            create_at: res.data.create_at,
            finish_at: res.data.finish_at,
            session: res.data.session,
            classnum: res.data.classnum,
          });
        }
      });
    },
    // 加载列表数据
    loadData() {
      var that = this;
      this.$POST("/attendance", {}, function(res) {
        console.log("attendance的vc:",res);
        if (res.code == 200) {
          that.tableData = res.data;
        }
      });
    },
    attendanaceShow(){
      this.$router.push({
        name:'attendanceDetail',
      })
    },
    // 删除考勤
    handleDelete(index, row) {
      console.log(index, row);
      var that = this;
      this.$confirm("你确定要删除该数据?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$POST("/delattendance", { id: row.id }, function(res) {
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
  }
};
</script>
<style lang="scss" scoped>
.el-card {
  margin: 10px;
  filter: opacity(90%);
}
.routerlink{
  margin-right: 10px;
}
</style>
