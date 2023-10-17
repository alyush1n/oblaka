# Лабораторная работа №1

## Цель работы

Написать два Dockerfile – плохой и хороший.

## Задачи

* Написать "плохой" Dockerfile, который должен запускаться и работать корректно, но в нём должно быть не менее 3 "bad
  practices"
* Написать "хороший" Dockerfile, в котором эти "bad practices" должны быть исправлены.
* В отчёте (формат .md) описать все плохие практики и почему они плохие, как они были исправлены.

## Ход работы

За основу был взят Dockerfile, который запускает образ Ubuntu:

```
FROM ubuntu:20.04

RUN apt-get update && apt-get install -y curl

CMD ["tail", "-f", "/dev/null"]
```

### "Плохой" Dockerfile

Далее были допущены следующие "bad practices":

* #### Использование latest вместо конкретной версии

```
# bad practice 1 - latest
FROM ubuntu:latest
```

#### Почему плохо?

"latest" означает использование самой последней доступной версии образа, что может привести к неожиданному его
обновлению, и, как следствие, непредсказуемому поведению сервиса.

* #### Запуск двух сервисов в одном Dockerfile

```
FROM ubuntu:latest

RUN apt-get update && apt-get install -y curl

CMD ["tail", "-f", "/dev/null"]

# bad practice 2 - multiple services inside 1 container
FROM ubuntu:20.04
```

#### Почему плохо?

Каждый контейнер должен соответствовать одному процессу. В случае, когда требуется запустить несколько контейнеров,
используют **docker-compose**.

* #### Выполнение POST-запроса в Dockerfile

```
# bad practice 3 - POST-request
CMD ["curl", "-X", "POST", "https://httpbin.org/post"]
```

#### Почему плохо?

Dockerfile должен описывать, как собрать контейнер для приложения, в котором оно будет работать. Выполнение сторонних
функций (POST-запросов, git-коммитов и т.д.) нарушает детерминированность сборки образа.

### "Хороший" Dockerfile

* #### Использование конкретной версии вместо "latest"

```
FROM ubuntu:20:04
```

* #### Использование docker-compose для запуска нескольких контейнеров

```
version: '3'
services:
  ubuntu1:
    image: ubuntu:20.04
    container_name: my-ubuntu1
    command: tail -f /dev/null
  ubuntu2:
    image: ubuntu:20.04
    container_name: my-ubuntu2
    command: tail -f /dev/null
```

* #### Выполнение POST-запросов в коде приложения или в работе процесса

```
# код выполнения POST-запроса удалён
CMD ["tail", "-f", "/dev/null"]
```