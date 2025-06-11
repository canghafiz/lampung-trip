import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import {createPinia} from "pinia";
import {createRouter, createWebHistory} from "vue-router";
import LoginPage from "./pages/LoginPage.vue";
import RegisterPage from "./pages/RegisterPage.vue";
import PasswordPage from "./pages/PasswordPage.vue";
import {middlewareSetToLogin, middlewareTackleAuth} from "./middleware/middleware.js";
import MainPage from "./pages/MainPage.vue";
import DetailWisataPage from "./pages/DetailWisataPage.vue";
import DetailVillaPage from "./pages/DetailVillaPage.vue";
import NotFoundPage from "./pages/NotFoundPage.vue";
import UlasanWisataPage from "./pages/UlasanWisataPage.vue";
import UlasanVillaPage from "./pages/UlasanVillaPage.vue";
import ProfilePage from "./pages/ProfilePage.vue";

const pinia = createPinia();

const routes = [
    {
        path: "/masuk",
        component: LoginPage,
        name: "masuk",
        beforeEnter: (to, from, next) => {
            middlewareTackleAuth(to, next)
        }
    },
    {
        path: "/daftar",
        component: RegisterPage,
        name: "daftar",
        beforeEnter: (to, from, next) => {
            middlewareTackleAuth(to, next)
        }
    },
    {
        path: "/ubah-sandi",
        component: PasswordPage,
        name: "ubahSandi",
        beforeEnter: (to, from, next) => {
            middlewareTackleAuth(to, next)
        }
    },
    {
        path: "/",
        component: MainPage,
        name: "main",
    },
    {
        path: "/wisata/:id/:slug",
        component: DetailWisataPage,
        name: "detailWisata",
    },
    {
        path: "/wisata/:id/:slug/ulasan",
        component: UlasanWisataPage,
        name: "ulasanWisata",
    },
    {
        path: "/villa/:id/:slug",
        component: DetailVillaPage,
        name: "detailVilla",
    },
    {
        path: "/villa/:id/:slug/ulasan",
        component: UlasanVillaPage,
        name: "ulasanVilla",
    },
    {
        path: "/profil",
        component: ProfilePage,
        name: "profilPage",
        beforeEnter: (to, from, next) => {
            middlewareSetToLogin(to, next, "/profil")
        }
    },
    {
        path: "/:pathMatch(.*)*",
        component: NotFoundPage,
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
});

createApp(App)
    .use(pinia)
    .use(router)
    .mount('#app');