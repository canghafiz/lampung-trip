import { getCookie } from "../utils/cookies.js";
import { authStore } from "../store/auth_store.js";
import { parseJwt } from "../utils/jwt.js";

export function authMiddleware() {
  const jwtToken = getCookie("token_lampung_trip_admin");
  if (jwtToken) {
    const decoded = parseJwt(jwtToken);

    // Set Store
    const store = authStore();

    store.setUser(decoded["user_id"], decoded["role"]);
  } else {
    console.log("JWT token not found");
  }
}
