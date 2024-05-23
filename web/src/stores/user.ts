import { ref } from "vue";
import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";
import axios from 'axios';

interface User {
    id: string
    name: string
    role: string
    token: string
}

const user = ref<User|null>(null);
const userRoles = {
    admin: "admin",
    user: "user",
}

function getSessionToken(): string | null {
    const meta = document.querySelector('meta[name="session"]');
    return meta ? meta.getAttribute('content') : null;
}

async function fetchUserData() {
    const token = getSessionToken();
    if (!token) {
       return null;
    }

    const response = await axios.get('/api/user-data', {
        headers: {
            'Session-Token': token,
        },
    });

    return response.data;
}

export const useUserStore = defineStore("user", () => {
  const data = useStorage(
      "user",
      user,
      localStorage,
      {
          serializer: {
              read: (v: any) => v ? JSON.parse(v) : null,
              write: (v: any) => JSON.stringify(v),
          },
      },
  );

    async function loadUserData() {
        try {
            const userData = await fetchUserData();
            setUser(userData);
        } catch (error) {
            console.error('Failed to load user data:', error);
        }
    }


  function setUser(newUser: any) {
      data.value = newUser;
  }

  function logout() {
      data.value = null;
  }

  function isAdmin(): boolean {
      return data.value?.role === userRoles.admin;
  }

  return { data, setUser, logout, isAdmin, loadUserData };
});