CREATE DATABASE IF NOT EXISTS api_dvbk;

USE api_dvbk;

DROP TABLE IF EXISTS posts;

CREATE TABLE posts(
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