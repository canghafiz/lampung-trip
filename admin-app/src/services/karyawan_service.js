import axios from "axios";
import { apiEndPoint, apiKey } from "../config.js";
import { Response } from "../models/response/response.js";

export class KaryawanService {
  static async Create(request) {
    const body = request.toJson();

    try {
      const json = (
        await axios.post(`${apiEndPoint}/user/karyawan/register`, body, {
          headers: {
            Authorization: apiKey,
            "ngrok-skip-browser-warning": "true",
          },
        })
      ).data;

      const response = Response.fromJson(json);

      // Check if success
      if (response.code === 200) {
        return null;
      }

      return response;
    } catch (err) {
      const response = Response.fromJson(err["response"]["data"]);

      let message = "Gagal";
      if (response.data === "user already exists") {
        message = "Karyawan sudah ada";
      }

      return message;
    }
  }

  static async Update(request) {
    try {
      const body = request.toJson();
      const json = (
        await axios.put(`${apiEndPoint}/karyawan/`, body, {
          headers: {
            Authorization: apiKey,
            "ngrok-skip-browser-warning": "true",
          },
        })
      ).data;

      const response = Response.fromJson(json);

      // Check if success
      if (response.code === 200) {
        return null;
      }

      return response;
    } catch (err) {
      return "Gagal update karyawan";
    }
  }

  static async GetAll() {
    try {
      const json = (
        await axios.get(`${apiEndPoint}/karyawan/`, {
          headers: {
            Authorization: apiKey,
            "ngrok-skip-browser-warning": "true",
          },
        })
      ).data;

      const response = Response.fromJson(json);

      // Check if success
      if (response.code === 200) {
        return response;
      }

      return response;
    } catch (err) {
      return "Gagal mendapatkan karyawan";
    }
  }

  static async Delete(userId) {
    try {
      const json = (
        await axios.delete(`${apiEndPoint}/user/karyawan/${userId}`, {
          headers: {
            Authorization: apiKey,
            "ngrok-skip-browser-warning": "true",
          },
        })
      ).data;

      const response = Response.fromJson(json);

      // Check if success
      if (response.code === 200) {
        window.location.reload();
        return null;
      }

      return response;
    } catch (err) {
      return "Gagal menghapus karyawan";
    }
  }
}
