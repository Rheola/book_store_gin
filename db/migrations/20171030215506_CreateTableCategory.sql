-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE category
(
  id   INT AUTO_INCREMENT,
  name VARCHAR(32) NOT NULL,
  CONSTRAINT `PRIMARY`
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX unique_name
  ON category (name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX unique_name
ON category;
DROP TABLE category;
