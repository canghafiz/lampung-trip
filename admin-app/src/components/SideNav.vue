<script setup>

import {sideNavStore} from "../store/side_nav.js";
import Navbar from "./Navbar.vue";
import Footer from "./Footer.vue";
import {authStore} from "../store/auth_store.js";

// Store
const store = sideNavStore()
const useAuthStore = authStore()

function navigate(value) {
  store.setPage(value)
}
</script>

<template>
  <div class="min-height-300 bg-dark position-absolute w-100"></div>
  <aside
      class="sidenav bg-white navbar navbar-vertical navbar-expand-xs border-0 border-radius-xl my-3 fixed-start ms-4"
      id="sidenav-main"
  >
    <div class="sidenav-header">
      <i
          class="fas fa-times p-3 cursor-pointer text-secondary opacity-5 position-absolute end-0 top-0 d-none d-xl-none"
          aria-hidden="true"
          id="iconSidenav"
      ></i>
      <a
          class="navbar-brand m-0"
          href=" https://demos.creative-tim.com/argon-dashboard/pages/dashboard.html "
          target="_blank"
      >
        <img
            src="../assets/img/logo.png"
            width="26px"
            height="26px"
            class="navbar-brand-img h-100"
            alt="main_logo"
        />
        <span class="ms-1 font-weight-bold">Lampung Trip</span>
      </a>
    </div>
    <hr class="horizontal dark mt-0" />
    <div class="collapse navbar-collapse w-auto" id="sidenav-collapse-main">
      <ul class="navbar-nav">
        <li class="nav-item">
          <a :class="`nav-link ${store.state.page === 'Wisata' || store.state.page === 'Kelola Wisata' ? 'active' : ''}`" href="javascript:void(0)"
             @click="navigate('Wisata')">
            <div
                class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center"
            >
              <i
                  class="ni ni-calendar-grid-58 text-dark text-sm opacity-10"
              ></i>
            </div>
            <span class="nav-link-text ms-1">Wisata</span>
          </a>
        </li>
        <li class="nav-item">
          <a :class="`nav-link ${store.state.page === 'Villa' || store.state.page === 'Kelola Villa' ? 'active' : ''}`" href="javascript:void(0)"
             @click="navigate('Villa')">
            <div
                class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center"
            >
              <i
                  class="ni ni-calendar-grid-58 text-dark text-sm opacity-10"
              ></i>
            </div>
            <span class="nav-link-text ms-1">Villa</span>
          </a>
        </li>
        <li v-if="useAuthStore.user.role === 'Admin'" class="nav-item">
          <a :class="`nav-link ${store.state.page === 'Karyawan' || store.state.page === 'Kelola Karyawan' ? 'active' : ''}`" href="javascript:void(0)"
             @click="navigate('Karyawan')">
            <div
                class="icon icon-shape icon-sm border-radius-md text-center me-2 d-flex align-items-center justify-content-center"
            >
              <i class="ni ni-single-02 text-dark text-sm opacity-10"></i>
            </div>
            <span class="nav-link-text ms-1">Karyawan</span>
          </a>
        </li>
      </ul>
    </div>
  </aside>
  <main class="main-content position-relative border-radius-lg">
    <Navbar :Halaman="`${store.state.page}`"/>
    <div class="container-fluid py-4">
      <div class="row">
        <div class="col-12">
          <slot />
        </div>
      </div>
    </div>
  </main>

  <Footer/>
</template>

<style scoped>

</style>