-- On Table Ulasan Wisata
CREATE TRIGGER update_rating_wisata_when_create
    AFTER INSERT ON ulasanwisata
    FOR EACH ROW
BEGIN
    UPDATE wisata
    SET wisata.rating = (SELECT AVG(rating) FROM ulasanwisata WHERE wisata_id = NEW.wisata_id)
    WHERE id = NEW.wisata_id;
END;

CREATE TRIGGER update_rating_wisata_when_delete
    AFTER DELETE ON ulasanwisata
    FOR EACH ROW
BEGIN
    UPDATE wisata
    SET rating = IFNULL((SELECT AVG(rating) FROM ulasanwisata WHERE wisata_id = OLD.wisata_id), 0)
    WHERE id = OLD.wisata_id;
END;

-- On Table Ulasan Villa
CREATE TRIGGER update_rating_villa_when_create
    AFTER INSERT ON ulasanvilla
    FOR EACH ROW
BEGIN
    UPDATE villa
    SET villa.rating = (SELECT AVG(rating) FROM ulasanvilla WHERE villa_id = NEW.villa_id)
    WHERE id = NEW.villa_id;
END;

CREATE TRIGGER update_rating_villa_when_delete
    AFTER DELETE ON ulasanvilla
    FOR EACH ROW
BEGIN
    UPDATE villa
    SET rating = IFNULL((SELECT AVG(rating) FROM ulasanvilla WHERE villa_id = OLD.villa_id), 0)
    WHERE id = OLD.villa_id;
END;