import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    // Tambahkan blok ini
    port: 3001, // Tentukan port yang diinginkan
  },
});
