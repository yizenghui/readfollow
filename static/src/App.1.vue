<template>
  <div id="app">
    <header-tab></header-tab>
    <router-view/>
        

    <!----> 
    <vm-back-top :bottom="90" :timing="'ease'">
        <svg class="icon" width="30px" height="30.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="rgba(0, 8, 255, 0.50)" d="M511.5648 358.4a12.7616 12.7616 0 0 0-9.6 3.8784L355.584 512.9088a12.8 12.8 0 1 0 18.368 17.8432L499.2 401.856V729.6a12.8 12.8 0 0 0 25.6 0V400.512l120.768 126.8736a12.8 12.8 0 1 0 18.5472-17.6512l-140.288-147.3536a12.7744 12.7744 0 0 0-10.5472-3.9168 12.9536 12.9536 0 0 0-1.7152-0.0512zM512 1024C229.2352 1024 0 794.7648 0 512S229.2352 0 512 0s512 229.2352 512 512-229.2352 512-512 512zM320 320h384a12.8 12.8 0 0 0 0-25.6H320a12.8 12.8 0 0 0 0 25.6z" /></svg>
        <!-- <svg class="icon" width="30px" height="30.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="rgba(0, 8, 255, 0.57)" d="M936.3 667.6L823.1 780.7 512 469.6 200.9 780.7 87.7 667.6 512 243.3l424.3 424.3z"  /></svg> -->
    </vm-back-top> 
    <!-- <div class="weui-footer">
        <p class="weui-footer__text">粤ICP备17094378号-1</p>
    </div> -->
  </div>
</template>
<style>
.weui-panel {
    background-color: #fff;
     margin-top: 0; 
    position: relative;
    overflow: hidden;
}
.weui-media-box_appmsg .weui-media-box__thumb {
    width: 100%;
    height: 100%;
    vertical-align: top;
}



.vm-back-top{
      line-height: 0.8;
}
.vm-back-top-inner {
    background-color: rgba(0, 0, 0, 0);
}
.vm-back-top-inner:hover {
    background-color: rgba(255, 0, 0, 0.17);
}
.vm-back-top i {
    color: rgba(0, 8, 255, 0.57);
}
</style>
<script>

// /**
//  */
//  // 微信公众平台分享 (目前没有权限)
// const wx = require('weixin-js-sdk')

// import api from './api';


//  api.get("/jssdk?url=http://readfollow.com/hot",function(err,data){
//     wx.config({
//       debug: true, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
//       appId: "wx267866e82ab809fc", // 必填，公众号的唯一标识
//       timestamp: data.timestamp, // 必填，生成签名的时间戳
//       nonceStr: data.nonceStr, // 必填，生成签名的随机串
//       signature: data.signature, // 必填，签名，见附录1
//       jsApiList: ['checkJsApi', 'onMenuShareTimeline', 'onMenuShareAppMessage'] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
//     })

    
//     wx.ready(function () {
//         wx.checkJsApi({jsApiList: ['checkJsApi', 'onMenuShareTimeline', 'onMenuShareAppMessage']});
//         wx.onMenuShareTimeline({
//             title: '跟读，优质微信文章聚合平台',
//             link: "http://readfollow.com/hot",
//             imgUrl: 'http://readfollow.com/logo.png'
//         });
//         wx.onMenuShareAppMessage({
//             title: '跟读，微信优质文章聚合平台',
//             desc: '在这里阅读优质的微信文章，能节省了很多翻看订阅号的时间。',
//             link: "http://readfollow.com/hot",
//             imgUrl: 'http://readfollow.com/logo.png'
//         })
//     })
//   })

// import BackTop from '@/components/BackTop'
import HeaderTab from '@/components/Header';
export default {
  name: 'app',
  components:{
    HeaderTab,
    // BackTop,
  }
}
</script>

