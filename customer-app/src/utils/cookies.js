export function setCookie(name, value, hours = 1, path = '/', secure = false, samesite = 'Strict') {
    const expires = new Date();
    expires.setTime(expires.getTime() + (hours * 60 * 60 * 1000));

    let cookieString = `${name}=${value}; expires=${expires.toUTCString()}; path=${path};`;

    if (secure) {
        cookieString += ' secure;';
    }

    cookieString += ` samesite=${samesite};`;

    document.cookie = cookieString;
}

export function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

export function deleteCookie(name) {
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
}