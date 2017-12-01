<template>
<div>
    <div class="weui-cells__title">自助服务 > 文章推广</div>


    <div class="weui-cells__title">请提交您的推广文章</div>

    <weui-textarea placeholder="如果您所发布的微信文章中，原文链接指向 http://readfollow.com/  任意地址。那么，该文章在我们平台将获得5%的排序优先加成" :value="text"></weui-textarea>




    <div class="weui-btn-area">
        <a class="weui-btn weui-btn_primary" href="javascript:" id="showTooltips">确定</a>
    </div>
</div>
</template>

<script>
import news from '../../api/news.js';
import WeuiTextarea from '@/components/weui/Textarea';

  export default {
    name: 'app',
    data(){
        return{
           url:"",
           text:""
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
      'weui-textarea':WeuiTextarea
    }
  }
</script>
