CREATE TABLE IF NOT EXISTS contacts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_contacts_name ON contacts (name);

INSERT INTO
    contacts (name, email, phone)
VALUES (
        'test',
        'test@mail.com',
        '083131314545'
    ),
    (
        'Aditya Pratama',
        'aditya.pratama@email.com',
        '081234567801'
    ),
    (
        'Bunga Lestari',
        'bunga.lestari@email.com',
        '081234567802'
    ),
    (
        'Cahyo Nugroho',
        'cahyo.nugroho@email.com',
        '081234567803'
    ),
    (
        'Dinda Sari',
        'dinda.sari@email.com',
        '081234567804'
    ),
    (
        'Eko Santoso',
        'eko.santoso@email.com',
        '081234567805'
    ),
    (
        'Fitri Ayu',
        'fitri.ayu@email.com',
        '081234567806'
    ),
    (
        'Galih Saputra',
        'galih.saputra@email.com',
        '081234567807'
    ),
    (
        'Hesti Anindya',
        'hesti.anindya@email.com',
        '081234567808'
    ),
    (
        'Irfan Maulana',
        'irfan.maulana@email.com',
        '081234567809'
    ),
    (
        'Joko Wibisono',
        'joko.wibisono@email.com',
        '081234567810'
    );