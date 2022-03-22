<!--
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:29:48
 * @LastEditTime: 2021-12-11 16:55:22
 * @Description: 
-->
<template>
  <div>
    <el-card>
      <el-tabs>
      <el-tab-pane label="旷课统计图" name="first">
        <div id="main" style="width: 600px;height:400px;"></div
      ></el-tab-pane>
      <el-tab-pane label="早退请假图" name="second">
        <div id="main2" style="width: 600px;height:400px;"></div>
      </el-tab-pane>
      <el-tab-pane label="请假统计图" name="third">
        <div id="main3" style="width: 600px;height:400px;"></div>
      </el-tab-pane>
    </el-tabs>
    </el-card>
    
  </div>
</template>
<script>
export default {
  name: "",
  created() {
    this.loadData();
  },
  methods: {
    loadData() {
      var that = this;
      this.$POST("/student", {}, function(res) {
        console.log("student", res);
        if (res.code == 200) {
          var temp = [];
          for (const item of res.data) {
            temp.push(item.name);
          }
          // console.log(123321);
          console.log(temp);
          that.drawChart(temp);
        }
      });
      // console.log("ssss", this.s);
    },
    drawChart(data) {
      // 基于准备好的dom，初始化echarts实例
      var myChart = this.$echarts.init(document.getElementById("main"));
      var myChart2 = this.$echarts.init(document.getElementById("main2"));
      var myChart3 = this.$echarts.init(document.getElementById("main3"));
      // 指定图表的配置项和数据
      var option = {
        title: {
          text: "旷课统计图"
        },
        color: ["blue"],

        tooltip: {},
        legend: {
          data: ["人数"]
        },
        yAxis: {
          type: "category",
          data: data
        },
        xAxis: {},
        series: [
          {
            name: "人数",
            type: "bar",
            data: [5, 20, 36, 10, 10, 20]
          },
          {}
        ]
      };
      var option2 = {
        title: { text: "早退统计图" },
        tooltip: {},
        color: ["red"],
        legend: {
          data: ["人数"]
        },
        yAxis: {
          data: data
        },
        xAxis: {},
        series: [
          {
            name: "人数",
            type: "bar",
            data: [10, 20, 30, 14, 50, 60]
          }
        ]
      };
      var option3 = {
        title: { text: "请假统计图" },
        tooltip: {},
        color: ["yellow"],
        legend: {
          data: ["节"]
        },
        yAxis: {
          data: data
        },
        xAxis: {},
        series: [
          {
            name: "节",
            type: "bar",
            data: [10, 2, 13, 1, 5, 6]
          }
        ]
      };
      // 使用刚指定的配置项和数据显示图表。
      myChart.setOption(option);
      myChart2.setOption(option2);
      myChart3.setOption(option3);
    }
  }
};
</script>


<style scoped>
.el-card {
  margin: 10px;
  filter: opacity(90%);
}
</style>
