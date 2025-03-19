
# Решение тестового задания на позицию Junior Golang Developer для Effective-Mobile

# Запуск:
```bash
git clone https://github.com/MaksimUlitin/effective-mobile.git
```
```bash
make all
```

# Документация по методам API

- [Вызов методов](#вызов-методов)

Методы:

- [Songs](#songs)
    - [Add song information](#add-song-information)
    - [List songs](#list-songs)
    - [List songs with filter](#list-songs-with-filter)
    - [Update an existing song](#update-an-existing-song)
    - [Delete a song](#delete-a-song)
    - [Get song text by ID with pagination](#get-song-text-by-id-with-pagination)

---

## Вызов методов

Методы вызываются через HTTP с использованием методов GET, POST, PATCH или DELETE.

Пример вызова метода:

```
http://localhost:8080/<resource>?<params>
```


---

## Songs

### Add song information

Добавляет информацию о песне.

#### URL

```
POST /info
```

#### Параметры

| Параметр | Тип   | Описание             | Обязательный |
|----------|-------|----------------------|--------------|
| group    | string | Название группы      | Да           |
| song     | string | Название песни       | Да           |

#### Пример запроса

```json
{
  "group": "Muse",
  "song": "Supermassive Black Hole"
}
```

#### Пример ответа

```json
{
  "group_name": "Muse",
  "song_name": "Supermassive Black Hole",
  "release_date": "16.07.2006",
  "text": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
  "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
}
```

---

### List songs with filter

Возвращает список песен с фильтрацией.

#### URL

```
GET /songs
```

#### Параметры фильтрации

| Параметр | Тип        | Описание                     | Обязательный |
|----------|------------|------------------------------|--------------|
| group     | string     | Фильтр по названию группы	   | Нет          |
| song    | string     | Фильтр по названию песни     | Нет          |
| release_date    | DD.MM.YYYY | Фильтр по названию группы    | Нет          |
| text     | string     | Фильтр по тексту песни       | Нет          |
| link     | string     | Фильтр по ссылке песни       | Нет          |
|page      | int        | Фильтр по номеру страницы    | нет          |    
|limit      | int        | Кол-во элементов на странице | нет          |    
#### Пример запроса

```
GET http://localhost:8080/songs?page=1&limit=10
```

#### Пример ответа

```json
[
  {
    "id": 1,
    "group_id": 1,
    "group_name": "Muse",
    "song": "Supermassive Black Hole",
    "release_date": "2006-07-16",
    "text": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
    "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
    "created_at": "2025-03-19T23:55:19.298701+03:00",
    "updated_at": "2025-03-19T23:55:19.298701+03:00"
  },
  {
    "id": 2,
    "group_id": 1,
    "group_name": "Muse",
    "song": "Supermassive Black Hole",
    "release_date": "2006-07-16",
    "text": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
    "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
    "created_at": "2025-03-20T00:04:08.920553+03:00",
    "updated_at": "2025-03-20T00:04:08.920553+03:00"
  }
]
```

---

### Update an existing song

Обновляет информацию о существующей песне.

#### URL

```
PATCH /songs/{id}
```

#### Параметры

| Параметр     | Тип   | Описание                    | Обязательный |
|--------------|-------|-----------------------------|--------------|
| id (в URL)   | int   | Идентификатор песни         | Да           |
| group        | string | Название группы            | Да           |
| song         | string | Название песни             | Да           |
| release_date | string | Дата выпуска               | Нет          |
| text         | string | Текст песни                | Нет          |
| link         | string | Ссылка на песню            | Нет          |

#### Пример 1-го запроса

```json
{
  "group_name": "ABBA",
  "link": "https://youtu.be/XEjLoHdbVeE?si=v-5B9ZNQs5b2G5-W",
  "release_date": "02.10.1979",
  "song": "Gimme! Gimme! Gimme!",
  "text": "Gimme, gimme, gimme a man after midnight\n Won't somebody help me chase the shadows away?\n Gimme, gimme, gimme a man after midnight\n Take me through the darkness to the break of the day"
}
```

#### Пример 1-го ответа

```json
{
  "message": "song updated successfully",
  "updated_fields": [
    "group_name",
    "title",
    "release_date",
    "text",
    "link"
  ]
}
```
#### GET /songs

```json
[
  {
    "id": 2,
    "group_id": 1,
    "group_name": "Muse",
    "song": "Supermassive Black Hole",
    "release_date": "2006-07-16",
    "text": "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
    "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
    "created_at": "2025-03-20T00:04:08.920553+03:00",
    "updated_at": "2025-03-20T00:04:08.920553+03:00"
  },
  {
    "id": 1,
    "group_id": 1,
    "group_name": "ABBA",
    "song": "Gimme! Gimme! Gimme!",
    "release_date": "1979-10-02",
    "text": "Gimme, gimme, gimme a man after midnight Won t somebody help me chase the shadows away? Gimme, gimme, gimme a man after midnight Take me through the darkness to the break of the day",
    "link": "https://youtu.be/XEjLoHdbVeE?si=v-5B9ZNQs5b2G5-W",
    "created_at": "2025-03-19T23:55:19.298701+03:00",
    "updated_at": "2025-03-20T00:17:46.09523+03:00"
  }
]
```

#### Пример 2-го запроса

```json
{
  "text": "la-la-la la-la-la la-la-la"
}
```
#### Пример 2-го ответа 
```json
{
  "message": "song updated successfully",
  "updated_fields": [
    "text"
  ]
}
```
#### GET /songs

```json
[
  {
    "id": 1,
    "group_id": 1,
    "group_name": "ABBA",
    "song": "Gimme! Gimme! Gimme!",
    "release_date": "1979-10-02",
    "text": "Gimme, gimme, gimme a man after midnight Won t somebody help me chase the shadows away? Gimme, gimme, gimme a man after midnight Take me through the darkness to the break of the day",
    "link": "https://youtu.be/XEjLoHdbVeE?si=v-5B9ZNQs5b2G5-W",
    "created_at": "2025-03-19T23:55:19.298701+03:00",
    "updated_at": "2025-03-20T00:17:46.09523+03:00"
  },
  {
    "id": 2,
    "group_id": 1,
    "group_name": "Muse",
    "song": "Supermassive Black Hole",
    "release_date": "2006-07-16T04:00:00+04:00",
    "text": "la-la-la la-la-la la-la-la",
    "link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
    "created_at": "2025-03-20T00:04:08.920553+03:00",
    "updated_at": "2025-03-20T00:24:11.574784+03:00"
  }
]
```

---

### Delete a song

Удаляет песню по ID.

#### URL

```
DELETE /songs/{id}
```

#### Параметры

| Параметр    | Тип   | Описание            | Обязательный |
|-------------|-------|---------------------|--------------|
| id   | int   | Идентификатор песни | Да           |

#### Пример запроса

```
DELETE http://localhost:8080/songs/1
```

#### Пример 1-го ответа

```json
{
  "message": "song deleted successfully"
}
```
#### Пример 2-го ответа
```json
{
  "message": "song already deleted or does not exist"
}
```
---

### Get song text by ID with pagination

Возвращает текст песни с поддержкой пагинации.

#### URL

```
GET /songs/{id}/text
```

#### Параметры

| Параметр     | Тип   | Описание                    | Обязательный |
|--------------|-------|-----------------------------|--------------|
| id    | int   | Идентификатор песни         | Да           |
| page         | int   | Номер страницы текста       | Нет          |
| limit        | int   | Количество строк на странице | Нет         |

#### Пример запроса

```
GET http://localhost:8080/songs/2/text?page=1&limit=10
```

#### Пример ответа

```json
{
  "limit": 10,
  "page": 1,
  "songId": 2,
  "text": [
    "la-la-la la-la-la la-la-la"
  ],
  "total": 1,
  "totalPage": 1
}
```