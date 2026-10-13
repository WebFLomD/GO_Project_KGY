-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1
-- Время создания: Окт 13 2026 г., 15:17
-- Версия сервера: 10.4.32-MariaDB
-- Версия PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `personal_expenses`
--

-- --------------------------------------------------------

--
-- Структура таблицы `expenses`
--

CREATE TABLE `expenses` (
  `id` int(11) NOT NULL COMMENT 'Индетификатор',
  `description` varchar(100) NOT NULL COMMENT 'Наименование',
  `amount` char(50) NOT NULL COMMENT 'Сумма',
  `date` date NOT NULL DEFAULT current_timestamp() COMMENT 'Дата'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Дамп данных таблицы `expenses`
--

INSERT INTO `expenses` (`id`, `description`, `amount`, `date`) VALUES
(1, 'Такси до дома', '450', '2026-10-12'),
(2, 'Продукты', '1500', '2026-10-13'),
(3, 'Подписка ВК', '99', '2026-10-09');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `expenses`
--
ALTER TABLE `expenses`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `expenses`
--
ALTER TABLE `expenses`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Индетификатор', AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
