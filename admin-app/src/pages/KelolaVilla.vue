<script setup>
import {onMounted, ref} from "vue"
import {sideNavStore} from "../store/side_nav.js";
import {
  FasilitasRequest,
  GambarRequest,
  InfoRequest,
} from "../models/request/villa_request.js";
import {FileService} from "../services/file_service.js";
import {deleteMultipleGambar} from "../utils/function.js";
import {villaStore} from "../store/villa_store.js";
import {VillaCreateRequest} from "../models/request/villa_request.js";
import {VillaService} from "../services/villa_service.js";

// Store
const useSideNavStore = sideNavStore()
const useVillaStore = villaStore()

// State
const judul = ref(  useVillaStore.state.editData?.judul ?? "")
const deskripsi = ref(useVillaStore.state.editData?.deskripsi ?? "")
const fasilitas = ref([{icon: "🏊", title: "Kolam Renang"}]) // Inisialisasi dengan satu fasilitas awal
const harga = ref(useVillaStore.state.editData?.info?.harga ?? "")
const urlLokasi = ref(useVillaStore.state.editData?.info?.lokasi ?? "")
const imagesUpload = ref([])

const loading = ref(false)

onMounted(() => {
  if (useVillaStore.state.editData?.fasilitas) {
    fasilitas.value = []

    for (let data of useVillaStore.state.editData.fasilitas) {
      fasilitas.value.push({icon: data.url_icon, title: data.judul})
    }
  }

  const gambarInput = document.getElementById("gambar")
  const gambarError = document.getElementById("gambar-error")
  const form = document.getElementById("tambahVillaForm")
  const hargaInput = document.getElementById("harga")
  let uploadedImages = []

  function formatRupiah(angka) {
    const numberFormat = new Intl.NumberFormat("id-ID", {
      style: "currency",
      currency: "IDR",
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    })
    return numberFormat.format(angka)
  }

  hargaInput.addEventListener("input", function () {
    let value = this.value.replace(/\D/g, "")
    this.value = value ? formatRupiah(parseInt(value)) : ""
  })

  gambarInput.addEventListener("change", function () {
    const files = this.files
    if (files.length > 5) {
      gambarError.style.display = "block"
      this.value = ""
      return
    } else {
      gambarError.style.display = "none"
    }

    for (let i = 0; i < files.length; i++) {
      const file = files[i]
      if (uploadedImages.length < 5) {
        uploadedImages.push(file)
        imagesUpload.value.push(file)
      } else {
        alert("Anda hanya dapat memilih maksimal 5 gambar.")
        break
      }
    }
    this.value = "" // Clear input setelah preview
    updateGambarInput()
  })

  form.addEventListener("submit", function (event) {
    if (uploadedImages.length > 5) {
      event.preventDefault()
      gambarError.style.display = "block"
    } else if (uploadedImages.length === 0 && gambarInput.files.length === 0) {
      // Optional: Handle case where no images are selected
    } else {
      // formData.append("gambar[]", file) akan dihandle oleh Vue dalam handleSubmit
      // event.preventDefault() juga dihandle oleh Vue di handleSubmit
    }
  })

  function updateGambarInput() {
    const dataTransfer = new DataTransfer()
    uploadedImages.forEach((file) => {
      dataTransfer.items.add(file)
    })
    gambarInput.files = dataTransfer.files
  }
})

// Metode untuk menambah fasilitas
function addFasilitas() {
  fasilitas.value.push({icon: "🏊", title: "Kolam Renang"})
}

// Metode untuk menghapus fasilitas
function removeFasilitas(index) {
  fasilitas.value.splice(index, 1)
}

async function handleSubmit(event) {
  loading.value = true
  event.preventDefault()

  const fasilitasData = fasilitas.value.map(f => new FasilitasRequest(f.icon, f.title))
  await FileService.Upload(imagesUpload.value).then(response => {
    const gambarData = response.data.map(img => new GambarRequest(img.FileURL))

    if (!useVillaStore.state.editData?.judul) {
      // Create
      const request = new VillaCreateRequest(
          judul.value,
          deskripsi.value,
          gambarData,
          new InfoRequest(
              harga.value,
              urlLokasi.value
          ),
          fasilitasData
      )

      VillaService.Create(request).then(_ => {
        console.log("Success Create Villa")
      }).finally(() => {
        loading.value = false
        window.location.reload()
        useSideNavStore.setPage("Villa")
      })
    } else {
      // Update
      deleteMultipleGambar(useVillaStore.state.editData.gambar)

      const request = new VillaUpdateRequest(
          useVillaStore.state.editData.id,
          judul.value,
          deskripsi.value,
          gambarData,
          new InfoRequest(
              harga.value,
              urlLokasi.value
          ),
          fasilitasData
      )

      VillaService.Update(request).then(_ => {
        console.log("Success Update Villa")
      }).finally(() => {
        loading.value = false
        window.location.reload()
        useSideNavStore.setPage("Villa")
      })
    }
  })
}

