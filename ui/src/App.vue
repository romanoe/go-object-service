<script setup lang="ts">
import GoelandObjectList from "@/components/GoelandObjectList.vue";

</script>



<template>
  <header>
  </header>

  <Suspense> 
    <template #default>
      <router-view :key="$route.fullPath" ></router-view>
    </template>
    <template #fallback>
      Loading...
    </template>
  </Suspense>

</template>

<style scoped>
</style>
