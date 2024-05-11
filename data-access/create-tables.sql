DROP TABLE IF EXISTS book;
CREATE TABLE book (
  id         SERIAL NOT NULL,
  title      VARCHAR(128) NOT NULL,
  author     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY ("id")
);

INSERT INTO book
  (title, author, price)
VALUES
  ('Book 1', 'Author 1', 56.99),
  ('Book 2', 'Author 1', 63.99),
  ('Book 3', 'Author 2', 17.99),
  ('Book 4', 'Author 3', 34.98);
