<script setup>
import {onMounted, ref} from "vue";
import {NCard} from "naive-ui";

const objects = ref(null)
const loading = ref(null)
const error = ref(null)


function fetchData() {
  loading.value = true

  return fetch("http://localhost:1323/objects", {
    method : 'get',
    headers : {
      'content-type' : 'application/json'
    }
  }).then(res => {
    if (!res.ok) {
      const error = new Error(res.statusText);
      error.json = res.json();
      throw error;
    }
    return res.json();
  }).then(json => {
    objects.value = json;
     })
     .catch(err =>  {
        error.value = err;
      })
      .then(() => {
        loading.value = false;
      })
}

onMounted(() => {
  fetchData();
});

</script>



<template>

  <n-card size="huge" v-for="object in objects" :key="object.id" :title="object.type">
    {{object.type}}
  </n-card>

</template>

<style scoped>

</style>
