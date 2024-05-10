<script setup lang="ts">
import {onMounted, ref, watch} from "vue";

import DefaultLayout from '@/layouts/DefaultLayout.vue'
import ServiceListTable from '@/components/Tables/ServiceListTable.vue';
import DefaultCard from "@/components/Forms/DefaultCard.vue";
import MultiSelect from "@/components/Forms/MultiSelect.vue";
import RadioCheckbox from "@/components/Forms/Checkboxes/RadioCheckbox.vue";
import SwitchOne from "@/components/Forms/Switchers/SwitchOne.vue";
import SwitchTwo from "@/components/Forms/Switchers/SwitchTwo.vue";
import SwitchThree from "@/components/Forms/Switchers/SwitchThree.vue";
import SwitchFour from "@/components/Forms/Switchers/SwitchFour.vue";

const filter = ref({
  isActive: ref<boolean>(false)
})

function isActiveFilterHandler(value: any) {
  filter.value.isActive = value
}

const mockEnvironments = [
  { id: 1, value: 1, text: 'Development', selected: false },
  { id: 2, value: 2, text: 'Staging', selected: false },
  { id: 3, value: 3, text: 'Production', selected: false }
]

</script>

<template>
  <DefaultLayout>
    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <div class="col-span-12">
        <DefaultCard cardTitle="Filter">
          <div class="flex flex-row items-center gap-5.5 p-6.5">
            <div>
              <label class="mb-3 block text-sm font-medium  text-black dark:text-white">
                Service ID
              </label>
              <input
                  type="text"
                  placeholder="Service ID"
                  class="w-full rounded-lg border-[1.5px] text-black border-stroke bg-transparent py-3 px-5 font-normal outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:text-white dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
              />
            </div>
            <div>
              <label class="mb-3 block text-sm font-medium text-black dark:text-white">
                Service Name
              </label>
              <input
                  type="text"
                  placeholder="Service Name"
                  class="w-full rounded-lg border-[1.5px] text-black border-stroke bg-transparent py-3 px-5 font-normal outline-none transition focus:border-primary active:border-primary disabled:cursor-default disabled:bg-whiter dark:text-white dark:border-form-strokedark dark:bg-form-input dark:focus:border-primary"
              />
            </div>
            <MultiSelect id="filter-environment" title="Environment" select-title="Select an Environment" :options="mockEnvironments"/>
            <RadioCheckbox  id="filter-active" label="Is Active"  @update="isActiveFilterHandler" />
          </div>
        </DefaultCard>
      </div>
    </div>
    <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
      <div class="col-span-12">
        <ServiceListTable  title="List" show-add-button />
      </div>
    </div>
  </DefaultLayout>
</template>
