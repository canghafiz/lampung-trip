<script setup>
import MainLayout from "../layouts/MainLayout.vue";
import {tabStore} from "../store/tab_store.js";
import {computed, onMounted, ref, watch} from "vue";
import {wisataStore} from "../store/wisata_store.js";
import ProductCard from "../components/ProductCard.vue";
import {villaStore} from "../store/villa_store.js";
import {searchStore} from "../store/search_store.js";

// Store
const useTabStore = tabStore()
const useWisataStore = wisataStore()
const useVillaStore = villaStore()
const useSearchStore = searchStore()

// State
const searchInput = ref("")

function setTab(value) {
  useTabStore.setType(value)
  searchInput.value = ""
}

onMounted(() => {
  document.title = "Lampung Trip"
})

watch(useTabStore.state, async (newValue) => {
  // Wisata
  if (newValue.type === "Wisata") {
    await useWisataStore.initListData()
  }

  // Villa
  if (newValue.type === "Villa") {
    await useVillaStore.initListData()
  }
},{ immediate: true })

watch(searchInput, async (newValue) => {
  if (newValue !== "") {
    // Wisata
    if (useTabStore.state.type === 'Wisata') {
      useSearchStore.setWisataResult(searchInput.value, useWisataStore.state.listData)
    }

    // Villa
    if (useTabStore.state.type === 'Villa') {
      useSearchStore.setVillaResult(searchInput.value, useVillaStore.state.listData)
    }
  }
})

const listData = computed(() => {
  if (searchInput.value.length > 0) {
    return useSearchStore.state.result
  }

  return (useTabStore.state.type === "Villa") ? useVillaStore.state.listData : useWisataStore.state.listData
})
</script>

<template>
  <MainLayout>
    <!-- Search & Tab -->
    <div class="d-flex my-4">
      <input
          class="form-control me-2 bg-white rounded-4 px-4 py-2"
          type="search"
          placeholder="Cari tujuan destinasi anda..."
          aria-label="Search"
          v-model="searchInput"
      />
      <div class="d-flex flex-row ms-4 align-items-center">
        <button
            @click="setTab('Wisata')"
            type="button"
            :class="`${useTabStore.state.type === 'Wisata' ? 'bg-color-primary' : 'bg-outline-primary'} rounded-4 px-3 me-2`"
            style="height: 36px"
        >
          Wisata
        </button>
        <button
            @click="setTab('Villa')"
            type="button"
            :class="`${useTabStore.state.type === 'Villa' ? 'bg-color-primary' : 'bg-outline-primary'} rounded-4 px-3`"
            style="height: 36px"
        >
          Villa
        </button>
      </div>
    </div>

    <!-- List -->
    <div class="row my-4">
      <div v-if="useTabStore.state.type === 'Wisata'" v-for="(wisata, index) in listData" :key="index" class="col-sm-6 col-md-4 col-lg-3 mb-4">
        <ProductCard :wisata="wisata" :villa="null" />
      </div>
      <div v-else v-for="(villa, index2) in listData" :key="index2" class="col-sm-6 col-md-4 col-lg-3 mb-4">
        <ProductCard :villa="villa"  :wisata="null" />
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>

</style>