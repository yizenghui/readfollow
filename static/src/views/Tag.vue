//index.vue
<template>

<div>




    <!-- <div class="page__hd">
        <p class="page__desc"><router-link class="weui-agree__text" :to="{ name: 'home'}">跟读微信文章</router-link></p>
    </div> -->
        <div class="weui-cells__title">分类列表</div>


          <router-link class="weui-cell weui-cell_access"  :to="{ name: 'home'}">
             <div class="weui-cell__bd">
                    <p>全部</p>
                </div>
                <div class="weui-cell__ft">
                </div>

          </router-link>

          <router-link class="weui-cell weui-cell_access" :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in category" :key="cate.ID">
             <div class="weui-cell__bd">
                    <p>{{cate.Title}}</p>
                </div>
                <div class="weui-cell__ft">
                </div>

          </router-link>

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
.weui-toast{
    background: none;
}
</style>
<script>

import InfiniteLoading from 'vue-infinite-loading'
import api from '../api';

export default {
  
  name: 'indexP',
  
  
    components: {
      InfiniteLoading
    },
    data () {
      return {
        tags: [],
        category: []
      }
    },
    mounted() {
      setTimeout(function(){
      }, 200);
        this.GetCate()
      // alert(1)
    },
    methods: {
        GetCate:function(){
          var site = this
          api.get("/tags?type=cate",function(err,data){
            site.category = data
          })
        },
    },

}
</script>



