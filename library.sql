-- Adminer 4.8.1 MySQL 8.1.0 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `Authors`;
CREATE TABLE `Authors` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Authors` (`id`, `name`) VALUES
(1,	'Иар Эльтеррус'),
(2,	'Екатерина Белецкая');

DROP TABLE IF EXISTS `Books`;
CREATE TABLE `Books` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Books` (`id`, `name`) VALUES
(1,	'Танцующий бог'),
(2,	'Затерянный замок'),
(3,	'Долгий солнечный день'),
(4,	'Адонай и Альтея'),
(5,	'Тингл твист'),
(6,	'Утро черных звезд'),
(7,	'День черных звезд'),
(8,	'Вечер черных звезд'),
(9,	'Ночь черных звезд');

DROP TABLE IF EXISTS `Relevant`;
CREATE TABLE `Relevant` (
  `id_book` int NOT NULL COMMENT 'Book primary Key',
  `id_author` int NOT NULL COMMENT 'Author primary Key'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Relevant` (`id_book`, `id_author`) VALUES
(1,	1),
(2,	1),
(6,	1),
(7,	1),
(8,	1),
(9,	1),
(3,	2),
(4,	2),
(5,	2),
(6,	2),
(7,	2),
(8,	2),
(9,	2);

-- 2023-09-07 10:57:23
