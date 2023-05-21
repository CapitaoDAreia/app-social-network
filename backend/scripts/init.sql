CREATE DATABASE IF NOT EXISTS social_network_database;

USE social_network_database;

CREATE TABLE IF NOT EXISTS users(
    id int auto_increment primary key,
    username varchar(100) not null,
    nick varchar(100) not null unique,
    email varchar(100) not null unique,
    password varchar(100) not null,
    CreatedAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS posts(
    id int auto_increment primary key,
    title  varchar(100) not null,
    content varchar(500) not null,
    
    authorId int not null,
    FOREIGN KEY (authorId)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS followers(
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