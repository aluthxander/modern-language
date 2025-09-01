-- Active: 1754623357006@@127.0.0.1@3306@siappak-shadow
SELECT * FROM category;

SELECT COUNT(*) AS jumlah_kolom
FROM information_schema.COLUMNS
    WHERE table_schema = 'siappak-shadow'
AND table_name   = 'wp_data_mahasiswa';