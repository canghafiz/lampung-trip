<script setup>

import {computed, ref, watchEffect} from "vue";
import logo from "../../public/logo.png";
import hero from "../../public/hero.jpg";
import {RegisterRequest} from "../model/request/register_request.js";
import {UserService} from "../service/user_service.js";
import Modal from "../components/Modal.vue";
import {useRouter} from "vue-router";

const pageTitle = ref('Daftar - Lampung Trip');
const router = useRouter()

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
const nama = ref("")
const nomor = ref("")
const username = ref("")
const password = ref("")
const errorMsg = ref("")
const loading = ref(false);
const success = ref(false);

async function submitHandle(event) {
  event.preventDefault();
  loading.value = true;

  const request = new RegisterRequest(username.value, password.value, nama.value, nomor.value)

  await UserService.Register(request).then((response) => {
    if (response === null) {
      success.value = true;
    } else {
      errorMsg.value = response;
    }
  }).finally(() => {
    loading.value = false;
  })
}

function closeModal() {
  errorMsg.value = '';
}

function tutupModalSuccess() {
  router.push("/masuk");
}
</script>

<template>
  <Modal title="Error!" :show="errorMsg !== ''" :message="errorMsg" @close="closeModal" />
  <div v-if="success" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close-button" @click="tutupModalSuccess">&times;</button>
      </div>
      <div class="modal-body">
        <p>Akun anda telah berhasil dibuat</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="modal-button btn-ok" @click="tutupModalSuccess">OK</button>
      </div>
    </div>
  </div>

  <div class="content-wrapper row w-100 vh-100 overflow-x-hidden m-0">
    <div class="col-md-5 d-flex flex-column align-items-center p-md-5 p-3">
      <div class="d-flex justify-content-between flex-row">
        <div>
          <img :src="`${logo}`" class="img-fluid h-25" alt="Logo" />
        </div>
        <p>Sudah punya akun ? <RouterLink :to="{ name: 'masuk' }">Masuk</RouterLink></p>
      </div>

      <form class="d-flex flex-column align-items-start w-100" @submit="submitHandle">
        <div class="p-2 rounded-4 bg-white my-4">
          <i class="material-icons fs-1 font-color-primary bg-transparent"
          >person</i
          >
        </div>
        <h3 class="font-color-primary fw-normal">Daftar</h3>
        <p class="font-color-primary fw-thin">
          Masukkan detail akun anda di sini
        </p>

        <div class="my-4 w-100">
          <div class="mb-3">
            <label for="nama" class="form-label font-color-secondary"
            >Nama</label
            >
            <input
                type="text"
                class="form-control py-2 rounded-5"
                id="nama"
                placeholder="Masukkan nama anda"
                v-model="nama"
                required
            />
          </div>
          <div class="mb-3">
            <label for="no" class="form-label font-color-secondary"
            >No Hp</label
            >
            <input
                type="text"
                class="form-control py-2 rounded-5"
                id="no"
                placeholder="Masukkan nomor anda"
                v-model="nomor"
                required
            />
          </div>
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
        </div>

        <button
            type="submit"
            class="py-1 fs-5 bg-color-primary mt-4 w-100 rounded-5"
        >
          <span v-if="loading" class="spinner-border" role="status">
            <span class="sr-only"></span>
          </span>

          <span v-else>
            Daftar
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
/* Modal Styles */
.modal-overlay {
  display: flex;
  justify-content: center;
  align-items: center;
  position: fixed;
  z-index: 9999; /* Lebih tinggi dari elemen lain */
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
  background-color: #fefefe;
  margin: auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 500px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.modal-title {
  font-size: 1.2em;
  font-weight: bold;
  margin: 0;
}

.close-button {
  color: #aaa;
  font-size: 28px;
  font-weight: bold;
  cursor: pointer;
  border: none;
  background: none;
  padding: 0;
}

.close-button:hover,
.close-button:focus {
  color: black;
  text-decoration: none;
}

.modal-body {
  padding: 20px 0;
}

.modal-footer {
  display: flex;
  justify-content: flex-end;
  padding-top: 15px;
  border-top: 1px solid #eee;
}

.modal-button {
  padding: 10px 15px;
  margin-left: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1em;
}

.btn-ok {
  background-color: #103d47;
  color: white;
}

.btn-batal:hover,
.btn-batal:focus {
  background-color: #103d47;
  color: #ffffff;
}

.btn-yakin:hover,
.btn-yakin:focus {
  background-color: #1b293a;
}

/* Responsiveness */
@media (max-width: 600px) {
  .modal-content {
    width: 95%;
  }
}
</style>