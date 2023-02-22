<script setup lang="ts">
import GoelandObjectList from "@/components/GoelandObjectList.vue";

</script>



<template>
  <header>
  </header>

  <Suspense> 
    <template #default>
      <router-view></router-view>
    </template>
    <template #fallback>
      Loading...
    </template>
  </Suspense>

</template>

<style scoped>
</style>
