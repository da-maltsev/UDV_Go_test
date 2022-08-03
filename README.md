В папке Sripts лежат два скрипта - на создание таблиц и на их заполнение.

Для наглядности есть ещё [схема БД](https://github.com/da-maltsev/UDV_Go_test/blob/master/scheme.png)

Сервер запускается через - [go run main.go](https://github.com/da-maltsev/UDV_Go_test/blob/master/main.go)

# UDV_Go_test
Тестовое задание

Разработка REST API Web-сервиса с использованием gorilla/mux и gorm

Приложение для каталогизации книг (библиотека), модель данных должна учитывать следующую информацию:

	- информацию о книге: заголовок, краткое описание, список автором, год, издание, издательство и т.п. (таблица-справочник)
	- информация об издательстве (таблица-справочник)
	- информация об экземплярах книги (количество и т.п.)
	- информация о размещении экземпляров книг (например, номер стеллажа и др. данные, которые позволяют идентифицировать размещение)
	- информация об пользователях библиотеки
	- информация о статусе (выдана или нет, кому в данный момент выдана) в данном случае, для уровня выше джуниор должна быть реализована история выдачи экземпляра книги
Что необходимо реализовать:

    1. Подготовить модель данных и sql-скрипты для инициализации начальными данными для тестирования и проверки
    2. Для уровня Junior должны быть реализованы следующие REST API эндпоинты:
* GET /api/book[/?page={page_number}&size={size}] метод для получения списка книг с использование пэйджинации
* GET /api/book **(да в ТЗ указано именно так, без ".{id}", иначе бы был реализован запрос к конкретной книге по её id по тому же принципу, как в "book/{id}/items")** метод для получения одного объекта книги из справочника
* GET /api/book/{id}/items – метод для получения информации об экземплярах книги по идентификатору справочника
