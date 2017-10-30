-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `order`
(
  id        INT AUTO_INCREMENT
    PRIMARY KEY,
  book_id   INT(11)  NOT NULL
  COMMENT 'Книга',
  client_id INT(11)  NOT NULL
  COMMENT 'Категория',
  DATE      DATETIME NOT NULL
  COMMENT 'Дата',
  reserved  INT(1)   NOT NULL
  COMMENT 'Зарезервирована'
);

CREATE INDEX book_id
  ON `order` (book_id);
CREATE INDEX client_id
  ON `order` (client_id);

ALTER TABLE `order`
  ADD CONSTRAINT `order_book_fk` FOREIGN KEY (`book_id`) REFERENCES `book` (`id`),
  ADD CONSTRAINT `order_client_fk` FOREIGN KEY (`client_id`) REFERENCES `client` (`id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `order`
  DROP FOREIGN KEY order_client_fk;
ALTER TABLE `order`
  DROP FOREIGN KEY order_book_fk;
DROP INDEX book_id
ON `order`;
DROP INDEX client_id
ON `order`;
DROP TABLE `order`;
