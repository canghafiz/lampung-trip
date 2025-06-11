export class WisataCreateRequest {
    constructor(judul, deskripsi, gambar, info, fasilitas) {
        this.judul = judul;
        this.deskripsi = deskripsi;
        this.gambar = gambar;
        this.info = info;
        this.fasilitas = fasilitas;
    }

    toJson() {
        return {
            judul: this.judul,
            deskripsi: this.deskripsi,
            gambar: GambarRequest.toArrayJson(this.gambar),
            info: this.info.toJson(),
            fasilitas: FasilitasRequest.toArrayJson(this.fasilitas),
        }
    }
}

export class WisataUpdateRequest {
    constructor(id, judul, deskripsi, gambar, info, fasilitas) {
        this.id = id;
        this.judul = judul;
        this.deskripsi = deskripsi;
        this.gambar = gambar;
        this.info = info;
        this.fasilitas = fasilitas;
    }

    toJson() {
        return {
            id: this.id,
            judul: this.judul,
            deskripsi: this.deskripsi,
            gambar: GambarRequest.toArrayJson(this.gambar),
            info: this.info.toJson(),
            fasilitas: FasilitasRequest.toArrayJson(this.fasilitas),
        }
    }
}

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

export class InfoRequest {
    constructor(harga, waktu_buka, waktu_tutup, hari_buka, hari_tutup, url_lokasi) {
        this.harga = harga;
        this.waktu_buka = waktu_buka;
        this.waktu_tutup = waktu_tutup;
        this.url_lokasi = url_lokasi;
        this.hari_buka = hari_buka;
        this.hari_tutup = hari_tutup;
    }

    toJson() {
        return {
            harga: this.harga,
            waktu_buka: this.waktu_buka,
            waktu_tutup: this.waktu_tutup,
            url_lokasi: this.url_lokasi,
            hari_buka: this.hari_buka,
            hari_tutup: this.hari_tutup
        }
    }
}

export class FasilitasRequest {
    constructor(url_icon, judul) {
        this.url_icon = url_icon;
        this.judul = judul;
    }

    static toArrayJson(infoRequests) {
        if (!Array.isArray(infoRequests)) {
            console.error("Input must be an array of GambarRequest objects.");
            return [];
        }
        return infoRequests.map(request => ({
            url_icon: request.url_icon,
            judul: request.judul,
        }));
    }
}