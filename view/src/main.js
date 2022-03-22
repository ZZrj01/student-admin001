/*
 * @Author: WUMIANHUA
 * @Date: 2021-10-31 16:29:48
 * @LastEditTime: 2021-12-11 16:55:02
 * @Description: 
 */
import Vue from 'vue'
import App from './App.vue'
import router from './router';
//引入并使用element-ui
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
Vue.use(ElementUI);
import * as echarts from "echarts";
Vue.prototype.$echarts = echarts;
import "@/style/style.scss"
Vue.config.productionTip = false
// 给Vue 添加一个属性
import {
  POST
} from './utils/http'
Vue.prototype.$POST = POST

new Vue({
  render: h => h(App),
  router
}).$mount('#app')