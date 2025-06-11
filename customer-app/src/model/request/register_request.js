export class RegisterRequest {
    constructor(username, password, nama, noHp) {
        this.username = username;
        this.password = password;
        this.nama = nama;
        this.noHp = noHp;
    }

    toJson() {
        return {
            username: this.username,
            password: this.password,
            customer: {
                nama: this.nama,
                photo: "",
                no_hp: this.noHp
            }
        }
    }
}