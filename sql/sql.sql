CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id serial primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(100) not null unique,
    senha varchar(32) not null unique,
    criadoEm timestamp default current_timestamp()

) ENGINE=INNODB;