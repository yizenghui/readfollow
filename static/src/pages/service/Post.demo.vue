<template>
<div>

<div  v-if="showSuccess">
    <div class="weui-mask"></div>
    <div class="weui-dialog">
        <div class="weui-dialog__hd"><strong class="weui-dialog__title">分享成功</strong></div>
        <div class="weui-dialog__bd">请耐心等候文章链接审核通过！</div>
        <div class="weui-dialog__ft">
            <a href="javascript:;" v-on:click="showSuccess = false" class="weui-dialog__btn weui-dialog__btn_primary">确定</a>
        </div>
    </div>
</div>
<div  v-if="showWarn">
    <div class="weui-mask"></div>
    <div class="weui-dialog">
        <div class="weui-dialog__hd"><strong class="weui-dialog__title">分享失败</strong></div>
        <div class="weui-dialog__bd">请检查分享的链接是否正常</div>
        <div class="weui-dialog__ft">
            <a href="javascript:;" v-on:click="showSuccess = false" class="weui-dialog__btn weui-dialog__btn_primary">确定</a>
        </div>
    </div>
</div>

<form @submit.prevent="submit">
    <div class="weui-cells__title">自助服务 > 分享公众号文章</div>
    <div class="weui-cells__title">请粘贴待分享文章链接地址</div>
    <weui-textarea placeholder="如果您有优质微信公众号文章要分享，请提交其文章链接，通过审核后将收录、展示。" v-model="url"></weui-textarea>

    
    <weui-checklist title="基本示例" :options="options" v-model="checkedItems"></weui-checklist>

    <div class="weui-btn-area">
      <button type="submit" class="weui-btn weui-btn_primary">分享</button>
    </div>
</form>
</div>
</template>

<script>
import api from '../../api';
import WeuiTextarea from '@/components/weui/Textarea';
import WeuiChecklist from '@/components/weui/Checklist';

  export default {
    name: 'app',
    data(){
        return{
                  options: [
          {
            label: '选项1',
            value: 'value1'
          },
          {
            label: '选项2',
            value: 'value2'
          },
          {
            label: '选项3',
            value: 'value3'
          },
          {
            label: '选项4（禁用）',
            value: 'value4',
            disabled: true
          }
        ],
        checkedItems: ['value1', 'value3'],

           url:"",
           showSuccess:false,
           showWarn:false,
        }
    },
    methods: {
       submit: function(event) {
          var site = this
          if(site.url != ''){
            setTimeout(function(){
              site.posting = false
            }, 200);

            site.posting = true
              api.post(site.url,function(err,data){
                if(data == "0"){
                  alert('提交失败')
                }else{
                    if(site.url == ''){
                      site.showSuccess = true
                    }
                    site.url = ''
                }
              });
          }else{
            site.showWarn = true
          }
        }
    },
    components: {
      'weui-textarea':WeuiTextarea,
      'weui-checklist':WeuiChecklist
    }
  }
</script>
