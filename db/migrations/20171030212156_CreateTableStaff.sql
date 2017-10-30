-- +goose Up
CREATE TABLE staff
(
  id       INT AUTO_INCREMENT,
  active   INT(1) DEFAULT '1' NOT NULL,
  login    VARCHAR(30)        NOT NULL,
  password VARCHAR(32)        NOT NULL,
  CONSTRAINT `PRIMARY`
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX unique_login
  ON staff (login);

-- +goose Down
DROP INDEX unique_login
ON staff;
DROP TABLE staff;
