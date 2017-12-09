//list.vue 文章列表
<template>

<div>

<div id="loadingToast" v-if="loading">
    <div class="weui-mask_transparent"></div>
    <div class="weui-toast">
        <i class="weui-loading weui-icon_toast"></i>
        <p class="weui-toast__content">数据加载中</p>
    </div>
</div>

<div class="weui-panel_access">
   <!--<div class="weui-panel__hd">
    <span>
      跟读微信文章
    </span>
    

    <label for="weuiAgree" class="weui-agree">
        
              <router-link class="weui-agree__text" :to="{ name: 'tags'}">全部分类 </router-link>

    </label>
  </div> -->

  
  <div class="weui-panel__bd">
    <ArticleList :articles="articles"></ArticleList>
  </div>
  
    <infinite-loading @infinite="infiniteHandler" :distance="distance" ref="infiniteLoading">

    <div slot="spinner" class="weui-loadmore">
        <i class="weui-loading"></i>
    </div>

    <div slot="no-results" class="weui-loadmore weui-loadmore_line">
      <span class="weui-loadmore__tips">加载不出内容!</span>
    </div>

    <div slot="no-more" class="weui-loadmore weui-loadmore_line">
      <span class="weui-loadmore__tips">到底了!</span>
    </div>


    </infinite-loading> 

</div>


</div>


</template>
<style>
 .placeholder {
    margin: 5px;
    padding: 0 10px;
    background-color: #ebebeb;
    height: 2.3em;
    line-height: 2.3em;
    text-align: center;
    color: #cfcfcf;
}
/* .weui-toast{
    background: none;
} */
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
        cate: 0,
        articles: [],
        distance: 200,
        page:0,
        showload:false,
        rank:0, // 当前下拉文章最低rank
        loading:false,
      }
    },
    mounted() {
      
      this.cate = parseInt(this.$route.params.id)
    },
    methods: {
        
        init:function(){ // 初始化数据
          var site = this
          this.loading = true
          // setTimeout(function(){
            // 有效期限时间
            var timestamp = Date.parse(new Date()) - 30 * 60 *1000; 
            var articles = JSON.parse(window.localStorage.getItem("article_cate_"+site.cate))||[]
            var rank = parseFloat(window.localStorage.getItem("rank_cate_"+site.cate)) || 0
            var cache_time = parseInt(window.localStorage.getItem("cache_time_cate_"+site.cate)) || 0
            if (timestamp<cache_time){
              site.articles= articles
              site.rank=rank
              console.log("载入缓存")
            }else{
              // site.articles= []
              // site.rank=0
              console.log("重新加载")
              site.refresh()
            }
            // site.$refs.infiniteLoading.stateChanger.reset()
          // }, 50);

          return true
        },

        cache:function(){
          var site = this
          var timestamp = Date.parse(new Date()); 
          localStorage.setItem("article_cate_"+site.cate, JSON.stringify(site.articles))
          localStorage.setItem("rank_cate_"+site.cate, site.rank)
          localStorage.setItem("cache_time_cate_"+site.cate, timestamp)
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
      
      "$route": function(){
          this.cate = this.$route.params.id
      },
      cate:function(){
        console.log(this.cate)
        // this.init()
      }
    }
}
</script>



