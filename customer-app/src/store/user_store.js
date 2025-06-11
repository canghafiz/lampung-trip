import {defineStore} from "pinia";
import {reactive} from "vue";
import {UserResponse} from "../model/response/user_response.js";

export const userStore = defineStore("active_user", () => {
    const user = reactive({
        userId: 0,
        role: "",
        biodata: new UserResponse()
    });

    function setUser(id, role) {
        user.userId = id;
        user.role = role;
    }

    function setBiodata(biodata) {
        user.biodata = biodata;
    }

    return {
        user,
        setUser,
        setBiodata,
    }
});