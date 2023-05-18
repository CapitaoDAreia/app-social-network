insert into users (username, nick, email, password)
values 
("user1", "user1", "user1@email.com", "$2a$10$gu28S32jsJybceQ/9gIUJu5l9xjUpr/7uduCdzD/54He57kiRYyfG"), 
("user2", "user2", "user2@email.com", "$2a$10$gu28S32jsJybceQ/9gIUJu5l9xjUpr/7uduCdzD/54He57kiRYyfG"),
("user3", "user3", "user3@email.com", "$2a$10$gu28S32jsJybceQ/9gIUJu5l9xjUpr/7uduCdzD/54He57kiRYyfG"),
("user4", "user4", "user4@email.com", "$2a$10$gu28S32jsJybceQ/9gIUJu5l9xjUpr/7uduCdzD/54He57kiRYyfG");

--password without hash: 123456

insert into followers(user_id, follower_id)
values
(1, 2),
(1, 3),
(2, 3);
