export function parseJwt(token) {
    const base64Url = token.split('.')[1]; // JWT token's payload part is at index 1
    const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/'); // Handle URL-safe base64 encoding
    const jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}