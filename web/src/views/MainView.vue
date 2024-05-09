<script setup lang="ts">
import { ref } from 'vue'

import DefaultLayout from '@/layouts/DefaultLayout.vue'
import ServiceListTable  from "@/components/Tables/ServiceListTable.vue";
import MonitoringArea from "@/components/Monitoring/MonitoringArea.vue";


const mockServices = ref([
  { id: 1, name: 'Core', type: 'Container', status: true, workTime: '2h', actions: { setting: true, run: true, stop: true, delete: true } },
  { id: 2, name: 'Order', type: 'Container',  status: true, workTime: '2h', actions: { setting: true, run: true, stop: true, delete: true } },
  { id: 3, name: 'Delivery', type: 'Container', status: false, actions: { setting: true, run: true, stop: true, delete: true } },
  { id: 4, name: 'Balance', type: 'Local', status: false, actions: { setting: true, run: true, stop: true, delete: true } },
])
</script>

<template>
  <DefaultLayout>
    <MonitoringArea />

    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <div class="col-span-12">
        <ServiceListTable :only-active=true title="Active Services" :items="mockServices"/>
      </div>
    </div>
  </DefaultLayout>
</template>
