<template>
  <h1 class="title">{{ props.value }}</h1>

  <div class="select">
    <el-select
      v-model="props.value"
      class="m-2"
      placeholder="Select"
      size="large"
      @change="chatTypeChange"
    >
      <el-option
        v-for="item in props.SeleteItem"
        :key="item.key"
        :label="item.title"
        :value="item.title"
      />
    </el-select>
  </div>
</template>

<script setup>
import { ref, defineProps } from "vue";

//   const  = ref('周报生成器')

const props = defineProps(["SeleteItem", "value"]);

const emit = defineEmits(["dosth"]);

const chatTypeChange = (val) => {
  emit("dosth", val);
};
</script>
