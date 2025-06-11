export class KaryawanRequest {
    constructor(username, password, nama, nomor, photo) {
        this.username = username;
        this.password = password;
        this.nama = nama;
        this.nomor = nomor;
        this.photo = photo;
    }

    toJson() {
        return {
            username: this.username,
            password: this.password,
            karyawan: {
                nama: this.nama,
                no_hp: this.nomor,
                photo: this.photo,
            }
        }
    }
}

export class KaryawanUpdateRequest {
    constructor(id, nama, nomor, photo) {
        this.id = id;
        this.nama = nama;
        this.nomor = nomor;
        this.photo = photo;
    }

    toJson() {
        return {
            id: this.id,
            nama: this.nama,
            photo: this.photo,
            no_hp: this.nomor,
        }
    }
}