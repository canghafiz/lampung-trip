<script setup>
import { ref, watchEffect } from 'vue';
import heroImage from '../assets/img/hero.jpg';
import {UserService} from "../services/user_service.js";
import {PasswordRequest} from "../models/request/password_request.js";
import Modal from "../components/Modal.vue";

const pageTitle = ref('Ubah Kata Sandi - Admin Lampung Trip');
const copyright = new Date().getFullYear();

watchEffect(() => {
  document.title = `${pageTitle.value}`;
})

// State
const username = ref("")
const password_lama = ref("")
const password_baru = ref("")
const loading = ref(false);
const errorMsg = ref('');
const successMsg = ref('');

async function formHandle(event) {
  event.preventDefault();
  loading.value = true;

  const request = new PasswordRequest(username.value, password_lama.value, password_baru.value);

  await UserService.Password(request).then((response) => {
    // Jika gagal
    if (response) {
      errorMsg.value = response;
    }

    if (response === null) {
      successMsg.value = "Password berhasil diganti";
    }
  }).finally(() => {
    loading.value = false;
  })
}

function closeModal() {
  errorMsg.value = '';
  successMsg.value = '';
}
</script>

<template>

  <Modal title="Berhasil!" :show="successMsg !== ''" :message="successMsg" @close="closeModal" />
  <Modal title="Error!" :show="errorMsg !== ''" :message="errorMsg" @close="closeModal" />

  <main class="main-content mt-0">
    <div
        class="page-header align-items-start min-vh-50 pt-5 pb-11 m-3 border-radius-lg"
        :style="{
          backgroundImage: `url(${heroImage})`,
          backgroundPosition: 'top'
        }"
    >
      <span class="mask bg-gradient-dark opacity-6"></span>
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-5 text-center mx-auto">
            <h1 class="text-white mb-2 mt-5">Selamat Datang!</h1>
            <p class="text-lead text-white">
              Jika anda lupa kata sandi, anda bisa merubah disini
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="container">
      <div class="row mt-lg-n10 mt-md-n11 mt-n10 justify-content-center">
        <div class="col-xl-4 col-lg-5 col-md-7 mx-auto">
          <div class="card z-index-0">
            <div class="card-header text-center pt-4">
              <h5>Ubah Kata Sandi</h5>
            </div>
            <div class="card-body">
              <form role="form" @submit="formHandle">
                <div class="mb-3">
                  <input
                      type="text"
                      class="form-control"
                      placeholder="Username"
                      aria-label="Username"
                      v-model="username"
                      required
                  />
                </div>
                <div class="mb-3">
                  <input
                      type="password"
                      class="form-control"
                      placeholder="Kata Sandi Lama"
                      aria-label="Kata Sandi Lama"
                      v-model="password_lama"
                      required
                  />
                </div>
                <div class="mb-3">
                  <input
                      type="password"
                      class="form-control"
                      placeholder="Kata Sandi Baru"
                      aria-label="Kata Sandi Baru"
                      v-model="password_baru"
                      required
                  />
                </div>
                <div class="text-center">
                  <button
                      type="submit"
                      class="btn bg-gradient-primary w-100 my-4 mb-2"
                  >
                    {{ loading ? 'Loading..' : 'Ganti' }}
                  </button>
                </div>
                <p class="text-sm mt-3 mb-0 text-center">
                  <RouterLink :to="{ name: 'login' }" class="text-primary font-weight-bolder">
                    Kembali Masuk
                  </RouterLink>
                </p>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
  <footer class="footer py-5">
    <div class="container">
      <div class="row">
        <div class="col-8 mx-auto text-center mt-1">
          <p class="mb-0 text-secondary">
            Copyright © {{ copyright }} - Lampung Trip
          </p>
        </div>
      </div>
    </div>
  </footer>
</template>

<style scoped>

</style>