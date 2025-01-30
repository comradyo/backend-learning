# Порядок работы с goose

Установка:

`go install github.com/pressly/goose/v3/cmd/goose@latest`

Переменные окружения (добавил в файле .env):

1. Драйвер базы <br />
`GOOSE_DRIVER=postgres`
2. Строка подключения к базе <br />
`GOOSE_DBSTRING="креды"` <br />
Пример: `host=go-db user=postgres password=postgres dbname=postgres sslmode=disable`
3. Папка с миграциями <br />
`GOOSE_MIGRATION_DIR=migrations` <br />

Создание новой миграции:

`goose create add_users_table sql`

Накатка всех ненакаченных миграции:

`goose up`

Накатка миграций по одной:

`goose up-by-one`

Откат последней миграции:

`goose down`

Статус миграций

`goose status`

###  Проверка работоспособности

Сначала запустил все контейнеры, проверил, что приложение отдаст ошибку, когда миграция не накачена, потом накатил миграцию (не выключая приложение), и опять сделал запрос, ошибка исчезла.
