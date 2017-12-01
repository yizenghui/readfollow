<template>
<div>



    <div class="weui-cells__title">自助服务</div>

     <router-link :to="{ name: 'servicepost'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>提交文章</p>
        </div> 
        <div class="weui-cell__ft">
            分享文章
        </div>
    </router-link>

     <router-link :to="{ name: 'servicedelete'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>删除收录</p>
        </div>
        <div class="weui-cell__ft">
            移除收录
        </div>
    </router-link>

     <router-link :to="{ name: 'servicespread'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>文章推荐</p>
        </div>
        <div class="weui-cell__ft">
            互推
        </div>
    </router-link>


<!--
    <div class="weui-cells__title">用户反馈</div>

     <router-link :to="{ name: 'feedbackprivacy'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>侵犯隐私</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>

    <router-link :to="{ name: 'feedbackreport'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>信息举报</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>



    <router-link :to="{ name: 'feedbackother'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>优化建议</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>
    
     <router-link :to="{ name: 'feedbackother'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>其它</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>
    -->
    
    <div class="weui-cells__title">相关</div>

     <router-link :to="{ name: 'feedbackmedia'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>收录</p>
        </div>
        <div class="weui-cell__ft"></div>
    </router-link>

     <router-link :to="{ name: 'feedbackother'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>排序</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>

    <router-link :to="{ name: 'feedbackother'}" class="weui-cell weui-cell_access">
        <div class="weui-cell__bd">
            <p>推广</p>
        </div> 
        <div class="weui-cell__ft"></div>
    </router-link>


    
</div>
</template>

<script>
import api from '../api';

  export default {
    name: 'app',
    data(){
        return{
           url:""
        }
    },
    methods: {
       submit: function(event) {
            
              api.post(this.url,function(err,data){
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
