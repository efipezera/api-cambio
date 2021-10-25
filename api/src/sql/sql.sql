CREATE DATABASE IF NOT EXISTS desafio;
USE desafio;

DROP TABLE IF EXISTS depositos;
CREATE TABLE depositos(
    valorDeposito float not null
) ENGINE=INNODB;