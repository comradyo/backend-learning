# Порядок действий

https://dev.to/francescoxx/build-a-crud-rest-api-in-go-using-mux-postgres-docker-and-docker-compose-2a75

1. Создаем папку (уже создана) <br />
`mkdir go-crud-api`
2. Переходим в папку <br />
`cd go-crud-api`
3. Создаём go-модуль <br />
`go mod init api`
4. Устанавливаем зависимости <br />
`go get github.com/gorilla/mux github.com/lib/pq`
5. Создаём нужные файлы (уже созданы) <br />
`touch main.go Dockerfile docker-compose.yml`
6. Копируем и вставляем код в файлы (уже сделано) <br />
7. Поднимаем контейнер с БД (образ берется из репозитория с образами postgres, подробнее описал в docker-compose.yml) <br />
`docker compose up -d go-db`
8. Смотрим, что контейнер поднялся и подключаемся к базе, чтобы убедиться, что с ней всё ок <br />
`docker ps -a` — список контейнеров и их статусы  <br />
`psql -h localhost -p 5432 -d postgres -U postgres -W` — подключение к БД через консоль
9. Собираем образ нашего go-приложения (подробнее об этом в docker-compose.yml и Dockerfile) <br />
`docker compose build`
10. Проверяем, что образ создался <br />
`docker images` — список образов (в локальном репозитории (?))
11. Поднимаем контейнер с go-приложением <br />
`docker compose up go-app`
12. Тестируем приложение через curl (с помощью postman, например) <br />
Список методов и url:
    1. `POST localhost:8000/users` — создание пользователя
    2. `GET localhost:8000/users` — список всех пользователей
    3. `GET localhost:8000/users/<user_id>` — получить конкретного пользователя (треугольные скобочки не нужны)
    4. `PUT localhost:8000/users/<user_id>` — обновление конкретного пользователя
    5. `DELETE localhost:8000/users/<user_id>` — удаление конкретного пользователя
