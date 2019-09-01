DROP TABLE IF EXISTS article;

CREATE TABLE article
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(160)  NOT NULL,
    content VARCHAR(2000) NOT NULL,
    author VARCHAR(160) NOT NULL
);


INSERT INTO article (title,content,author) VALUES ('AC/DC','test','test');
INSERT INTO article (title,content,author) VALUES ('Accept','test','test');
INSERT INTO article (title,content,author) VALUES ('Aerosmith','test','test');
INSERT INTO article (title,content,author) VALUES ('Alanis Morissette','test','test');
INSERT INTO article (title,content,author) VALUES ('Alice In Chains','test','test');
INSERT INTO article (title,content,author) VALUES ('Apocalyptica','test','test');
INSERT INTO article (title,content,author) VALUES ('Audioslave','test','test');
INSERT INTO article (title,content,author) VALUES ('BackBeat','test','test');
INSERT INTO article (title,content,author) VALUES ('Billy Cobham','test','test');
INSERT INTO article (title,content,author) VALUES ('Black Label Society','test','test');
INSERT INTO article (title,content,author) VALUES ('Black Sabbath','test','test');


