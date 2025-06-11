export class UlasanVillaRequest {
    constructor(villaId, userId, rating, keterangan) {
        this.villaId = villaId;
        this.userId = userId;
        this.rating = rating;
        this.keterangan = keterangan;
    }

    toJson() {
        return {
            villa_id: this.villaId,
            customer_id: this.userId,
            rating: this.rating,
            keterangan: this.keterangan,
        }
    }
}