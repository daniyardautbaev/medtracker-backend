# MedTracker

**MedTracker** — это веб-приложение, предназначенное для управления данными о пациентах, медикаментах, напоминаниях и журналах приема медикаментов. Приложение предоставляет набор RESTful эндпоинтов для выполнения различных операций с данными и обеспечения функционала для отслеживания приема медикаментов и напоминаний.

Кроме того, проект включает **Frontend**, разработанный с использованием **Next.js**, который взаимодействует с API и предоставляет удобный интерфейс для управления данными.

## Ключевые особенности

### Управление пациентами:
- Создание новых пациентов.
- Получение информации о пациентах по их ID.
- Обновление информации о пациентах.
- Удаление пациентов.

### Управление медикаментами:
- Создание новых медикаментов.
- Получение информации о медикаментах по их ID.
- Обновление информации о медикаментах.
- Удаление медикаментов.

### Управление напоминаниями:
- Создание новых напоминаний.
- Получение информации о напоминаниях по их ID.
- Обновление информации о напоминаниях.
- Удаление напоминаний.

### Управление журналами приема:
- Создание новых записей в журнале приема.
- Получение информации о записях по их ID.
- Обновление информации о записях.
- Удаление записей из журнала приема.

## API Endpoints

### Пациенты (Patients)

- `POST /patients` — Создание нового пациента.
- `GET /patients/:id` — Получение информации о пациенте по ID.
- `PUT /patients/:id` — Обновление информации о пациенте.
- `DELETE /patients/:id` — Удаление пациента.

### Медикаменты (Medications)

- `POST /medications` — Создание нового медикамента.
- `GET /medications/:id` — Получение информации о медикаменте по ID.
- `PUT /medications/:id` — Обновление информации о медикаменте.
- `DELETE /medications/:id` — Удаление медикамента.

### Напоминания (Reminders)

- `POST /reminders` — Создание нового напоминания для пациента.
- `GET /reminders/:id` — Получение информации о напоминании по ID.
- `PUT /reminders/:id` — Обновление информации о напоминании.
- `DELETE /reminders/:id` — Удаление напоминания.

### Журнал приема (Intake Logs)

- `POST /intake_logs` — Создание новой записи в журнале приема.
- `GET /intake_logs/:id` — Получение информации о записи по ID.
- `PUT /intake_logs/:id` — Обновление информации о записи.
- `DELETE /intake_logs/:id` — Удаление записи из журнала приема.

## Структура базы данных

### Таблица пользователей (users)

TABLE `users` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100),
  `email` varchar(100),
  `password_hash` varchar(255),
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);

Таблица пациентов (patients)

TABLE `patients` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100),
  `age` int,
  `condition` varchar(255),
  PRIMARY KEY (`id`)
) ;


Таблица медикаментов (medications)

TABLE `medications` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36),
  `name` varchar(100),
  `dosage` varchar(50),
  `frequency` varchar(50),
  `start_date` date,
  `end_date` date,
  PRIMARY KEY (`id`)
) ;

Таблица напоминаний (reminders)

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
) ;

Таблица журналов приёма (intake_logs)

CREATE TABLE `intake_logs` (
  `id` varchar(36) NOT NULL,
  `medication_id` varchar(36),
  `taken_at` datetime,
  `was_taken` tinyint(1),
  PRIMARY KEY (`id`),
  KEY `medication_id` (`medication_id`),
  CONSTRAINT `intake_logs_ibfk_1` FOREIGN KEY (`medication_id`) REFERENCES `medications` (`id`) ON DELETE CASCADE
);

Frontend

Frontend часть проекта разработана с использованием Next.js. Это современное одностраничное приложение с поддержкой SSR (Server-Side Rendering) и API маршрутов в Next.js.

GitHub репозиторий Frontend части:
MedTracker Frontend

Установка и запуск

Клонируйте репозиторий:
git clone https://github.com/daniyardautbaev/medtracker-backend.git
cd medtracker-backend
Установите зависимости:
go mod tidy
Запустите сервер:
go run main.go
Перейдите на http://localhost:8080 для доступа к API.
