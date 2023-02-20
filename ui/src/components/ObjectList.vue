<script lang="ts" setup>
import { useFetch } from "@/composables/FetchData";
import { ref, defineAsyncComponent } from 'vue';

interface GoelandObject {
  created_at: Date;
  fk_type: number;
  id: number
}

const goelandObjects = ref<GoelandObject[] | null>([])
const goelandObjectsError = ref(false)

  const {hasError, data} = await useFetch<GoelandObject[]>(import.meta.env.VITE_API_URL + '/objects');
  goelandObjects.value = data.value;
  goelandObjectsError.value = hasError.value;
</script>



  <template>


    <h1>Objets</h1>
    <ul>
      <li v-for="object in data">
        {{object.fk_type}}
      </li>
    </ul>


  </template>

