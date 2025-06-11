import {FileService} from "../services/file_service.js";

export async function deleteMultipleGambar(gambars) {
    for (let gambar of gambars) {
        await FileService.Delete(gambar.url)
    }
}