import {defineStore} from "pinia";
import {reactive} from "vue";
import {WisataService} from "../service/wisata_service.js";
import {WisataResponse} from "../model/response/wisata_response.js";

export const wisataStore = defineStore("wisata_store", () => {
    const state = reactive({
        "listData": [],
        "loadData": false,
        "singleData": null,
    });

    async function initListData() {
        if (state.listData.length === 0) {
            setLoadData(true)
            await WisataService.GetAll().then((response) => {
                state.listData = WisataResponse.fromArray(response.data)

                console.log("Successfully getting list data wisata")
            }).finally(() => {
                setLoadData(false)
            })
        }
    }

    function setLoadData(data) {
        state.loadData = data;
    }

    async function setSingleData(id, slug) {
        const idToFind = id.toString();

        await initListData();

        state.singleData = state.listData.find(item => {
            return item.id.toString() === idToFind;
        });

        if (state.singleData) {
            console.log("Item ditemukan dan disimpan ke singleData:", state.singleData);
        } else {
            console.log(`Item dengan ID '${id}' dan slug '${slug}' tidak ditemukan.`);
            state.singleData = null
        }
    }

    return {
        state,
        initListData,
        setLoadData,
        setSingleData,
    }
})