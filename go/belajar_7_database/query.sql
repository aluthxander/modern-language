-- Active: 1754386007649@@localhost@3306@golang_db
DELETE FROM customer;

ALTER TABLE customer
ADD COLUMN email VARCHAR(100),
ADD COLUMN balance INT DEFAULT 0,
ADD COLUMN rating DOUBLE DEFAULT 0.0,
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN birth_date DATE,
ADD COLUMN married BOOLEAN DEFAULT FALSE;


INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES
('diman', 'Diman', 'diman@go.dev', 1500000, 2.5, '2000-11-11', true),
('budi', 'Budi', 'budi@go.dev', 2500000, 3, '2001-11-11', FALSE),
('lutfan', 'Lutfan', 'lutfan@go.dev', 1000000, 4.5, '1999-11-11', true);

INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES
('riyan', 'Riyan', NULL, 2500000, 1.5, NULL, true);


CREATE TABLE user(
    username VARCHAR(100) PRIMARY KEY NOT NULL,
    password VARCHAR(100) NOT NULL
) ENGINE=InnoDB;

SELECT * FROM `user`;

INSERT INTO `user`(username, password) VALUES('diman', 'diman');


CREATE TABLE comments (
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(100) NOT NULL,
    comment TEXT
) ENGINE=InnoDB;

DESC comments;

SELECT COUNT(*) FROM comments;