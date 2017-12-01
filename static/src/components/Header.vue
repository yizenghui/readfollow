<template>

<div>
    <div class="weui-panel__hd hd_no_boder">

        <div class="weui-flex navber">
          <router-link class="weui-flex__item"  :to="{ name: 'hot'}">
            热门
          </router-link>
          <router-link class="weui-flex__item" :to="{ name: 'cate', params: { id: cate.ID }}" v-for="cate in category" :key="cate.ID">
             {{cate.Title}}
          </router-link>
          <!-- <router-link class="weui-flex__item"  :to="{ name: 'servicepost'}">
            Share
          </router-link> -->
        </div>
    </div>
    <!-- <div class="weui-cells">
            <a class="weui-cell weui-cell_access" href="javascript:;">
                <div class="weui-cell__bd">
                    <p>cell standard</p>
                </div>
                <div class="weui-cell__ft">说明文字</div>
            </a>

        </div> -->
</div>
</template>
<style>

.navber{
    text-align: center;
}
.weui-panel__hd:after{
    border: none;
}

.navber a{
    color: #999
}
.router-link-active{
    font-weight: bold;  
}
.navber .router-link-active{
    color: #00a06a;
    font-weight: bold;  
}

.page__hd {
    padding: 40px;
}
.weui-navbar__item.weui-bar__item_on{
  background-color:white;
  color: red;
}


.weui-media-box_appmsg .weui-media-box__hd{
      width: 80px;
      margin-left:0;
      margin-right:0;
}
.weui-navbar__item {
    position: relative;
    display: block;
    -webkit-box-flex: 1;
    -webkit-flex: 1;
    flex: 1;
    padding: 13px 0;
    text-align: center;
    font-size: 15px;
    -webkit-tap-highlight-color: rgba(0,0,0,0);
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
        tags: [],
        category: [],
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
          api.get("/tags?type=cate",function(err,data){
            site.category = data
          })
        },
      back () {
        this.$router.push(this.backTo)
      }
    },
        
    mounted() {
        this.GetCate()
    },
  }
</script>