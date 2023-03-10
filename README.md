# Social Net Service

Сервис для социальной сети. В рамках небольшого учебного проекта научиться

- писать микросервис и использовать proxy
- тестировать написанное приложение 

- работать с запросами POST, GET, PUT, DELETE;
- применять принципы написания обработчиков HTTP-запросов

## Пользователь

Данные пользователя социальной сети:

- Имя
- Возраст
- Массив друзей

## Функционал сервиса

- создание пользователя 
- получение информации о пользователе
- подружить двух пользователей
- удалить пользователя
- получений информации о всех друзьях пользователя
- обновление возраста пользователя

## Обработчики сервиса

#### 1. Создание пользователя

Обработчик должен создавать пользователя

Пример запроса

```http
POST /create HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
```

```json
{	
    "name":"some name",
    "age":"24",
    "friends":[]
}
```

Данный запрос должен возвращать ID пользователя и статус 201

#### 2. Сделать друзьями двух пользователей

Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья. Пример запроса:

```http
POST /make_friends HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
```

```json
{
	"source_id":"1",
	"target_id":"2"
}
```

Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».

#### 3. Удаление пользователя

Обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей. Пример запроса:

```http
DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
```

```json
{
    "target_id":"1"
}
```

Данный запрос должен возвращать 200 и имя удалённого пользователя.

#### 4. Друзья пользователя

Обработчик, который возвращает всех друзей пользователя. Пример запроса:

```http
GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
```

После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.

#### 5. Обновить возраст пользователя.

Пример запроса:

```http
PUT /user_id HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
```

```json
{
    "new age":"28"
}
```

Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».

#### 6. Информация о пользователе

Обработчик, который возвращает информацию о пользователе. Пример запроса:

```http
GET /user_id HTTP/1.1
Host: localhost:8080
Connection: close
```

После / указывается id пользователя, информацию которого мы хотим увидеть.

## Развертывание и запуск сервиса

### Docker

Сервис использует Docker

### Nginx

Используется proxy nginx для маршрутизации на три реплики данного приложения

### PostgreSQL

Данные пользователей хранятся в базе данных PostgreSQL

### Сборка и запуск

#### Запуск одной реплики с PostgreSQL хранилищем

Создание и запуск контейнеров

```makefile
make build_app_db
make up_app_db 
```

Остановить и удалить контейнеры

```makefile
make down_app_db
```

#### Запуск трех реплик с Nginx, PostgreSQL хранилищем

Создание и запуск контейнеров

```makefile
make build
make up
```

Остановить и удалить контейнеры

```makefile
make down
```

## Документирование

Для документирования используется [Swagger](https://github.com/swaggo/http-swagger)

Swagger работает с одной репликой

```makefile
make up_app_db_test
```

Генерация документации:

```makefile
make swag
```

Просмотреть документации и протестировать:

```http
http://localhost:8080/swagger/index.html#/
```

## Тестирование

Запустить тесты на тестовых данных:

```makefile
make up_app_db_test
```

```go
go test SocialNetHTTPService/tests
```

