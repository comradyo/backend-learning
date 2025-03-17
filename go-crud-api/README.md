# Базовое приложение с контейнерами

## Установка

1. Установка docker: TODO
2. Установка docker compose: <br />
Если следующая команда не сработает (а она, скорее всего, не сработает):
```
sudo apt-get install docker-compose-plugin
```
Тогда нужно подключиться сервер с обновлениями плагина docker compose к менеджеру пакетов (вроде) <br />
https://labex.io/tutorials/docker-how-to-fix-unable-to-locate-docker-compose-plugin-error-413758 <br />

```
sudo apt-get update

sudo apt-get install -y ca-certificates curl gnupg

sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update

sudo apt-get install -y docker-compose-plugin
```

Проверить, что плагин установился:
```
docker compose version
```

Добавить пользователя в группу sudo докера:
```
sudo usermod -aG docker $USER
sudo reboot
```
Или, чтобы не перезагружать систему, выполнить (нужно будет предварительно установить setfacl):
```
sudo setfacl --modify user:<user name or ID>:rw /var/run/docker.sock
```
3.

## Создание и запуск приложения
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

Вопросы:
1. Так, во-первых, куда-то потерял инструкцию, как поднимать контейнер через команду `docker` (без compose), вряд ли я её не писал
2. Опытным путем выяснил, что переменные окружения контейнера не инкапсулированы в нём, а берутся из docker-compose файла.
3. У одного контейнера свой localhost, поэтому не получится через localhost обратиться к другому контейнеру.

UPD (26.02.2025): redis <br />

https://habr.com/ru/articles/823936/ <br />
https://www.geeksforgeeks.org/complete-tutorial-of-configuration-in-redis/ <br />

7. Поднимаем контейнер с БД (образ берется из репозитория с образами postgres, подробнее описал в docker-compose.yml) <br />
`docker compose up -d go-redis`
8. Смотрим, что контейнер поднялся и подключаемся к базе, чтобы убедиться, что с ней всё ок <br />
`docker ps -a` — список контейнеров и их статусы  <br />
`psql -h localhost -p 5432 -d postgres -U postgres -W` — подключение к БД через консоль
`redis-cli`

# Команды

Информация о контейнере: <br />
`docker inspect container_name` <br />
`docker inspect -f "{{ .NetworkSettings.IPAddress }}" container_name` <br />

Подключение к терминалу контейнера: <br />

Выполнение команды в контейнере: <br />
`docker exec -it my_container sh -c "echo a && echo b"` <br />
Проверяем, что файл конфигурации nginx нормально скопировался: <br />
`docker exec -it nginx-test-server sh -c "cat /etc/nginx/conf.d/default.conf"` <br />

# Вопросы:
1. Что такое docker volume? <br />
Возможно, понял: volume - это какие-то данные, которые передаются в контейнер на этапе его поднятия. Например, данные для базы данных.
2. Как сделать новый докер образ, от которого потом буду создавать контейнеры? Например, образ с мигратором должен хранить в себе goose. `docker build` создает готовый образ?
3. Как передавать в контейнер какие-то файлы при его старте? Например, папку с миграциями в контейнер-мигратор. Возможно, в контейнер с nginx тоже стоит передавать конфигурационный файл?
4. Как передавать файлы c хоста на этапе build? Как правильно использовать copy?
5. Как сделать контейнер, который можно использовать как утилиту (например, мигратор)? То есть, чтобы его можно было "вызвать", он отработал и завершился? Возможно, через флаги в docker-compose? <br />
Возможно, это связано с понятие docker utility container <br />
https://github.com/actionanand/node-util <br />
Вот здесь и становится понятно, что такое `ENTRYPOINT` - это команда, которая дергается при вызове контейнера. <br />
Например: <br />
```
ENTRYPOINT [ "npm" ]
```
При вызове контейнера с названием `npm-util-container` и таким `ENTRYPOINT`, будет выполняться команда `npm` с переданными аргументами/флагами: 
```
docker-compose run --rm npm-util-container init
```
Сделал мигратор, работает. <br />
Как теперь сделать так, чтобы работала команда создания миграции? Или это не нужно вовсе? Миграцию буду создавать у себя на компе, а на сервере нужна только накатка, по сути. <br />
6. Как сделать так, чтобы контейнер слал данные в какой-то внешний файл, и чтобы он мог его подгрузить при перезапуске? Например, postgres.
