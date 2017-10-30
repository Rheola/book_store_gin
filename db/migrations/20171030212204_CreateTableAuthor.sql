-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE author
(
  id   INT AUTO_INCREMENT,
  name VARCHAR(32) NOT NULL,
  CONSTRAINT `PRIMARY`
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX unique_name
  ON author (name);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX unique_name
ON author;
DROP TABLE author;

