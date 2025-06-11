export class UserResponse {
    constructor(id, userId, username, nama, photo, noHp) {
        this.id = id;
        this.userId = userId;
        this.username = username;
        this.nama = nama;
        this.photo = photo;
        this.noHp = noHp;
    }

    static fromJSON(json= {}) {
        const id = json.id;
        const userId = json.user_id;
        const username = json.username;
        const nama = json.nama;
        const photo = json.photo;
        const noHp = json.no_hp;

        return new UserResponse(id, userId, username, nama, photo, noHp, noHp);
    }
}