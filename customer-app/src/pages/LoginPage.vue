<script setup>

import {computed, ref, watchEffect} from "vue";
import logo from "../../public/logo.png";
import hero from "../../public/hero.jpg";
import {LoginRequest} from "../model/request/login_request.js";
import {UserService} from "../service/user_service.js";
import Modal from "../components/Modal.vue";

const pageTitle = ref('Masuk - Lampung Trip');

watchEffect(() => {
  document.title = `${pageTitle.value}`;
})

const heroBackgroundStyle = computed(() => {
  return {
    backgroundImage: `url(${hero})`, // Gunakan variabel 'hero' langsung
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover',
    backgroundPosition: 'center',
  };
});

// State
const username = ref("")
const password = ref("")
const errorMsg = ref("")
const loading = ref(false);

async function submitHandle(event) {
  event.preventDefault();
  loading.value = true;

  const request = new LoginRequest(username.value, password.value)

  await UserService.Login(request).then((response) => {
    if (response !== null) {
      errorMsg.value = response;
    } else {
      window.location.reload();
    }
  }).finally(() => {
    loading.value = false;
  })
}

function closeModal() {
  errorMsg.value = '';
}
</script>

<template>
  <Modal title="Error!" :show="errorMsg !== ''" :message="errorMsg" @close="closeModal" />

  <div class="content-wrapper row w-100 vh-100 overflow-hidden m-0">
    <div class="col-md-5 d-flex flex-column align-items-center p-md-5 p-3">
      <div class="d-flex justify-content-between flex-row">
        <div>
          <img :src="`${logo}`" class="img-fluid h-25" alt="Logo" />
        </div>
        <p>Belum punya akun ? <RouterLink :to="{ name: 'daftar' }">Daftar</RouterLink></p>
      </div>

      <form class="d-flex flex-column align-items-start w-100" @submit="submitHandle">
        <div class="p-2 rounded-4 bg-white my-4">
          <i class="material-icons fs-1 font-color-primary bg-transparent"
          >person</i
          >
        </div>
        <h3 class="font-color-primary fw-normal">Masuk</h3>
        <p class="font-color-primary fw-thin">
          Masukkan detail akun anda di sini
        </p>

        <div class="my-4 w-100">
          <div class="mb-3">
            <label for="username" class="form-label font-color-secondary"
            >Username</label
            >
            <input
                type="text"
                class="form-control py-2 rounded-5"
                id="username"
                placeholder="Masukkan username"
                v-model="username"
                required
            />
          </div>
          <div class="mb-3">
            <label for="password" class="form-label font-color-secondary"
            >Kata Sandi</label
            >
            <input
                type="password"
                class="form-control py-2 rounded-5"
                id="password"
                placeholder="Masukkan kata sandi"
                v-model="password"
                required
            />
          </div>
          <div class="d-flex justify-content-end mb-3">
            <RouterLink :to="{ name: 'ubahSandi' }" class="text-decoration-none">Lupa Sandi ?</RouterLink>
          </div>
        </div>

        <button
            type="submit"
            class="py-1 fs-5 bg-color-primary mt-4 w-100 rounded-5"
        >
          <span v-if="loading" class="spinner-border" role="status">
            <span class="sr-only"></span>
          </span>

          <span v-else>
            Masuk
          </span>
        </button>
      </form>
    </div>
    <div
        class="col-md d-none d-md-flex flex-column justify-content-center align-items-center p-0 bg-transparent rounded-start-4"
        :style="heroBackgroundStyle"
    >
      <div
          class="d-flex flex-column justify-content-center align-items-center text-white text-center p-4 h-100 rounded-start-4"
          style="background-color: rgba(0, 0, 0, 0.5)"
      >
        <p class="mb-0 fs-4 fst-italic">
          Lampung Trip: Lebih dari Sekedar Destinasi, Melainkan Sebuah Kisah
          Perjalanan di Surga Sumatera.
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>