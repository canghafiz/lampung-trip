-- Table: User
CREATE TABLE User (
                      id INT(10) PRIMARY KEY AUTO_INCREMENT,
                      username TEXT UNIQUE,
                      password TEXT,
                      role ENUM('Admin', 'Karyawan', 'Customer'),
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: Karyawan
CREATE TABLE Karyawan (
                          id INT(10) PRIMARY KEY AUTO_INCREMENT,
                          user_id INT(10),
                          nama TEXT UNIQUE,
                          photo TEXT,
                          no_hp TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Table: Customer
CREATE TABLE Customer (
                          id INT(10) PRIMARY KEY AUTO_INCREMENT,
                          user_id INT(10),
                          nama TEXT UNIQUE,
                          photo TEXT,
                          no_hp TEXT,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES User(id)
);

-- Table: Wisata
CREATE TABLE Wisata (
                        id INT(10) PRIMARY KEY AUTO_INCREMENT,
                        judul TEXT,
                        deskripsi TEXT,
                        rating FLOAT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Table: GambarWisata
CREATE TABLE GambarWisata (
                              id INT(10) PRIMARY KEY AUTO_INCREMENT,
                              wisata_id INT(10),
                              url TEXT,
                              FOREIGN KEY (wisata_id) REFERENCES Wisata(id)
);

-- Table: FasilitasWisata
CREATE TABLE FasilitasWisata (
                                 id INT(10) PRIMARY KEY AUTO_INCREMENT,
                                 wisata_id INT(10),
                                 url_icon TEXT,
                                 judul TEXT,
                                 FOREIGN KEY (wisata_id) REFERENCES Wisata(id)
);

-- Table: InfoWisata
CREATE TABLE InfoWisata (
                            id INT(10) PRIMARY KEY AUTO_INCREMENT,
                            wisata_id INT(10),
                            harga TEXT,
                            waktu_buka TEXT,
                            waktu_tutup TEXT,
                            hari_buka TEXT,
                            hari_tutup TEXT,
                            url_lokasi TEXT,
                            FOREIGN KEY (wisata_id) REFERENCES Wisata(id)
);

-- Table: UlasanWisata
CREATE TABLE UlasanWisata (
                              wisata_id INT(10),
                              customer_id INT(10),
                              rating FLOAT,
                              keterangan TEXT,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              PRIMARY KEY (wisata_id, customer_id),
                              FOREIGN KEY (wisata_id) REFERENCES Wisata(id),
                              FOREIGN KEY (customer_id) REFERENCES Customer(id)
);

-- Table: Villa
CREATE TABLE Villa (
                       id INT(10) PRIMARY KEY AUTO_INCREMENT,
                       judul TEXT,
                       deskripsi TEXT,
                       rating FLOAT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Table: GambarVilla
CREATE TABLE GambarVilla (
                             id INT(10) PRIMARY KEY AUTO_INCREMENT,
                             villa_id INT(10),
                             url TEXT,
                             FOREIGN KEY (villa_id) REFERENCES Villa(id)
);

-- Table: FasilitasVilla
CREATE TABLE FasilitasVilla (
                                id INT(10) PRIMARY KEY AUTO_INCREMENT,
                                villa_id INT(10),
                                url_icon TEXT,
                                judul TEXT,
                                FOREIGN KEY (villa_id) REFERENCES Villa(id)
);

-- Table: InfoVilla
CREATE TABLE InfoVilla (
                           id INT(10) PRIMARY KEY AUTO_INCREMENT,
                           villa_id INT(10),
                           harga TEXT,
                           lokasi text,
                           FOREIGN KEY (villa_id) REFERENCES Villa(id)
);

-- Table: UlasanVilla
CREATE TABLE UlasanVilla (
                             villa_id INT(10),
                             customer_id INT(10),
                             rating FLOAT,
                             keterangan TEXT,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             PRIMARY KEY (villa_id, customer_id),
                             FOREIGN KEY (villa_id) REFERENCES Villa(id),
                             FOREIGN KEY (customer_id) REFERENCES Customer(id)
);