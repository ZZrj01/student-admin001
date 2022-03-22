/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:29:48
 * @LastEditTime: 2021-12-17 07:57:44
 * @Description: 
 */

import axios from 'axios'
import router from '@/router'
import {
    Message,
    MessageBox
} from 'element-ui';


axios.defaults.baseURL = 'http://localhost/api';
// axios.defaults.baseURL = 'http://127.0.0.1:80/api';
axios.defaults.timeout = 5000; // 超时
axios.defaults.withCredentials = true; // 带cookie
//
axios.interceptors.request.use(function (config) {
    // console.log(config);
    if (config.url != "/login") {
        config.headers.Authorization = sessionStorage.getItem("token")
    }
    return config;
})
// 拦截响应信息
axios.interceptors.response.use(function (res) {
    // 对响应数据做点什么
    if (res.status == 200) { // 浏览器响应状态 非就是失败
        // res.data.code 是接口返回的状态值
        if (res.data.code == 200) {
            return res.data
        } else if (res.data.code == 1008) { //用户没登录
            sessionStorage.setItem("token", "")
            //1008用户未登录
            MessageBox({
                title: "提示!",
                message: "登录超时，请重新登录!",
                type: "error"
            }).then(function () {
                //点击确定
                //通知路由跳转
                router.replace("/login")
            })

        } else {
            return res.data
        }
    } else {
        // 提升错误
        Message.error(res);
        return Promise.reject(res);
    }
}, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
});
// 封装一个post 方法
/**
 * @description: 封装post方法
 * @param {*} url 请求地址
 * @param {*} data 请求参数
 * @param {*} cb 回调方法
 * @return {*}
 */
export function POST(url, data, cb) {
    // 调用全局axios post 方法
    axios.post(url, data).then(function (res) { // 成功回调
        // 如果cb是function 就回调
        if (typeof cb === "function") {
            cb(res)
        }
    }).catch(function (err) { // 失败回调
        if (typeof cb === "function") {
            cb(err)
        }
    })
}