<script setup>
import { ref, watchEffect } from 'vue';
import heroImage from '../assets/img/hero.jpg';
import Footer from '../components/Footer.vue';
import {UserService} from "../services/user_service.js";
import {LoginRequest} from "../models/request/login_request.js";
import Modal from "../components/Modal.vue";

const pageTitle = ref('Login - Admin Lampung Trip');

watchEffect(() => {
  document.title = `${pageTitle.value}`;
})

// State
const username = ref("")
const password = ref("")
const loading = ref(false);
const errorMsg = ref('');

async function loginHandle(event) {
  event.preventDefault();
  loading.value = true;

  const request = new LoginRequest(username.value, password.value)

  await UserService.Login(request).then((response) => {
    // Jika login gagal
    if (response) {
      errorMsg.value = response;

      console.log(response)
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
              Untuk masuk pastikan masukkan informasi akun kamu username dan
              kata sandi
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
              <h5>Masuk</h5>
            </div>
            <div class="card-body">
              <form role="form" @submit="loginHandle">
                <div class="mb-3">
                  <input
                      type="text"
                      class="form-control"
                      placeholder="Username"
                      aria-label="Username"
                      required
                      v-model="username"
                  />
                </div>
                <div class="mb-3">
                  <input
                      type="password"
                      class="form-control"
                      placeholder="Kata Sandi"
                      aria-label="Kata Sandi"
                      v-model="password"
                      required
                  />
                </div>
                <div class="text-center">
                  <button
                      type="submit"
                      class="btn bg-gradient-primary w-100 my-4 mb-2"
                  >
                    {{ loading ? 'Loading..' : 'Masuk' }}
                  </button>
                </div>
                <p class="text-sm mt-3 mb-0 text-center">
                  Lupa kata sandi?
                  <RouterLink :to="{ name: 'password' }" class="text-primary font-weight-bolder">
                    Ganti
                  </RouterLink>
                </p>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>

  <Footer/>
</template>

<style scoped>
/* Your component styles */
</style>