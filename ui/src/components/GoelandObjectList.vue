<script lang="ts" setup>
import { useFetch } from "@/composables/FetchData";

interface GoelandObject {
  created_at: Date;
  fk_type: number;
  id: string | number
}

const url = import.meta.env.VITE_API_URL + '/objects'

const { data } = await useFetch<GoelandObject[]>(url);

</script>



<template>
    <h1>Objets</h1>
    <ul>
        <li v-for="object in data" :id?="object.id" >
          <router-link :to="'/' + object.id">{{object.id}}</router-link>
        </li>
    </ul>
</template>