import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from "pinia";
import { createRouter, createWebHistory } from "vue-router";

import LoginPage from "./pages/LoginPage.vue";
import PasswordPage from "./pages/PasswordPage.vue";
import MainPage from "./pages/MainPage.vue";
import {middlewareSetToLogin, middlewareTackleAuth} from "./middleware/middleware.js";

// Create a Pinia instance for state management
const pinia = createPinia();

// Define your application routes
const routes = [
    {
        path: "/login",
        component: LoginPage,
        name: "login",
        beforeEnter: (to, from, next) => {
            middlewareTackleAuth(to, next)
        }
    },
    {
        path: "/password",
        component: PasswordPage,
        name: "password",
        beforeEnter: (to, from, next) => {
            middlewareTackleAuth(to, next)
        }
    },
    {
        path: "/",
        component: MainPage,
        name: "main",
        beforeEnter: (to, from, next) => {
            middlewareSetToLogin(to, next)
        }
    }
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

createApp(App)
    .use(pinia)
    .use(router)
    .mount('#app');