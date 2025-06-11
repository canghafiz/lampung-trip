-- Indeks untuk tabel Wisata
-- Indeks pada 'judul' untuk pencarian atau filter berdasarkan judul wisata
CREATE INDEX idx_wisata_judul ON Wisata(judul(255));
-- Indeks pada 'rating' untuk filter atau pengurutan berdasarkan rating
CREATE INDEX idx_wisata_rating ON Wisata(rating);
-- Indeks pada 'created_at' untuk pengurutan atau filter berdasarkan waktu pembuatan
CREATE INDEX idx_wisata_created_at ON Wisata(created_at);
-- Indeks pada 'updated_at' untuk pengurutan atau filter berdasarkan waktu pembaruan
CREATE INDEX idx_wisata_updated_at ON Wisata(updated_at);

-- Indeks untuk tabel GambarWisata
-- Indeks pada 'wisata_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_gambarwisata_wisata_id ON GambarWisata(wisata_id);

-- Indeks untuk tabel FasilitasWisata
-- Indeks pada 'wisata_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_fasilitaswisata_wisata_id ON FasilitasWisata(wisata_id);

-- Indeks untuk tabel InfoWisata
-- Indeks pada 'wisata_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_infowisata_wisata_id ON InfoWisata(wisata_id);

-- Indeks untuk tabel UlasanWisata
-- Primary Key (wisata_id, customer_id) sudah otomatis terindeks.
-- Indeks tambahan pada 'customer_id' jika sering ada query yang mencari ulasan hanya berdasarkan customer_id
CREATE INDEX idx_ulasanwisata_customer_id ON UlasanWisata(customer_id);
-- Indeks pada 'rating' untuk filter atau pengurutan ulasan berdasarkan rating
CREATE INDEX idx_ulasanwisata_rating ON UlasanWisata(rating);
-- Indeks pada 'created_at' untuk pengurutan ulasan berdasarkan waktu pembuatan
CREATE INDEX idx_ulasanwisata_created_at ON UlasanWisata(created_at);


-- Indeks untuk tabel Villa
-- Indeks pada 'judul' untuk pencarian atau filter berdasarkan judul villa
CREATE INDEX idx_villa_judul ON Villa(judul(255));
-- Indeks pada 'rating' untuk filter atau pengurutan berdasarkan rating
CREATE INDEX idx_villa_rating ON Villa(rating);
-- Indeks pada 'created_at' untuk pengurutan atau filter berdasarkan waktu pembuatan
CREATE INDEX idx_villa_created_at ON Villa(created_at);
-- Indeks pada 'updated_at' untuk pengurutan atau filter berdasarkan waktu pembaruan
CREATE INDEX idx_villa_updated_at ON Villa(updated_at);

-- Indeks untuk tabel GambarVilla
-- Indeks pada 'villa_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_gambarvilla_villa_id ON GambarVilla(villa_id);

-- Indeks untuk tabel FasilitasVilla
-- Indeks pada 'villa_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_fasilitasvilla_villa_id ON FasilitasVilla(villa_id);

-- Indeks untuk tabel InfoVilla
-- Indeks pada 'villa_id' karena ini adalah Foreign Key dan akan sering digunakan dalam JOIN
CREATE INDEX idx_infovilla_villa_id ON InfoVilla(villa_id);

-- Indeks untuk tabel UlasanVilla
-- Primary Key (villa_id, customer_id) sudah otomatis terindeks.
-- Indeks tambahan pada 'customer_id' jika sering ada query yang mencari ulasan hanya berdasarkan customer_id
CREATE INDEX idx_ulasanvilla_customer_id ON UlasanVilla(customer_id);
-- Indeks pada 'rating' untuk filter atau pengurutan ulasan berdasarkan rating
CREATE INDEX idx_ulasanvilla_rating ON UlasanVilla(rating);
-- Indeks pada 'created_at' untuk pengurutan ulasan berdasarkan waktu pembuatan
CREATE INDEX idx_ulasanvilla_created_at ON UlasanVilla(created_at);