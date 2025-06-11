export class PasswordRequest {
    constructor(username, old_password, new_password) {
        this.username = username;
        this.old_password = old_password;
        this.new_password = new_password;
    }

    toJson() {
        return {
            username: this.username,
            old_password: this.old_password,
            new_password: this.new_password
        }
    }
}