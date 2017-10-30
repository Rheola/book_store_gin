-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE book
(
  id          INT AUTO_INCREMENT
    PRIMARY KEY,
  author_id   INT          NOT NULL
  COMMENT 'Автор',
  category_id INT          NOT NULL
  COMMENT 'Категория',
  title       VARCHAR(255) NOT NULL
  COMMENT 'Название',
  count       INT          NOT NULL
  COMMENT 'Остаток'
);

CREATE INDEX author_id
  ON book (author_id);
CREATE INDEX category_id
  ON book (category_id);

ALTER TABLE `book`
  ADD CONSTRAINT `book_author_fk` FOREIGN KEY (`author_id`) REFERENCES `author` (`id`),
  ADD CONSTRAINT `book_category_fk` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE book
  DROP FOREIGN KEY book_category_fk;
ALTER TABLE book
  DROP FOREIGN KEY book_author_fk;
DROP INDEX author_id
ON book;
DROP INDEX category_id
ON book;
DROP TABLE book;
