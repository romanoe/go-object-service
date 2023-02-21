<script lang="ts" setup>
import { useFetch } from "@/composables/FetchData";

interface GoelandObject {
  created_at: Date;
  fk_type: number;
  id: string
}

  const url = import.meta.env.VITE_API_URL + '/objects'

  const {data} = await useFetch<GoelandObject[]>(import.meta.env.VITE_API_URL + '/objects');
</script>



<template>
    <h1>Objets</h1>
    <ul>
        <li v-for="object in data">
          <a :href="'/' + object.id">{{object.fk_type}}</a>
      </li>
    </ul>
</template>

