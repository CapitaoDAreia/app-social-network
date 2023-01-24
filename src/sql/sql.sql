CREATE DATABASE IF NOT EXISTS api_dvbk;

USE api_dvbk;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    username varchar(100) not null,
    nick varchar(100) not null unique,
    email varchar(100) not null unique,
    password varchar(50) not null unique,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;