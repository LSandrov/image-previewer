# Image Previewer

![workflow](https://github.com/LSandrov/image-previewer/actions/workflows/tests.yml/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/LSandrov/image-previewer/branch/master/graph/badge.svg?token=TAAM8J01Y9)](https://codecov.io/gh/LSandrov/image-previewer)


Сервис предназначен для изготовления preview (создания изображения с новыми размерами на основе имеющегося изображения).

## Команды для работы с сервисом

- ``make run`` - билд сервиса + поднятие nginx контейнера для работы с сервисом (по умолчанию 127.0.0.1 порт 80)
- ``make stop`` - остановка контейнеров
- ``make build`` - билд сервиса с локального окружения
- ``make clear`` - очистка билдов
- ``make lint`` - запуск линтера кода
- ``make test`` - запуск юнит тестов
- ``make test-e2e`` - запуск интеграционных тестов (end to end) (обязательно перед выполнением выполнить ``make run``)

## Описание работы
Отправка параметров на url /fill/ возвращает масштабирование изображение

### Описание параметров:
http://127.0.0.1/fill/{width}/{height}/{url}

Где: width - ширина, height - высота, url - ссылка на изображение (без схемы)

### Пример обращения к сервису
```http://127.0.0.1/fill/200/200/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_1024x252.jpg```

### Возможные варианты ответов:
 - Код 200, тело ответа: масштабированное изображение
 - Код 400, тело ответа: Ошибка при валидации входных данных
 - Код 502, тело ответа: Невозможно обработать изображение
 - Код 500, тело ответа: Проблемы с обработкой ответа