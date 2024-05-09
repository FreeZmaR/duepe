<script setup lang="ts">

import DefaultLayout from "@/layouts/DefaultLayout.vue"
import EnvironmentListTable from "@/components/Tables/EnvironmentListTable.vue"

</script>

<template>
  <DefaultLayout>
    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <div class="col-span-12">
        <EnvironmentListTable />
      </div>
    </div>
  </DefaultLayout>
</template>
