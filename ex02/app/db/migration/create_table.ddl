USE mixi;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(255) NOT NULL DEFAULT '',
    PRIMARY KEY (id),
    UNIQUE KEY uniq_user_id (user_id)
);

DROP TABLE IF EXISTS friend_links;
CREATE TABLE friend_links (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user1_id) REFERENCES users(user_id),
    FOREIGN KEY (user2_id) REFERENCES users(user_id)
);

DROP TABLE IF EXISTS block_lists;
CREATE TABLE block_lists (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user1_id INT NOT NULL,
    user2_id INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user1_id) REFERENCES users(user_id),
    FOREIGN KEY (user2_id) REFERENCES users(user_id)
);
