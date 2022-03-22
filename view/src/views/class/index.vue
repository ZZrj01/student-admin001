<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-26 17:52:54
 * @LastEditTime: 2021-11-09 19:44:18
 * @Description: 
-->
<template>
  <div>
    <el-card class="box-card">
      <el-button type="primary" @click="dialogVisible = true"
        >添加班级</el-button
      >
      <el-table :data="tableData" style="width: 100%">
        <el-table-column type="index" width="50" label="序号"> </el-table-column>
        <!-- <el-table-column prop="id" label="班级ID"> </el-table-column> -->
        <el-table-column prop="name" label="班级名称"> </el-table-column>
        <el-table-column prop="createAt" label="创建时间"> </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="danger"
              @click="handleDelete(scope.$index, scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <!-- 弹窗 -->
    <el-dialog
      v-loading="loading"
      title="添加班级"
      :visible.sync="dialogVisible"
      width="30%"
    >
      <el-input v-model="className" placeholder="请输入班级名称"></el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="commitClassHandler">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 编辑弹窗 -->
    <el-dialog
      v-loading="loading"
      title="编辑班级"
      :visible.sync="modifyDialog"
      width="30%"
    >
      <el-input
        v-model="modifyForm.name"
        placeholder="请输入班级名称"
      ></el-input>
      <span slot="footer" class="dialog-footer">
        <el-button @click="modifyDialog = false">取 消</el-button>
        <el-button type="primary" @click="modifyClassHandler">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
export default {
  name: "ClassAdmin",
  data() {
    return {
      loading: false, // dialog loading
      className: "",
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
    commitClassHandler() {
      this.loading = true;
      var that = this;
      // 提交到後台
      this.$POST("/addclass", { name: this.className }, function(res) {
        // 收到服务器反馈时候就停止loading
        that.loading = false;
        if (res.code != 200) {
          that.$message.error(res.message);
        } else {
          that.$message.success(res.message);
          // 关闭弹窗
          that.dialogVisible = false;
          //
          that.tableData.push({
            id: res.data.id,
            name: res.data.name,
            createAt: res.data.createAt,
          });
        }
      },
      // 清空输入对象
      this.className=""
      );
    },
    // 加载列表数据
    loadData() {
      var that = this;
      this.$POST("/class", {}, function(res) {
        console.log("class的vc:",res);
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
          this.$POST("/delclass", { id: row.id }, function(res) {
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
      this.modifyIndex = index;
      //assign合并两个对象生成一个新的对象
      this.modifyForm = Object.assign({}, row);
      this.modifyDialog = true;
    },
    //提交编辑班级
    modifyClassHandler() {
      var that = this;
      this.$POST("/modifyclass", this.modifyForm, function(res) {
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
