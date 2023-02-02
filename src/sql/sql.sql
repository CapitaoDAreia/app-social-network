CREATE DATABASE IF NOT EXISTS api_dvbk;

USE api_dvbk;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    username varchar(100) not null,
    nick varchar(100) not null unique,
    email varchar(100) not null unique,
    password varchar(100) not null,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;



DROP TABLE IF EXISTS followers;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB;