import {defineStore} from "pinia";
import {reactive} from "vue";
import {VillaService} from "../service/villa_service.js";
import {VillaResponse} from "../model/response/villa_response.js";
import {slugToText} from "../utils/slug.js";

export const villaStore = defineStore("villa_store", () => {
    const state = reactive({
        "listData": [],
        "loadData": false,
        "singleData": null,
    });

    async function initListData() {
        if (state.listData.length === 0) {
            setLoadData(true)
            await VillaService.GetAll().then((response) => {
                state.listData = VillaResponse.fromArray(response.data)

                console.log("Successfully getting list data villa")
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
            return item.id.toString() === idToFind
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