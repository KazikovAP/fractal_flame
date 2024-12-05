[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)

# fractal_flame
# Генерация изображений фрактального пламени

---
## Описание проекта
Проект представляет собой алгоритм генерации изображения фрактального пламени, основанного на идее Chaos Game.

---
## Технологии
* Go 1.23.0
* DDD (Domain Driven Design)
* Tests

---
## Запуск проекта

**1. Клонировать репозиторий:**
```
git clone https://github.com/KazikovAP/fractal_flame
```

**2. Запустить генерацию изображения:**
```
go run cmd/fractal_flame/main.go
```

## Пример ответа
```
Запуск генерации изображения фрактального пламени с параметрами:
Ширина: 3000
Высота: 3000
Итерации: 200000
Функция трансформации: waves
Количество трансформаций: 3
Число симметрий: 80
Гамма-коррекция: 1.000000
Режим выполнения: single
Время выполнения: 4.2369746s
Фрактальное пламя сгенерировано и сохранено в файл: data\fractal_waves.png
```
![fractal_waves](data/fractal_waves.png)

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)
