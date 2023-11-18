CREATE DATABASE IF NOT EXISTS edu;

CREATE TABLE IF NOT EXISTS students (
    username VARCHAR(20) NOT NULL,
    email VARCHAR(30),
    phone numeric(10),
    address VARCHAR(40),
    password varchar(2000),
    role VARCHAR(30),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS website (
    name VARCHAR(200) NOT NULL,
    domain VARCHAR(200),
     address VARCHAR(40),
     created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS admin (
    username VARCHAR(20) NOT NULL,
    password varchar(2000),
    email VARCHAR(30),
    id numeric(20),
    role varchar(30),
    created_at TIMESTAMP DEFAULT NOW()
);


INSERT INTO students (username, email, phone, address, password, role) VALUES
    ('MahtabKokabi', 'kokabi.mahtab@gmail.com', '921392820', 'Tehran', 'moonmkdnh', 'User'),
    ('AlirezaKokabi', 'alireza.kokabi@gmail.com', '9213902734', 'Tehran', 'alirz23', 'User');

INSERT INTO website (name, domain, address) VALUES
    ('Google', 'www.google.com', 'USA')

INSERT INTO admin (username, password, email, id, role) VALUES
    ('Mozhdhkokabi', '00232192', ' kokabi.mozhdeh@gmail.com', '34', 'Admin')
SELECT * FROM students;
