# Обновленное API

![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414204013.png)
Примеры запросов в api/bannerify.postman_collection.json
ссылка на api

https://app.swaggerhub.com/apis/guresq/banners/1.0.0

## Authorization  
### Register 
Для регистрации пользователя используется следующий запрос. На вход принимается JSON, содержащий следующие поля:```
```  
"name": "<string>",   
"email": "<email>",   
"password": "<string>",   
"passwordConfirm": "<string>",   
"role": "<string>" 
```
Полученный в ответ токен действует в течение часа и сразу записывается в cookies. Полученный  токен можно передавать в заголовке `token`

Реализована проверка на повторяющий email или неправильный повторный пароль

role может быть "admin" и "user", в случае если поле указано пустым то оно по умолчанию будет user, так сделано чтобы была возможность расширить список ролей

### Login

Для авторизации уже зарегистрированного пользователя используется следующий запрос. На вход принимается JSON, содержащий следующие поля:

```
 "email": "<email>",
 "password": "<string>" 
 ```

В ответ приходит токен, который также действует один час.

### Logout

Выход из текущей сессии. При вызове этой функции удаляется текущий токен.

Все остальное API не изменено.

### Примеры использования api
1. Зарегистрируем нового пользователя с правами администратора
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414210452.png)
2. Создаем новый баннер ![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414210810.png) 

Баннер будет активный для всех пользователей, в случае если параметр is_active был бы false, тогда пользователь с правами user не смог получить бы к нему доступ
3. Получим баннер для пользователя, если параметр use_last_revision true, мы будем брать информацию о баннере напрямую из базы данных в ином случае мы получим ее и запишем в кеш 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414210938.png)
4. Обновим ранее созданный баннер, обновлять можно вписывая любые поля описанные в api, так же в БД хранится дата создание и обновления баннера 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414211424.png)
5. Получим весь список баннеров, так же можно получить баннеры по фиче или тегу или указать offset и limit для баннеров 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414211038.png)
6. Удалим баннер c id = 1 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414212158.png)
7. Проверим весь список  баннеров 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414212233.png)
8. Выполним выход из сессии 
![alt text](https://github.com/mirustal/banners/blob/develop/image/Pasted_image_20240414212351.png)

# Запуск 
	make start - полный запуск
	make build - сборка Docker образа
	make up - запуск всех сервисов через Docker Compose
	make down - остановка всех сервисов
	make restart - перезапуск всех сервисов
	make clean - удаление всех созданных Docker образов и контейнеров
	make logs - вывод логов приложения
## Линтер
	make lint


