//ListItem.vue 信息卡风格
<template>


      <div class="weui-media-box weui-media-box_appmsg" >

          <div class="weui-media-box__hd">
            <!-- <a :href="article.URL"><img class="weui-media-box__thumb"   v-lazy="article.Cover" ></a> -->
            <a :href="article.URL"><img class="weui-media-box__thumb"   :src="article.Cover" ></a>
          </div>  
          <div style="width:30px">
            <ul class="weui-media-box__info article-tags">
              <li class="weui-media-box__info__meta"><span v-on:click="like(article)">
                <svg class="icon" width="30px" height="20.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="#cecece" d="M509.927514 387.159081C517.168621 379.168894 527.507586 379.262447 534.709532 387.493244L805.278364 696.714765C813.036915 705.581679 826.514517 706.480186 835.381431 698.721636 844.248346 690.963085 845.146852 677.485483 837.388303 668.618569L566.819471 359.397045C542.819471 331.968474 502.692194 331.60538 478.31207 358.507586L197.525612 668.340919C189.61372 677.071283 190.277222 690.562496 199.007586 698.474389 207.737949 706.386281 221.229163 705.722778 229.141056 696.992414L509.927514 387.159081Z" /></svg>
                </span>
                </li> 
              <li class="weui-media-box__info__meta" style="font-size: 10px;text-align: center;width: 30px;padding: 0;">
                {{article.Like-article.Hate}}
              </li>
              <li class="weui-media-box__info__meta">
                
                <span v-on:click="hate(article)">
                  <svg class="icon" width="30px" height="20.00px" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg"><path fill="#cecece" d="M478.31207 644.159081C502.692194 671.061286 542.819471 670.698193 566.819471 643.269621L837.388303 334.048098C845.146852 325.181184 844.248346 311.703582 835.381431 303.94503 826.514517 296.186481 813.036915 297.084988 805.278364 305.951902L534.709532 615.173423C527.507586 623.40422 517.168621 623.497773 509.927514 615.507586L229.141056 305.674253C221.229163 296.943889 207.737949 296.280386 199.007586 304.192277 190.277222 312.104171 189.61372 325.595383 197.525612 334.325747L478.31207 644.159081Z" /></svg>                
                </span>
              </li>
            </ul>
          </div>          
           <div class="weui-media-box__bd">
            <a :href="article.URL"><p class="weui-media-box__desc article-title">{{article.Title}}</p></a>
            <ul class="weui-media-box__info article-tags">
              
              <li class="weui-media-box__info__meta">
                <router-link  :to="{ name: 'cate', params: { id: article.MediaTagID }}" style="color: #cecece">{{article.MediaTagTitle}}</router-link>
                <!-- {{article.Media.AppName}} -->
              </li> 
              

              <li class="weui-media-box__info__meta"><timeago :since="article.PubAt" class="text-muted" locale="zh-CN"></timeago></li>
            </ul>
          </div>
      </div>

       

</template>
<style>
.item-title{
  
    line-height: 1.2;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    font-size: 13px;
    color: #000;

}
.item-tags{
  margin-top: 10px;
  font-size: 12px;
}
</style>
<script>


import api from '../api';

export default {
  
  name: 'ListItem',


    props: ['article'],
    // props: ['article','swiperOption'],

    data () {
      return {
        
      }
    },
   
    methods: {
      
        like:function(article){
          var site = this
          site.showload = true
          api.get('/like/'+article.ID,function(err,data){
            article.Like = data.Like
            site.showload = false
          })
        },
        hate:function(article){
          var site = this
          site.showload = true
          api.get('/hate/'+article.ID,function(err,data){
            article.Hate = data.Hate     
            site.showload = false
          }) 
        },

    }
}
</script>



