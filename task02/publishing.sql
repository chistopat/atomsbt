CREATE TABLE IF NOT EXISTS `authors` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `created_at` timestamp DEFAULT NOW(),
    `full_name` varchar(255)
    );

CREATE TABLE IF NOT EXISTS `books` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `created_at` timestamp DEFAULT NOW(),
    `title` varchar(1000)
    );

CREATE TABLE IF NOT EXISTS `publications` (
    `isbn` int PRIMARY KEY AUTO_INCREMENT,
    `created_at` timestamp DEFAULT NOW(),
    `year` int not null,
    `book_id` int,
    `author_id` int,
    `fee` int not null ,
    `currency` ENUM('RUB', 'USD', 'EUR')
    );

ALTER TABLE `publications` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);
ALTER TABLE `publications` ADD FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`);


INSERT INTO `authors` (full_name)
VALUES
    ('Илья Ильф'),
    ('Евгений Петров'),
    ('Нил Стивенсон');


INSERT INTO `books` (title)
VALUES
    ('Двенадцать стульев'),
    ('Золотой теленок'),
    ('Лавина'),
    ('Криптономикон'),
    ('Семиевие');

INSERT INTO `publications` (`year`, `book_id`, `author_id`, `fee`, `currency`)
VALUES
    (1928, 1, 1, 1000, 'RUB'),
    (1928, 1, 2, 1000, 'RUB'),
    (1931, 2, 1, 800, 'RUB'),
    (1931, 2, 2, 800, 'RUB'),
    (1999, 3, 3, 100, 'USD'),
    (1992, 4, 3, 200, 'USD'),
    (2015, 5, 3, 300, 'USD'),
    (2021, 5, 3, 500, 'USD');


SELECT
    SUM(fee) as fee
     , a.full_name as author
FROM publications as p
JOIN authors as a on p.author_id = a.id
WHERE 1=1
    and year BETWEEN 1990 AND 2020
    and a.full_name = 'Нил Стивенсон'
GROUP BY author_id
