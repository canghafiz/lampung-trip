<script setup>
import {onMounted, onUnmounted, watch} from 'vue';
import {sideNavStore} from "../store/side_nav.js";
import { ref } from 'vue';
import {UserService} from "../services/user_service.js";
import {searchStore} from "../store/search_store.js";
import {wisataStore} from "../store/wisata_store.js";
import {villaStore} from "../store/villa_store.js";
import {karyawanStore} from "../store/karyawan_store.js";

const props = defineProps({
  Halaman: String,
});

const toggleSidenav = () => {
  const body = document.getElementsByTagName('body')[0];
  const className = 'g-sidenav-pinned';

  if (body.classList.contains(className)) {
    body.classList.remove(className);
  } else {
    body.classList.add(className);
  }
};

const handleResize = () => {
  const body = document.getElementsByTagName('body')[0];
  const className = 'g-sidenav-pinned';
  const breakpoint = 1200;

  if (window.innerWidth < breakpoint) {
    if (body.classList.contains(className)) {
      body.classList.remove(className);
    }
  } else {
  }
};

onMounted(() => {
  handleResize();
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

// store
const useSideNavStore = sideNavStore()
const useSearchStore = searchStore()
const useWisataStore = wisataStore()
const useVillaStore = villaStore()
const useKaryawanStore  = karyawanStore()

const search = ref("")
watch(search, (newValue) => {
  if (useSideNavStore.state.page === "Wisata") {
    useSearchStore.setWisataResult(newValue, useWisataStore.state.listData)
  }

  if (useSideNavStore.state.page === "Villa") {
    useSearchStore.setVillaResult(newValue, useVillaStore.state.listData)
  }

  if (useSideNavStore.state.page === "Karyawan") {
    useSearchStore.setKaryawanResult(newValue, useKaryawanStore.state.listData)
  }

  useSearchStore.setInput(newValue)
});

const tampilModal = ref(false);

const tampilkanModal = () => {
  tampilModal.value = true;
};

const sembunyikanModal = () => {
  tampilModal.value = false;
};

const keluarAkun = () => {
  UserService.Logout()
};
</script>

<template>
  <!-- Modal Keluar Akun -->
  <div v-if="tampilModal" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Konfirmasi Keluar</h5>
        <button type="button" class="close-button" @click="sembunyikanModal">&times;</button>
      </div>
      <div class="modal-body">
        <p>Apakah Anda yakin ingin keluar dari akun ini?</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="modal-button btn-batal" @click="sembunyikanModal">Batal</button>
        <button type="button" class="modal-button btn-yakin" @click="keluarAkun">Yakin</button>
      </div>
    </div>
  </div>

  <nav
      class="navbar navbar-main navbar-expand-lg px-0 mx-4 shadow-none border-radius-xl"
      id="navbarBlur"
      data-scroll="false"
  >
    <div class="container-fluid py-1 px-3">
      <div class="d-flex align-items-center">
        <nav aria-label="breadcrumb">
          <ol
              class="breadcrumb bg-transparent mb-0 pb-0 pt-1 px-0 me-sm-6 me-5"
          >
            <li class="breadcrumb-item text-sm">
              <a class="opacity-5 text-white" href="javascript:">Halaman</a>
            </li>
            <li
                class="breadcrumb-item text-sm text-white active"
                aria-current="page"
            >
              {{ props.Halaman }}
            </li>
          </ol>
          <h6 class="font-weight-bolder text-white mb-0">{{ props.Halaman }}</h6>
        </nav>
      </div>
      <div v-if="useSideNavStore.state.page !== 'Kelola Wisata' && useSideNavStore.state.page !== 'Kelola Villa' && useSideNavStore.state.page !== 'Kelola Karyawan' " class="ms-md-auto pe-md-3 d-flex align-items-center">
        <div class="input-group">
          <span class="input-group-text text-body">
            <i class="fas fa-search" aria-hidden="true"></i>
          </span>
          <input
              type="text"
              class="form-control"
              placeholder="Cari..."
              aria-label="Cari"
              v-model="search"
          />
        </div>
      </div>
      <div class="ms-md-auto mt-sm-0 mt-2" id="navbar">
        <ul class="navbar-nav justify-content-end">
          <li
              class="nav-item d-xl-none ps-3 d-flex align-items-center me-3"
          >
            <a
                href="javascript:"
                class="nav-link text-white p-0"
                id="iconNavbarSidenav"
                @click="toggleSidenav"
            >
              <div class="sidenav-toggler-inner">
                <i class="sidenav-toggler-line bg-white"></i>
                <i class="sidenav-toggler-line bg-white"></i>
                <i class="sidenav-toggler-line bg-white"></i>
              </div>
            </a>
          </li>
          <li class="nav-item d-flex align-items-center">
            <a
                @click="tampilkanModal"
                href="javascript:"
                class="nav-link text-white font-weight-bold px-0"
            >
              <i class="fa fa-sign-out me-sm-1"></i>
              <span class="d-sm-inline d-none">Keluar</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </nav>
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

.btn-batal {
  background: none;
  border: 1px solid #344767;
  color: #344767;
}

.btn-yakin {
  background-color: #344767;
  color: white;
}

.btn-batal:hover,
.btn-batal:focus {
  background-color: #344767;
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
