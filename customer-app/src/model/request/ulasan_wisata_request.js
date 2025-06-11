export class UlasanWisataRequest {
    constructor(wisataId, userId, rating, keterangan) {
        this.wisataId = wisataId;
        this.userId = userId;
        this.rating = rating;
        this.keterangan = keterangan;
    }

    toJson() {
        return {
            wisata_id: this.wisataId,
            customer_id: this.userId,
            rating: this.rating,
            keterangan: this.keterangan,
        }
    }
}