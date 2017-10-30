-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE client
(
  id       INT AUTO_INCREMENT,
  email    VARCHAR(50) NOT NULL,
  password VARCHAR(32) NOT NULL,
  status   INTEGER(1)  NOT NULL,
  CONSTRAINT `PRIMARY`
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX unique_email
  ON client (email);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX unique_email
ON client;
DROP TABLE client;