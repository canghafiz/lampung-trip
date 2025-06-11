import {authMiddleware} from "./auth_middleware.js";
import {authStore} from "../store/auth_store.js";

export function middlewareSetToLogin(to, next, currentPath) {
    authMiddleware();

    const store = authStore();

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
        next("/login");
    }
}

export function middlewareTackleAuth(to, next) {
    authMiddleware()

    const store = authStore();

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