Запуск с помощью ```docker-compose up```

----

Доступные запросы через gateway

- ```GET localhost:8080/api/profile/:id```

Аггрегация данных пользователя. Кэшируется через redis

Требует заголовок Authorization (типичный, с JWT)

- ```POST localhost:8080/api/signup```

```
{
    "username":"test",
    "password":"test"
}
```


- ```POST localhost:8080/api/login```

```
{
    "username":"test",
    "password":"test"
}
```

Добавить в базу новый продукт. требует jwt
- ```POST localhost:8080/api/products```

```
{
    "name": "Some product",
    "price": 234.12
}
```

Добавить в базу новый заказ. Требует jwt и добавляет заказ от имени пользователя, которому выдан jwt
- ```POST localhost:8080/api/orders```

```
{
    "products": [
        {
            "product_id": 1,
            "quantity": 5666
        },
        {
            "product_id": 2,
            "quantity": 727
        }
    ]
}
```

-----
по адресу ```localhost:3000/dashboards``` можно получить доступ к инстансу графаны с графиком одной базовой метрики (кол-во запросов на каждый эндпоинт)