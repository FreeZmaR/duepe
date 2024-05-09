<script setup lang="ts">
import TableOne from '@/components/Tables/TableOne.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import CardResources  from "@/components/Monitoring/CardResources.vue";
import EnvGraph from "@/components/Monitoring/EnvGraph.vue";
import EnvCycleGraph from "@/components/Monitoring/EnvCycleGraph.vue";
</script>

<template>
  <DefaultLayout>
    <div class="grid grid-cols-1 gap-4 md:grid-cols-2 md:gap-6 xl:grid-cols-4 2xl:gap-7.5">
      <CardResources />
    </div>

    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <EnvGraph />

      <EnvCycleGraph />

      <div class="col-span-12">
        <TableOne />
      </div>
    </div>
  </DefaultLayout>
</template>
