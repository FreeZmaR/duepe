<script setup lang="ts">
import { ref } from 'vue'
import UserListTable from "@/components/Tables/UserListTable.vue";

const formData = ref({
  maxEnv: 5,
  maxUserEnv: 3,
  maxServices: 5,
  maxUserServices: 3,
  maxRanServices: 5,
  maxUserRanServices: 3,
  ramLimit: 500,
  ramServiceLimit: 100,
  diskLimit: 500,
  diskServiceLimit: 100,
})


const handleSubmit = () => {
}

const handleCancel = () => {
}

</script>

<template>
  <div class="grid grid-cols-5 gap-8">
    <!-- Personal Information Section -->
    <div class="col-span-5 xl:col-span-3">
      <div
        class="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
      >
        <div class="border-b border-stroke py-4 px-7 dark:border-strokedark">
          <h3 class="font-medium text-black dark:text-white">Duepe Setting</h3>
        </div>
        <div class="p-7">
          <form @submit.prevent="handleSubmit">
            <!-- Max Env -->
            <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
              <div class="w-full sm:w-1/2">
                <label
                  class="mb-3 block text-sm font-medium text-black dark:text-white"
                  for="maxEnv"
                  >Max Environments</label
                >
                <div class="relative">
                  <input
                    v-model="formData.maxEnv"
                    class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                    type="number"
                    name="maxEnv"
                    id="maxEnv"
                    placeholder="Total environments allowed"
                  />
                </div>
              </div>

              <!-- Phone Number Section -->
              <div class="w-full sm:w-1/2">
                <label
                  class="mb-3 block text-sm font-medium text-black dark:text-white"
                  for="maxUserEnv"
                  >Max Environments for User</label
                >
                <input
                  v-model="formData.maxUserEnv"
                  class="w-full rounded border border-stroke bg-gray py-3 px-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                  type="number"
                  name="maxUserEnv"
                  id="maxUserEnv"
                  placeholder="Environments allowed for user"
                />
              </div>
            </div>
            <!-- Max Env -->

            <!-- Max Services -->
            <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="maxService"
                >Max Services</label
                >
                <div class="relative">
                  <input
                      v-model="formData.maxServices"
                      class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                      type="number"
                      name="maxService"
                      id="maxService"
                      placeholder="Total services allowed"
                  />
                </div>
              </div>

              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="maxUserServices"
                >Max Services for User</label
                >
                <input
                    v-model="formData.maxUserServices"
                    class="w-full rounded border border-stroke bg-gray py-3 px-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                    type="number"
                    name="maxUserServices"
                    id="maxUserServices"
                    placeholder="Services allowed for user"
                />
              </div>
            </div>
            <!-- Max Services -->

            <!-- RAM Limit -->
            <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="ramLimit"
                >RAM Limit MB</label
                >
                <div class="relative">
                  <input
                      v-model="formData.ramLimit"
                      class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                      type="number"
                      name="ramLimit"
                      id="ramLimit"
                      placeholder="RAM Limit MB"
                  />
                </div>
              </div>

              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="ramServiceLimit"
                >RAM Service Limit MB</label
                >
                <input
                    v-model="formData.ramServiceLimit"
                    class="w-full rounded border border-stroke bg-gray py-3 px-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                    type="number"
                    name="ramServiceLimit"
                    id="ramServiceLimit"
                    placeholder="RAM Limit MB for service"
                />
              </div>
            </div>
            <!-- RAM Limit -->

            <!-- DISK Limit -->
            <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="diskLimit"
                >DISK Limit MB</label
                >
                <div class="relative">
                  <input
                      v-model="formData.diskLimit"
                      class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                      type="number"
                      name="diskLimit"
                      id="diskLimit"
                      placeholder="DISK Limit MB"
                  />
                </div>
              </div>

              <div class="w-full sm:w-1/2">
                <label
                    class="mb-3 block text-sm font-medium text-black dark:text-white"
                    for="diskServiceLimit"
                >DISK Service Limit MB</label
                >
                <input
                    v-model="formData.diskServiceLimit"
                    class="w-full rounded border border-stroke bg-gray py-3 px-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                    type="number"
                    name="diskServiceLimit"
                    id="diskServiceLimit"
                    placeholder="DISK Limit MB for service"
                />
              </div>
            </div>
            <!-- DISK Limit -->

            <!-- Save and Cancel Buttons -->
            <div class="flex justify-end gap-4.5">
              <button
                class="flex justify-center rounded border border-stroke py-2 px-6 font-medium text-black hover:shadow-1 dark:border-strokedark dark:text-white"
                type="button"
                @click="handleCancel"
              >
                Cancel
              </button>
              <button
                class="flex justify-center rounded bg-primary py-2 px-6 font-medium text-gray hover:bg-opacity-90"
                type="submit"
              >
                Save
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <div class="col-span-5 xl:col-span-2">
      <div
          class="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
      >
        <div class="border-b border-stroke py-4 px-7 dark:border-strokedark">
          <h3 class="font-medium text-black dark:text-white">Duepe Actions</h3>
        </div>
      </div>
    </div>
  </div>



  <!-- User List Table Section -->
  <div class="mt-4 grid grid-cols-12 gap-4 md:mt-6 md:gap-6 2xl:mt-7.5 2xl:gap-7.5">
    <div class="col-span-12">
      <UserListTable />
    </div>
  </div>
  <!-- User List Table Section -->

</template>
