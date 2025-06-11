export class ApiResponse {
    constructor(code, status, data) {
        this.code = code;
        this.status = status;
        this.data = data;
    }

    static fromJson(json = {}) {
        const code = json.code;
        const status = json.status;
        const data = json.data;

        return new ApiResponse(code, status, data);
    }
}