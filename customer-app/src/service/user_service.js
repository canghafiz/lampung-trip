import {apiEndPoint, apiKey} from "../config.js";
import axios from "axios";
import {ApiResponse} from "../model/response/api_response.js";
import {deleteCookie, setCookie} from "../utils/cookies.js";
import {userStore} from "../store/user_store.js";

export class UserService {
    static async Register(request) {
        try {
            const body = request.toJson();

            const json = (await axios.post(`${apiEndPoint}/user/customer/register`,
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
            const response = ApiResponse.fromJson(error["response"]["data"]);

            console.error("Daftar error:", response.data);

            let message = "Gagal mendaftar"
            if (response.data === 'user already exists') {
                message = "Pengguna sudah ada"
            }

            return message;
        }
    }
    static async Login(request) {
        try {
            const body = request.toJson();

            const json = (await axios.post(`${apiEndPoint}/user/login`,
                body,
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': apiKey,
                    }
                },
            )).data

            const response = ApiResponse.fromJson(json);

            if (response.code === 200) {
                // Save to cookie
                setCookie("token", response.data, 1);

                // Set Store
                const useUserStore = userStore()
            }

            return null
        } catch (error) {
            const response = ApiResponse.fromJson(error["response"]["data"]);

            console.error("Daftar error:", response.data);

            let message = "Gagal mendaftar"
            if (response.data === 'user not found') {
                message = "Pengguna tidak ditemukan"
            }
            if (response.data === 'password wrong') {
                message = "Kata sandi salah"
            }

            return message;
        }
    }
    static async Password(request) {
        try {
            const body = request.toJson();

            const json = (await axios.put(`${apiEndPoint}/user/password`,
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
            const response = ApiResponse.fromJson(error["response"]["data"]);

            console.error("Password error:", response.data);

            let message = "Gagal ganti pw"
            if (response.data === 'user not found') {
                message = "Pengguna tidak ditemukan"
            }
            if (response.data === 'password wrong') {
                message = "Kata sandi salah"
            }

            return message;
        }
    }
    static async GetData(userId) {
        try {
            const json = (await axios.get(`${apiEndPoint}/customer/${userId}`,
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': apiKey,
                    }
                },
            )).data;

            console.log(json);

            const response = ApiResponse.fromJson(json);

            // Check if success
            if (response.code === 200) {

                return response;
            }

            return response;
        } catch (err) {
            return "Gagal mendapatkan user"
        }
    }
    static async UpdateData(request) {
        try {
            const body = request.toJson();

            const json = (await axios.put(`${apiEndPoint}/customer/`,
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
            return "Gagal update data user";
        }
    }
    static Logout() {
        // Cookie
        deleteCookie("token");

        // Reload Page
        window.location.reload();
    }
}