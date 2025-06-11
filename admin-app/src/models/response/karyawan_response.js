export class KaryawanResponse {
    constructor({ userId = "", username = "", nama = "", photo = "", noHp = "" } = {}) {
        this.userId = userId;
        this.username = username;
        this.nama = nama;
        this.photo = photo;
        this.noHp = noHp;
    }

    static fromJson(json = {}) {
        return new KaryawanResponse({
            userId: json.user_id,
            username: json.username,
            nama: json.nama,
            photo: json.photo,
            noHp: json.no_hp,
        });
    }

    static fromArrayJson(data = []) {
        if (!Array.isArray(data)) {
            console.warn("Input for fromArrayJson is not an array. Returning empty array.");
            return [];
        }
        return data.map(item => KaryawanResponse.fromJson(item));
    }
}