export class VillaResponse {
    constructor (id, judul, deskripsi, rating, gambar, fasilitas, info, ulasan) {
        this.id = id;
        this.judul = judul;
        this.deskripsi = deskripsi;
        this.rating = rating;
        this.gambar = gambar; // This will be an array of Gambar instances
        this.fasilitas = fasilitas; // This will be an array of Fasilitas instances
        this.info = info; // This will be an Info instance
        this.ulasan = ulasan;
    }

    static fromJson (json = {}) { // <--- Made static
        const id = json.id || null;
        const judul = json.judul || '';
        const deskripsi = json.deskripsi || '';
        const rating = json.rating || 0;

        // Safely parse nested objects/arrays using their STATIC fromArray/fromJson methods
        const gambar = (json.gambar && Array.isArray(json.gambar)) ? Gambar.fromArray(json.gambar) : [];
        const fasilitas = (json.fasilitas && Array.isArray(json.fasilitas)) ? Fasilitas.fromArray(json.fasilitas) : [];
        const info = (json.info) ? Info.fromJson(json.info) : null; // Defaults to null if info is missing
        const ulasan = (json.ulasan && Array.isArray(json.ulasan)) ? Ulasan.fromArray(json.ulasan) : [];

        return new VillaResponse(id, judul, deskripsi, rating, gambar, fasilitas, info, ulasan);
    }

    static fromArray(data) {
        // Robust check to ensure data is an array
        if (!Array.isArray(data)) {
            console.warn("Input to VillaResponse.fromArray is not an array. Returning empty array.");
            return [];
        }

        let newData = [];
        for (let i = 0; i < data.length; i++) {
            // Correctly call the static fromJson method using 'this' (which refers to the class itself)
            newData.push(this.fromJson(data[i]));
        }
        // console.log(newData); // Consider removing this console.log in production code
        return newData;
    }

    toJson() {
        return {
            id: this.id,
            judul: this.judul,
            deskripsi: this.deskripsi,
            rating: this.rating,
            gambar: this.gambar.map(g => g.toJson()), // Assuming Gambar has a toJson method
            fasilitas: this.fasilitas.map(f => f.toJson()), // Assuming Fasilitas has a toJson method
            info: this.info ? this.info.toJson() : null // Assuming Info has a toJson method
        };
    }
}
class Gambar {
    constructor(url) {
        this.url = url
    }

    static fromJson(json = {}) {
        const url = json.url || '';
        return new Gambar(url);
    }

    static fromArray(data) {
        if (!Array.isArray(data)) {
            console.warn("Input to Gambar.fromArray is not an array. Returning empty array.");
            return [];
        }

        let newData = [];
        for (let i = 0; i < data.length; i++) {
            // Correctly call the static fromJson method using 'this'
            newData.push(this.fromJson(data[i]));
        }
        return newData;
    }

    toJson() {
        return { url: this.url };
    }
}

class Fasilitas {
    constructor(url_icon, judul) {
        this.url_icon = url_icon;
        this.judul = judul;
    }

    static fromJson(json = {}) { // <--- Made static
        const url_icon = json.url_icon || '';
        const judul = json.judul || '';
        return new Fasilitas(url_icon, judul);
    }

    static fromArray(data) {
        if (!Array.isArray(data)) {
            console.warn("Input to Fasilitas.fromArray is not an array. Returning empty array.");
            return [];
        }

        let newData = [];
        for (let i = 0; i < data.length; i++) {
            // Correctly call the static fromJson method using 'this'
            newData.push(this.fromJson(data[i]));
        }
        return newData;
    }

    toJson() {
        return {
            url_icon: this.url_icon,
            judul: this.judul
        };
    }
}


class Info {
    constructor(harga, lokasi) {
        this.harga = harga;
        this.lokasi = lokasi;
    }

    static fromJson(json = {}) { // <--- Made static
        // Provide sensible defaults for each property
        const harga = json.harga || 0; // Assuming harga is a number
        const lokasi = json.lokasi || "";

        return new Info(harga, lokasi);
    }

    toJson() {
        return {
            harga: this.harga,
            lokasi: this.lokasi,
        };
    }
}

class Ulasan {
    constructor(userId, nama, photo, rating, keterangan) {
        this.userId = userId;
        this.nama = nama;
        this.photo = photo;
        this.rating = rating;
        this.keterangan = keterangan;
    }

    static fromJson(json = {}) {
        const userId = json.customer.user_id || null;
        const nama = json.customer.nama || '';
        const photo = json.customer.photo || '';
        const rating = json.rating || '';
        const keterangan = json.keterangan || '';

        return new Ulasan(userId, nama, photo, rating, keterangan);
    }

    static fromArray(data) {
        if (!Array.isArray(data)) {
            console.warn("Input to Fasilitas.fromArray is not an array. Returning empty array.");
            return [];
        }

        let newData = [];
        for (let i = 0; i < data.length; i++) {
            // Correctly call the static fromJson method using 'this'
            newData.push(this.fromJson(data[i]));
        }
        return newData;
    }
}