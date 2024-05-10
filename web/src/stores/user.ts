import { ref } from "vue";
import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";

interface User {
    id: string
    name: string
    role: string
    token: string
}

const user = ref<User|null>(null);

export const useUserStore = defineStore("user", () => {
  const data = useStorage("user", user);

  function setUser(newUser: any) {
      data.value = newUser;
  }

  function logout() {
      data.value = null;
  }

  return { data, setUser, logout };
});