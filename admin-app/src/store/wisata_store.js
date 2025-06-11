import {defineStore} from "pinia";
import {reactive} from "vue";
import {WisataResponse} from "../models/response/wisata_response.js";
import {WisataUpdateRequest} from "../models/request/wisata_request.js";

export const wisataStore = defineStore("wisata_store", () => {
    const state = reactive({
        "listData": [],
        "loadData": false,
        "editData": new WisataUpdateRequest(),
    });

    function initListData(data) {
        state.listData = data;
    }

    function setLoadData(data) {
        state.loadData = data;
    }

    function setEditData(data) {
        state.editData = data;
    }

    function clearEditData() {
        state.editData = new WisataUpdateRequest();
    }

    return {
        state,
        initListData,
        setLoadData,
        setEditData,
        clearEditData
    }
})