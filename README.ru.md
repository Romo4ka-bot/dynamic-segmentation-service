# Dynamic Segmentation Service

Читать на других языках: [English](README.md), [Русский](README.ru.md)

Это сервис, предназначен для хранения пользователя и сегментов в которых он состоит. 
С возможностью создания, изменения, удаления сегментов, также добавление и удаление пользователей в сегмент и создание пользователя.

## Локальный запуск сервиса

1) В сервисе используется Docker, поэтому для запуска достаточно прописать команду ниже.

```bash
docker-compose up --build -d
```

У вас поднимутся два контейнера(приложение и бд PostgreSQL), описанные в файле docker-compose.yml в корне проекта. 
Для проверки, что контейнеры успешно поднялись, можете написать команду ниже и посмотрите статус(Up).

```bash
docker ps
```

2) Вы можете остановить сервис с помощью команду:

```bash
docker-compose stop
```

## Основной функционал:
1) Создание сегмента
    * урл: /api/v1/segments POST
    * входящее значение:
   ```go
   package model
   type Segment struct {
	    Slug string
   }
   ```
    * возвращаемое значение: идентичная модель дополненная некоторыми значениями
2) Удаление сегмента по id
    * урл: /api/v1/segments/{id} DELETE
    * входящее значение: id сегмента
    * возвращаемое значение: При успешном(200) ответе вернется null
3) Добавление пользователей в сегмент и удаление из него
   * урл: /api/v1/user-segments PUT
   * входящее значение:
   ```go
   package model
   type UserSegments struct {
    AddSlugs []string
    RemoveSlugs []string
    UserId int
   }
   ```
   * возвращаемое значение:
   ```go
   package model
   type User struct {
	   Id int
	   FirstName string
	   SecondName string
	   Username string
	   HashPassword string
   }
   ```
4) Получение сегментов в которых состоит пользователь
   * урл: /api/v1/user-segments/users/{user_id} GET
   * входящее значение: id пользователя(user_id)
   * возвращаемое значение: модель User, которая описана в пункте 3

Более подробное описание можно посмотреть в свагере и протестировать запросы. Swagger Url: http://localhost:8080/swagger/index.html.

Также в файле [requests.http](requests.http), находящийся в корне проекта, были написаны примеры запросов, 
которые можно запустить удобной для вас среде разработки или любым другим способом.

### Список библиотек и технологий, которые использовались для создания сервиса:
   * [gin](https://github.com/gin-gonic/gin) для описания эндпоинтов API и обработки запросов
   * [net/http](https://pkg.go.dev/net/http) встроенная библиотека для запуска сервера
   * [gorm](https://github.com/go-gorm/gorm) в качестве ORM
   * [swag](https://github.com/swaggo/swag) swagger для описания API
   * [logrus](https://github.com/sirupsen/logrus) для логирования ошибок в формате JSON
   * СУБД PostgreSQL
   * Docker, docker-compose
