import { apiEndPoint, apiKey } from "../config.js";
import { Response } from "../models/response/response.js";
import axios from "axios";
import { deleteCookie, setCookie } from "../utils/cookies.js";

export class UserService {
  static async Login(request) {
    try {
      const body = request.toJson();

      const json = (
        await axios.post(`${apiEndPoint}/user/login-root`, body, {
          headers: {
            Authorization: apiKey,
            "ngrok-skip-browser-warning": "true",
          },
        })
      ).data;

      const response = Response.fromJson(json);

      // Check if success
      if (response.code === 200) {
        // Save to cookie
        setCookie("token_lampung_trip_admin", response.data, 1);

        window.location.reload();
        return null;
      }

      return response;
    } catch (err) {
      const response = Response.fromJson(err["response"]["data"]);

      console.error("Login error:", response.data);

      let message = "Gagal login";
      if (response.data === "password wrong") {
        message = "Password salah";
      }

      if (response.data === "user not found") {
        message = "Pengguna tidak ditemukan";
      }

      return message;
    }
  }

  static async Password(request) {
    try {
      const body = request.toJson();

      const json = (
        await axios.put(`${apiEndPoint}/user/password`, body, {
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

      console.error("Login error:", response.data);

      let message = "Gagal login";
      if (response.data === "password wrong") {
        message = "Password salah";
      }

      return message;
    }
  }

  static async Logout() {
    // Cookie
    deleteCookie("token_lampung_trip_admin");

    // Reload Page
    window.location.reload();
  }
}
