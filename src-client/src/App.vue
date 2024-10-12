<script setup lang="ts">
import {RouterLink} from "vue-router";
</script>

<template>
  <div class="p-4 bg-gray-800 text-white min-h-screen">
    <div>
      <RouterLink
          to="/"
          class="text-lg font-semibold hover:text-gray-300 transition duration-300"
      >
        Dashboard
      </RouterLink>
    </div>
    <div class="mt-8">
      <RouterView/>
    </div>
  </div>
</template>

<style scoped>

</style>