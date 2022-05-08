### Сервис для записи и чтения данных из Clickhouse


Создайте .env файл в корне проекта (используется godotenv для простейшей реализации конфигурирования, если нужно сделать грамотнее можно использовать Viper)
```dotenv
PORT = 8000

DB_HOST=127.0.0.1
DB_PORT=9000
DB_NAME=information
DB_USERNAME=""
DB_PASSWORD=""
```

### Миграции базы данных
```
migrate -path ./migration -database clickhouse://:@localhost:9000/information -verbose up
```
Для сидинга рандомных данных в репозитории clickhouse есть метод Information.SeedData
Генерирует 10_000 рандомных "пользователей" с рандомными юзерагентами,айпи и айди

## Docker-compose 
Докер компоуз разворачивает сервис и базу данных
В .env вписать в качестве хоста для дб "ch_server"


