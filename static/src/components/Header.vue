<template>

<div>
    <!-- <div class="weui-panel__hd hd_no_boder">

        <div class="weui-flex navber">
          <router-link class="weui-flex__item"  :to="{ name: 'hot'}">
            热门
          </router-link>
          <router-link class="weui-flex__item" :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in showCategory" :key="cate.ID">
             {{cate.Title}}
          </router-link>
        </div>
    </div>
    <div class="weui-grids">
        <router-link class="weui-grid"  :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in hideCategory" :key="cate.ID">
            <p class="weui-grid__label">{{cate.Title}}</p>
        </router-link>
    </div>
    <div class="weui-cells">
            <a class="weui-cell weui-cell_access" href="javascript:;">
                <div class="weui-cell__bd">
                    <p>cell standard</p>
                </div>
                <div class="weui-cell__ft">说明文字</div>
            </a>

        </div> -->


    <div class="weui-grids nav-grids">
          <router-link class="weui-grid"  :to="{ name: 'hot'}">
            <p class="weui-grid__label">热门</p>
          </router-link>
          <router-link class="weui-grid" :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in showCategory" :key="cate.ID">
              <p class="weui-grid__label">{{cate.Title}}</p>
          </router-link>
        <template v-if="hideCategory.length>0">
          <div class="weui-grid" v-on:click="showHideCategoryStatus = !showHideCategoryStatus" >
             <p class="weui-grid__label">
               {{showHideCategoryStatus?'<<':'>>'}}
               </p>
          </div>
        </template>
        <template v-if="showHideCategoryStatus">
          <router-link class="weui-grid" :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in hideCategory" :key="cate.ID">
              <p class="weui-grid__label">{{cate.Title}}</p>
          </router-link>
        </template>
    </div>


</div>
</template>
<style>
.nav-grids .weui-grid{
  width:16.66%;
  border: none;
  padding: 10px 0px;
}
.nav-grids .weui-grid:before{
  
  border: none;
}
.nav-grids{
  padding: 0 1em;
}

.nav-grids:before{
  border: none;
}
.weui-grid:active {
    background-color: #ffffff;
}
.nav-grids .weui-grid:after {
   border: none;
}
.nav-grids .router-link-active{
  background-color: #ececec;
}
.router-link-active .weui-grid__label{
    font-weight: bold;  
    color: #0f88eb;
}
</style>

<script>
import api from '../api';

  export default {

    props: {
      title: {
        type: String
      },
      backTo: {
        type: String,
        default: '/hot'
      }
    },  
    data () {
      return {
        tag: [],
        tags: [],
        category: [],
        listCount:4,
        selectCategoryInShowList:false,
        showCategory: [],
        hideCategory: [],
        showHideCategoryStatus:false,
        swiperOption2: {
          pagination: '.swiper-pagination',
          slidesPerView: 6,
          initialSlide: 6,
          paginationClickable: true,
          spaceBetween: 0
        }
      }
    },
    methods: {
      GetCate(){
        var site = this
        var timestamp = (new Date()).valueOf();
        var cate = this.cate
        var time = parseInt(window.localStorage.getItem("cate_cache_time")) || 0
        if (time>timestamp){
          site.category  = JSON.parse(window.localStorage.getItem("all_category"))||[]
          // console.log(window.localStorage.getItem("hide_category_status"))
          if(window.localStorage.getItem("hide_category_status")==='true'){ //布尔不为空的字符串都是 true
              site.showHideCategoryStatus  = true
          }
          site.manageCategory()
        }else{
          api.get("/tags?type=cate",function(err,data){
            localStorage.setItem("all_category",JSON.stringify(data))
            localStorage.setItem("cate_cache_time",timestamp+1800000) //1800000
            site.category = data
            site.manageCategory()
          })
        }
      },

      back () {
        this.$router.push(this.backTo)
      },
      manageCategory(){
        var category = this.category
        var routeName = this.$route.name
        if(routeName == 'cate'){
          var id = this.$route.params.id
        }
        for(var i=0;i<category.length;i++){
          if(i<this.listCount){
            this.showCategory.push(category[i])
          }else{
            this.hideCategory.push(category[i])
          }
        }
        // console.log('data', category,this.showCategory, this.hideCategory)
      }
    },
    watch:{
      showHideCategoryStatus:function(){
        // console.log(this.showHideCategoryStatus)
        localStorage.setItem("hide_category_status",this.showHideCategoryStatus)
      }
    },
    mounted() {
        this.GetCate()
    },
  }
</script>