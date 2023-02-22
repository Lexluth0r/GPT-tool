<template>
  <div class="home flex_col">
    <div class="header flex_row">
      <Head @dosth="chatTypeChange" :SeleteItem="SeleteItem" :value="typeValue">
      </Head>
    </div>
    <div class="section">
      <div>
        <h1>{{ bodyObj.section1 }}</h1>
      </div>
      <div>
        <h1>{{ bodyObj.section2 }}</h1>
      </div>
    </div>
    <div class="desc">{{ bodyObj.desc }}</div>
    <div class="content">
      <div class="question">
        <h4>{{ bodyObj.question }}</h4>
      </div>
      <div class="textarea el-textarea">
        <el-input
          v-model="textarea"
          :rows="2"
          type="textarea"
          :placeholder="bodyObj.placeholder"
        />
      </div>
    </div>
    <div class="button">
      <button
        type="button"
        class="el-button el-button--primary"
        @click="submit"
      >
        <span>{{ bodyObj.buttonText }}</span>
      </button>
    </div>
    <div class="response"  >
      <el-skeleton :rows="5" :loading="loading" animated >
        <article v-html="response" />
      </el-skeleton>
     
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";

import { ElMessage } from 'element-plus';

import Head from "@/components/Head.vue";

import SeleteItem from "@/data/item.js";

import { submitQuestion } from "@/api/chat";

import {marked} from 'marked'

const textarea = ref("");

const bodyObj = ref({});

const response = ref("")

const loading = ref(false)


const typeValue = ref("周报生成器");

onMounted(() => {
  SeleteItem &&
    SeleteItem.map((_) => {
      if (typeValue) {
        _.key === typeValue.value && (bodyObj.value = _);
      }
    });
});

const submit = () => {

  let question = textarea.value;

  if(isEmpty(question)){
    ElMessage.error('请输入问题')
    return
  }

  loading.value = true

  let data = {
        question:bodyObj.value.basePrompt + question,
        task : bodyObj.value.key
      }

  submitQuestion(data)
        .then((result) => {
          loading.value = false
          if (result.code === 200) {
            response.value = marked(result.data)
          } else {
            ElMessage({
              message: result.msg,
              type: 'warning',
            })
          }
        })
        .catch((err) => {}
        );

};

const isEmpty = (str) => {
  if (str instanceof Array) {
    if (str.length === 0) {
      return true;
    } else {
      return false;
    }
  } else if (str == null || typeof str == "undefined" || str == "") {
    return true;
  }
  return false;
};

const chatTypeChange = (val) => {
  typeValue.value = val;

  SeleteItem.map((_) => _.key === val && (bodyObj.value = _));
};
</script>
