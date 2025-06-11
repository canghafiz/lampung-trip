export class GambarRequest {
    constructor(url) {
        this.url = url;
    }

    static toArrayJson(gambarRequests) {
        if (!Array.isArray(gambarRequests)) {
            console.error("Input must be an array of GambarRequest objects.");
            return [];
        }
        return gambarRequests.map(request => ({
            url: request.url
        }));
    }
}