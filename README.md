# DataLab Drinks Backend

## Run the server

```bash
GIN_MODE=release go run src/main.go
```

The server will start listening on port 5000

## Run the unit and integration test

```bash
GIN_MODE=test go test ./...
```

## API

- GET: `/v1/menus`
    - return Cama's menu in json format

```json
{
    "status_message": "ok",
    "payload": {
        "menu_version": "2020W",
        "menu": [
            {
                "series": "現烘義式",
                "items": [
                    {
                        "item_id": 1,
                        "item": "黑咖啡",
                        "medium_price": 45,
                        "large_price": 60,
                        "cold": true,
                        "hot": true
                    }
                ]
            }
        ],
        "sugar": [
            {
                "sugar_id": 1,
                "sugar_tag": "無糖"
            },
            {
                "sugar_id": 2,
                "sugar_tag": "微糖"
            },
            {
                "sugar_id": 3,
                "sugar_tag": "半糖"
            },
            {
                "sugar_id": 4,
                "sugar_tag": "正常糖"
            }
        ],
        "ice": [
            {
                "ice_id": 1,
                "ice_tag": "熱"
            },
            {
                "ice_id": 2,
                "ice_tag": "去冰"
            },
            {
                "ice_id": 3,
                "ice_tag": "少冰"
            },
            {
                "ice_id": 4,
                "ice_tag": "正常冰"
            }
        ]
    }
}
```

- GET: `/v1/orders`
    - return the orders within `this week`!
    - the new week is defined if time pasts at 16:00 Friday
        - 15:59 Friday (this week)
        - 16:00 Friday (next week)
    - in json format
    - order_time follows RFC3339 format

```json
{
    "status_message": "ok",
    "payload": {
        "week_orders": [
            {
                "order_by": "王平郁",
                "item": "紅茶拿鐵",
                "size": "large",
                "price": 80,
                "sugar_tag": "微糖",
                "ice_tag": "熱",
                "order_time": "2021-02-10T23:46:49+08:00"
            },
            {
                "order_by": "何星緯",
                "item": "特調咖啡",
                "size": "large",
                "price": 70,
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "order_time": "2021-02-10T23:47:23+08:00"
            }
        ]
    }
}
```

- POST: `/v1/oders`
    - requires json body
    - if you want to update the order, just re-post
        - The server considers the value of `order_by` as key

```json
{
    "order_by": "平郁",
    "size": "medium",
    "item_id": 10,
    "sugar_id": 3,
    "ice_id": 2
}
```

## TODO:
- order validation (done)
- complete integrateion test (done)
- return an order report (in development)
- dockerize (in development)
- get history orders, not only just get this week's order (in development)

