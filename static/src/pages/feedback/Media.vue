<template>
<div>
    <div class="weui-cells__title">用户反馈 > 提交公众号</div>


    <div class="weui-cells__title">文本域</div>

    <weui-textarea placeholder="如果您有优质微信公众号要分享，请提交其一篇文章的链接，通过审核后将收录其后续文章。" :value="text"></weui-textarea>

    <!-- <div class="weui-cells weui-cells_form">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <textarea class="weui-textarea" placeholder="  如果您有优质微信公众号要分享，请提交其一篇文章的链接，通过审核后将收录其后续文章。" rows="3"></textarea>
                <div class="weui-textarea-counter"><span>0</span>/200</div>
            </div>
        </div>
    </div> -->



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
    <label for="weuiAgree" class="weui-agree">
        <input id="weuiAgree" type="checkbox" class="weui-agree__checkbox">
        <span class="weui-agree__text">
            阅读并同意<a href="javascript:void(0);">《相关条款》</a>
        </span>
    </label>
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
