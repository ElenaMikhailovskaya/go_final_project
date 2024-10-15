# Итоговый проект
# TODO лист

В проекте реализуется простейший планировщик задач.
Доступные функции:
- Добавление задачи (в том числе с периодичностью)
- Редактирование задачи
- Удаление задачи

## Архитектура
- основной сервис  [находится в /cmd/app/main.go]
- БД sqlLite [internal/interfaces/database/source/scheduler.db]
- приложение реализовано на фреймворке Fiber

## Требования к окружению
- Go 1.22^

### Переменные окружения
TODO_DBFILE - путь к файлу базы данных [значение по умолчанию "./internal/interfaces/database/source/scheduler.db"]
TODO_PORT - порт сервиса [значение по умолчанию "localhost:7540"]

### Запуск тестов
тесты можно запускать через test-coverage в файле Makefile в корне проекта

параметры в tests/settings.go

var Port = 7540
var DBFile = "../internal/interfaces/database/source/scheduler.db"
var FullNextDate = false
var Search = false
var Token = ``

### Запуск контейнера
Dockerfile находится в docker/ci/Dockerfile

Для сборки контейнера:
- скопировать файл Dockerfile в корень проекта
- выполнить команду на сборку контейнера [docker build --tag elenamikhailovskaya/golang_final_project:final .]
- либо скачать с докер хаба [docker pull elenamikhailovskaya/golang_final_project]
- выполнить команду на запуск контейнера с портом 7540 [docker run -d -p 7540:7540 elenamikhailovskaya/golang_final_project:final]