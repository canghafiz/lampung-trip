<script setup>

import {karyawanStore} from "../store/karyawan_store.js";
import {computed, onMounted, ref} from "vue";
import {KaryawanService} from "../services/karyawan_service.js";
import TableData from "../components/Loading/TableData.vue";
import {KaryawanResponse} from "../models/response/karyawan_response.js";
import {sideNavStore} from "../store/side_nav.js";
import {FileService} from "../services/file_service.js";
import {searchStore} from "../store/search_store.js";

const useSideNavStore = sideNavStore()
const useKaryawanStore = karyawanStore()
const useSearchStore = searchStore()

onMounted(async () => {
  if (useKaryawanStore.state.listData.length === 0) {
    useKaryawanStore.setLoadData(true)
    await KaryawanService.GetAll().then((response) => {
      const result = KaryawanResponse.fromArrayJson(response.data)
      console.log(response)

      // Update Store
      useKaryawanStore.initListData(result)

      console.log("Successfully getting list data karyawan")
    }).finally(() => {
      useKaryawanStore.setLoadData(false)
    })
  }
})

const tampilModalHapus = ref(false);
const hapusId = ref("")
const gambar = ref("")

const tampilkanModalHapus = (response) => {
  tampilModalHapus.value = true;
  hapusId.value = response.userId;
  gambar.value = response.photo;

  console.log(response)
};

const editKaryawan = async (karyawan) => {
  useKaryawanStore.setEditData(karyawan)
  useSideNavStore.setPage("Kelola Karyawan")

  console.log(karyawan)
}

const sembunyikanModalHapus = () => {
  tampilModalHapus.value = false;
  hapusId.value = "";
};

const hapusModal = async () => {
  await FileService.Delete(gambar.value)
  await KaryawanService.Delete(hapusId.value)
};

const dataList = computed(() => {
  if (useSearchStore.state.input !== "") {
    return useSearchStore.state.result
  }

  return useKaryawanStore.state.listData
});
</script>

<template>
  <!-- Modal Hapus -->
  <div v-if="tampilModalHapus" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Hapus Data</h5>
        <button type="button" class="close-button" @click="sembunyikanModalHapus">&times;</button>
      </div>
      <div class="modal-body">
        <p>Apakah Anda yakin ingin menghapus data ini?</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="modal-button btn-batal" @click="sembunyikanModalHapus">Batal</button>
        <button type="button" class="modal-button btn-yakin" @click="hapusModal">Yakin</button>
      </div>
    </div>
  </div>

  <TableData v-if="useKaryawanStore.state.loadData" />
  <div v-if="!useKaryawanStore.state.loadData" class="card mb-4">
    <div class="card-header pb-0">
      <h6>Data Karyawan</h6>
    </div>
    <div class="card-body px-0 pt-0 pb-2">
      <div class="table-responsive p-0">
        <table class="table align-items-center mb-0">
          <thead>
          <tr>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7"
            >
              Foto
            </th>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7 ps-2"
            >
              Nama
            </th>
            <th
                class="text-uppercase text-secondary text-xxs font-weight-bolder opacity-7 ps-2"
            >
              No. HP
            </th>
            <th class="text-secondary opacity-7"></th>
          </tr>
          </thead>
          <tbody>
          <tr v-for="karyawan in dataList" :key="karyawan.id">
            <td>
              <div class="d-flex px-2 py-1">
                <div>
                  <img
                      :src="`${karyawan.photo}`"
                      class="avatar avatar-sm me-3 border-radius-lg"
                      :alt="`Foto ${karyawan.nama}`"
                  />
                </div>
              </div>
            </td>
            <td>
              <p class="text-xs font-weight-bold mb-0">
                {{karyawan.nama}}
              </p>
            </td>
            <td>
              <p class="text-xs text-secondary mb-0">
                {{karyawan.noHp}}
              </p>
            </td>
            <td class="align-middle">
              <a
                  @click="editKaryawan(karyawan)"
                  href="javascript:"
                  class="btn btn-link text-dark px-3 mb-0"
                  data-toggle="tooltip"
                  data-original-title="Edit Karyawan"
              >
                <i
                    class="fas fa-pencil-alt text-sm me-2"
                    aria-hidden="true"
                ></i
                >Edit
              </a>
              <a
                  @click="tampilkanModalHapus(karyawan)"
                  href="javascript:"
                  class="btn btn-link text-danger px-3 mb-0"
                  data-toggle="tooltip"
                  data-original-title="Hapus Karyawan"
              >
                <i
                    class="far fa-trash-alt text-sm me-2"
                    aria-hidden="true"
                ></i
                >Hapus
              </a>
            </td>
          </tr>
          </tbody>
        </table>
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