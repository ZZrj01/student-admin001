/*
 * @Author: WUMIANHUA 
 * @Date: 2021-10-17 22:41:27 
 * @Last Modified by: WUMIANHUA
 * @Last Modified time: 2021-10-24 22:45:15
 */
import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

import main from '@/views/main/index.vue'
// 路由成员
var routes = [{
    path: "/login",
    name: "login",
    hide: true,
    component: () => import('@/views/login/index'),
},{
    path:'/studentMain',
    name:"studentMain",
    hide: true,
	component: () => import('@/views/studentMain/studentMain'),
},{
    path:'/studentRegister',
    name:"studentRegister",
    hide: true,
	component: () => import('@/views/student/studentRegister'),
},{
    path: "/studentDetails",
    name: "studentDetails",
    hide: true,
    component: () => import('@/views/studentMain/studentDetails'),
},{
    path: "/",
    name: "main",
    component: main,
    children: [{
            path: "/statistics",
            name: "statistics",
            isturn:true,
            component: () => import('@/views/statistics/index'),
            meta: {
                title: '统计',
            },
            icon: "el-icon-eleme",
        }, {
            path: "/student",
            name: "student",
            isturn:true,
            component: () => import('@/views/student/index'),
            meta: {
                title: "学生管理",
                // isturn:true
            },
            icon: "el-icon-user",
        },
        {
            path: "/attendance",
            name: "attendance",
            isturn:true,
            component: () => import('@/views/attendance/index'),
            meta: {
                title: "考勤管理",
                // isturn:true
            },
            icon: "el-icon-user",
        },
        {
            path: "/class",
            name: "class",
            isturn:true,
            component: () => import('@/views/class/index'),
            meta: {
                title: "班级管理",
                // isturn:true
            },
            icon: "el-icon-user",
        },
        {
            path: "/system",
            name: "system",
            isturn:true,
            component: () => import('@/views/system/index'),
            meta: {
                title: "系统管理",
                // isturn:true
            },
            icon: "el-icon-user",
        },
        {
            path: "/attendanceDetail",
            name: "attendanceDetail",
            // hide: true,
            isturn:false,
            component: () => import('@/views/attendance/attendanceDetail'),
            meta: {
                title: "考勤详情",
                // isturn:false
            },
            icon: "el-icon-user",
        },
    ],
}]
// 声明路由
const router = new VueRouter({
    routes
})

//路由前置守卫
router.beforeEach((to, from, next) => {
    //调转path不是登录
    if (to.path != "/login") {
        var token = sessionStorage.getItem("token")
        // 空
        if (token == "" || token == undefined) {
            router.replace("/login")
        }
    }
    next()
})
// 导出路由
export default router