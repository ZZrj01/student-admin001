/*
 * @Author: WUMIANHUA
 * @Date: 2021-11-02 20:54:08
 * @LastEditTime: 2021-11-07 16:26:20
 * @Description: 
 */
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
const store = new Vuex.Store({
    state: {
        token: "",
    },
    mutations: {
        setToken(state, token) {
            AudioContext.commit('setToken', token)
        }
    }
})
export default store