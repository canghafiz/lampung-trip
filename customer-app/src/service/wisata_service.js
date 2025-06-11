import axios from "axios";
import {apiEndPoint, apiKey} from "../config.js";
import {ApiResponse} from "../model/response/api_response.js";

export class WisataService {
    static async GetAll() {
        try {
            const json = (await axios.get(`${apiEndPoint}/wisata/`,
                {
                    headers: {
                        Authorization: apiKey,
                    }
                },
            )).data;

            const response = ApiResponse.fromJson(json);

            // Check if success
            if (response.code === 200) {

                return response;
            }

            return response;
        } catch (err) {
            return "Gagal mendapatkan wisata"
        }
    }

    static async CreateUlasan(request) {
        try {
            const body = request.toJson();

            const json = (await axios.post(`${apiEndPoint}/ulasan-wisata/`,
                body,
                {
                    headers: {
                        'Content-Type': 'application/json', // Pastikan header Content-Type ada
                        'Authorization': apiKey, // Menggunakan quotes untuk key string
                    }
                },
            )).data

            const response = ApiResponse.fromJson(json);

            if (response.code === 200) {
                return null;
            }

            return null
        } catch (error) {
            return "Gagal membuat ulasan";
        }
    }
}