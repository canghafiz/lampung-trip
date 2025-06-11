<script setup>
import { onMounted, ref } from "vue";
import {
  FasilitasRequest,
  GambarRequest,
  InfoRequest,
  WisataCreateRequest,
  WisataUpdateRequest,
} from "../models/request/wisata_request.js";
import { FileService } from "../services/file_service.js";
import { sideNavStore } from "../store/side_nav.js";
import { WisataService } from "../services/wisata_service.js";
import { wisataStore } from "../store/wisata_store.js";
import { deleteMultipleGambar } from "../utils/function.js";

// Store
const useSideNavStore = sideNavStore();
const useWisataStore = wisataStore();

// State
const judul = ref(useWisataStore.state.editData.judul ?? "");
const deskripsi = ref(useWisataStore.state.editData.deskripsi);
const fasilitas = ref([{ icon: "🚽", title: "" }]); // Inisialisasi dengan satu fasilitas awal
const hargaTiket = ref(useWisataStore.state.editData.info?.harga ?? "");
const waktuBuka = ref(useWisataStore.state.editData.info?.waktu_buka ?? "");
const waktuTutup = ref(useWisataStore.state.editData.info?.waktu_tutup ?? "");
const hariBuka = ref(useWisataStore.state.editData.info?.hari_buka ?? "");
const hariTutup = ref(useWisataStore.state.editData.info?.hari_tutup ?? "");
const urlLokasi = ref(useWisataStore.state.editData.info?.url_lokasi ?? "");
const imagesUpload = ref([]);

const loading = ref(false);

onMounted(() => {
  if (useWisataStore.state.editData?.fasilitas) {
    fasilitas.value = [];

    for (let data of useWisataStore.state.editData.fasilitas) {
      fasilitas.value.push({ icon: data.url_icon, title: data.judul });
    }
  }

  const gambarInput = document.getElementById("gambar");
  const gambarError = document.getElementById("gambar-error");
  const form = document.getElementById("tambahWisataForm");
  const hargaInput = document.getElementById("harga");
  let uploadedImages = [];

  function formatRupiah(angka) {
    const numberFormat = new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    });
    return numberFormat.format(angka);
  }

  hargaInput.addEventListener("input", function () {
    let value = this.value.replace(/\D/g, "");
    this.value = value ? formatRupiah(parseInt(value)) : "";
  });

  gambarInput.addEventListener("change", function () {
    const files = this.files;
    if (files.length > 5) {
      gambarError.style.display = "block";
      this.value = "";
      return;
    } else {
      gambarError.style.display = "none";
    }

    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      if (uploadedImages.length < 5) {
        uploadedImages.push(file);
        imagesUpload.value.push(file);
      } else {
        alert("Anda hanya dapat memilih maksimal 5 gambar.");
        break;
      }
    }
    this.value = ""; // Clear input setelah preview
    updateGambarInput();
  });

  form.addEventListener("submit", function (event) {
    if (uploadedImages.length > 5) {
      event.preventDefault();
      gambarError.style.display = "block";
    } else if (uploadedImages.length === 0 && gambarInput.files.length === 0) {
      // Optional: Handle case where no images are selected
    } else {
      // formData.append("gambar[]", file) akan dihandle oleh Vue dalam handleSubmit
      // event.preventDefault() juga dihandle oleh Vue di handleSubmit
    }
  });

  function updateGambarInput() {
    const dataTransfer = new DataTransfer();
    uploadedImages.forEach((file) => {
      dataTransfer.items.add(file);
    });
    gambarInput.files = dataTransfer.files;
  }
});

// Metode untuk menambah fasilitas
function addFasilitas() {
  fasilitas.value.push({ icon: "🚽", title: "" });
}

// Metode untuk menghapus fasilitas
function removeFasilitas(index) {
  fasilitas.value.splice(index, 1);
}

