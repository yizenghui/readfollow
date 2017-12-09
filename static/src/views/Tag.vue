//index.vue
<template>

<div>



    <!-- <div class="page__hd">
        <p class="page__desc"><router-link class="weui-agree__text" :to="{ name: 'home'}">跟读微信文章</router-link></p>
    </div> -->

        
        <div class="weui-cells__title">个性化定制{{tags}}</div>
        <div class="weui-cells weui-cells_checkbox">
            <label class="weui-cell weui-check__label"  v-for="cate in category" :key="cate.ID">
                <div class="weui-cell__hd">
                    <input type="checkbox" class="weui-check" :value="cate.ID" v-model="tags">
                    <i class="weui-icon-checked"></i>
                </div>
                <div class="weui-cell__bd">
                    <p>{{cate.Title}}</p>
                </div>
            </label>
        </div>

        <weui-search v-model="value" :result="filterResult" @result-click="resultClick"></weui-search>

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
import WeuiSearch from '@/components/weui/Search';
import WeuiCell from '@/components/weui/Cell';
export default {
  
  name: 'indexP',
  
  
    components: {
      InfiniteLoading,
      WeuiSearch,
      WeuiCell
    },
    data () {
      return {
        tags: [1,2],
        category: [],
        value: '',
        defaultResult: [
          'Apple',
          'Banana',
          'Orange',
          'Durian',
          'Lemon',
          'Peach',
          'Cherry',
          'Berry',
          'Core',
          'Fig',
          'Haw',
          'Melon',
          'Plum',
          'Pear',
          'Peanut',
          'Other'
        ]
      }
    },
    mounted() {
      
      setTimeout(function(){
      }, 200);
        this.GetCate()
      // alert(1)
    },
    methods: {
      resultClick (item) {
        window.alert('you click the result item: ' + JSON.stringify(item))
      },
        GetCate:function(){
          var site = this
          api.get("/tags?type=cate",function(err,data){
            site.category = data
          })
        },
    },
    computed: {
      filterResult () {
        console.log(this.value)
        return this.defaultResult.filter(value => new RegExp(this.value, 'i').test(value))
      }
    }

}
</script>



