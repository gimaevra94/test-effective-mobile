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
### 1. Перейдите в папку test-effective-mobile

### 2. Создайте файл .env и вставьте туда:

CONNECTION_CFG="postgresql://admin:admin@localhost:5432/db?sslmode=disable"

POSTGRES_PASSWORD="admin" # :)

### 3. Соберите и запустите контейнеры:

```
docker compose -f docker-compose.yml --env-file .env up -d
```
### 4. Запустите сервер

```
go run main.go
```
<br>

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

## Пример стуктуры для POST запроса
```json
{
“service_name”: “Yandex Plus”,
“price”: 400,
“user_id”: “60601fee-2bf1-4721-ae6f-7636e79a0cba”,
“start_date”: “07-2025”
}
```