async function handleSubmit(event) {
  loading.value = true;
  event.preventDefault();

  const fasilitasData = fasilitas.value.map(
    (f) => new FasilitasRequest(f.icon, f.title)
  );
  await FileService.Upload(imagesUpload.value).then((response) => {
    const gambarData = response.data.map(
      (img) => new GambarRequest(img.FileURL)
    );

    if (!useWisataStore.state.editData?.judul) {
      // Create
      const request = new WisataCreateRequest(
        judul.value,
        deskripsi.value,
        gambarData,
        new InfoRequest(
          hargaTiket.value,
          waktuBuka.value,
          waktuTutup.value,
          hariBuka.value,
          hariTutup.value,
          urlLokasi.value
        ),
        fasilitasData
      );

      WisataService.Create(request)
        .then((_) => {
          console.log("Success Create Wisata");
        })
        .finally(() => {
          loading.value = false;
          window.location.reload();
          useSideNavStore.setPage("Wisata");
        });
    } else {
      // Update
      deleteMultipleGambar(useWisataStore.state.editData.gambar);

      const request = new WisataUpdateRequest(
        useWisataStore.state.editData.id,
        judul.value,
        deskripsi.value,
        gambarData,
        new InfoRequest(
          hargaTiket.value,
          waktuBuka.value,
          waktuTutup.value,
          hariBuka.value,
          hariTutup.value,
          urlLokasi.value
        ),
        fasilitasData
      );

      WisataService.Update(request)
        .then((_) => {
          console.log("Success Update Wisata");
        })
        .finally(() => {
          loading.value = false;
          window.location.reload();
          useSideNavStore.setPage("Wisata");
        });
    }
  });
}

function cancelSubmit() {
  useSideNavStore.setPage("Wisata");
}
</script>

