export class PasswordRequest {
    constructor(username, oldPassword, newPassword) {
        this.username = username;
        this.oldPassword = oldPassword;
        this.newPassword = newPassword;
    }

    toJson() {
        return {
            username: this.username,
            old_password: this.oldPassword,
            new_password: this.newPassword,
        }
    }
}