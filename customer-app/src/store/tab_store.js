import {defineStore} from "pinia";
import {reactive} from "vue";

export const tabStore = defineStore("tab_store", () => {
    const state = reactive({
        type: 'Wisata',
    });

    function setType(type) {
        state.type = type;
    }

    return {
        state,
        setType,
    }
})