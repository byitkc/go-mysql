CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(255),
    firstName VARCHAR(255),
    lastName VARCHAR(255),
    createdAt TIME,
    lastLogin TIME,
    PRIMARY KEY (id)
)