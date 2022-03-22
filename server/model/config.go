/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:28:58
 * @LastEditTime: 2021-11-07 16:24:28
 * @Description:
 */
package model

// 基础根对象
type Config struct {
	Mysql  Mysql
	Server Server
}

// mysql 配置
type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DataName string `yaml:"dataName"`
	UserName string `yaml:"userName"`
	UserPwd  string `yaml:"userPwd"`
	Charset  string `yaml:"charset"`
}

// 服务基本参数
type Server struct {
	Port int `yaml:"port"`
}
