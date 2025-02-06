# Реализация онлайн библиотеки песен 🎶

## Описание задания
Необходимо реализовать REST API для онлайн библиотеки песен с использованием PostgreSQL.

### Требования:
1. **REST API методы:**
   - Получение данных библиотеки с фильтрацией по всем полям и пагинацией
   - Получение текста песни с пагинацией по куплетам
   - Удаление песни
   - Изменение данных песни
   - Добавление новой песни в формате JSON:
     ```json
     {
       "group": "Muse",
       "song": "Supermassive Black Hole"
     }
     ```

2. **Интеграция с внешним API**
   При добавлении песни необходимо сделать запрос в сторонний API, описанный с помощью OpenAPI (swagger). Реализовывать его не нужно, он будет поднят при проверке задания.

   **OpenAPI Specification:**
   ```yaml
   openapi: 3.0.3
   info:
     title: Music info
     version: 0.0.1
   paths:
     /info:
       get:
         parameters:
           - name: group
             in: query
             required: true
             schema:
               type: string
           - name: song
             in: query
             required: true
             schema:
               type: string
         responses:
           '200':
             description: Ok
             content:
               application/json:
                 schema:
                   $ref: '#/components/schemas/SongDetail'
           '400':
             description: Bad request
           '500':
             description: Internal server error
   components:
     schemas:
       SongDetail:
         required:
           - releaseDate
           - text
           - link
         type: object
         properties:
           releaseDate:
             type: string
             example: 16.07.2006
           text:
             type: string
             example: |
               Ooh baby, don't you know I suffer?
               Ooh baby, can you hear me moan?
               You caught me under false pretenses
               How long before you let me go?

               Ooh
               You set my soul alight
               Ooh
               You set my soul alight
           link:
             type: string
             example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
   ```

3. **Сохранение данных в базу PostgreSQL:**
   - Структура базы данных должна быть создана с помощью миграций при старте сервиса.

4. **Логирование:**
   - Покрыть код debug- и info-логами.

5. **Конфигурация:**
   - Все конфигурационные данные должны быть вынесены в `.env` файл.

6. **Swagger-документация:**
   - Сгенерировать swagger-документацию для реализованного API.

---

## Инструкция по запуску с Docker Compose

1. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/artyom-kalman/go-song-library.git
   ```

2. Создайте `.env` файл для настройки переменных окружения (пример `.env` файла):
   ```env
    # Database connection config
    DB_HOST=db
    DB_PORT=5432
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=password
    POSTGRES_DB=song_lib

    # Server config
    APP_PORT=:3030

    # SongInfoAPI config
    SONG_INFO_API=someapi
   ```

3. Запустите Docker Compose:
   ```sh
   docker-compose up --build
   ```
