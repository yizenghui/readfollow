<template>
<div>
    <div class="weui-cells__title">自助服务 > 移除收录</div>



    <div class="weui-cells__title">请粘贴待移除收录的文章地址</div>

    <weui-textarea placeholder="" :value="text"></weui-textarea>


    <!-- <div class="weui-cells__title">如需反馈，请留下联系方式</div>
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
    </label> -->
    <div class="weui-cells__title">满足任意条件可移除</div>
    <div class="weui-cells">
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <p>将移除已不存在的文章收录</p>
            </div>
            <div class="weui-cell__ft"></div>
        </div>
        
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <p>将移除含<code>limit.readfollw.com</code>的文章收录</p>
            </div>
            <div class="weui-cell__ft"></div>
        </div>
        
        <div class="weui-cell">
            <div class="weui-cell__bd">
                <p>停止收录并移除含<code>ban.readfollw.com</code>的公众号所有文章</p>
            </div>
            <div class="weui-cell__ft"></div>
        </div>
        
    </div>
<!--     
    <div class="weui-cells__title">
        <ol>
            <li>将移除已不存在的文章收录</li>
            <li>将移除含<code>limit.readfollw.com</code>字符的文章收录</li>
            <li>移除并停止收录含<code>ban.readfollw.com</code>字符的公众号文章</li>
        </ol>
    </div> -->
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
