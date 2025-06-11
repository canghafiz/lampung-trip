import axios from "axios";
import { apiEndPoint, apiKey } from "../config.js";
import { Response } from "../models/response/response.js";

export class WisataService {
  static async Create(request) {
    try {
      const body = request.toJson();
      const json = (
        await axios.post(`${apiEndPoint}/wisata/`, body, {
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
      return "Gagal membuat wisata";
    }
  }

  static async Update(request) {
    try {
      const body = request.toJson();
      const json = (
        await axios.put(`${apiEndPoint}/wisata/`, body, {
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
      return "Gagal update wisata";
    }
  }

  static async GetAll() {
    try {
      const json = (
        await axios.get(`${apiEndPoint}/wisata/`, {
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
      return "Gagal mendapatkan wisata";
    }
  }

  static async Delete(id) {
    try {
      const json = (
        await axios.delete(`${apiEndPoint}/wisata/${id}`, {
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
      return "Gagal menghapus wisata";
    }
  }
}
