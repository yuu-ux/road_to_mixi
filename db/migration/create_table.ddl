USE mixi;
CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

CREATE TABLE friend_link (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE block_list (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    PRIMARY KEY (id)
);
