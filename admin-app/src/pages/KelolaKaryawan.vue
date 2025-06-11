<script setup>
import {onMounted, ref} from "vue";
import {KaryawanService} from "../services/karyawan_service.js";
import {KaryawanRequest, KaryawanUpdateRequest} from "../models/request/karyawan_request.js";
import {FileService} from "../services/file_service.js";
import {GambarRequest} from "../models/request/villa_request.js";
import {karyawanStore} from "../store/karyawan_store.js";
import {sideNavStore} from "../store/side_nav.js";

const useKaryawanStore = karyawanStore()
const useSideNavStore = sideNavStore()

const nama = ref(useKaryawanStore.state.editData?.nama ?? "")
const no = ref(useKaryawanStore.state.editData?.noHp ?? "")
const files = ref([])
const username = ref("")
const password = ref("")

const loading = ref(false)
const errMsg = ref("")

onMounted(() => {
  const photoInput = document.getElementById("photo");
  const photoError = document.getElementById("photo-error");
  const photoPreviewContainer = document.getElementById(
      "photo-preview-container"
  );
  let uploadedPhoto = null;

  photoInput.addEventListener("change", function () {
    const file = this.files[0];
    files.value.push(file);
    photoPreviewContainer.innerHTML = ""; // Clear previous preview

    if (this.files.length > 1) {
      photoError.style.display = "block";
      this.value = "";
      uploadedPhoto = null;
      return;
    } else {
      photoError.style.display = "none";
    }

    if (file) {
      const reader = new FileReader();
      reader.onload = function (e) {
        const previewImg = document.createElement("img");
        previewImg.src = e.target.result;
        previewImg.alt = "Preview Foto Karyawan";
        previewImg.style.maxWidth = "100px";
        previewImg.style.maxHeight = "100px";
        photoPreviewContainer.appendChild(previewImg);
        uploadedPhoto = file;
      };
      reader.readAsDataURL(file);
    } else {
      uploadedPhoto = null;
    }
  });
})

async function handleSubmit(event) {
  event.preventDefault()
  loading.value = true

  if (!useKaryawanStore.state.editData?.nama) {
    // Create
    await FileService.Upload(files.value).then(response => {
      const gambarData = response.data.map(img => new GambarRequest(img.FileURL))

      const request = new KaryawanRequest(username.value, password.value, nama.value, no.value, gambarData[0].url)

      KaryawanService.Create(request).then(response => {
        console.log("Success Create Karyawan")

        // Jika error
        if (response) {
          errMsg.value = response

          FileService.Delete(gambarData[0].url)
        } else {
          window.location.reload()
        }
      }).finally(() => {
        loading.value = false
      })
    })
  } else {
    // Update
    await FileService.Delete(useKaryawanStore.state.editData.photo).then(_ => {
      FileService.Upload(files.value).then(response => {
        const gambarData = response.data.map(img => new GambarRequest(img.FileURL))

        const request = new KaryawanUpdateRequest(useKaryawanStore.state.editData.userId, nama.value, no.value, gambarData[0].url)

        KaryawanService.Update(request).then(_ => {
          console.log("Success Update Karyawan")
        }).finally(() => {
          loading.value = false
          window.location.reload()
        })
      })
    })
  }
}

function tutupModalError() {
  errMsg.value = ""
}

function cancelSubmit() {
  useSideNavStore.setPage("Karyawan")
}
</script>

<template>
  <!-- Modal Hapus -->
  <div v-if="errMsg" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Hapus Data</h5>
        <button type="button" class="close-button" @click="tutupModalError">&times;</button>
      </div>
      <div class="modal-body">
        <p>{{errMsg}}</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="modal-button btn-ok" @click="tutupModalError">Batal</button>
      </div>
    </div>
  </div>

  <div class="card mb-4">
    <div class="card-body">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-lg-8">
            <form role="form" @submit="handleSubmit">
              <div class="mb-3">
                <label for="nama" class="form-label"
                >Nama Karyawan</label
                >
                <input
                    type="text"
                    class="form-control"
                    placeholder="Masukkan nama karyawan"
                    id="nama"
                    aria-label="Nama"
                    aria-describedby="nama-addon"
                    name="nama"
                    v-model="nama"
                    required
                />
              </div>
              <div class="mb-3">
                <label for="no_hp" class="form-label">Nomor HP</label>
                <input
                    type="text"
                    class="form-control"
                    placeholder="Masukkan nomor HP karyawan"
                    id="no_hp"
                    aria-label="Nomor HP"
                    aria-describedby="no_hp-addon"
                    name="no_hp"
                    v-model="no"
                    required
                />
              </div>

              <hr class="horizontal dark my-4" />
              <h6>Foto Karyawan (Maksimal 1)</h6>
              <div class="mb-3">
                <label for="photo" class="form-label"
                >Foto Karyawan</label
                >
                <input
                    type="file"
                    class="form-control"
                    id="photo"
                    aria-label="Foto"
                    aria-describedby="photo-addon"
                    name="photo"
                    accept="image/*"
                    required
                />
                <small class="text-muted"
                >Pilih maksimal 1 foto untuk karyawan ini.</small
                >
                <div
                    id="photo-error"
                    class="text-danger"
                    style="display: none"
                >
                  Anda hanya dapat memilih maksimal 1 foto.
                </div>
              </div>

              <div class="mb-3" id="photo-preview-container"></div>

              <hr v-if="!useKaryawanStore.state.editData?.nama" class="horizontal dark my-4" />
              <h6 v-if="!useKaryawanStore.state.editData?.nama">Informasi Akun</h6>
              <div v-if="!useKaryawanStore.state.editData?.nama" class="mb-3">
                <label for="username" class="form-label"
                >Username</label
                >
                <input
                    type="text"
                    class="form-control"
                    placeholder="Masukkan username"
                    id="username"
                    aria-label="Username"
                    aria-describedby="username-addon"
                    name="username"
                    v-model="username"
                    required
                />
              </div>
              <div  v-if="!useKaryawanStore.state.editData?.nama" class="mb-3">
                <label for="password" class="form-label"
                >Password</label
                >
                <input
                    type="password"
                    class="form-control"
                    placeholder="Masukkan password"
                    id="password"
                    aria-label="Password"
                    aria-describedby="password-addon"
                    name="password"
                    v-model="password"
                    required
                />
              </div>

              <div class="text-center">
                <button
                    type="submit"
                    class="btn bg-gradient-primary w-100 my-4 mb-2"
                >
                  Simpan Data Karyawan
                </button>
              </div>
              <div class="text-center">
                <button
                    @click="cancelSubmit"
                    class="btn bg-white w-100 my-2"
                >
                  Cancel
                </button>
              </div>
            </form>
          </div>
        </div>
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
  background-color: #344767;
  color: white;
}

.btn-ok:hover,
.btn-ok:focus {
  background-color: #0d1422;
  color: #ffffff;
}

/* Responsiveness */
@media (max-width: 600px) {
  .modal-content {
    width: 95%;
  }
}
</style>