<p align="center">
  <a href="" rel="noopener">
 <img width=843px height=173px src="https://github.com/honyshyota/l0-wb-test/blob/master/images/logo.png" alt="Project logo"></a>
</p>


# WB L0 test

## Task Description

В БД:
Развернуть локально postgresql
Создать свою бд
Настроить своего пользователя.
Создать таблицы для хранения полученных данных.
В сервисе:
1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать просте ши интерфе с отображения полученных данных, для
их запроса по id
Доп инфо:
• Данные статичны, исходя из этого подума те насчет модели хранения
в Кеше и в pg. Модель в файле model.json
• В канал могут закинуть что угодно, подума те как избежать проблем
из-за этого
• Чтобы проверить работает ли подписка онла н, сдела те себе
отдельны скрипт, для публикации данных в канал
• Подума те как не терять данные в случае ошибок или проблем с
сервисом
• Nats-streaming разверните локально ( не путать с Nats )

## Built using

- [PostgresqlDB](https://www.postgresql.org/) - Database
- [Nats-streaming](https://nats.io/) - Data streaming
- [Docker-compose](https://www.docker.com/) - Container service
- [Chi](https://github.com/go-chi/chi) - Router

## How to

* Запускать с помощью ```make docker```
* Запуск скрипта
```cd script```
```go run scipt.go```
* Проверка данных по адрессу
```localhost:8080/orders/{id}``` для прошедших валидацию данных
```localhost:8080/bad/{id}``` для не прошедших валидацию данных


## Look here

![alt text](https://github.com/honyshyota/l0-wb-test/blob/master/images/docker_run.png)
![alt text](https://github.com/honyshyota/l0-wb-test/blob/master/images/web.png)