<template>
  <div class="card mb-4">
    <div class="card-body">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <form role="form" id="tambahWisataForm" @submit="handleSubmit">
              <div class="mb-3">
                <label for="judul" class="form-label">Judul Wisata</label>
                <input
                  type="text"
                  class="form-control"
                  placeholder="Masukkan judul wisata"
                  id="judul"
                  aria-label="Judul"
                  aria-describedby="judul-addon"
                  name="judul"
                  v-model="judul"
                  required
                />
              </div>
              <div class="mb-3">
                <label for="deskripsi" class="form-label">Deskripsi</label>
                <textarea
                  class="form-control"
                  placeholder="Masukkan deskripsi wisata"
                  id="deskripsi"
                  rows="4"
                  aria-label="Deskripsi"
                  aria-describedby="deskripsi-addon"
                  name="deskripsi"
                  v-model="deskripsi"
                  required
                ></textarea>
              </div>

              <hr class="horizontal dark my-4" />
              <h6>Informasi Wisata</h6>
              <div class="row">
                <div class="col-md-6">
                  <div class="mb-3">
                    <label ref="gambarLabel" for="gambar" class="form-label">
                      Gambar Utama (Maksimal 5)</label
                    >
                    <input
                      type="file"
                      class="form-control"
                      id="gambar"
                      multiple
                      aria-label="Gambar"
                      aria-describedby="gambar-addon"
                      name="gambar[]"
                      accept="image/*"
                      required
                    />
                    <small class="text-muted"
                      >Pilih maksimal 5 gambar untuk slide utama wisata
                      ini.</small
                    >
                    <div
                      id="gambar-error"
                      class="text-danger"
                      style="display: none"
                    >
                      Anda hanya dapat memilih maksimal 5 gambar.
                    </div>
                  </div>
                  <div class="mb-3" id="gambar-preview-container"></div>
                </div>
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="harga" class="form-label">Harga Tiket</label>
                    <div class="input-group">
                      <input
                        type="text"
                        class="form-control"
                        placeholder="Masukkan harga tiket"
                        id="harga"
                        aria-label="Harga Tiket"
                        aria-describedby="harga-addon"
                        name="harga"
                        v-model="hargaTiket"
                        required
                      />
                    </div>
                    <small class="text-muted">Masukkan angka saja.</small>
                  </div>
                </div>
              </div>

              <div class="row">
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="waktu_buka" class="form-label"
                      >Waktu Buka</label
                    >
                    <input
                      type="time"
                      class="form-control"
                      id="waktu_buka"
                      aria-label="Waktu Buka"
                      aria-describedby="waktu_buka-addon"
                      name="waktu_buka"
                      v-model="waktuBuka"
                      required
                    />
                  </div>
                </div>
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="waktu_tutup" class="form-label"
                      >Waktu Tutup</label
                    >
                    <input
                      type="time"
                      class="form-control"
                      id="waktu_tutup"
                      aria-label="Waktu Tutup"
                      aria-describedby="waktu_tutup-addon"
                      name="waktu_tutup"
                      v-model="waktuTutup"
                      required
                    />
                  </div>
                </div>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="hari_buka" class="form-label">Hari Buka</label>
                    <select
                      class="form-control"
                      id="hari_buka"
                      aria-label="Hari Buka"
                      aria-describedby="hari_buka-addon"
                      name="hari_buka"
                      v-model="hariBuka"
                      required
                    >
                      <option value="">Pilih Hari Buka</option>
                      <option value="Senin">Senin</option>
                      <option value="Selasa">Selasa</option>
                      <option value="Rabu">Rabu</option>
                      <option value="Kamis">Kamis</option>
                      <option value="Jumat">Jumat</option>
                      <option value="Sabtu">Sabtu</option>
                      <option value="Minggu">Minggu</option>
                      <option value="Setiap Hari">Setiap Hari</option>
                      <option value="Sesuai Jadwal">Sesuai Jadwal</option>
                    </select>
                  </div>
                </div>
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="hari_tutup" class="form-label"
                      >Hari Tutup</label
                    >
                    <select
                      class="form-control"
                      id="hari_tutup"
                      aria-label="Hari Tutup"
                      aria-describedby="hari_tutup-addon"
                      name="hari_tutup"
                      v-model="hariTutup"
                      required
                    >
                      <option value="">Pilih Hari Tutup</option>
                      <option value="Senin">Senin</option>
                      <option value="Selasa">Selasa</option>
                      <option value="Rabu">Rabu</option>
                      <option value="Kamis">Kamis</option>
                      <option value="Jumat">Jumat</option>
                      <option value="Sabtu">Sabtu</option>
                      <option value="Minggu">Minggu</option>
                      <option value="Libur Nasional">Libur Nasional</option>
                      <option value="Sesuai Jadwal">Sesuai Jadwal</option>
                    </select>
                  </div>
                </div>
              </div>
              <div class="mb-3">
                <label for="url_lokasi" class="form-label"
                  >URL Lokasi (Google Maps)</label
                >
                <input
                  type="url"
                  class="form-control"
                  placeholder="Masukkan URL Google Maps"
                  id="url_lokasi"
                  aria-label="URL Lokasi"
                  aria-describedby="url_lokasi-addon"
                  name="url_lokasi"
                  v-model="urlLokasi"
                  required
                />
              </div>

              <hr class="horizontal dark my-4" />
              <h6>Fasilitas Wisata</h6>
              <div id="fasilitas-container">
                <div
                  class="row mb-3 fasilitas-item"
                  v-for="(item, index) in fasilitas"
                  :key="index"
                >
                  <div class="col-md-4">
                    <label :for="'ikon_fasilitas_' + index" class="form-label"
                      >Ikon Fasilitas</label
                    >
                    <select
                      class="form-control"
                      :id="'ikon_fasilitas_' + index"
                      :aria-label="'Ikon Fasilitas ' + index"
                      :aria-describedby="'ikon_fasilitas_' + index + '-addon'"
                      :name="'ikon_fasilitas[' + index + ']'"
                      v-model="item.icon"
                      required
                    >
                      <option value="🚽">🚽 Toilet Umum</option>
                      <option value="🚿">🚿 Kamar Bilas</option>
                      <option value="🗑️">🗑️ Tempat Sampah</option>
                      <option value="🅿️">🅿️ Area Parkir</option>
                      <option value="⛱️">⛱️ Gazebo / Tempat Berte-duh</option>
                      <option value="🪑">🪑 Bangku dan Meja Umum</option>
                      <option value="👮">
                        👮 Petugas Keamanan / Lifeguard
                      </option>
                      <option value="⚕️">⚕️ Pos Kesehatan / P3K</option>
                      <option value="♿">♿ Aksesibilitas untuk Difabel</option>
                    </select>
                  </div>
                  <div class="col-md-6">
                    <label :for="'judul_fasilitas_' + index" class="form-label"
                      >Judul Fasilitas</label
                    >
                    <input
                      type="text"
                      class="form-control"
                      placeholder="Masukkan nama fasilitas"
                      :id="'judul_fasilitas_' + index"
                      :aria-label="'Judul Fasilitas ' + index"
                      :aria-describedby="'judul_fasilitas_' + index + '-addon'"
                      :name="'judul_fasilitas[' + index + ']'"
                      v-model="item.title"
                      required
                    />
                  </div>
                  <div class="col-md-2 d-flex align-items-end">
                    <button
                      type="button"
                      class="btn btn-outline-danger btn-sm mb-0 remove-fasilitas"
                      @click="removeFasilitas(index)"
                    >
                      Hapus
                    </button>
                  </div>
                </div>
              </div>
              <button
                type="button"
                class="btn btn-outline-primary btn-sm mb-0"
                id="add-fasilitas"
                @click="addFasilitas"
              >
                Tambah Fasilitas Lain
              </button>

              <hr class="horizontal dark my-4" />
              <div class="text-center">
                <button
                  type="submit"
                  class="btn bg-gradient-primary w-100 my-4 mb-2"
                >
                  {{ loading ? "Loading.." : "Simpan Data Wisata" }}
                </button>
              </div>
            </form>
            <div class="text-center">
              <button @click="cancelSubmit" class="btn bg-white w-100 my-2">
                Cancel
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
