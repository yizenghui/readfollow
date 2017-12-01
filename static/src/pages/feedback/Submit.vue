<template>
<div>
    <div class="weui-cells__title">分类列表</div>
     <a href="/" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>提交公众号</p>
        </div>
        <div class="weui-cell__ft"></div>
    </a> 
    <a href="/cate/1" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>文章未收录</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </a>
    <a href="/cate/2" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>侵犯隐私</p>
        </div>
        <div class="weui-cell__ft"></div>
    </a>
    <a href="/cate/3" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>信息举报</p>
        </div>
        <div class="weui-cell__ft"></div>
    </a>
    <a href="/cate/129" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>其它</p>
        </div>
        <div class="weui-cell__ft"></div>
    </a>
             
</div>
</template>

<script>
import news from '../../api/news.js';

  export default {
    name: 'app',
    data(){
        return{
           url:""
        }
    },
    methods: {
       submit: function(event) {
            
              news.post(this.url,function(err,data){
                if(data == "0"){
                  alert("收录失败")
                }else{
                  alert("收录成功")
                }
              });

           
        }
        
    },
    components: {
      
    }
  }
</script>
