// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.

import 'weui/dist/style/weui.min.css'
import 'swiper/dist/css/swiper.css'
// import '../assets/icon/inonfont.css'
// import '@assets/iconfont.css'

import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false


import axios from 'axios'
import VueAxios from 'vue-axios'

// import VueLazyload from 'vue-lazyload'

// import quickMenu from 'vue-quick-menu'
// Vue.component(quickMenu.name,quickMenu)

import VueTimeago from 'vue-timeago'


// var VueAwesomeSwiper = require('vue-awesome-swiper')

import VmBackTop from 'vue-multiple-back-top'


/**
 */
 // 微信公众平台分享 (目前没有权限)
//  const wx = require('weixin-js-sdk')
 
 import api from './api';
 

// import VueScrollPicker from 'vue-scroll-picker'

// Vue.use(VueScrollPicker)

Vue.use(VueTimeago, {
    name: 'timeago', // component name, `timeago` by default
    locale: 'zh-CN',
    locales: {
        // you will need json-loader in webpack 1
        // 'zh-CN': require('vue-timeago/locales/zh-CN.json')
        // 自定义时间格式
        'zh-CN':[
            "",
            "",
            "",
            "%sh",
            "%sd",
            "%sw",
            "%sm",
            "%sy"
            ]
    }
})

Vue.use(VueAxios, axios)
// Vue.use(VueLazyload)




Vue.component(VmBackTop.name, VmBackTop)

// mount with global
// Vue.use(VueAwesomeSwiper)

// router.afterEach(function () {
//     alert("after");
//     });

    router.beforeEach((to, from, next) => {
 
        // setTimeout(function(){
        //     api.get("/jssdk?url=http://readfollow.com"+to.path,function(err,data){
        //         wx.config({
        //           debug: false, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
        //           appId: "wx267866e82ab809fc", // 必填，公众号的唯一标识
        //           timestamp: data.timestamp, // 必填，生成签名的时间戳
        //           nonceStr: data.nonceStr, // 必填，生成签名的随机串
        //           signature: data.signature, // 必填，签名，见附录1
        //           jsApiList: ['checkJsApi', 'onMenuShareTimeline', 'onMenuShareAppMessage'] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
        //         })
            
                
        //         wx.ready(function () {
        //             wx.checkJsApi({jsApiList: ['checkJsApi', 'onMenuShareTimeline', 'onMenuShareAppMessage']});
        //             wx.onMenuShareTimeline({
        //                 title: '跟读，优质微信文章聚合平台',
        //                 link: "http://readfollow.com"+to.path,
        //                 imgUrl: 'http://readfollow.com/logo.png'
        //             });
        //             wx.onMenuShareAppMessage({
        //                 title: '跟读，微信优质文章聚合平台',
        //                 desc: '阅读优质的微信文章，节省翻看订阅号的时间。',
        //                 link: "http://readfollow.com"+to.path,
        //                 imgUrl: 'http://readfollow.com/logo.png'
        //             })
        //         })
        //       })
        // }, 200);

        // var __to = localStorage.getItem("__to")||''
        // // to 和 from 都是 路由信息对象
        
        // // console.log(to.path ,from.path , __to ,to.path )
        // if(to.path != from.path && __to !=to.path ){
        //   localStorage.removeItem("articles")
        //   localStorage.removeItem("rank")
        // //   console.log("clear cache")
        // }
        
        // localStorage.setItem("__to",to.path)
        // console.log(to.path, from.path)
        next()
      })

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
