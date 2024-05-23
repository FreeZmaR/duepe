<script setup lang="ts">
import {onMounted, ref} from 'vue';
import { RouterView } from 'vue-router'
import { useUserStore } from "@/stores/user";
import DefaultLoader from "@/components/Loaders/DefaultLoader.vue";

const loaded = ref(false);

onMounted(async () => {
  await useUserStore().loadUserData();
  loaded.value = true;
});
</script>

<template>
  <DefaultLoader v-if="!loaded" />
  <RouterView v-else/>
</template>
