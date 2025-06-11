import {defineStore} from "pinia";
import {reactive} from "vue";

export const authStore = defineStore("active_user", () => {
    const user = reactive({
        userId: 0,
        role: ""
    });

    function setUser(id, role) {
        user.userId = id;
        user.role = role;
    }

    return {
        user,
        setUser,
    }
});