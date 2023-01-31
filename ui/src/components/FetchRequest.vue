<template>

</template>

<script setup lang="ts">

const request = (url: string): Promise<TResponse> => {
  return fetch(url).then((response) => response.json)
      .then((data) => data as TResponse);
}




</script>

<style scoped>

</style>
