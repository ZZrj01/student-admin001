/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-17 22:48:06
 * @Last Modified by: WUMIANHUA
 * @Last Modified time: 2021-10-17 23:03:40
 */
package api

import (
	"fmt"
	"server/model"
	"server/serve"
	"server/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 添加学生
 * @param {*gin.Context} c
 * @return {*}
 */

func AddStudent(c *gin.Context) {
	var query model.Student
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	// 继续执行
	query.ID = utils.GetUUID()
	t := time.Now().Unix()
	query.CreateAt = time.Unix(t, 0).Format("2006-01-02")
	err = serve.AddStudent(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	// 返回新建的id
	utils.Success(c, query)
}

/**
 * @description: 查询所有的学生
 * @param {*gin.Context} c
 * @return {*}
 */
func GetStudent(c *gin.Context) {
	// 调用数据层方法SelectStudent
	token := c.Request.Header.Get("Authorization")
	fmt.Println(token)
	list, err := serve.SelectStudent()
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
func DelStudent(c *gin.Context) {
	var query model.Student
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	err = serve.DelStudent(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, nil)

}

/**
 * @description: 更新学生
 * @param {*}
 * @return {*}
 */
func ModifyStudent(c *gin.Context) {
	var query model.Student
	err := c.ShouldBindJSON(&query)
	if err != nil {
		utils.Error(c, 1001, "参数错误！")
		return
	}
	//判断name不能为空
	//服务端必须验证数据，不能完全相信前端数据效验
	//严格写法
	//if strings.Trim(query.Name, "") == "" || len(query.Name) == 0 {
	if len(strings.Trim(query.Name, "")) == 0 {
		utils.Error(c, 1001, "学生名称不能为空")
		return
	}
	//判断一下学生名字是否重复了，重复了就不允许提交
	err = serve.ModifyStudent(query)
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	utils.Success(c, nil)
}
