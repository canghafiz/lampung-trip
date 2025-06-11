export class VillaCreateRequest {
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

export class VillaUpdateRequest {
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
    constructor(harga, lokasi) {
        this.harga = harga;
        this.lokasi = lokasi;
    }

    toJson() {
        return {
            harga: this.harga,
            lokasi: this.lokasi
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