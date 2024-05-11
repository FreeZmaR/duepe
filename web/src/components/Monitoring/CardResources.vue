<script setup lang="ts">
import {ref} from 'vue'
import {useUserStore} from "@/stores/user";

interface IItem {
  title: string
  total: string
  growthRate: string
  isGrowth: boolean
  metaNumber: string
  forAdmin: boolean
}

const cardItems = ref<IItem[]>([
  {
    title: 'CPU usage',
    total: '25.17%',
    growthRate: "0.43%",
    isGrowth: true,
    metaNumber: "5",
    forAdmin: false,
  },
  {
    title: 'RAM usage',
    total: '3.4GB',
    growthRate: "4.35%",
    isGrowth: true,
    metaNumber: "5",
    forAdmin: false,
  },
  {
    title: 'Disk usage',
    total: '1GB',
    growthRate: "2.59%",
    isGrowth: true,
    metaNumber: "5",
    forAdmin: false,
  },
  {
    title: 'Total Environments',
    total: '4',
    growthRate: "-1",
    isGrowth: false,
    metaNumber: "",
    forAdmin: true,
  },
  {
    title: 'Total Active Services',
    total: '36',
    growthRate: "-2",
    isGrowth: false,
    metaNumber: "",
    forAdmin: true,
  }
])

const user = useUserStore()

function canShow(item: IItem): boolean {
  if (user.isAdmin()) {
    return true
  }

  return !item.forAdmin
}

function getColorClassForItem(prefix: string, metaNumber: string): string {
  if (metaNumber === "") {
    return ""
  }

  return `${prefix}-${metaNumber}`
}

</script>

<template>
  <!-- Card Item Start -->
  <div
      v-for="(item, index) in cardItems"
      :key="index"
      class="rounded-sm border border-stroke bg-white py-6 px-7.5 shadow-default dark:border-strokedark dark:bg-boxdark"
  >

    <div v-show="canShow(item)" class="mt-3 flex items-center justify-between">
      <div>
        <span class="text-sm font-medium">{{ item.title }}</span>
      </div>

      <div>
        <h4 class="text-title-md font-bold text-black dark:text-white">{{ item.total }}</h4>
        <span
          class="flex items-center gap-1 text-sm font-medium"
          :class="getColorClassForItem('text-meta', item.metaNumber)"
      >
        {{ item.growthRate }}
        <svg
            v-if="item.isGrowth"
            :class="getColorClassForItem('fill-meta', item.metaNumber)"
            width="10"
            height="11"
            viewBox="0 0 10 11"
            :fill="getColorClassForItem('fill-meta', item.metaNumber) === '' ? 'currentColor' : 'none'"
            xmlns="http://www.w3.org/2000/svg"
        >
          <path
              d="M4.35716 2.47737L0.908974 5.82987L5.0443e-07 4.94612L5 0.0848689L10 4.94612L9.09103 5.82987L5.64284 2.47737L5.64284 10.0849L4.35716 10.0849L4.35716 2.47737Z"
              fill=""
          />
        </svg>

        <svg
            v-if="!item.isGrowth"
            :class="getColorClassForItem('fill-meta', item.metaNumber)"
            width="10"
            height="11"
            viewBox="0 0 10 11"
            :fill="getColorClassForItem('fill-meta', item.metaNumber) === '' ? 'currentColor' : 'none'"
            xmlns="http://www.w3.org/2000/svg"
        >
          <path
              d="M5.64284 7.69237L9.09102 4.33987L10 5.22362L5 10.0849L-8.98488e-07 5.22362L0.908973 4.33987L4.35716 7.69237L4.35716 0.0848701L5.64284 0.0848704L5.64284 7.69237Z"
              fill=""
          />
        </svg>
      </span>
      </div>
    </div>
  </div>
  <!-- Card Item End -->
</template>



