CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id serial primary key,
    nome varchar(255) not null,
    nick varchar(255) not null unique,
    email varchar(255) not null unique,
    senha varchar(255) not null,
    criadoEm timestamp default current_timestamp

)

DROP TABLE IF EXISTS seguidores;

CREATE TABLE seguidores (
    usuario_id integer not null,
    seguidor_id integer not null,
    criadoEm timestamp default current_timestamp,
    foreign key (usuario_id) references usuarios(id) on delete cascade,
    foreign key (seguidor_id) references usuarios(id) on delete cascade,
    primary key (usuario_id, seguidor_id)
);

CREATE TABLE publicacoes (
    id serial primary key,
    autor_id integer not null,
    titulo varchar(255) not null,
    conteudo varchar(300) not null,
    criadoEm timestamp default current_timestamp,
    curtidas integer default 0,
    foreign key (usuario_id) references usuarios(id) on delete cascade
);