MedTracker

MedTracker — это веб-приложение, разработанное на языке Golang, предназначенное для управления данными о пациентах, медикаментах, напоминаниях и журналах приема медикаментов. Оно предоставляет набор RESTful эндпоинтов для выполнения различных операций с данными и обеспечения функционала для отслеживания приема медикаментов и напоминаний. Приложение поддерживает CRUD операции для пациентов, медикаментов, напоминаний и журналов приема.

Кроме того, проект включает Frontend часть, разработанную с использованием Next.js, которая взаимодействует с API и предоставляет пользователю удобный интерфейс для управления данными.

Ключевые особенности:

Управление пациентами:
Создание новых пациентов.
Получение информации о пациентах по их ID.
Обновление информации о пациентах.
Удаление пациентов.
Управление медикаментами:
Создание новых медикаментов.
Получение информации о медикаментах по их ID.
Обновление информации о медикаментах.
Удаление медикаментов.
Управление напоминаниями:
Создание новых напоминаний.
Получение информации о напоминаниях по их ID.
Обновление информации о напоминаниях.
Удаление напоминаний.
Управление журналами приема:
Создание новых записей в журнале приема.
Получение информации о записях по их ID.
Обновление информации о записях.
Удаление записей из журнала приема.


API Endpoints

Пациенты (Patients)
POST /patients
Описание: Создание нового пациента.
GET /patients/:id
PUT /patients/:id
DELETE /patients/:id

Медикаменты (Medications)
POST /medications
Описание: Создание нового медикамента.
GET /medications/:id
PUT /medications/:id
DELETE /medications/:id
Напоминания (Reminders)
POST /reminders
Описание: Создание нового напоминания для пациента.
GET /reminders/:id
PUT /reminders/:id
DELETE /reminders/:id

Журнал приема (Intake Logs)
POST /intake_logs
Описание: Создание новой записи в журнале приема.

GET /intake_logs/:id
PUT /intake_logs/:id
DELETE /intake_logs/:id

Структура базы данных

-- Таблица пользователей
CREATE TABLE `users` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100),
  `email` varchar(100),
  `password_hash` varchar(255),
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Таблица пациентов
CREATE TABLE `patients` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100),
  `age` int,
  `condition` varchar(255),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Таблица медикаментов
CREATE TABLE `medications` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36),
  `name` varchar(100),
  `dosage` varchar(50),
  `frequency` varchar(50),
  `start_date` date,
  `end_date` date,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Таблица напоминаний
CREATE TABLE `reminders` (
  `id` varchar(36) NOT NULL,
  `medication_id` varchar(36) NOT NULL,
  `time` varchar(5),
  `medication_name` varchar(100),
  `user_id` varchar(36) NOT NULL,
  `frequency` varchar(20) DEFAULT 'daily',
  PRIMARY KEY (`id`),
  KEY `medication_id` (`medication_id`),
  KEY `fk_reminders_user` (`user_id`),
  CONSTRAINT `fk_reminders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `reminders_ibfk_1` FOREIGN KEY (`medication_id`) REFERENCES `medications` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Таблица журналов приёма
CREATE TABLE `intake_logs` (
  `id` varchar(36) NOT NULL,
  `medication_id` varchar(36),
  `taken_at` datetime,
  `was_taken` tinyint(1),
  PRIMARY KEY (`id`),
  KEY `medication_id` (`medication_id`),
  CONSTRAINT `intake_logs_ibfk_1` FOREIGN KEY (`medication_id`) REFERENCES `medications` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



Frontend

Frontend часть проекта разработана с использованием Next.js. Она предоставляет пользователю удобный интерфейс для управления пациентами, медикаментами, напоминаниями и журналами приема. Это современное одностраничное приложение с поддержкой SSR (Server-Side Rendering) и API маршрутов в Next.js.

GitHub репозиторий Frontend части: https://github.com/daniyardautbaev/medtracker-frontend.git
