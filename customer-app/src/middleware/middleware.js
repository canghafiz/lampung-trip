import {authMiddleware} from "./auth_middleware.js";
import {userStore} from "../store/user_store.js";

export function middlewareSetToLogin(to, next, currentPath) {
    authMiddleware();

    const store = userStore();

    // Check if authenticated or not
    if (store.user.userId !== 0) {
        // Authenticated and check if the path is on current path or not
        if (to.path !== currentPath) {
            // If not, make it to current path
            next(currentPath);
        } else {
            // Process Normally
            next();
        }

    } else {
        // Not authenticated set to login page
        next("/masuk");
    }
}

export function middlewareTackleAuth(to, next) {
    authMiddleware()

    const store = userStore();

    if (store.user.userId !== 0) {
        if (to.path !== "/") {
            next("/");
        } else {
            next();
        }
    } else {
        // If not authenticated,
        next();
    }
}