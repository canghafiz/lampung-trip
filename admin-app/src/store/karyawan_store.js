import {defineStore} from "pinia";
import {reactive} from "vue";
import {KaryawanUpdateRequest} from "../models/request/karyawan_request.js";

export const karyawanStore = defineStore("karyawan_store", () => {
    const state = reactive({
        "listData": [],
        "loadData": false,
        "editData": new KaryawanUpdateRequest(),
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
        state.editData = new KaryawanUpdateRequest()
    }

    return {
        state,
        initListData,
        setLoadData,
        setEditData,
        clearEditData
    }
})