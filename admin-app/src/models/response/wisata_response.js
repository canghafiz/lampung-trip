export class WisataResponse {
    constructor (id, judul, deskripsi, rating, gambar, fasilitas, info) {
        this.id = id;
        this.judul = judul;
        this.deskripsi = deskripsi;
        this.rating = rating;
        this.gambar = gambar; // This will be an array of Gambar instances
        this.fasilitas = fasilitas; // This will be an array of Fasilitas instances
        this.info = info; // This will be an Info instance
    }

    /**
     * Creates a WisataResponse instance from a JSON object.
     * This is a static factory method.
     * @param {object} json The JSON object representing the WisataResponse.
     * @returns {WisataResponse} A new WisataResponse instance.
     */
    static fromJson (json = {}) { // <--- Made static
        const id = json.id || null;
        const judul = json.judul || '';
        const deskripsi = json.deskripsi || '';
        const rating = json.rating || 0;

        // Safely parse nested objects/arrays using their STATIC fromArray/fromJson methods
        const gambar = (json.gambar && Array.isArray(json.gambar)) ? Gambar.fromArray(json.gambar) : [];
        const fasilitas = (json.fasilitas && Array.isArray(json.fasilitas)) ? Fasilitas.fromArray(json.fasilitas) : [];
        const info = (json.info) ? Info.fromJson(json.info) : null; // Defaults to null if info is missing

        return new WisataResponse(id, judul, deskripsi, rating, gambar, fasilitas, info);
    }

    /**
     * Creates an array of WisataResponse instances from an array of JSON objects.
     * This is a static factory method.
     * @param {Array<object>} data An array of JSON objects, each representing a WisataResponse.
     * @returns {Array<WisataResponse>} An array of WisataResponse instances.
     */
    static fromArray(data) {
        // Robust check to ensure data is an array
        if (!Array.isArray(data)) {
            console.warn("Input to WisataResponse.fromArray is not an array. Returning empty array.");
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

    /**
     * Converts the WisataResponse instance back into a plain JavaScript object.
     * Useful for serialization (e.g., sending to an API or storing).
     * @returns {object} A plain object representation of the WisataResponse.
     */
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

        /**
         * Creates a Gambar instance from a JSON object.
         * This is a static factory method.
         * @param {object} json The JSON object containing the 'url' property.
         * @returns {Gambar} A new Gambar instance.
         */
        static fromJson(json = {}) { // <--- Made static
            // Correctly defaults to an empty string if json.url is missing or falsy
            const url = json.url || '';
            return new Gambar(url);
        }

        /**
         * Creates an array of Gambar instances from an array of JSON objects.
         * This is a static factory method.
         * @param {Array<object>} data An array of JSON objects, each representing a Gambar.
         * @returns {Array<Gambar>} An array of Gambar instances.
         */
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

        /**
         * Converts the Gambar instance back into a plain JavaScript object.
         * @returns {object} A plain object representation.
         */
        toJson() {
            return { url: this.url };
        }
    }

    class Fasilitas {
        constructor(url_icon, judul) {
            this.url_icon = url_icon;
            this.judul = judul;
        }

        /**
         * Creates a Fasilitas instance from a JSON object.
         * This is a static factory method.
         * @param {object} json The JSON object containing 'url_icon' and 'judul'.
         * @returns {Fasilitas} A new Fasilitas instance.
         */
        static fromJson(json = {}) { // <--- Made static
            const url_icon = json.url_icon || '';
            const judul = json.judul || '';
            return new Fasilitas(url_icon, judul);
        }

        /**
         * Creates an array of Fasilitas instances from an array of JSON objects.
         * This is a static factory method.
         * @param {Array<object>} data An array of JSON objects.
         * @returns {Array<Fasilitas>} An array of Fasilitas instances.
         */
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

        /**
         * Converts the Fasilitas instance back into a plain JavaScript object.
         * @returns {object} A plain object representation.
         */
        toJson() {
            return {
                url_icon: this.url_icon,
                judul: this.judul
            };
        }
    }


    class Info {
        constructor(harga, waktuBuka, waktuTutup, hariBuka, hariTutup, urlLokasi) {
            this.harga = harga;
            this.waktu_buka = waktuBuka;
            this.waktu_tutup = waktuTutup;
            this.hari_buka = hariBuka;
            this.hari_tutup = hariTutup;
            this.url_lokasi = urlLokasi;
        }

        /**
         * Creates an Info instance from a JSON object.
         * This is a static factory method.
         * @param {object} json The JSON object containing info properties.
         * @returns {Info} A new Info instance.
         */
        static fromJson(json = {}) { // <--- Made static
            // Provide sensible defaults for each property
            const harga = json.harga || 0; // Assuming harga is a number
            const waktu_buka = json.waktu_buka || '';
            const waktu_tutup = json.waktu_tutup || '';
            const hari_buka = json.hari_buka || '';
            const hari_tutup = json.hari_tutup || '';
            const url_lokasi = json.url_lokasi || '';

            return new Info(harga, waktu_buka, waktu_tutup, hari_buka, hari_tutup, url_lokasi);
        }

        /**
         * Converts the Info instance back into a plain JavaScript object.
         * @returns {object} A plain object representation.
         */
        toJson() {
            return {
                harga: this.harga,
                waktu_buka: this.waktu_buka,
                waktu_tutup: this.waktu_tutup,
                hari_buka: this.hari_buka,
                hari_tutup: this.hari_tutup,
                url_lokasi: this.url_lokasi
            };
        }
    }