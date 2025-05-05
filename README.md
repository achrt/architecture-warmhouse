# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удалённо включать/выключать отопление в своих домах.
- Система поддерживает датчики и реле определенного вида.
- Датчики могут быть добавлены в систему только при помощи выездного специалиста компании
- Датчики регистрируются в системе при включении

**Мониторинг температуры:**

- Пользователи могут проверять температуру в доме через web приложение.
- Система запрашивает данные с датчиков синхронно во время запроса от пользователя к web приложению. Т.е. пользователь заходит в интерфейс просмотра температуры и в этот момент отправляется запрос с сервера на датчики.

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- База данных: PostgreSQL
- Архитектура: Монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения.
- Взаимодействие: Синхронное, запросы обрабатываются последовательно. 
- Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям.
- Развертывание: Требует остановки всего приложения.

### 3. Определение доменов и границы контекстов

__Домен: Управление отоплением__ 
* контроль температуры
* установка режима отопления
* получение данных с датчиков

__Домен: Установка и подключение__ 

_В текущей архитектуре выделен как отдельный домен, т.к. подключение возможно лишь через выездного инженера компании и требует организационного процесса._
 * организация выездов
 * установка оборудования
 * подключение оборудования

__Домен: Пользовательский доступ__ 
 * управление температурой
 * просмотр подключенных датчиков
 * просмотр счетов на оплату (гипотетически)

__Домен: Оплата__ 

_Точной информации об этом нет, но могу предположить что у компании присутствует определенная платежная модель. Не ясно, предоплатная она или постоплатная, но предположим что расчет постоплатный и зависит от кол-ва датчиков подключенных к системе._
 * учет использования ресурсов
 * генерация счетов
 * оплата и отслеживание платежей




### **4. Проблемы монолитного решения**

- Монолитная архитектура мешает масштабированию, т.к. нет возможности увеличить количество инстансов только для определенного функционала системы
- Сложности с разработкой и поддержкой: высокая стоимость ошибки - можно положить всю систему целиком после выкладки обновления 
- Если развертывание требует остановки всего приложения, то есть вероятность что в такие моменты система "лежит" не обрабатывает никакие запросы
- Синхронная последовательная обработка может порождать задержки в обработке запросов


### 5. Визуализация контекста системы — диаграмма С4

