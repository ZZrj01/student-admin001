/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:24:23
 * @Description:
 */
package main

import (
	"fmt"
	"os"
	"server/config"
	"server/source"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
)

// 初始化
func init() {
	// 初始化的时候读取config 配置到缓存里
	// 读取config.yaml 文件字节
	by, err := os.ReadFile("config.yaml")
	// 如果错误等于nil,说明是成功
	if err == nil {
		// 解析by字节到config 对象
		err = yaml.Unmarshal(by, &config.Conf)
		if err != nil {
			// 解析错误就直接终止程序
			fmt.Println(err)
			// 非0就终止程序
			os.Exit(1)
		}
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	// 数据库链接
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.Conf.Mysql.UserName, config.Conf.Mysql.UserPwd, config.Conf.Mysql.Host, config.Conf.Mysql.Port, config.Conf.Mysql.DataName, config.Conf.Mysql.Charset))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// gorm debug
	db.LogMode(true)
	config.DB = db
}

func main() {
	fmt.Println("测试调试")
	// fmt.Println(utils.GetUUID())
	// 调用路由
	source.CreateRouter()

}
