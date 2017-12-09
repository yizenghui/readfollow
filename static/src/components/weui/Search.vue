<template>
  <div>
    <div class="weui-search-bar">
      <div class="weui-search-bar__form">
        <div class="weui-search-bar__box">
          <i class="weui-icon-search"></i>
          <input class="weui-search-bar__input" :placeholder="placeholder" v-model="currentValue" ref="searchInput">
          <a class="weui-icon-clear" @click="searchClear"></a>
        </div>
        <label class="weui-search-bar__label" @click="textClick" v-show="!isActive">
          <i class="weui-icon-search"></i>
          <span v-text="placeholder"></span>
        </label>
      </div>
      <a class="weui-search-bar__cancel-btn" @click="searchCancel" v-show="isActive" v-text="cancelText"></a>
    </div>

     <slot>
      <div class="weui-cells searchbar-result" v-show="show || currentValue">
        <div class="weui-cell weui-cell_access" v-for="(item, key, index)  in result" :key="key" @click="handleResultClick(item)">
        <div class="weui-cell__bd weui-cell_primary">
          <p>{{item}}</p>
        </div>
        </div>
        <!-- <weui-cell v-for="(item, key, index) in result" :key="key" :title="item" @click="handleResultClick(item)"></weui-cell> -->
      </div>
    </slot>
  </div>
</template>

<script>
  import WeuiCell from './Cell.vue'
  export default {
    name: 'weui-search',
    components: {
      WeuiCell
    },
    props: {
      value: String,
      autofocus: Boolean,
      show: Boolean,
      placeholder: {
        type: String,
        default: '搜索'
      },
      cancelText: {
        type: String,
        default: '取消'
      },
      result: Array
    },
    data () {
      return {
        isActive: false,
        currentValue: this.value
      }
    },
    mounted () {
      if (this.autofocus) {
        this.$refs.searchInput.focus()
        this.isActive = true
      }
    },
    methods: {
      textClick (e) {
        // focus the input
        this.$refs.searchInput.focus()
        this.isActive = true
      },
      handleResultClick (item) {
        console.log(item)
        this.$emit('result-click', item) // just for compatibility
        this.$emit('on-result-click', item)
      },
      // 清除输入
      searchClear () {
        this.currentValue = ''
      },
      // 取消搜索
      searchCancel () {
        this.searchClear()
        this.isActive = false
      }
    },
    
    watch: {
      currentValue (val) {
        this.$emit('input', val)
      },
      value (val) {
        this.currentValue = val
      }
    }
  }
</script>