function cancelSubmit() {
  useSideNavStore.setPage("Villa")
}
</script>

<template>
  <div class="card mb-4">
    <div class="card-body">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <form role="form" id="tambahVillaForm" @submit="handleSubmit">
              <div class="mb-3">
                <label for="judul" class="form-label"
                >Judul Villa</label
                >
                <input
                    type="text"
                    class="form-control"
                    placeholder="Masukkan judul villa"
                    id="judul"
                    aria-label="Judul"
                    aria-describedby="judul-addon"
                    name="judul"
                    v-model="judul"
                    required
                />
              </div>
              <div class="mb-3">
                <label for="deskripsi" class="form-label"
                >Deskripsi</label
                >
                <textarea
                    class="form-control"
                    placeholder="Masukkan deskripsi villa"
                    id="deskripsi"
                    rows="4"
                    aria-label="Deskripsi"
                    aria-describedby="deskripsi-addon"
                    name="deskripsi"
                    v-model="deskripsi"
                    required
                ></textarea>
              </div>

              <hr class="horizontal dark my-4"/>
              <h6>Informasi Villa</h6>
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
                    >Pilih maksimal 5 gambar untuk slide utama
                      villa ini.</small
                    >
                    <div
                        id="gambar-error"
                        class="text-danger"
                        style="display: none"
                    >
                      Anda hanya dapat memilih maksimal 5 gambar.
                    </div>
                  </div>
                  <div
                      class="mb-3"
                      id="gambar-preview-container"
                  ></div>
                </div>
                <div class="col-md-6">
                  <div class="mb-3">
                    <label for="harga" class="form-label"
                    >Harga</label
                    >
                    <div class="input-group">
                      <input
                          type="text"
                          class="form-control"
                          placeholder="Masukkan harga"
                          id="harga"
                          aria-label="Harga"
                          aria-describedby="harga-addon"
                          name="harga"
                          v-model="harga"
                          required
                      />
                    </div>
                    <small class="text-muted"
                    >Masukkan angka saja.</small
                    >
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

              <hr class="horizontal dark my-4"/>
              <h6>Fasilitas Villa</h6>
              <div id="fasilitas-container">
                <div
                    class="row mb-3 fasilitas-item"
                    v-for="(item, index) in fasilitas"
                    :key="index"
                >
                  <div class="col-md-4">
                    <label
                        for="url_icon_fasilitas_1"
                        class="form-label"
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
                      <option value="">Pilih Ikon</option>
                      <option value="🏊">🏊 Kolam Renang</option>
                      <option value="📶">📶 Wi-Fi</option>
                      <option value="📺">📺 TV</option>
                      <option value="❄️">❄️ AC</option>
                      <option value="🛁">🛁 Kamar Mandi</option>
                      <option value="🛏️">🛏️ Tempat Tidur</option>
                      <option value="🛋️">🛋️ Ruang Keluarga</option>
                      <option value="🍳">🍳 Dapur</option>
                      <option value="🅿️">🅿️ Parkir</option>
                      <option value="⛰️">
                        ⛰️ Pemandangan Gunung
                      </option>
                      <option value="🌊">🌊 Pemandangan Laut</option>
                      <option value="🌳">🌳 Taman</option>
                      <option value="🔥">🔥 Peralatan BBQ</option>
                      <option value="♿">
                        ♿ Aksesibilitas Difabel
                      </option>
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

              <hr class="horizontal dark my-4"/>
              <div class="text-center">
                <button
                    type="submit"
                    class="btn bg-gradient-primary w-100 my-4 mb-2"
                >
                  {{ loading ? 'Loading..' : 'Simpan Data Villa' }}
                </button>
              </div>
            </form>
            <div class="text-center">
              <button
                  @click="cancelSubmit"
                  class="btn bg-white w-100 my-2"
              >
                Cancel
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>