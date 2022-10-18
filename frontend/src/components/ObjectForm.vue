<script setup>
import { ref } from 'vue'
import { NForm, NInput, NButton } from 'naive-ui'

</script>

<template>

  <h1>Objets</h1>
  <n-form
      ref="formRef"
      inline
      :label-width="80"
      :model="formValue"
      :rules="rules"
      :size="size"
  >
    <n-form-item label="Type d'objet">
      <n-input />
    </n-form-item>
    <n-form-item label="Age">
      <n-input  />
    </n-form-item>
    <n-form-item label="Phone">
      <n-input />
    </n-form-item>
    <n-form-item>
      <n-button>
        Validate
      </n-button>
    </n-form-item>
  </n-form>

</template>

<style scoped>

</style>
