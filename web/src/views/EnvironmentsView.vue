<script setup lang="ts">
import {ref} from "vue";

import DefaultLayout from '@/layouts/DefaultLayout.vue'
import EnvironmentListTable from "@/components/Tables/EnvironmentListTable.vue"

const mockEnvironments = ref([
  { id: 1, name: "Production", user: "John Doe", runningServices: 5, totalServices: 5 },
  { id: 2, name: "Staging", user: "Admin", runningServices: 3, totalServices: 5 },
  { id: 3, name: "Development", user: "Ban" }
])

</script>

<template>
  <DefaultLayout>
    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <div class="col-span-12">
        <EnvironmentListTable title="List" :items="mockEnvironments"/>
      </div>
    </div>
  </DefaultLayout>
</template>
