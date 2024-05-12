<script setup lang="ts">
import { ref } from 'vue'
import BreadcrumbDefault from '@/components/Breadcrumbs/BreadcrumbDefault.vue'
import ProfileCard from '@/components/ProfileCard.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import {useUserStore} from "@/stores/user";

const formData = ref({
  name: '',
  oldName: '',
  newPassword: '',
  repeatNewPassword: '',
})

const user = useUserStore()

formData.value.name = user.data.name
formData.value.oldName = user.data.name

const handleSubmit = () => {
}

const handleCancel = () => {
  formData.value.name = formData.value.oldName
  formData.value.newPassword = ''
  formData.value.repeatNewPassword = ''
}

</script>

<template>
  <DefaultLayout>
    <div class="mx-auto max-w-242.5">
      <div
          class="overflow-hidden rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
      >
        <div class="relative z-20 h-35 md:h-65">
          <img
              src="@/assets/images/cover/cover-01.png"
              alt="profile cover"
              class="h-full w-full rounded-tl-sm rounded-tr-sm object-cover object-center"
          />
        </div>
        <div class="px-4 pb-6 text-center lg:pb-8 xl:pb-11.5">
          <div class="mt-4">
            <h3 class="mb-1.5 text-2xl font-medium text-black dark:text-white">{{ user.data.name }}</h3>
            <p class="font-medium">{{ user.data.role }}</p>
            <div
                class="mx-auto mt-4.5 mb-5.5 grid max-w-94 grid-cols-2 rounded-md border border-stroke py-2.5 shadow-1 dark:border-strokedark dark:bg-[#37404F]"
            >
              <div
                  class="flex flex-col items-center justify-center gap-1 border-r border-stroke px-4 dark:border-strokedark xsm:flex-row"
              >
                <span class="font-semibold text-black dark:text-white">4</span>
                <span class="text-sm">Environments</span>
              </div>
              <div
                  class="flex flex-col items-center justify-center gap-1 border-r border-stroke px-4 dark:border-strokedark xsm:flex-row"
              >
                <span class="font-semibold text-black dark:text-white">10</span>
                <span class="text-sm">Services</span>
              </div>
            </div>

            <div class="grid grid-cols-5">
              <!-- Personal Information Section -->
              <div class="col-span-5 xl:col-span-5">
                <div
                    class="rounded-sm border border-stroke bg-white shadow-default dark:border-strokedark dark:bg-boxdark"
                >
                  <div class="border-b border-stroke py-4 px-7 dark:border-strokedark">
                    <h3 class="font-medium text-black dark:text-white">Change Data</h3>
                  </div>
                  <div class="p-7">
                    <form @submit.prevent="handleSubmit">
                      <!-- Name -->
                      <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
                        <div class="w-full sm:w-1/2">
                          <label
                              class="mb-3 block text-sm font-medium text-black dark:text-white"
                              for="name"
                          >Names</label
                          >
                          <div class="relative">
                            <input
                                v-model="formData.name"
                                class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                                type="text"
                                name="name"
                                id="name"
                                placeholder="User Name"
                            />
                          </div>
                        </div>
                        <!-- Name -->

                        <!-- Phone New Password  -->
                        <div class="w-full sm:w-1/2">
                          <label
                              class="mb-3 block text-sm font-medium text-black dark:text-white"
                              for="newPassword"
                          >New Password</label
                          >
                          <input
                              v-model="formData.newPassword"
                              class="w-full rounded border border-stroke bg-gray py-3 px-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                              type="password"
                              name="newPassword"
                              id="newPassword"
                              placeholder="New Password"
                          />
                        </div>
                      </div>
                      <!-- Phone New Password  -->

                      <!-- Repeat New Password -->
                      <div class="mb-5.5 flex flex-col gap-5.5 sm:flex-row">
                        <div class="w-full sm:w-1/2">
                          <label
                              class="mb-3 block text-sm font-medium text-black dark:text-white"
                              for="reapedNewPassword"
                          >Repeat New Password</label
                          >
                          <div class="relative">
                            <input
                                v-model="formData.repeatNewPassword"
                                class="w-full rounded border border-stroke bg-gray py-3 pl-11.5 pr-4.5 font-normal text-black focus:border-primary focus-visible:outline-none dark:border-strokedark dark:bg-meta-4 dark:text-white dark:focus:border-primary"
                                type="password"
                                name="reapedNewPassword"
                                id="reapedNewPassword"
                                placeholder="Repeat New Password"
                            />
                          </div>
                        </div>
                        <!-- Repeat New Password -->
                      </div>
                      <!-- Repeat New Password -->

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
            </div>


          </div>
        </div>
      </div>

    </div>
  </DefaultLayout>
</template>
