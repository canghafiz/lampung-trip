export function generateSlug(text) {
    if (typeof text !== 'string' || text.trim() === '') {
        // Mengembalikan string kosong jika input bukan string valid atau kosong
        return '';
    }

    // 1. Ubah ke huruf kecil
    let slug = text.toLowerCase();

    // 2. Ganti karakter non-alfanumerik (kecuali spasi dan tanda hubung) dengan spasi
    //    Contoh: "Hello, World!" -> "hello world"
    slug = slug.replace(/[^a-z0-9 -]/g, '');

    // 3. Ganti spasi dengan tanda hubung
    //    Contoh: "hello world" -> "hello-world"
    slug = slug.replace(/\s+/g, '-');

    // 4. Hapus tanda hubung ganda atau di awal/akhir string
    //    Contoh: "hello--world" -> "hello-world"
    //    Contoh: "-hello-world-" -> "hello-world"
    slug = slug.replace(/-+/g, '-'); // Hapus tanda hubung ganda
    slug = slug.replace(/^-/, '');    // Hapus tanda hubung di awal
    slug = slug.replace(/-$/, '');    // Hapus tanda hubung di akhir

    return slug;
}

export function slugToText(slug) {
    // Validasi input: pastikan itu string dan tidak kosong (setelah trim)
    if (typeof slug !== 'string' || slug.trim() === '') {
        return '';
    }

    // 1. Ganti semua tanda hubung (-) dengan spasi
    let text = slug.replace(/-/g, ' ');

    // 2. Kapitalkan huruf pertama setiap kata
    //    Contoh: "hello world" -> "Hello World"
    text = text.replace(/\b\w/g, char => char.toUpperCase());

    return text;
}