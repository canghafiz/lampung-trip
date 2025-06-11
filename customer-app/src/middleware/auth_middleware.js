import {getCookie} from "../utils/cookies.js";
import {userStore} from "../store/user_store.js";
import {parseJwt} from "../utils/jwt.js";
import {UserResponse} from "../model/response/user_response.js";
import {UserService} from "../service/user_service.js";

export async function authMiddleware() {
    const jwtToken = getCookie('token');
    if (jwtToken) {
        const decoded = parseJwt(jwtToken);

        // Set Store
        const useUserStore = userStore()

        useUserStore.setUser(decoded["user_id"], decoded["role"]);
        await UserService.GetData(decoded["user_id"]).then((response) => {
            if (response.code === 200) {
                const user = UserResponse.fromJSON(response.data)

                useUserStore.setBiodata(user)
            }
        })
    } else {
        console.log('JWT token not found');
    }
}