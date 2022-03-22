/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-12-14 19:31:26
 * @Description:
 */
package api

import (
	"fmt"
	"net"
	"server/model"
	"server/serve"
	"server/utils"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

/** ctrl+alt+t
 * @description: 声明一个登陆的接口
 * @param {*gin.Context} c
 * @return {*}
 */
func Login(c *gin.Context) {
	// 声明一个model.User类型的对象
	var query model.User
	// 把传递的参数解析到query对象里
	err := c.ShouldBindJSON(&query)
	// 假设err != nil 该解析过程出错了，就返回错误，如果 == nil 就是成功地
	if err != nil {
		// c.JSON(c)
		utils.Error(c, 1001, "参数错误！")
		return
	}
	// 拿用户名和密码进行匹配
	user, err := serve.IsHasUser(query)
	// 如果错误存在就提示错误终止请求
	if err != nil {
		utils.Error(c, 1001, err.Error())
		return
	}
	user.Passwd = ""
	//登录成功
	store := ginsession.FromContext(c)
	// 保存用户id到session
	store.Set(user.ID, user.ID)
	err = store.Save()
	if err != nil {
		fmt.Println("存储用户", 1001, "参数错误!")
		return
	}
	// utils.Success(c, user)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}
	ip := ""
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println((ipnet.IP.String()))
						ipnet.IP.String()
						ip = ipnet.IP.String()
					}
				}
			}
		}
	}
	query.ID = utils.GetUUID()
	if ip != "" {
		query.Ip = ip

		err = serve.ModifyIp(query)
		if err != nil {
			utils.Error(c, 1001, err.Error())
			return
		}
	}
	user.Ip = ip
	fmt.Println(user)
	utils.Success(c, user)
}

/**
 * @description: 退出登录
 * @param {*gin.Context} c
 * @return {*}
 */
func LogOut(c *gin.Context) {
	obj, b := c.Get("token")
	if b {
		//把interface转成string
		token := fmt.Sprintf("%v", obj)
		// 把token从session里删除
		store := ginsession.FromContext(c)
		store.Delete(token)
		err := store.Save()
		if err != nil {
			utils.Error(c, 1001, err.Error())
			return
		}
		utils.Success(c, nil)
	} else {
		//不带token属于非法请求
		utils.Error(c, 1001, "非法请求")
		return
	}
}
