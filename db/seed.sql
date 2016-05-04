DROP TABLE users;

CREATE TABLE users (
  email varchar(255),
  password varchar(255)
);

INSERT INTO users (email, password)
VALUES ("sam", "password");

INSERT INTO users (email, password)
VALUES ("ringo", "starr");
