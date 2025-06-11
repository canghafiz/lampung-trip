import { defineStore } from "pinia";
import { reactive } from "vue";

export const searchStore = defineStore("search_store", () => {
  const state = reactive({
    result: [],
    input: "",
  });

  function setWisataResult(input, data) {
    const lowerCaseInput = input.toLowerCase();

    state.result = data.filter(
      (item) =>
        item.judul &&
        typeof item.judul === "string" &&
        item.judul.toLowerCase().includes(lowerCaseInput)
    );
  }

  function setVillaResult(input, data) {
    const lowerCaseInput = input.toLowerCase();

    state.result = data.filter(
      (item) =>
        item.judul &&
        typeof item.judul === "string" &&
        item.judul.toLowerCase().includes(lowerCaseInput)
    );
  }

  function setKaryawanResult(input, data) {
    const lowerCaseInput = input.toLowerCase();

    state.result = data.filter(
      (item) =>
        item.nama &&
        typeof item.nama === "string" &&
        item.nama.toLowerCase().includes(lowerCaseInput)
    );
  }

  function setInput(value) {
    state.input = value;
  }

  return {
    state,
    setWisataResult,
    setVillaResult,
    setKaryawanResult,
    setInput,
  };
});
