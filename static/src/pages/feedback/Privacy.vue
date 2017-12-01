<template>
<div>
    <div class="weui-cells__title">用户反馈 > 侵犯隐私</div>


    <div class="weui-cells__title">待举报公众号</div>

    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input class="weui-input" type="text" placeholder="请输入公众号名称">
            </div>
        </div>
    </div>

    <div class="weui-cells__title">待举报公众号文章</div>

    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <input class="weui-input" type="text" placeholder="https://mp.weixin.qq.com/s...">
            </div>
        </div>
    </div>

    <div class="weui-cells__title">补充说明</div>
    <weui-textarea placeholder="请详细描述您遇到的问题" :value="text"></weui-textarea>

    <div class="weui-cells__title">如需反馈，请留下联系方式</div>
    <div class="weui-cells">
        <div class="weui-cell weui-cell_select weui-cell_select-before">
            <div class="weui-cell__hd">
                <select class="weui-select" name="select2">
                    <option value="1">邮箱</option>
                    <option value="2">手机</option>
                </select>
            </div>
            <div class="weui-cell__bd">
                <input class="weui-input" type="number" pattern="[0-9]*" placeholder="请输入号码">
            </div>
        </div>
    </div>
   
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
