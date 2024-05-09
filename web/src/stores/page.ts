import { ref } from 'vue'
import { defineStore } from 'pinia'
import { useStorage } from '@vueuse/core'

export const usePageStore = defineStore('page', () => {
  const title = useStorage('pageTitle', ref('Main'))

  function setTitle(newTitle: string) {
    title.value = newTitle
  }

  return { title, setTitle }
})