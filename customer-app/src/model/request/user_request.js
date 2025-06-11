export class UpdateDataUserRequest {
    constructor(id, nama, photo, noHp) {
        this.id = id;
        this.nama = nama;
        this.photo = photo;
        this.noHp = noHp;
    }

    toJson() {
        return {
            id: this.id,
            nama: this.nama,
            photo: this.photo,
            no_hp: this.noHp,
        }
    }
}