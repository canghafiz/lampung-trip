<script setup>
import {ref, onMounted, watch} from 'vue'; // Import ref dan onMounted
import MainLayout from '../layouts/MainLayout.vue';
import { userStore } from '../store/user_store.js';
import {UserService} from "../service/user_service.js";
import {UpdateDataUserRequest} from "../model/request/user_request.js";
import {FileService} from "../service/file_service.js";
import {GambarRequest} from "../model/request/gambar_request.js";

// Store
const useUserStore = userStore();

onMounted(() => {
  document.title = 'Profil';
});

// State reaktif
const isEditing = ref(false);
const profile = ref({
  photo: 'https://placehold.co/400', // Default atau URL foto profil
  name:  'Nama Pengguna',
  phoneNumber: '081234567890',
});
const tempProfile = ref({}); // Digunakan untuk menyimpan perubahan sementara saat editing
const tempPhoto = ref([]); // Digunakan untuk pratinjau foto yang baru diunggah

watch(useUserStore.user, (newVal) => {
  if(useUserStore.user.biodata.photo.length > 0) {
    profile.value.photo = useUserStore.user.biodata.photo
  }
  profile.value.name = newVal.biodata.nama;
  profile.value.phoneNumber = newVal.biodata.noHp;
})

// Metode
const startEditing = () => {
  isEditing.value = true;
  // Salin data profil saat ini ke tempProfile untuk diedit
  tempProfile.value = { ...profile.value };
  tempPhoto.value = null; // Reset pratinjau foto saat memulai edit
};

const handlePhotoUpload = (event) => {
  const file = event.target.files[0];
  if (file) {
    const reader = new FileReader();
    reader.onload = (e) => {
      tempPhoto.value = e.target.result; // Tampilkan pratinjau foto yang baru diunggah
      // Simpan data base64 ke tempProfile.value.photo sebagai array
      tempProfile.value.photo = [file];
    };
    reader.readAsDataURL(file);
  }
};

const saveProfile = async () => {
  // Delete File
  if (useUserStore.user.biodata.photo.length > 0) {
    await FileService.Delete(useUserStore.user.biodata.photo);
  }

  console.log(tempProfile.value.photo)

  await FileService.Upload(tempProfile.value.photo).then((response) => {
    const gambarData = response.data.map(img => new GambarRequest(img.FileURL))

    const request = new UpdateDataUserRequest(useUserStore.user.userId, tempProfile.value.name, gambarData[0].url, tempProfile.value.phoneNumber)

    UserService.UpdateData(request).then((response) => {
      if (response === null) {
        // Simpan perubahan dari tempProfile ke profile
        profile.value = { ...tempProfile.value };
        isEditing.value = false;
        alert('Profil berhasil diperbarui!');
        window.location.reload();
      }
    })
  })
};

const cancelEditing = () => {
  isEditing.value = false;
  tempProfile.value = {}; // Bersihkan tempProfile
  tempPhoto.value = null; // Bersihkan pratinjau foto
};
</script>

<template>
  <MainLayout>
    <div class="profile-container">
      <h1>Profil Pengguna</h1>

      <div v-if="!isEditing" class="profile-preview">
        <div class="profile-photo-wrapper">
          <img :src="profile.photo" alt="Foto Profil" class="profile-photo" />
        </div>
        <p><strong>Nama:</strong> {{ profile.name }}</p>
        <p><strong>No. HP:</strong> {{ profile.phoneNumber }}</p>
        <button class="bg-color-primary" @click="startEditing">Edit Profil</button>
      </div>

      <div v-else class="profile-edit">
        <h2>Edit Profil</h2>
        <form @submit.prevent="saveProfile">
          <div class="form-group">
            <label for="photo">Foto Profil:</label>
            <input type="file" id="photo" @change="handlePhotoUpload" accept="image/*" required/>
            <img v-if="tempPhoto" :src="tempPhoto" alt="Pratinjau Foto Baru" class="profile-photo-preview" />
          </div>
          <div class="form-group">
            <label for="name">Nama:</label>
            <input type="text" id="name" v-model="tempProfile.name" required/>
          </div>
          <div class="form-group">
            <label for="phoneNumber">No. HP:</label>
            <input type="text" id="phoneNumber" v-model="tempProfile.phoneNumber" required/>
          </div>
          <div class="actions">
            <button type="submit">Simpan</button>
            <button @click="cancelEditing">Batal</button>
          </div>
        </form>
      </div>
    </div>
  </MainLayout>
</template>

<style scoped>
.profile-container {
  max-width: 600px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  font-family: Arial, sans-serif;
}

h1,
h2 {
  text-align: center;
  color: #333;
}

/* Preview Profile Styles */
.profile-preview {
  text-align: center;
  margin-top: 20px;
}

.profile-photo-wrapper {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  overflow: hidden;
  margin: 0 auto 20px;
  border: 2px solid #eee;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.profile-photo {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.profile-preview p {
  margin: 10px 0;
  font-size: 1.1em;
  color: #555;
}

.profile-preview strong {
  color: #333;
}

/* Edit Profile Styles */
.profile-edit {
  margin-top: 20px;
}

.form-group {
  margin-bottom: 15px;
}

.form-group label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #333;
}

.form-group input[type='text'],
.form-group input[type='file'] {
  width: calc(100% - 22px);
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1em;
}

.profile-photo-preview {
  display: block;
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  margin-top: 10px;
  border: 1px dashed #ccc;
}

.actions {
  text-align: right;
  margin-top: 20px;
}

button {
  padding: 10px 20px;
  margin-left: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1em;
  transition: background-color 0.3s ease;
}

button:first-child {
  background-color: #103d47;
  color: white;
}

button:first-child:hover {
  background-color: #0b2c34; /* Warna sedikit lebih gelap saat hover */
}

/* Perbaikan untuk styling tombol terakhir */
button:last-child {
  border: 1px solid #103d47;
  background-color: transparent;
  color: #103d47;
}

button:last-child:hover {
  background-color: #103d47;
  color: white;
}

/* Styling untuk tombol di bagian preview profile */
.profile-preview button {
  background-color: #103d47;
  color: white;
  margin-top: 15px;
}

.profile-preview button:hover {
  background-color: #0b2c34; /* Warna sedikit lebih gelap saat hover */
}
</style>