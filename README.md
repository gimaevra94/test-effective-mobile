## Описание
RESTful API для управления подписками пользователей
<br>
<br>
## Поддерживаемые операции
* создание;
* чтение;
* обновление;
* удаление;
* получение списка подписок;
* агрегация стоимости подписок за указанный период
<br>

## Стек технологий

* Go-Chi — легковесный HTTP-роутер;
* GORM — ORM для работы с базой данных;
* PostgreSQL — реляционная база данных;
* Swaggo — генерация документации в формате Swagger/OpenAPI;
* Docker — контейнеризация сервисов
<br>

## Запуск
### 1. Создайте файл .env в корне проекта и вставьте туда:

CONNECTION_CFG="postgresql://admin:admin@localhost:5432/db?sslmode=disable"

POSTGRES_PASSWORD="admin" # :)

### 2. Находясь в корне, соберите и запустите контейнеры:

docker compose -f docker-compose.yml --env-file .env up -d

**После запуска сервер будет доступен по адресу**
http://localhost:8080

**Документация Swagger доступна по адресу**
http://localhost:8080/swagger/index.html
<br>
<br>

## Доступные эндпоинты
| Метод | Путь | Описание |
|-------|-----------------------------------------------|--------------------------------------------------------|
| POST | /api/v1/subscription | Создать новую подписку |
| GET | /api/v1/subscription | Получить список всех подписок (с фильтрацией) |
| GET | /api/v1/subscription/{service_name}/{user_id} | Получить конкретную подписку |
| PATCH | /api/v1/subscription/{service_name}/{user_id} | Обновить цену подписки |
| DELETE | /api/v1/subscription/{service_name}/{user_id} | Удалить подписку |
| GET | /api/v1/subscription/totalPrice | Получить общую стоимость подписок за период |
<br>

## Структура проекта
| consts | Константы (SQL-запросы, сообщения)
|--|--|
| database | Работа с PostgreSQL через sql.DB |
| docs | Сгенерированная Swagger-документация |
| errs | Обработка ошибок и ответов |
| handlers | HTTP-обработчики |
| structs | Структуры данных |
| main.go | Точка входа |
| go.mod | Зависимости |
