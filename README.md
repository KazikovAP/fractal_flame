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

## Пример генерации изображения
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

### Доступные параметры для генерации изображения
```
--width       ширина изображения
--height      высота изображения
--iterations  количество итераций для генерации фрактала
--trans       функция трансформации
--trans_count количество трансформаций
--symmetry    число симметрий
--gamma       гамма коррекция
--mode        режим выполнения программы 
```

## Ещё примеры
```
go run cmd/fractal_flame/main.go --width 4000 --height 4000 --iterations 180000 --trans polar --trans_count 10 --symmetry 30 --gamma 0.9 --mode multi
```
```
Запуск генерации изображения фрактального пламени с параметрами:
Ширина: 4000
Высота: 4000
Итерации: 180000
Функция трансформации: polar
Количество трансформаций: 10
Число симметрий: 30
Гамма-коррекция: 1.000000
Режим выполнения: multi
Время выполнения: 1.60567s
Фрактальное пламя сгенерировано и сохранено в файл: data\fractal_polar.png
```
![fractal_waves](data/fractal_polar.png)

```
go run cmd/fractal_flame/main.go --width 4000 --height 4000 --iterations 210000 --trans bubble --trans_count 5 --symmetry 60 --gamma 1.0 --mode single
```
```
Запуск генерации изображения фрактального пламени с параметрами:
Ширина: 4000
Высота: 4000
Итерации: 210000
Функция трансформации: bubble
Количество трансформаций: 5
Число симметрий: 60
Гамма-коррекция: 1.000000
Режим выполнения: single
Время выполнения: 5.0768386s
Фрактальное пламя сгенерировано и сохранено в файл: data\fractal_bubble.png
```
![fractal_waves](data/fractal_bubble.png)

```
go run cmd/fractal_flame/main.go --width 3500 --height 3500 --iterations 200000 --trans spherical --trans_count 7 --symmetry 90 --gamma 1.0 --mode multi
```
```
Запуск генерации изображения фрактального пламени с параметрами:
Ширина: 3500
Высота: 3500
Итерации: 200000
Функция трансформации: spherical
Количество трансформаций: 7
Число симметрий: 90
Гамма-коррекция: 1.000000
Режим выполнения: multi
Время выполнения: 1.3746694s
Фрактальное пламя сгенерировано и сохранено в файл: data\fractal_spherical.png
```
![fractal_waves](data/fractal_spherical.png)

```
go run cmd/fractal_flame/main.go --width 3000 --height 3000 --iterations 150000 --trans sinusoidal --trans_count 15 --symmetry 75 --gamma 1.5 --mode multi
```
```
Запуск генерации изображения фрактального пламени с параметрами:
Ширина: 3000
Высота: 3000
Итерации: 150000
Функция трансформации: sinusoidal
Количество трансформаций: 15
Число симметрий: 75
Гамма-коррекция: 1.500000
Режим выполнения: multi
Время выполнения: 2.3328728s
Фрактальное пламя сгенерировано и сохранено в файл: data\fractal_sinusoidal.png
```
![fractal_waves](data/fractal_sinusoidal.png)

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)
