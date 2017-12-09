//index.vue
<template>

<div>


<div class="weui-panel_access">
   <!--<div class="weui-panel__hd">
    <span>
      跟读微信文章
    </span>
    

    <label for="weuiAgree" class="weui-agree">
        
              <router-link class="weui-agree__text" :to="{ name: 'tags'}">全部分类 </router-link>

    </label>
  </div> -->

  
  <ArticleList :articles="articles"></ArticleList>

    <infinite-loading @infinite="infiniteHandler" :distance="distance" ref="infiniteLoading">

    <div slot="spinner" class="weui-loadmore">
        <i class="weui-loading"></i>
    </div>

    <div slot="no-results" class="weui-loadmore weui-loadmore_line">
      <span class="weui-loadmore__tips">载入失败!</span>
    </div>

    <div slot="no-more" class="weui-loadmore weui-loadmore_line">
      <span class="weui-loadmore__tips">到底了!</span>
    </div>

  </infinite-loading> 

</div>
    
    <div class="refresh"  v-on:click="refresh">
      <svg class="icon" width="30px" height="30.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="rgba(0, 8, 255, 0.50)" d="M838.695385 374.153846A354.619077 354.619077 0 0 0 512 157.538462a354.461538 354.461538 0 1 0 0 708.923076 354.579692 354.579692 0 0 0 330.161231-225.20123 39.384615 39.384615 0 1 1 73.334154 28.750769A433.309538 433.309538 0 0 1 512 945.230769C272.738462 945.230769 78.769231 751.261538 78.769231 512S272.738462 78.769231 512 78.769231c144.423385 0 275.140923 71.286154 354.461538 183.965538V177.230769a39.384615 39.384615 0 0 1 78.769231 0v236.307693a39.266462 39.266462 0 0 1-39.384615 39.384615h-196.923077a39.384615 39.384615 0 0 1 0-78.769231h129.772308z" /></svg>
      <!-- <svg class="icon" width="30px" height="30.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="rgba(0, 8, 255, 0.57)" d="M958.174805 481.820654c-1.017166-35.661198-5.421476-91.061972-32.872688-144.119375L774.270309 509.142929l125.009132 0c0.002047 0.827855 0.00307 1.650593 0.00307 2.51017 0 213.726709-173.881213 387.607922-387.607922 387.607922-213.727733 0-387.608945-173.881213-387.608945-387.607922 0-210.105229 181.584662-387.608945 396.521942-387.608945 39.323611 0 78.106916 5.861498 115.27442 17.421788l17.693987-56.888661c-42.902112-13.343913-87.639013-20.10899-132.968407-20.10899-60.358693 0-119.340015 11.80486-175.303608 35.0861-54.135968 22.520923-102.939487 54.68139-145.05263 95.588055-42.355667 41.141002-75.673493 88.934518-99.028411 142.051273-24.360827 55.404867-36.713156 114.100688-36.713156 174.459381 0 119.447463 46.515407 231.745085 130.977054 316.207755C279.929504 912.322501 392.227126 958.837908 511.674589 958.837908c119.447463 0 231.745085-46.515407 316.207755-130.977054s130.977054-196.760293 130.977054-316.207755C958.859397 499.757165 958.672132 490.187205 958.174805 481.820654z" /></svg> -->
    </div>

</div>


</template>
<style>
.article-title{
  color: #000;
}
.refresh{
    width:30px;
    height: 30px;
    line-height: 30px;
    bottom: 50px;
    right: 30px;
    z-index: 10;
    position: fixed;
    cursor: pointer;
}
.category{
    width:30px;
    height: 30px;
    line-height: 30px;
    bottom: 10px;
    right: 30px;
    z-index: 10;
    position: fixed;
    cursor: pointer;
} 
</style>
<script>

import ArticleList from '@/components/ArticleList';
import InfiniteLoading from 'vue-infinite-loading'
import api from '../api';

