import {reactive} from "vue";
import {WisataUpdateRequest} from "../models/request/wisata_request.js";
import {defineStore} from "pinia";

export const villaStore = defineStore("villa_store", () => {
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