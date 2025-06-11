<script setup>

import SideNav from "../components/SideNav.vue";
import {ref, watch, watchEffect} from "vue";
import WisataPage from "./WisataPage.vue";
import {sideNavStore} from "../store/side_nav.js";
import VillaPage from "./VillaPage.vue";
import KaryawanPage from "./KaryawanPage.vue";
import KelolaWisata from "./KelolaWisata.vue";
import KelolaVilla from "./KelolaVilla.vue";
import KelolaKaryawan from "./KelolaKaryawan.vue";
import {wisataStore} from "../store/wisata_store.js";

const pageTitle = ref('Wisata - Admin Lampung Trip');

// Store
const sidenavStore = sideNavStore()
const useWisataStore = wisataStore()

watch(() => sidenavStore.state.page, (newVal) => {
  pageTitle.value = newVal;
})

watchEffect(() => {
  document.title = `${pageTitle.value} - Admin Lampung Trip`;
})

const handleFabClick = () => {
  if (sidenavStore.state.page === "Wisata") {
    sidenavStore.setPage("Kelola Wisata")
    useWisataStore.clearEditData()
  }

  if (sidenavStore.state.page === "Villa") {
    sidenavStore.setPage("Kelola Villa")
  }

  if (sidenavStore.state.page === "Karyawan") {
    sidenavStore.setPage("Kelola Karyawan")
  }
};

// Store
const useSideNavStore = sideNavStore()
</script>

<template>
  <SideNav>
    <WisataPage v-if="sidenavStore.state.page === 'Wisata'" />
    <KelolaWisata v-if="sidenavStore.state.page === 'Kelola Wisata'" />

    <VillaPage v-if="sidenavStore.state.page === 'Villa'" />
    <KelolaVilla v-if="sidenavStore.state.page === 'Kelola Villa'" />

    <KaryawanPage v-if="sidenavStore.state.page === 'Karyawan'" />
    <KelolaKaryawan v-if="sidenavStore.state.page === 'Kelola Karyawan'" />
  </SideNav>

  <button
      v-if="useSideNavStore.state.page !== 'Kelola Wisata' && useSideNavStore.state.page !== 'Kelola Villa' && useSideNavStore.state.page !== 'Kelola Karyawan' "
      @click="handleFabClick"
      class="fab-button"
      aria-label="Add new item"
      title="Add new item"
  >
    <i class="fa fa-plus"></i>
  </button>
</template>

<style scoped>
.fab-button {
  position: fixed;
  bottom: 24px;
  right: 24px;
  width: 64px;
  height: 64px;
  background-color: #344767;
  color: white;
  border: none;
  border-radius: 50%;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 1000;
  transition: transform 0.2s ease-in-out;
}

.fab-button:hover {
  transform: scale(1.05);
}

.fab-button i {
  font-size: 24px;
}
</style>