[Диаграмма](https://www.plantuml.com/plantuml/uml/nLJBJXHD5DxtKzG_ctz8W0d4bIiOlCG4ar5ZM9BgcH8gzdHFkgk3n9Wmu6M1qQ0iJ5nWTE5sk8mK0pIlSEeLV1A_asRQQHWtRjpGrAdpNSwbPYunCZPfE_3-qs4pI5igZGEnQamdkLchnN9jUaMRrTHF4nKtez2eq4mtetQj4yYGOLE9QGNQhotSc6fmwF7znLfR9aR522ql025rgEBf3aju3HKdKLXbiabHeJtAwDIzfMDy3wZdDgdFvyg4zsWT5EtbsyzCTJLggndEtaNMcPYfyFzVE9FwWgoW8qQwDsHfW2EJq1bTKDzj45zikGstVProzZVbxDzI3ZdhKHgsPBnUJQ8dPat6AeypwrUt0-2fswOJfcMvdnkxmdN9kkubSK-uBRfWRZfW2tIEo_UL2V7SyuGeUbXTK_vyfmFM9UMpPVfC5YRFS_81uFja0UwxJRwcJxXwIHc8eU5UCwcWJC1F1Yi3k7e7EEIOmhqgbzr3e3R66FN7zWEK9-Z8g5tPqRdTkcm-LM5Bu1JefZGw2jdVtOZ_qWuuZubTTmL4CZhCzSjij_ngInhfoqG1r8WIiu8sFrYi_76Fo4wMAyBPTRbWtFS4MNoE8FB2y_uNz07GZFQv8Ff1lNmkwDULGL76qsJIqNtoTH3eSAMOUZq_Ypp05kaZTuyRLaoYy0BFSCfjjgm2P_c9Qs70eAJBQ_TG1PTlg5W9shk26B-qa-CdIXdI8L6408KGwmmLYcdED-u9Ed3Rt6t-8g-kENEOCJ8d-WxUGiGAO7fGuOxoAnz5NUbXlXUdF4R03w-937yLblJK7SqTfc_ZwYfjDDfXiTF4B_qm7-l8F4gvgEno9AYrz5GisR5cPl-TcJayhhGT_080)

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

*To-Be домены

__Домен: Управление устройствами__ 
 * регистрация устройств
 * авторизация устройств
 * просмотр состояния
 * создание пользовательских сценариев
 * планировщики и триггеры

__Домен: Контроль доступа__ 
 * регистрация пользователей
 * авторизация пользователей
 * права и роли
 * связи пользователь ↔ устройство ↔ дом

__Домен: Оплата__ 
 * генерация/отображение счетов
 * проведение оплат
 * отслеживание платежей
 * тарифы и лимиты

__Домен: Телеметрия__ 
 * история событий и состояния
 * отображение графиков

__Домен: Автоматизация (сценарии)__
 * создание пользовательских сценариев
 * планировщики и триггеры
 * запуск действий над устройствами по условиям


__Домен: Уведомления__
 * алерты по датчикам
 * напоминания об оплате
 * триггеры по событиям


**Диаграмма контейнеров (Containers)**

[Диаграмма](https://www.plantuml.com/plantuml/uml/dLTRRnD757xthvZw1IE5M44yzOdWc42RG1eReJv5u_LWhzYBjJiBZQfA42vjvKfH4L8VAfLU-bh9AN5nBd7-mkm_wdV6wxringGgoDXpvfnlpdSkPuQReUI1Z5o7DMxiD7nFSjiJGUKJsxESg2kYm65zAGVXf_Lwm9ykzspPZpfHA08BkiAJovRlrWSEzmZZMYYxZjsfDsvSsoBHWysDkij3AO9w7drv0DrAPMFro_iFsZiFxxRlx6mqrvltrcfNAvKj4OI-LwCZbbZr0Rwg-5Rlr5ITn3-e3_Z-MutYFJMcDOkVgqdyN2zFuo6BNwXJTGQTzshArA6QOZcg0ZOGbkHUpn70UalYbIhxki9OQnVkkJlDhsJDzkMQU69ROeKE_8DGusTGFQ9VE7Iw0Y2pnVK55jTXyKsbaZCPy5qNOOE4JDx0rt7ydJgBz-DzfixL2L6BVrJ_O6C4KPuU4Sb32PVRJai4v021_GQePpWL1anDONEEvGY2OpLXMhbe_yXoONRR3aI7Xu8r73_giaqHXhndUxqawZE7NWBl6NcWOqeEfS23AEnVYrzWDPx56oPZU3zHHpWkvNxBZxmk3tPheLuJ-5-mdnO9ClKxo2CCyL0TiLcYgbUHACGzBP_QKz5P7Gm0yL1qoCrjmIsf_VqPODyYcjhw63uUm0yAyHvyoKEuVWAnwNTi50KM173RYIIgzNAi0XWG74GiIyMjcO0b4e9RzpNedqaW9_XN5xAEr1Xvnbc09k4I2e6sav2E4nK8alg0mWb4XBA7caSIA1ymcSGlus6HwLfdxjxQBJYmvOUo5uZM5nlaJhAZBGgCU2Jx6PrLh0nSVeLhxsVPeqABN-FtFdb3jK5Ko2Fr8TivD_KmFaTbb_Nv1dz9ZEQnBIvjtokdeQjRCNDxrffiatky9wZf_Yk79JC3oAahiowdgi34ARTzcKtYM6WWbEoKcKZXm64Px6PasdDHdiJd_D5ZNftr_W5vgslXbMxxixHmqZYVKeLfIXDG8T89nTcWcIQ2_LdtBfTvP6vf9Wt7jXxVyQDGkr8-X401KAugtzMfAzHM8cDD3tCdImwHcYBuXscJW1pICfpLKKRGN5XOWzCn_7X5tVMQjAb9oCAGeymtGvf2Ax07yaAQF5_QZsohnEXUJcgek5yAWtb4fNP2WtlU--PHIZE1CaO1ljYkjY_NkHHF-IvEkEktMRAY0-a3j1FAyKjTpHCQDMbiBglV6X_O6VZgmCw3RpTRxVfwm0Tz1AeKdSTKd5cHaYst-eAbw-OJwibG1eAxD2thpDMNZs1uOEYBK_-vmjGRwa4zyQPw82QEKerLjeMJF1RIQm6scTJrCwaMpzKoNfa9b8rxUZClB0v8Q9GLIdU2KSVK3tGOAHPoMfvi2yVMCLMDxfV3UBZ4Q9_QWN8Cu39iaY6wwgO4Hqdvh7N_NgAQR5yOHelq3NGAXQb-WcVePfJVH4zIOpVFertYd9-Xsm_dIkKe9-jKowXKV8Az--XLkm32tFdfYSRj_7FDE0ZI1bySpqnKe6YuC599KR5GjxJ8jXTxZeUYoM7zygDx0tyFuk_JLnzqxxJRMwtaIBBEkb2zBRmVHnzzFzABhB4971D9Snek7x7wka2hw3jgBpyDrIawHfCPwerzf09ulUsj1lbm9NcGhbnefDB_7EQ7fqp6XYaB1JpWWBiGxFh_0AlS55utSfr_0G00)

**Диаграмма компонентов (Components)**

Добавьте диаграмму для каждого из выделенных микросервисов.

**Диаграмма кода (Code)**

Добавьте одну диаграмму или несколько.

# Задание 3. Разработка ER-диаграммы

Добавьте сюда ER-диаграмму. Она должна отражать ключевые сущности системы, их атрибуты и тип связей между ними.

# Задание 4. Создание и документирование API

### 1. Тип API

Для взаимодействия между микросервисами будет использоваться gRPC. Миросервисы на Go, удобнее будет использовать gRPC, т.к. для него есть нативные плагины для кодогенерации, что удобнее чем go-swagger. Также gRPC использует http2, то есть работает быстрее и создает меньше трафика, что будет преимуществом при большом количестве мелких запросов, которые предполагает специфика приложения (большой поток данных с разных датчиков). Есть поддержка стриминга из коробки.

Проксирующий api-gateway для мобильного и web приложений будет иметь rest api.

### 2. Документация API

Здесь приложите ссылки на документацию API для микросервисов, которые вы спроектировали в первой части проектной работы. Для документирования используйте Swagger/OpenAPI или AsyncAPI.

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


