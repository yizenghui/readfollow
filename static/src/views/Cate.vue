//index.vue
<template>

<div>




<div class="weui-panel_access">
   <!-- <div class="weui-panel__hd">
    <span>跟读微信文章</span>
     


    <label for="weuiAgree" class="weui-agree">
        
        
              <router-link class="weui-agree__text" :to="{ name: 'tags'}">{{tag.Title}} </router-link>
    </label>
  </div> 
-->
  <List :cate="cate"></List>
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

import List from '@/components/List';
import InfiniteLoading from 'vue-infinite-loading'
import api from '../api';

export default {
  
  name: 'cateP',
  
  
    components: {
      InfiniteLoading,
      List,
      cate:[],
    },
    data () {
      return {
        cate: 0,
        tag:[]
      }
    },
    mounted() {
      this.cate = this.$route.params.id
      // this.GetTagInfo()
    },
    methods: {
      GetTagInfo:function(){
          var site = this
          api.get("/tag/"+site.cate,function(err,data){
            site.tag = data
          })
      }
    },
    

    watch:{
      // cate:function(){
      //   // console.log("cate::",this.cate)
      // },
      "$route": function(){
          this.cate = this.$route.params.id
      }
    }
}
</script>



