<script lang="ts" setup>
import { useFetch } from "@/composables/FetchData";
import {useRoute} from 'vue-router';
import {ref, watch} from 'vue';

interface GoelandObject {
  created_at: Date;
  fk_type: number;
  id: string
}

const route = useRoute(); 
const url = import.meta.env.VITE_API_URL + '/objects/';
const { data } = await useFetch<GoelandObject>(url + route.params.id);

 
</script>



<template>
    <h1>Formulaire Objet</h1>
    {{ data }}
</template>

