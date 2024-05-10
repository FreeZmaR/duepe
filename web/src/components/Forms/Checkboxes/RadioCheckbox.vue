<script setup lang="ts">
import { ref } from 'vue'

interface IProps {
  id: any
  label: string
}

const props = defineProps<IProps>()
const checkboxToggle = ref<boolean>(false)

const emit = defineEmits<{
  (e: 'update', value: boolean): void
}>()

function toggleCheckbox() {
  checkboxToggle.value = !checkboxToggle.value
  emit('update', checkboxToggle.value)
}

</script>

<template>
  <div>
    <label :for="props.id" class="flex cursor-pointer select-none items-center">
      <div class="mr-2">
        <span class="text-black dark:text-white">{{ props.label }}</span>
      </div>
      <div class="relative">
        <input
            type="checkbox"
            :id="props.id"
            class="sr-only"
            @change="toggleCheckbox"
        />
        <div
            :class="checkboxToggle && 'border-primary'"
            class="mr-4 flex h-5 w-5 items-center justify-center rounded-full border"
        >
          <span
              :class="checkboxToggle && '!bg-primary'"
              class="h-2.5 w-2.5 rounded-full bg-transparent"
          >
          </span>
        </div>
      </div>

    </label>
  </div>
</template>
