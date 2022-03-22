/*
 * @Author: WUMIANHUA
 * @Date: 2021-12-11 16:25:08
 * @LastEditTime: 2022-03-21 10:46:19
 * @Description:
 */
package api

import (
	"fmt"
	"server/model"
	"server/serve"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 添加考勤
 * @param {*gin.Context} c
 * @return {*}
 */

func AddAttendance(c *gin.Context) {
	var query model.Attendance
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	// 继续执行
	query.ID = utils.GetUUID()
	t := time.Now().Unix()
	// t1 := 60 * int64(query.attendanceTime)
	t1 := 60 * 10
	t2 := t + int64(t1)
	query.Finish_at = time.Unix(t2, 0).Format("2006-01-02 03:04")
	query.Create_at = time.Unix(t, 0).Format("2006-01-02 03:04")
	err = serve.AddAttendance(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	// 返回新建的id
	utils.Success(c, query)
}

/**
 * @description: 查询所有考勤信息
 * @param {*gin.Context} c
 * @return {*}
 */
func GetAttendance(c *gin.Context) {
	// 调用数据层方法SelectAttendance
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	list, err := serve.SelectAttendance()
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, list)
	// c.JSON(200, map[string]interface{}{
	// 	"code": 1008,
	// 	"data": list,
	// })
}

/**
 * @description: 删除学生
 * @param {*gin.Context} c
 * @return {*}
 */
func DelAttendance(c *gin.Context) {
	var query model.Attendance
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	err = serve.DelAttendance(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, nil)

}

/**
 * @description: 更新考勤（暂时不用）
 * @param {*}
 * @return {*}
 */
//  func ModifyAttendance(c *gin.Context) {
// 	 var query model.Attendance
// 	 err := c.ShouldBindJSON(&query)
// 	 if err != nil {
// 		 utils.Error(c, 1001, "参数错误！")
// 		 return
// 	 }
// 	 //判断name不能为空
// 	 //服务端必须验证数据，不能完全相信前端数据效验
// 	 //严格写法
// 	 //if strings.Trim(query.Name, "") == "" || len(query.Name) == 0 {
// 	 if len(strings.Trim(query.Name, "")) == 0 {
// 		 utils.Error(c, 1001, "学生名称不能为空")
// 		 return
// 	 }
// 	 //判断一下学生名字是否重复了，重复了就不允许提交
// 	 err = serve.ModifyStudent(query)
// 	 if err != nil {
// 		 utils.Error(c, 1001, err.Error())
// 		 return
// 	 }
// 	 utils.Success(c, nil)
//  }
