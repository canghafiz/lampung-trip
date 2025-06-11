<script setup>
import logo from '../../public/logo.png'
import {UserService} from "../service/user_service.js";
import {userStore} from "../store/user_store.js";
import {useRouter} from "vue-router";
import {onMounted} from "vue";
import {authMiddleware} from "../middleware/auth_middleware.js";

const router = useRouter()

// Store
const useUserStore = userStore()

function logout() {
  UserService.Logout()
}

function masuk() {
  router.push('/masuk')
}

onMounted(async () => {
  await authMiddleware()
})
</script>

<template>
  <nav class="navbar navbar-expand-lg">
    <div class="container-fluid">
      <RouterLink to="/"><img
          :src="`${logo}`"
          class="navbar-brand img-fluid"
          style="height: 54px"
          alt="Logo"
      /></RouterLink>
      <div class="d-flex align-items-center">
        <div v-if="useUserStore.user.userId !== 0" class="dropdown">
          <a
              class="nav-link dropdown-toggle d-flex align-items-center p-0"
              href="#"
              role="button"
              id="profileDropdown"
              data-bs-toggle="dropdown"
              aria-expanded="false"
          >
            <i class="material-icons fs-1 font-color-primary">account_circle</i>
          </a>
          <ul
              class="dropdown-menu dropdown-menu-end"
              aria-labelledby="profileDropdown"
          >
            <li><RouterLink class="dropdown-item" to="/profil">Profil</RouterLink></li>
            <li><a class="dropdown-item" href="#" @click="logout">Logout</a></li>
          </ul>
        </div>

        <button
            v-if="useUserStore.user.userId === 0"
            @click="masuk"
            class="px-2 py-1 rounded-4 btn bg-warning ms-3"
            style="font-size: 16px"
        >
          Masuk / Daftar
        </button>
      </div>
    </div>
  </nav>
</template>

<style scoped>

</style>