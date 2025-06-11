import {defineStore} from "pinia";
import {reactive} from "vue";

export const sideNavStore = defineStore("side_nav", () => {
    const state = reactive({
        "page": "Wisata"
    });

    function setPage(value) {
        state.page = value;
    }

    return {
        state,
        setPage
    }
})