<template>
<h1>Objets</h1>
  <ul>
    <li v-for="object in objects">
      {{object.id}}
    </li>
  </ul>
</template>

<script lang="ts" setup>
import {ref} from 'vue'

interface GoelandObject {
  created_at : Date;

  fk_type : number;

  id: number
}

const objects = ref<GoelandObject[]>([])

async function getData() {
  const res = await fetch(import.meta.env.VITE_API_URL + '/objects');
  objects.value = await res.json();
}

getData();
</script>

<style scoped>

</style>