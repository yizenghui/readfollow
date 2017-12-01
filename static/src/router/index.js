import Vue from 'vue'
import Router from 'vue-router'
// import homePage from '@/views/Home'
import hotPage from '@/views/Hot'
import CatePage from '@/views/Cate' // 进入某个分类
// import HelloWorld from '@/components/HelloWorld'
// import mediaPage from '@/views/Media'
// import newPage from '@/views/New'
// import aboutPage from '@/views/about'
// import guidePage from '@/views/guide'
// import postPage from '@/views/Post'
// import TagPage from '@/views/Tag' // 展示所有标签

// // 反馈模块
// import fbArticlePage from '@/pages/feedback/Article'
// import fbMediaPage from '@/pages/feedback/Media'
// import fbPrivacyPage from '@/pages/feedback/Privacy'
// import fbOtherPage from '@/pages/feedback/Other'
// import fbReportPage from '@/pages/feedback/Report'


// // 自助服务
import SVPostPage from '@/pages/service/Post'
// import SVDeletePage from '@/pages/service/Delete'
// import SVSpreadPage from '@/pages/service/spread'


// 测试 

// import swiperPage from '@/views/Swiper'

Vue.use(Router)

export default new Router({
    mode: 'history',
    linkExactActiveClass:"weui-bar__item_on",

    routes: [
        {
            path:'/',
            name: 'home',
            redirect: '/hot'
        },      
        {
            path:'/hot',
            name: 'hot',
            component:hotPage
        }, 
        {
            path:'/t/:id', 
            name: 'cate',
            component:CatePage
        },      
        // {
        //     path:'/swiper',
        //     name: 'swiper',
        //     component:swiperPage
        // },
        // {
        //     path:'/new',
        //     component:newPage
        // },
        // {
        //     path:'/post',
        //     component:postPage
        // },
        // {
        //     path:'/tags', 
        //     name: 'tags',
        //     component:TagPage
        // },
        // {
        //     path:'/p/:id',
        //     name: 'media',
        //     component:mediaPage
        // },
        // {
        //     path:'/about', 
        //     name: 'about',
        //     component:aboutPage
        // },
        // {
        //     path:'/guide', 
        //     name: 'guide',
        //     component:guidePage
        // },
        // {
        //     path:'/feedback/article', 
        //     name: 'feedbackarticle',
        //     component:fbArticlePage
        // },
        // {
        //     path:'/feedback/media', 
        //     name: 'feedbackmedia',
        //     component:fbMediaPage
        // },       
        // {
        //     path:'/feedback/report', 
        //     name: 'feedbackreport',
        //     component:fbReportPage
        // },    
        // {
        //     path:'/feedback/other', 
        //     name: 'feedbackother',
        //     component:fbOtherPage
        // },
        // {
        //     path:'/feedback/privacy', 
        //     name: 'feedbackprivacy',
        //     component:fbPrivacyPage
        // },     
        {
            path:'/service/post', 
            name: 'servicepost',
            component:SVPostPage
        }, 
        // {
        //     path:'/service/delete', 
        //     name: 'servicedelete',
        //     component:SVDeletePage
        // },
        // {
        //     path:'/service/spread', 
        //     name: 'servicespread',
        //     component:SVSpreadPage
        // },
  ]
})
