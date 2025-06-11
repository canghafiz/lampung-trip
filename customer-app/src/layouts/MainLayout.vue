<script setup>

import Footer from "../components/Footer.vue";
import Navbar from "../components/Navbar.vue";
</script>

<template>
  <div class="content-wrapper d-flex flex-column align-items-center">
    <div class="container-fluid container-sm px-md-5">
      <navbar/>
      <slot/>
    </div>
  </div>
  <Footer/>
</template>

<style scoped>

</style>