export default {
  
  name: 'indexP',
  
  
    components: {
      InfiniteLoading,
      ArticleList
    },
    data () {
      return {
        articles: [],
        distance: 200,
        page:0,
        showload:false,
        rank:0, // 当前下拉文章最低rank
        loading:false,
        cate: 0,
        err:'',

      }
    },
    // beforeRouteEnter (to, from, next) {
    //   console.log('beforeRouteEnter',to.params.id)
    //   this.refresh()
    //   var uri = '/hot?limit=10&rank=0&tag='+to.params.id;
    //   api.get(uri, (err, post) => {
    //     next(vm => vm.setData(err, post))
    //   })
    // },
    // // 路由改变前，组件就已经渲染完了
    // // 逻辑稍稍不同
    // beforeRouteUpdate (to, from, next) {
      
    //   console.log('beforeRouteUpdate',to.params.id)
    //   // this.post = null
    //   // getPost(to.params.id, (err, post) => {
    //   //   this.setData(err, post)
    //   //   next()
    //   // })
    // },

    // beforeRouteUpdate (to, from, next) {
    // // react to route changes...
    // // don't forget to call next()
    // console.log('beforeRouteUpdate',to.params.id, from, next)
    // },
    mounted() {
      if (this.$route.params.id != undefined){
        this.cate = this.$route.params.id
      }else{
        this.getCache()
      }
    },
    methods: {

      // setData (err, articles) {
      //   if (err) {
      //     this.error = err.toString()
      //   } else {
      //     this.articles = articles
      //   }
      // },

      // 从客户端缓存中恢复
      getCache:function(){

        var timestamp = (new Date()).valueOf();
        var cate = this.cate
        var time = parseInt(window.localStorage.getItem("cache_time"+cate)) || 0
        if (time>timestamp){
          this.articles = JSON.parse(window.localStorage.getItem("articles"+cate))||[]
          this.rank = parseFloat(window.localStorage.getItem("rank"+cate)) || 0
          this.$refs.infiniteLoading.stateChanger.reset()
        }else{
          this.refresh()
        }
        // console.log('getCache',cate)
       
      },
      // 设置缓存
      setCache:function(){
        var cate = this.cate
        var rank = this.rank
        var articles = this.articles
        var timestamp = (new Date()).valueOf();
        localStorage.setItem("articles"+cate,JSON.stringify(articles))
        localStorage.setItem("rank"+cate,rank)
        localStorage.setItem("cache_time"+cate,timestamp+1800000) //1800000
        localStorage.setItem("cate",cate)

      },

        refresh:function(){
          var site = this
          setTimeout(function(){
            site.articles= []
            site.rank=0
            site.$refs.infiniteLoading.stateChanger.reset()
          }, 50);
          return true
        },

        infiniteHandler: function ($state) {
          var site = this
          if(!site.loading){
            if (this.articles.length > 500) {
              $state.complete();
            } else {
              /** load data start */
              setTimeout(function(){
                site.loading = true
                var uri = '/hot?limit=10&rank='+site.rank+'&tag='+site.cate;
                api.get(uri,function(err,data){
                  if(data && data.length>0){
                    for(var t=0;t<data.length;t++){
                      site.articles.push(data[t])
                    }
                    site.rank = data[(data.length-1)].Rank
                    $state.loaded();
                    site.page ++
                    site.setCache()
                  }else{
                    $state.complete();
                  }
                  site.loading = false // 加载完成
                })
              }, 200);
              /** load data end */
            }
          }
        },
       

    },


   
    watch:{
      cate:function(){
        // console.log(this.cate)
        // this.refresh()
        this.getCache()
      },
      '$route' (to, from) {
        // console.log(to.params.id)
        if(to.params.id !=undefined){
          this.cate = to.params.id
        }else{
          this.cate = 0
        }
        // 对路由变化作出响应...
      }
    }
}
</script>



