import axios from "axios";
import {apiEndPoint, apiKey} from "../config.js";

export class FileService {
    static async Upload(files) {
        try {
            const formData = new FormData();
            for (const file of files) {
                formData.append('files', file);
            }
            const response = await axios.post(`${apiEndPoint}/file/upload`, formData, {
                headers: {
                    Authorization: apiKey,
                    'Content-Type': 'multipart/form-data',
                }
            });

            return response.data;
        } catch (err) {
            console.error(err);
            return "Gagal upload file";
        }
    }

    static async Delete(url) {
        try {
            await axios.delete(`${apiEndPoint}/file/delete?url=${url}`, {
                headers: {
                    Authorization: apiKey,
                }
            });

            return null;
        } catch (err) {
            console.error(err);
            return "Gagal upload file";
        }
    }
}