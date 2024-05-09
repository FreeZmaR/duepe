<script setup lang="ts">

interface IProps {
  columns: number
}

const props = defineProps<IProps>()
</script>

<template>
  <tr>
    <template v-for="i in props.columns" :key="i">
      <td v-if="i === 0" class="py-5 px-4 pl-9 xl:pl-11">
        <p class="text-sm"> - </p>
      </td>
      <td v-else class="py-5 px-4">
        <p class="text-sm"> - </p>
      </td>
    </template>
  </tr>
</template>
