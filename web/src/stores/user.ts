import { ref } from "vue";
import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";

interface User {
    id: string
    name: string
    role: string
}


export const useUserStore = defineStore("user", () => {
  const user = useStorage("user", ref<User>({id: "Admin", name: "Admin", role: "Admin"}));

  function setUser(newUser: any) {
    user.value = newUser;
  }

  return { user, setUser };
});